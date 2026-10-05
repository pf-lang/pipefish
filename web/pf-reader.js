// The base visual element that goes inside the `pf-editor` which goes inside
// the `pf-coder` which goes inside the `pf-ide`, etc.

class PipefishReader extends HTMLElement {
    constructor() {
        super();

        this.highlighter =
            document.createElement("pf-highlighter");

        this.ready = this.highlighter.ready;

        this.files = [];
        this.currentFile = null;

        const shadow =
            this.attachShadow({ mode: "open" });

        const style =
            document.createElement("link");
        style.rel = "stylesheet";
        style.href =
            new URL("./assets/pf-reader.css", import.meta.url);

        const syntax =
            document.createElement("link");
        syntax.rel = "stylesheet";
        syntax.href =
            new URL("./assets/syntax.css", import.meta.url);

        const tabs =
            document.createElement("div");
        tabs.classList.add("tabs");

        const box =
            document.createElement("pre");
        box.classList.add("reader");

        shadow.append(style, syntax, tabs, box);

        this.tabs = tabs;
        this.box = box;
    }

    connectedCallback() {
        this.initialize(this.textContent);
    }

    async initialize(source) {
        const files = await this.loadFiles(source);

        this.files = files;

        if (!files.length) {
            this.currentFile = null;
            await this.display("");
            return "";
        }

        this.currentFile = files[0].path;
        this.makeTabs();
        await this.display(files[0].data);

        return files[0].data;
    }

    async loadFiles(source) {
        const lines =
            source.replace(/\r\n/g, "\n").split("\n");

        if (
            lines.length === 0 ||
            !lines[0].trim().startsWith("./")
        ) {
            return [
                {
                    path: "main.pf",
                    data: normalizeSource(source)
                }
            ];
        }

        const base =
            new URL(lines[0].trim(), document.baseURI);

        const rootFiles =
            this.getRootFiles(lines.slice(1));

        const files = [];

        for (const name of rootFiles) {
            const response =
                await fetch(new URL(name, base), {
                    cache: "no-store"
                });;

            if (!response.ok) {
                throw new Error(
                    `Unable to load ${name}: ` +
                    `${response.status} ${response.statusText}`
                );
            }

            files.push({
                path: name,
                data: await response.text()
            });
        }

        return files;
    }

    getRootFiles(lines) {
        let rootIndent = Infinity;

        for (const line of lines) {
            if (!line.trim()) {
                continue;
            }

            const match = line.match(/^[ \t]*/);
            const indent =
                match[0].replace(/\t/g, "    ").length;

            rootIndent =
                Math.min(rootIndent, indent);
        }

        if (rootIndent === Infinity) {
            return [];
        }

        const files = [];

        for (const line of lines) {
            if (!line.trim()) {
                continue;
            }

            const match = line.match(/^([ \t]*)(.*)$/);
            const indent =
                match[1].replace(/\t/g, "    ").length;

            if (indent !== rootIndent) {
                continue;
            }

            const name = match[2].trim();

            if (!name.endsWith("/")) {
                files.push(name);
            }
        }

        return files;
    }

    makeTabs() {
        this.tabs.innerHTML = "";

        for (const file of this.files) {
            const tab =
                document.createElement("button");

            tab.classList.add("tab");
            tab.textContent = file.path;

            if (file.path === this.currentFile) {
                tab.classList.add("selected");
            }

            tab.addEventListener("click", () => {
                this.selectFile(file.path);
            });

            this.tabs.append(tab);
        }
    }

    async selectFile(path) {
        const file =
            this.files.find(file => file.path === path);

        if (!file) {
            return;
        }

        this.currentFile = path;

        for (const tab of this.tabs.children) {
            tab.classList.toggle(
                "selected",
                tab.textContent === path
            );
        }

        await this.display(file.data);

        this.dispatchEvent(new CustomEvent("filechange", {
            detail: {
                path: file.path,
                data: file.data
            }
        }));
    }

    async display(source) {
        await this.ready;

        const extension =
            this.currentFile
                .slice(this.currentFile.lastIndexOf("."))
                .toLowerCase();

        const highlighters = {
            ".pf": source => this.highlighter.highlight(source)
        };

        const highlighter =
            highlighters[extension] || (source => source);

        this.box.innerHTML =
            await highlighter(source) +
            (source.endsWith("\n") ? "\u00a0" : "");
    }

    get scrollTop() {
        return this.box.scrollTop;
    }

    set scrollTop(value) {
        this.box.scrollTop = value;
    }

    get scrollLeft() {
        return this.box.scrollLeft;
    }

    set scrollLeft(value) {
        this.box.scrollLeft = value;
    }
}

// This strips superfluous indentation from the source code given between
// `<pf-reader></pf-reader>` tags.

function normalizeSource(source) {
    const lines = source.split("\n");

    // Remove blank lines at the beginning and end.
    while (lines.length && !lines[0].trim()) {
        lines.shift();
    }

    while (lines.length && !lines[lines.length - 1].trim()) {
        lines.pop();
    }

    if (!lines.length) {
        return "";
    }

    // Find the common indentation of the non-blank lines.
    let indent = Infinity;

    for (const line of lines) {
        if (!line.trim()) {
            continue;
        }

        const match = line.match(/^[ \t]*/);
        indent = Math.min(indent, match[0].length);
    }

    // Remove the common indentation.
    for (let i = 0; i < lines.length; i++) {
        lines[i] = lines[i].slice(indent);
    }

    return lines.join("\n");
}

customElements.define("pf-reader", PipefishReader);