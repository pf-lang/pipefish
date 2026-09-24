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
        console.log("Coder constructed.")
    }

    async initialize(source) {
        await this.ready;
        console.log("coder.initialize called");
        await this.service.compile(source);
        await this.editor.initialize(source);
    }

    async compile() {
        await this.ready;
        console.log("coder.compile called");
        console.trace();
        return this.service.compileMain();
    }

    async do(line) {
        await this.ready;
        console.log("coder.do called");
        await this.service.compileMain();

        return this.service.do(line);
    }
}

customElements.define(
    "pf-coder",
    PipefishCoder
);