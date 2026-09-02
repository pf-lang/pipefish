class PipefishService extends HTMLElement {
    constructor() {
        super();

        this.ready = this.loadWasm();

        this.history = [];
        this.historyIndex = 0;

        const shadow = this.attachShadow({ mode: "open" });

        const style = document.createElement("style");
        style.textContent = `
            .service {
                border: 1px solid #888;
                padding: 1rem;
                max-width: 600px;
            }

            textarea {
                display: block;
                width: 100%;
                box-sizing: border-box;
                min-height: 100px;
                margin-bottom: 1rem;
            }

            button {
                margin-bottom: 1rem;
            }

            .repl {
                font-family: monospace;
                line-height: 1.4;
            }

            .transcript {
                margin: 0;
                white-space: pre-wrap;
            }

            .input-line {
                margin: 0;
                display: flex;
                font-family: monospace;
            }

            .transcript,
            .input-line {
                line-height: 1.4;
            }

            .prompt {
                margin-right: 0.5rem;
            }
            
            .input {
                flex: 1;
                font-family: inherit;
                border: none;
                outline: none;
                padding: 0;
            }
        `;

        const wrapper = document.createElement("div");
        wrapper.classList.add("service");

        const heading = document.createElement("h2");
        heading.textContent = "Pipefish service";

        const code = document.createElement("textarea");
        code.placeholder = "Type Pipefish code here...";

        const compileButton = document.createElement("button");
        compileButton.textContent = "Compile";

        const repl = document.createElement("div");
        repl.classList.add("repl");

        const transcript = document.createElement("div");
        transcript.classList.add("transcript");
        this.transcript = transcript;

        const inputLine = document.createElement("div");
        inputLine.classList.add("input-line");

        const prompt = document.createElement("span");
        prompt.classList.add("prompt");
        prompt.textContent = "→";

        const input = document.createElement("input");
        input.classList.add("input");
        input.type = "text";
        input.autocomplete = "off";
        input.spellcheck = false;

        inputLine.append(prompt, input);
        repl.append(transcript, inputLine);

        wrapper.append(
            heading,
            code,
            compileButton,
            repl
        );

        shadow.append(style, wrapper);

        compileButton.addEventListener("click", async () => {
            try {
                await this.compile(code.value);
                this.write("Compiled successfully.");
                input.focus();
            } catch (error) {
                this.write("Error: " + error);
            }
        });

        input.addEventListener("keydown", async (event) => {
            if (event.key === "Enter") {
                event.preventDefault();
                await this.submit(input);
            }
        });
    }

    async submit(input) {
        const line = input.value;

        if (!line.trim()) {
            return;
        }

        this.history.push(line);
        this.historyIndex = this.history.length;

        this.write("→ " + line);
        input.value = "";

        try {
            const result = await this.do(line);

            if (result !== undefined && result !== "") {
                this.write(result);
            }
        } catch (error) {
            this.write("Error: " + error);
        }
    }

    write(text) {
        const line = document.createElement("div");
        line.textContent = text;
        this.transcript.appendChild(line);
    }

    async loadWasm() {
        const wasmURL = "pipefish.wasm";
        const execURL = "wasm_exec.js";

        if (!window.Go) {
            await this.loadScript(execURL);
        }

        this.go = new Go();

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

    async compile(source) {
        await this.ready;
        return window.pipefishCompile(source);
    }

    async do(line) {
        await this.ready;
        return window.pipefishDo(line);
    }
}

customElements.define("pipefish-service", PipefishService);