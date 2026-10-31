export const MAX_TEXT_FILE_SIZE = 512 * 1024;

const IMAGE_TYPES = {
    ".bmp": "image/bmp",
    ".jpeg": "image/jpeg",
    ".png": "image/png"
};

const TEXT_TYPES = {
    ".md": {
        highlighting: "markdown",
        editable: true,
        wrap: true
    },

    ".pf": {
        highlighting: "pipefish",
        editable: true,
        wrap: false
    },

    ".txt": {
        highlighting: null,
        editable: true,
        wrap: true
    }
};


function getExtension(path) {
    const filename = path.split("/").pop();
    const index = filename.lastIndexOf(".");

    return index < 0
        ? ""
        : filename.slice(index).toLowerCase();
}


function getSize(data) {
    if (data == null) {
        return 0;
    }

    if (typeof data === "string") {
        return new TextEncoder().encode(data).byteLength;
    }

    return data.byteLength;
}


export function classifyFile(path, data) {
    if (!path) {
        return {
            type: "empty"
        };
    }

    const extension = getExtension(path);
    const size = getSize(data);

    if (IMAGE_TYPES[extension]) {
        return {
            type: "image",
            extension,
            mimeType: IMAGE_TYPES[extension],
            editable: false,
            size
        };
    }

    const knownTextType = TEXT_TYPES[extension];

    if (knownTextType) {
        return {
            type: "text",
            extension,
            editable: knownTextType.editable,
            wrap: knownTextType.wrap,
            highlighting: knownTextType.highlighting,
            size
        };
    }

    if (size > MAX_TEXT_FILE_SIZE) {
        return {
            type: "too-large",
            extension,
            editable: false,
            size
        };
    }

    return {
        type: "text",
        extension,
        editable: true,
        wrap: false,
        highlighting: null,
        size
    };
}


export function decodeFileText(data) {
    if (typeof data === "string") {
        return data;
    }

    return new TextDecoder("utf-8").decode(data);
}