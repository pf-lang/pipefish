const BOOK_CSS =
    "https://pf-lang.github.io/pipefish/assets/pf-book.css";

const FONTS_CSS =
    "https://pf-lang.github.io/pipefish/assets/fonts.css";

const BOOK_COMPONENT =
    "https://pf-lang.github.io/pipefish/pf-book.js";

function normalizePath(path) {
    return path
        .replaceAll("\\", "/")
        .split("/")
        .filter(part => part && part !== ".")
        .join("/");
}

function asBytes(data) {
    if (data instanceof Uint8Array) {
        return data;
    }

    if (data instanceof ArrayBuffer) {
        return new Uint8Array(data);
    }

    if (ArrayBuffer.isView(data)) {
        return new Uint8Array(
            data.buffer,
            data.byteOffset,
            data.byteLength
        );
    }

    if (typeof data === "string") {
        return new TextEncoder().encode(data);
    }

    throw new TypeError("Book file data is not a byte array or string.");
}

function escapeHTML(value) {
    return String(value).replace(/[&<>"']/g, character => ({
        "&": "&amp;",
        "<": "&lt;",
        ">": "&gt;",
        '"': "&quot;",
        "'": "&#39;"
    })[character]);
}

function makeIndexHTML(bookTitle) {
    return `<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>${escapeHTML(bookTitle)}</title>

    <link rel="stylesheet" href="${BOOK_CSS}">
    <link rel="stylesheet" href="${FONTS_CSS}">
</head>
<body>

    <pf-book src="book/"></pf-book>

    <script type="module" src="${BOOK_COMPONENT}"></script>

</body>
</html>
`;
}

function collectFiles(directory, output) {
    if (directory.type === "file") {
        output.push(directory);
        return;
    }

    for (const child of directory.children ?? []) {
        collectFiles(child, output);
    }
}

/**
 * Build the canonical exported representation of the currently open book.
 *
 * Returns:
 * {
 *     directory: "book-title",
 *     files: [
 *         { path: "index.html", data: Uint8Array },
 *         { path: "book/Contents.md", data: Uint8Array },
 *         ...
 *     ]
 * }
 *
 * All paths are relative to the exported book directory.
 */
export async function makeBookExport(bookTitle) {
    const tree = window.pipefishGetBookFileTree();

    if (!tree || tree.type !== "folder") {
        throw new Error("Could not read the current book's file tree.");
    }

    const bookFiles = [];
    collectFiles(tree, bookFiles);

    const files = [{
        path: "index.html",
        data: new TextEncoder().encode(makeIndexHTML(bookTitle))
    }];

    for (const file of bookFiles) {
        const relativePath = normalizePath(file.path);

        if (!relativePath) {
            continue;
        }

        const result = window.pipefishReadBookFile(relativePath);

        if (!result?.ok) {
            throw new Error(
                `Could not read ${relativePath}: ${result?.error ?? "Unknown error"}`
            );
        }

        files.push({
            path: `book/${relativePath}`,
            data: asBytes(result.data)
        });
    }

    return files;
}

function crc32(bytes) {
    let crc = 0xffffffff;

    for (const byte of bytes) {
        crc ^= byte;

        for (let bit = 0; bit < 8; bit++) {
            crc = (crc >>> 1) ^
                (0xedb88320 & -(crc & 1));
        }
    }

    return (crc ^ 0xffffffff) >>> 0;
}

function write16(view, offset, value) {
    view.setUint16(offset, value, true);
}

function write32(view, offset, value) {
    view.setUint32(offset, value >>> 0, true);
}

function concatenate(parts) {
    const length = parts.reduce(
        (total, part) => total + part.byteLength,
        0
    );

    const result = new Uint8Array(length);
    let offset = 0;

    for (const part of parts) {
        result.set(part, offset);
        offset += part.byteLength;
    }

    return result;
}

function makeZip(files) {
    const encoder = new TextEncoder();
    const localParts = [];
    const centralParts = [];

    let localOffset = 0;

    // DOS timestamps: use the current local time.
    const now = new Date();
    const dosTime =
        (now.getHours() << 11) |
        (now.getMinutes() << 5) |
        Math.floor(now.getSeconds() / 2);

    const dosDate =
        ((Math.max(1980, now.getFullYear()) - 1980) << 9) |
        ((now.getMonth() + 1) << 5) |
        now.getDate();

    for (const file of files) {
        const name = encoder.encode(file.path);
        const data = asBytes(file.data);
        const crc = crc32(data);

        // Local file header.
        const localHeader = new Uint8Array(30 + name.length);
        const localView = new DataView(localHeader.buffer);

        write32(localView, 0, 0x04034b50);
        write16(localView, 4, 20);
        write16(localView, 6, 0x0800); // UTF-8 filenames
        write16(localView, 8, 0);      // Stored, no compression
        write16(localView, 10, dosTime);
        write16(localView, 12, dosDate);
        write32(localView, 14, crc);
        write32(localView, 18, data.length);
        write32(localView, 22, data.length);
        write16(localView, 26, name.length);
        write16(localView, 28, 0);
        localHeader.set(name, 30);

        localParts.push(localHeader, data);

        // Central directory entry.
        const centralHeader = new Uint8Array(46 + name.length);
        const centralView = new DataView(centralHeader.buffer);

        write32(centralView, 0, 0x02014b50);
        write16(centralView, 4, 20);
        write16(centralView, 6, 20);
        write16(centralView, 8, 0x0800);
        write16(centralView, 10, 0);
        write16(centralView, 12, dosTime);
        write16(centralView, 14, dosDate);
        write32(centralView, 16, crc);
        write32(centralView, 20, data.length);
        write32(centralView, 24, data.length);
        write16(centralView, 28, name.length);
        write16(centralView, 30, 0);
        write16(centralView, 32, 0);
        write16(centralView, 34, 0);
        write16(centralView, 36, 0);
        write32(centralView, 38, 0);
        write32(centralView, 42, localOffset);
        centralHeader.set(name, 46);

        centralParts.push(centralHeader);

        localOffset += localHeader.length + data.length;
    }

    const centralDirectory = concatenate(centralParts);

    const end = new Uint8Array(22);
    const endView = new DataView(end.buffer);

    write32(endView, 0, 0x06054b50);
    write16(endView, 4, 0);
    write16(endView, 6, 0);
    write16(endView, 8, files.length);
    write16(endView, 10, files.length);
    write32(endView, 12, centralDirectory.length);
    write32(endView, 16, localOffset);
    write16(endView, 20, 0);

    return concatenate([
        ...localParts,
        centralDirectory,
        end
    ]);
}

export async function downloadBookExport(files, directory) {
    const zipFiles = files.map(file => ({
        path: `${directory}/${file.path}`,
        data: file.data
    }));

    const zip = makeZip(zipFiles);
    const blob = new Blob([zip], {
        type: "application/zip"
    });

    const url = URL.createObjectURL(blob);
    const anchor = document.createElement("a");

    anchor.href = url;
    anchor.download = `${directory}.zip`;

    document.body.append(anchor);
    anchor.click();
    anchor.remove();

    setTimeout(() => URL.revokeObjectURL(url), 1000);
}