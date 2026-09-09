import "./pf-highlighter.js";

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
            new URL("./pf-reader.css", import.meta.url);

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

    set padding(value) {
        this.style.setProperty("--reader-padding", value);
    }

    async display(source) {
        await this.ready;

        this.box.innerHTML =
            await this.highlighter.highlight(source);
    }
}

customElements.define("pf-reader", PipefishReader);