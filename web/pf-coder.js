import "./pf-editor.js";
import "./pf-service.js";
import { setFlags } from "./assets/flags.js";

class PipefishCoder extends HTMLElement {
    constructor() {
        super();

        const shadow =
            this.attachShadow({ mode: "open" });

        const service =
            document.createElement("pf-service");

        const editor =
            document.createElement("pf-editor");

        editor.setAttribute("managed", "");

        shadow.append(
            service,
            editor
        );

        this.service = service;
        this.editor = editor;
        this.dirty = false;
        this.source = null;

        editor.addEventListener("change", () => {
            this.dirty = true;
        });

        this.ready = Promise.all([
            service.ready,
            editor.ready
        ]);
    }

    initialize(source, flags) {
        this.initialization = (async () => {
            setFlags(this, flags);

            await this.ready;

            this.source = source;

            await this.service.initialize(source, this.flags);
            await this.editor.initialize(source, this.flags);

            this.dirty = false;
        })();

        return this.initialization;
    }

    async build(fileExtractor) {
        await this.ready;

        // A standalone pf-coder is initialized on its first build.
        if (!this.initialization) {
            await this.initialize(this.textContent);
        } else {
            await this.initialization;
        }

        await this.editor.build(fileExtractor);
        await this.service.build(fileExtractor);
    }

    async compile() {
        await this.ready;

        if (!this.initialization) {
            throw new Error(
                "pf-coder must be initialized before compiling"
            );
        }

        await this.initialization;

        if (!this.dirty) {
            return;
        }

        await this.service.compileMain();
        this.dirty = false;
    }

    async getErrorReport() {
        await this.ready;
        return this.compileTimeErrors;
    }

    async do(line) {
        await this.ready;
        await this.compile();
        return this.service.do(line);
    }

    async getFilepath() {
        await this.ready;
        await this.compile();
        return this.service.getFilepath();
    }
}

customElements.define(
    "pf-coder",
    PipefishCoder
);
