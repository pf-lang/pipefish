import "./pf-coder.js";
import "./pf-tui.js";

class PipefishIde extends HTMLElement {
    constructor() {
        super();

        const tui =
            document.createElement("pf-tui");

        const coder =
            document.createElement("pf-coder");

        this.coder = coder;
        this.tui = tui;

        tui.service = coder;

        const shadow =
            this.attachShadow({ mode: "open" });

        const style =
            document.createElement("link");
        style.rel = "stylesheet";
        style.href = "assets/pf-page.css";

        const wrapper =
            document.createElement("div");
        wrapper.classList.add("service");

        wrapper.append(
            coder,
            tui
        );

        shadow.append(
            style,
            wrapper
        );

        this.ready = Promise.all([
            coder.ready,
            tui.ready
        ]);
    }

    connectedCallback() {
        this.initialize(this.textContent);
    }

    async initialize(source) {
        await this.ready;

        await this.coder.initialize(source);
        await this.tui.initialize(source);
    }
}

customElements.define(
    "pf-ide",
    PipefishIde
);