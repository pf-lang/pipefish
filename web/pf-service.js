import { loadWasm } from "./assets/pf-wasm.js";

class PipefishService extends HTMLElement {
    constructor() {
        super();
        this.ready = loadWasm();
    }

    async compile(source) {
        await this.ready;
        return window.pipefishCompile(source);
    }

    async do(line) {
        await this.ready;
        return window.pipefishDo(line);
    }

}

customElements.define("pf-service", PipefishService);
