import "./pf-reader.js";
import { FileManager } from "../assets/file-manager.js";
import "../assets/file-menu.js";

class PipefishEditor extends HTMLElement {
    constructor() {
        super();

        const shadow = this.attachShadow({ mode: "open" });

        const style = document.createElement("link");
        style.rel = "stylesheet";
        style.href = new URL("../assets/pf-editor.css", import.meta.url);

        const editor = document.createElement("div");
        editor.classList.add("editor");

        const logo = document.createElement("img");
        logo.classList.add("logo");
        logo.src = new URL("../assets/icon.ico", import.meta.url);
        logo.alt = "";

        const projectMenu = document.createElement("pf-file-menu");
        projectMenu.setAttribute("size", "28");

        const controls = document.createElement("div");
        controls.classList.add("top-controls");

        const minimize = document.createElement("button");
        minimize.classList.add("top-control", "minimize");
        minimize.innerHTML = `
            <svg viewBox="0 0 16 16" aria-hidden="true">
                <path d="M4 3h8M4 11l4-4 4 4"
                    fill="none"
                    stroke="currentColor"
                    stroke-opacity=".8"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"/>
            </svg>
        `;
        minimize.setAttribute("aria-label", "Minimize");

        controls.append(logo, projectMenu, minimize);

        const reader = document.createElement("pf-reader");
        reader.text.editable = true;

        const files = document.createElement("pf-files");

        const manager = new FileManager(files, {
            getTree: () => window.pipefishGetFileTree(),
            readFile: path => window.pipefishReadFile(path),
            writeFile: (path, data) => window.pipefishUpdateFile(path, data),
            createDirectory: path => window.pipefishCreateDirectory(path),
            renamePath: (oldPath, newPath) =>
                window.pipefishRenamePath(oldPath, newPath),
            deleteFile: path => window.pipefishDeleteFile(path),
            deleteDirectory: path => window.pipefishDeleteDirectory(path),
            revertFile: path => window.pipefishRevertFile(path),
            zipDirectory: path => window.pipefishZipDirectory(path)
        }, {
            openPaths: paths => this.reader.openPaths(paths),

            renamePaths: async (oldPath, newPath) => {
                const oldPrefix = oldPath + "/";

                const renamed = path =>
                    path === oldPath
                        ? newPath
                        : path.startsWith(oldPrefix)
                            ? newPath + path.slice(oldPath.length)
                            : path;

                for (const file of this.reader.files) {
                    file.path = renamed(file.path);
                }

                this.reader.currentFile =
                    renamed(this.reader.currentFile);

                this.reader.makeTabs();

                if (this.reader.currentFile) {
                    await this.reader.selectFile(this.reader.currentFile);
                }
            },

            removePaths: paths => this.reader.removePaths(paths),

            revertPaths: async (target, type) => {
                const affectsPath = path => type === "file"
                    ? path === target
                    : target === "." ||
                        path === target ||
                        path.startsWith(target + "/");

                const currentPath = this.reader.currentFile;
                const removedPaths = [];

                for (const file of this.reader.files) {
                    if (!affectsPath(file.path)) {
                        continue;
                    }

                    const result = window.pipefishReadFile(file.path);

                    if (result.ok) {
                        file.data = result.data;
                    } else {
                        removedPaths.push(file.path);
                    }
                }

                if (removedPaths.length) {
                    await this.reader.removePaths(removedPaths);
                }

                if (
                    currentPath &&
                    affectsPath(currentPath) &&
                    !removedPaths.includes(currentPath)
                ) {
                    await this.reader.selectFile(currentPath);
                }
            }
        });

        projectMenu.addEventListener("menu-select", event => {
            manager.open(event.detail.toLowerCase());
        });

        editor.append(reader);
        shadow.append(style, controls, editor, files);

        this.reader = reader;
        this.ready = reader.ready;
        this.editor = editor;
        this.menu = projectMenu;
        this.files = files;
        this.fileManager = manager;
        this.minimizeButton = minimize;
        this.initialization = null;

        minimize.addEventListener("click", () => {
            const minimized =
                this.editor.classList.toggle("minimized");

            this.classList.toggle("minimized", minimized);

            this.minimizeButton.setAttribute(
                "aria-label",
                minimized ? "Maximize" : "Minimize"
            );

            this.dispatchEvent(
                new CustomEvent(
                    minimized ? "minimize" : "maximize",
                    {
                        bubbles: true,
                        composed: true
                    }
                )
            );
        });

        reader.addEventListener("input", async () => {
            const path = this.reader.currentFile;

            if (!path) {
                return;
            }

            const source = this.reader.text.value;
            const bytes = new TextEncoder().encode(source);

            const file = this.reader.files.find(
                file => file.path === path
            );

            if (file) {
                file.data = bytes;
            }

            await window.pipefishUpdateFile(path, bytes);

            this.dispatchEvent(new CustomEvent("change"));
        });

        reader.addEventListener("tabschange", () => {
            this.updateScrollbar();
        });

        reader.addEventListener("filechange", () => {
            this.setInitialHeight();
        });
    }

    connectedCallback() {
        this.updateScrollbar();
    }

    updateScrollbar() {
        this.classList.toggle(
            "no-open-files",
            this.reader.files.length === 0
        );
    }

    get value() {
        return this.reader.text.value;
    }

    set value(value) {
        this.reader.text.value = String(value ?? "");
        this.setInitialHeight();
    }

    async initialize(source) {
        this.initialization = (async () => {
            await this.ready;

            const firstLine = source.trim().split(/\r?\n/, 1)[0].trim();

            this.menu.projectName = firstLine.startsWith("./")
                ? firstLine.replace(/\/+$/, "").split("/").pop()
                : "⋮";

            await this.reader.initialize(source);
        })();

        return this.initialization;
    }

    async build(fileExtractor) {
        await this.ready;

        if (!this.initialization) {
            await this.initialize(this.textContent);
        } else {
            await this.initialization;
        }

        const bytes = await this.reader.build(fileExtractor);

        // A standalone editor owns the initial VFS setup.
        // A coder-managed editor leaves that to pf-service.
        if (!this.hasAttribute("managed")) {
            await window.pipefishCompile(this.reader.files);
        }

        this.setInitialHeight();

        return bytes;
    }

    async display(source) {
        await this.ready;

        await this.reader.display(
            new TextEncoder().encode(source)
        );

        this.setInitialHeight();
    }

    setInitialHeight() {
        if (!this.reader.text?.input) {
            return;
        }

        const input = this.reader.text.input;
        const lines = Math.max(1, input.value.split("\n").length);
        const visibleLines = Math.min(lines, 20);

        const inputStyle = getComputedStyle(input);
        const lineHeight = parseFloat(inputStyle.lineHeight);
        const padding =
            parseFloat(inputStyle.paddingTop) +
            parseFloat(inputStyle.paddingBottom);

        const editorStyle = getComputedStyle(this.editor);
        const borders =
            parseFloat(editorStyle.borderTopWidth) +
            parseFloat(editorStyle.borderBottomWidth);

        const tabsStyle = getComputedStyle(this.reader.tabs);
        const tabsHeight = parseFloat(tabsStyle.height) || 0;

        this.editor.style.height =
            `${tabsHeight + visibleLines * lineHeight + padding + borders}px`;
    }
}

customElements.define("pf-editor", PipefishEditor);