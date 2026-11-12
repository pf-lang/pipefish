import { PfBook, parseContents } from "./pf-book.js";
import "./pf-writer.js";
import "./assets/file-menu.js";
import { FileManager } from "./assets/file-manager.js";
import { AuthorVFS } from "./assets/author-vfs.js";
import { loadWasm } from "./assets/pf-wasm.js";
import { makeAuthorFileExtractor } from "./assets/files.js";
import {
    load,
    save
} from "./assets/author-storage.js";
import "./pf-text.js";
import {
    classifyFile,
    decodeFileText,
    displayFile
} from "./assets/file-viewer.js";
import "./assets/publish.js";

class PfAuthor extends PfBook {
    constructor() {
        super();

        this.vfs = new AuthorVFS();
        this.fileManager = null;

        this.addEventListener("file-open", event => {
            this.openFile(event.detail.path, event.detail.data);
        });
    }

    async connectedCallback() {
        await loadWasm();
        await this.vfs.initialize();

        this.bookRoot = await load("openBook") ?? "using-pf-author";
        this.vfs.openBook(this.bookRoot);

        const result = window.pipefishReadBookFile("Contents.md");

        if (!result.ok) {
            throw new Error(result.error);
        }
        this.initializeBook(result.data);
        this.addUnusedPages();
        this.applyTheme();
        this.render();
        await this.showInitialPage();

        window.addEventListener("hashchange", () => {
            this.showPage(location.hash.slice(1));
        });
    }

    addUnusedPages() {
        const usedFiles = new Set(
            this.contents.flatMap(section =>
                section.pages.map(page => page.filename)
            )
        );

        const tree = window.pipefishGetBookFileTree();

        const unusedPages = tree.children
            .filter(file =>
                file.type === "file" &&
                file.name.endsWith(".md") &&
                !usedFiles.has(file.name)
            )
            .map(file => ({
                title: file.name.slice(0, -3).replace(/-/g, " "),
                filename: file.name,
                hash: file.name.slice(0, -3).toLowerCase()
            }));

        if (unusedPages.length > 0) {
            this.contents.push({
                title: "Unused pages",
                pages: unusedPages
            });
        }
    }

    get pageElementName() {
        return "pf-writer";
    }

    async loadText(src) {
        const result = window.pipefishReadBookFile("Contents.md");

        if (!result.ok) {
            throw new Error(result.error);
        }

        return result.data;
    }

    pageSrc(page) {
        return `bookfs:${page.filename}`;
    }

