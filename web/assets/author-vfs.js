import {
    load,
    save
} from "./author-storage.js";

const SAMPLE_BOOK_FILES = [
    "Contents.md",
    "Components-overview.md",
    "Initializing-components-from-code.md",
    "Initializing-components-from-files.md",
    "Introduction.md",
    "Markdown.md",
    "The-Books-menu.md",
    "The-Files-menu.md",
    "The-structure-of-a-book.md",
    "project/main.pf",
    "project/other.pf",
    "project/README.md",
    "project/text/text.txt",
    "other-book/Contents.md",
    "other-book/Jabberwocky.md",
];

export class AuthorVFS {
    async loadSampleFiles() {
        return Promise.all(
            SAMPLE_BOOK_FILES.map(async name => {
                const path = `using-pf-author/${name}`;
                const url = new URL(
                    `using-pf-author/${name}`,
                    import.meta.url
                );

                const response = await fetch(url, {
                    cache: "no-store"
                });

                if (!response.ok) {
                    throw new Error(
                        `Failed to load ${path}: ${response.status}`
                    );
                }

                return {
                    path,
                    data: new TextEncoder().encode(
                        await response.text()
                    )
                };
            })
        );
    }

    async initialize() {
        await window.pipefishReady;

        const savedFiles = await load("authorFS");

        if (savedFiles !== undefined) {
            const result = window.pipefishInitAuthorFiles(savedFiles);

            if (!result?.ok) {
                throw new Error(
                    result?.error ?? "Failed to restore author files."
                );
            }
            this.installPersistence();
            return;
        }

        const files = await this.loadSampleFiles();

        const result = window.pipefishInitAuthorFiles(files);

        if (!result?.ok) {
            throw new Error(
                result?.error ?? "Failed to initialize author files."
            );
        }

        await this.save();
        this.installPersistence();
    }

    installPersistence() {
        window.addEventListener("pipefish-book-file-written", () => {
            this.save().catch(error => {
                console.error("Failed to persist author files:", error);
            });
        });
    }

    async revertSampleBook() {
        const sampleFiles = await this.loadSampleFiles();

        // Preserve every file outside the sample-book directory.
        const otherFiles = this.getFiles().filter(file =>
            file.path !== "using-pf-author" &&
            !file.path.startsWith("using-pf-author/")
        );

        // Reinitialize the author VFS with the preserved files and
        // the freshly loaded sample files.
        const result = window.pipefishInitAuthorFiles([
            ...otherFiles,
            ...sampleFiles
        ]);

        if (!result?.ok) {
            throw new Error(
                result?.error ?? "Failed to revert the sample book."
            );
        }

        await this.save();
    }


    getTree() {
        const result = window.pipefishGetAuthorFileTree();

        if (result?.error) {
            throw new Error(result.error);
        }

        return result;
    }

    openBook(path) {
        const result = window.pipefishOpenBook(path);

        if (!result?.ok) {
            throw new Error(
                result?.error ?? `Failed to open book: ${path}`
            );
        }

        return window.pipefishGetBookFileTree();
    }

    getFiles() {
        const result = window.pipefishGetAuthorFiles();

        if (!result?.ok) {
            throw new Error(
                result?.error ?? "Failed to read author files."
            );
        }

        return result.files;
    }

    async save() {
        const files = this.getFiles();
        await save("authorFS", files);
    }

    writeFile(path, data) {
        return this.mutate(
            window.pipefishWriteAuthorFile,
            path,
            data
        );
    }

    createDirectory(path) {
        return this.mutate(
            window.pipefishCreateAuthorDirectory,
            path
        );
    }

    deleteFile(path) {
        return this.mutate(
            window.pipefishDeleteAuthorFile,
            path
        );
    }

    deleteDirectory(path) {
        return this.mutate(
            window.pipefishDeleteAuthorDirectory,
            path
        );
    }

    renamePath(oldPath, newPath) {
        return this.mutate(
            window.pipefishRenameAuthorPath,
            oldPath,
            newPath
        );
    }

    async mutate(operation, ...args) {
        const result = operation(...args);

        if (!result?.ok) {
            throw new Error(
                result?.error ?? "Author filesystem operation failed."
            );
        }

        await this.save();
    }

    readFile(path) {
        const result = window.pipefishReadAuthorFile(path);

        if (!result?.ok) {
            throw new Error(
                result?.error ?? `Failed to read author file: ${path}`
            );
        }

        return result.data;
    }

}