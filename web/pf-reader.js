// The base visual element that goes inside the `pf-editor` which goes inside the `pf-coder`
// which goes insde the `pf-ide`, etc.

class PipefishReader extends HTMLElement {
    constructor() {
        super();

        this.highlighter =
            document.createElement("pf-highlighter");

        this.ready = this.highlighter.ready;

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

        const box =
            document.createElement("pre");
        box.classList.add("reader");

        shadow.append(style, syntax, box);

        this.box = box;
    }

    connectedCallback() {
        this.display(this.textContent);
    }

    async initialize(source) {
        const normalized = normalizeSource(source)
        await this.display(
            normalized
        );
        return normalized
    }

    async display(source) {
        await this.ready;

        this.box.innerHTML =
            await this.highlighter.highlight(source);
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