import "./pf-editor.js";
import "./pf-service.js";

class PipefishCoder extends HTMLElement {
    constructor() {
        super();

        const shadow =
            this.attachShadow({ mode: "open" });

        const service =
            document.createElement("pf-service");

        const editor =
            document.createElement("pf-editor");

        shadow.append(
            service,
            editor
        );

        this.service = service;
        this.editor = editor;

        this.ready = Promise.all([
            service.ready,
            editor.ready
        ]);
    }

    async initialize(source) {
        await this.ready;

        await this.editor.initialize(source);
        await this.service.compile(this.editor.value);
    }

    async compile() {
        await this.ready;

        return this.service.compile(
            this.editor.value
        );
    }

    async do(line) {
        await this.ready;

        await this.service.compile(
            this.editor.value
        );

        return this.service.do(line);
    }
}

customElements.define(
    "pf-coder",
    PipefishCoder
);