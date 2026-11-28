import { loadWasm } from "./assets/pf-wasm.js";
import { setFlags } from "./assets/flags.js";

class PipefishService extends HTMLElement {
    constructor() {
        super();

        this.ready = loadWasm();
        this.source = null;
        this.serviceIndex = null;

        this.fileExtractor = async (path, filename) => {
            const url = new URL(
                filename,
                new URL(
                    path,
                    new URL("book/", document.baseURI)
                )
            );

            const response = await fetch(url, {
                cache: "no-store"
            });

            if (!response.ok) {
                throw new Error(
                    `Unable to load ${url}: ` +
                    `${response.status} ${response.statusText}`
                );
            }

            return new Uint8Array(await response.arrayBuffer());
        };
    }

    async initialize(source, flags) {
        this.initialization = (async () => {
            setFlags(this, flags);
            this.source = source;
        })();

        return this.initialization;
    }

    async build(fileExtractor) {
        await this.initialization;

        if (fileExtractor) {
            this.fileExtractor = fileExtractor;
        }

        return await this.compile(this.source);
    }

    async getFilepath() {
        await this.ready;

        if (this.serviceIndex === null) {
            return "";
        }

        return window.pipefishGetFilepath(this.serviceIndex);
    }

    async compile(source) {
        await this.ready;

        let files;

        const lines = source.replace(/\r\n/g, "\n").split("\n");

        if (lines[0].trim().startsWith("./")) {
            files = await this.loadDirectory(lines);
        } else {
            files = [
                {
                    path: "main.pf",
                    data: new TextEncoder().encode(source)
                }
            ];
        }

        const result = await window.pipefishCompile(files);

        if (Number.isInteger(result)) {
            this.serviceIndex = result;
        } else {
            return "failed to compile service"
        }

        const errorReport = await window.pipefishGetErrorReport(result)
        if (errorReport != "Current service has no recent errors.") {
            this.compileTimeErrors = await window.pipefishRenderMdAsHtml(errorReport);
        }

        this.source = source;
        return result;
    }

    async getErrorReport() {
        await this.ready;
        return this.compileTimeErrors;
    }

    async loadDirectory(lines) {
        const base = lines[0].trim();
        const entries = this.parseDirectory(lines.slice(1));
        const files = [];

        for (const path of entries) {
            const data = await this.fileExtractor(base, path);

            files.push({
                path,
                data
            });
        }

        return files;
    }

    async compileMain() {
        console.log("Called compileMain.")
        await this.ready;
        await window.pipefishCompileMain(this.serviceIndex);
        const errorReport = await window.pipefishGetErrorReport(this.serviceIndex)
        if (errorReport != "Current service has no recent errors.") {
            console.log("Created error report", errorReport)
            this.compileTimeErrors = await window.pipefishRenderMdAsHtml(errorReport);
        }
    }

    parseDirectory(lines) {
        const entries = [];
        const stack = [];

        for (const line of lines) {
            if (!line.trim()) {
                continue;
            }

            const match = line.match(/^(\s*)(.*)$/);
            const indent = match[1].replace(/\t/g, "    ").length;
            const name = match[2].trim();

            while (
                stack.length > 0 &&
                indent <= stack[stack.length - 1].indent
            ) {
                stack.pop();
            }

            const path =
                stack.length > 0
                    ? stack[stack.length - 1].path + name
                    : name;

            if (name.endsWith("/")) {
                stack.push({
                    indent,
                    path
                });
            } else {
                entries.push(path);
            }
        }

        return entries;
    }

    async do(line) {
        await this.ready;

        if (this.serviceIndex === null) {
            throw new Error("Service has not been compiled");
        }

        return window.pipefishDo(this.serviceIndex, line);
    }

}

customElements.define(
    "pf-service",
    PipefishService
);