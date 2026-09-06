import { loadWasm } from "./assets/pf-wasm.js";

class PipefishHighlighter extends HTMLElement {
    constructor() {
        super();
        this.ready = loadWasm();
    }

    async highlight(source) {
        await this.ready;
        return window.pipefishHighlighter(source);
    }

}

customElements.define("pf-highlighter", PipefishHighlighter);
