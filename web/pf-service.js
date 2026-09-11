class PipefishService extends HTMLElement {
    constructor() {
        super();

        this.ready = loadWasm();
        this.source = null;
    }

    async initialize(source) {
        await this.compile(source);
    }

    async compile(source) {
        await this.ready;

        if (source === this.source) {
            return;
        }

        await window.pipefishCompile(source);

        this.source = source;
    }

    async do(line) {
        await this.compile(this.source);

        return window.pipefishDo(line);
    }
}

customElements.define(
    "pf-service",
    PipefishService
);