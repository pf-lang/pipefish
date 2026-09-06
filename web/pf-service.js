class PipefishService extends HTMLElement {
    constructor() {
        super();

        this.ready = Promise.all([
            this.loadWasm(),
        ]);
    }

    async compile(source) {
        await this.ready;
        return window.pipefishCompile(source);
    }

    async do(line) {
        await this.ready;
        return window.pipefishDo(line);
    }

    async loadWasm() {
        const wasmURL = "assets/pipefish.wasm";
        const execURL = "assets/wasm_exec.js";

        if (!window.Go) {
            await this.loadScript(execURL);
        }

        this.go = new window.Go();

        const result = await WebAssembly.instantiateStreaming(
            fetch(wasmURL),
            this.go.importObject
        );

        this.go.run(result.instance);
    }

    loadScript(url) {
        return new Promise((resolve, reject) => {
            const script = document.createElement("script");

            script.src = url;
            script.onload = resolve;
            script.onerror = reject;

            document.head.appendChild(script);
        });
    }
}

customElements.define("pf-service", PipefishService);