    async openBook(path) {
        path = path.replace(/^\.\//, "");

        this.vfs.openBook(path);
        this.bookRoot = path;

        const result = window.pipefishReadBookFile("Contents.md");

        if (!result.ok) {
            throw new Error(result.error);
        }

        const parsed = parseContents(result.data);

        this.header = parsed.header;
        this.bookTitle = parsed.bookTitle;
        this.contents = parsed.sections;

        const usedFiles = new Set(
            this.contents.flatMap(section =>
                section.pages.map(page => page.filename)
            )
        );

        const tree = window.pipefishGetBookFileTree();

        const unusedPages = tree.children
            .filter(file =>
                file.type === "file" &&
                file.name.endsWith(".md") &&
                !usedFiles.has(file.name)
            )
            .map(file => ({
                title: file.name.slice(0, -3).replace(/-/g, " "),
                filename: file.name,
                hash: file.name.slice(0, -3).toLowerCase()
            }));

        if (unusedPages.length > 0) {
            this.contents.push({
                title: "Unused pages",
                pages: unusedPages
            });
        }

        const firstPage = this.contents[0]?.pages[0];

        if (firstPage) {
            if (location.hash !== `#${firstPage.hash}`) {
                location.hash = firstPage.hash;
            }

            this.render();
            await this.showPage(firstPage.hash);
        } else {
            this.render();
        }

        window.scrollTo(0, 0);

        await save("openBook", this.bookRoot);
    }

    async showPage(hash) {
        const writer = this.querySelector("pf-writer");
        const navigation = this.querySelector(".page-navigation");
        const fileView = this.getFileView();

        fileView.container.hidden = true;
        navigation.hidden = false;
        writer.hidden = false;

        writer.fileExtractor =
            makeAuthorFileExtractor(this.bookRoot);

        await super.showPage(hash);
    }

    async openFile(path, data) {
        const info = classifyFile(path, data);

        // Is this Markdown file one of the pages in the current book?
        if (info.type === "text" && info.highlighting === "markdown") {
            const page = this.contents
                .flatMap(section => section.pages)
                .find(page => page.filename === path);

            if (page) {
                await this.showPage(page.hash);
                return;
            }

            // Otherwise, display it as standalone Markdown.
            this.showStandaloneMarkdown(path, data);
            return;
        }

        this.showOtherFile(path, data);
    }

    showStandaloneMarkdown(path, data) {
        const writer = this.querySelector("pf-writer");
        const navigation = this.querySelector(".page-navigation");
        const fileView = this.getFileView();

        fileView.container.hidden = true;
        navigation.hidden = true;

        writer.hidden = false;
        writer.fileExtractor = makeAuthorFileExtractor(this.bookRoot);

        writer.setContent(
            decodeFileText(data),
            async markdown => {
                const result = this.vfs.writeFile(
                    path,
                    new TextEncoder().encode(markdown)
                );

                if (!result?.ok) {
                    throw new Error(result?.error ?? "Failed to save file");
                }
            }
        );
    }

    showOtherFile(path, data) {
        const writer = this.querySelector("pf-writer");
        const navigation = this.querySelector(".page-navigation");
        const fileView = this.getFileView();

        writer.hidden = true;
        navigation.hidden = true;
        fileView.container.hidden = false;

        if (this.fileViewImageURL) {
            URL.revokeObjectURL(this.fileViewImageURL);
            this.fileViewImageURL = null;
        }

        const result = displayFile({
            path,
            data,
            text: fileView.text,
            image: fileView.image,
            emptyState: fileView.emptyState,
            emptyMessage: fileView.emptyMessage
        });

        this.fileViewImageURL = result.imageURL;
    }

    getFileView() {
        if (this.fileView) {
            return this.fileView;
        }

        const main = this.querySelector(".book-main");

        const container = document.createElement("div");
        container.className = "author-file-view";
        container.hidden = true;

        const text = document.createElement("pf-text");
        const image = document.createElement("img");

        image.alt = "";

        const emptyState = document.createElement("div");
        emptyState.className = "author-file-empty";

        const emptyMessage = document.createElement("p");

        emptyState.append(emptyMessage);
        container.append(text, image, emptyState);
        main.append(container);

        this.fileView = {
            container,
            text,
            image,
            emptyState,
            emptyMessage
        };

        return this.fileView;
    }

    renderHeaderExtras(controls, themeToggle) {
        controls.classList.add("author-header-controls");

        const size = parseFloat(
            getComputedStyle(themeToggle).height
        );

        const filesMenu = document.createElement("pf-file-menu");
        filesMenu.setAttribute("size", String(size));
        filesMenu.projectName = "Files";

        const booksMenu = document.createElement("pf-menu");

        booksMenu.configure({
            label: "Books",
            icon: `
                <svg viewBox="0 0 24 24" aria-hidden="true">
                    <path
                        d="M12 6.5C9.5 4.5 6 4 2.5 5v14
                        c3.5-1 7-.5 9.5 1.5
                        c2.5-2 6-2.5 9.5-1.5V5
                        c-3.5-1-7-.5-9.5 1.5Z"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="1.6"
                        stroke-linejoin="round"
                    />
                    <path
                        d="M12 6.5v14"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="1.6"
                    />
                </svg>
            `,
            items: [
                "Open folder as book",
                "Publish book",
                "Download book",
                "Revert sample book",
            ],
            size
        });

        const files = document.createElement("pf-files");
        this.appendChild(files);

        this.fileManager = new FileManager(
            files,
            {
                getTree: () => this.vfs.getTree(),
                readFile: path => this.vfs.readFile(path),
                writeFile: (path, data) => this.vfs.writeFile(path, data),
                createDirectory: path => this.vfs.createDirectory(path),
                renamePath: (oldPath, newPath) =>
                    this.vfs.renamePath(oldPath, newPath),
                deleteFile: path => this.vfs.deleteFile(path),
                deleteDirectory: path => this.vfs.deleteDirectory(path),
                revertFile: path => this.vfs.revertFile(path),
                zipDirectory: path => this.vfs.zipDirectory(path),
                openBook: path => this.openBook(path)
            },
            {
                openPaths: paths => {
                    for (const path of paths) {
                        const result = this.vfs.readFile(path);

                        if (result.ok) {
                            this.dispatchEvent(new CustomEvent("file-open", {
                                detail: {
                                    path,
                                    data: result.data
                                },
                                bubbles: true,
                                composed: true
                            }));
                        }
                    }
                },
                openBook: path => this.openBook(path)
            }
        );

        filesMenu.addEventListener("menu-select", event => {
            this.fileManager.open(event.detail.toLowerCase());
        });

        booksMenu.addEventListener("menu-select", async event => {
            try {
                switch (event.detail) {
                    case "Open folder as book":
                        this.fileManager.open("open-book");
                        break;

                    case "Revert sample book": {
                        const refresh =
                            this.bookRoot === "using-pf-author";

                        await this.vfs.revertSampleBook();

                        if (refresh) {
                            await this.openBook(this.bookRoot);
                        }

                        break;
                    }

                    case "Publish book":
                        this.publisher.open({
                            bookTitle: this.bookTitle || this.bookRoot
                        });
                        break;

                    case "Download book":
                        break;
                }
            } catch (error) {
                console.error("Books menu operation failed:", error);
                alert(error.message);
            }
        });

        controls.append(filesMenu, booksMenu, themeToggle);

        const publisher = document.createElement("pf-publish");
        this.appendChild(publisher);
        this.publisher = publisher;
    }
}

customElements.define("pf-author", PfAuthor);