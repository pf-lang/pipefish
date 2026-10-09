const TITLES = {
open: "Open",
download: "Download",
new: "New",
rename: "Rename",
delete: "Delete",
revert: "Revert",
};

const DEMO_VFS = {
type: "folder", name: "", children: [
{ type: "file", name: "main.pf" },
{ type: "folder", name: "examples", children: [
{ type: "file", name: "hello.pf" },
{ type: "file", name: "fibonacci.pf" },
{ type: "folder", name: "more_examples", children: [
{ type: "file", name: "hello_2.pf" },
{ type: "file", name: "fibonacci_2.pf" },
]},
] },
{ type: "folder", name: "libraries", children: [
{ type: "file", name: "math.pf" },
{ type: "file", name: "strings.pf" },
] },
{ type: "file", name: "README.md" },
],
};

class PFFiles extends HTMLElement {
constructor() {
super();
    this.attachShadow({ mode: "open" });

    this.vfs = null;
    this.operation = "open";
    this.selected = null;
    this.selectedPath = null;
}

connectedCallback() {
    if (!this.vfs) this.vfs = DEMO_VFS;

    if (this.hasAttribute("demo")) {
        this.open(this.getAttribute("operation") || "open");
    }
}

setVFS(vfs) {
    this.vfs = vfs;
    this.render();
}

open(operation = "open", selectedPath = null) {
    this.operation = TITLES[operation] ? operation : "open";
    this.selectedPath = selectedPath;
    this.selected = null;

    this.render();

    this.shadowRoot.querySelector(".backdrop").hidden = false;
}

close() {
    const backdrop = this.shadowRoot.querySelector(".backdrop");

    if (backdrop) {
        backdrop.hidden = true;
    }
}

render() {
    const title = TITLES[this.operation] || "Open";

    this.shadowRoot.innerHTML = `
        <link rel="stylesheet" href="${new URL("./files.css", import.meta.url)}">

        <div class="backdrop" hidden>
            <section class="dialog"
                role="dialog"
                aria-modal="true"
                aria-labelledby="dialog-title">

                <header class="dialog-header">
                    ${title}
                    ${this.operation === "new" ? `
                        <div class="kind-choice">
                            <button type="button"
                                class="kind-button selected"
                                data-kind="file">▤</button>
                            <button type="button"
                                class="kind-button"
                                data-kind="folder">▱</button>
                        </div>
                    ` : ""}

                    <button class="icon-button close"
                        type="button"
                        aria-label="Close">×</button>
                </header>

                <div class="tree" role="tree"></div>

                ${this.operation === "new" || this.operation === "rename" ? `
                    <div class="name-field">
                        <input id="entity-name"
                            class="text-field"
                            type="text"
                            autocomplete="off"
                            placeholder="${this.operation === "new" ? "Name" : "New name"}">
                    </div>
                ` : ""}

                <footer class="dialog-footer">
                    <button class="cancel" type="button">Cancel</button>
                    <button class="confirm" type="button">
                        ${this.confirmLabel()}
                    </button>
                </footer>
            </section>
        </div>`;

    const tree = this.shadowRoot.querySelector(".tree");

    this.renderNodes(this.vfs.children || [], tree, "");

    this.shadowRoot.querySelector(".close")
        .addEventListener("click", () => this.close());

    this.shadowRoot.querySelector(".cancel")
        .addEventListener("click", () => this.close());

    this.shadowRoot.querySelector(".backdrop")
        .addEventListener("click", event => {
            if (event.target === event.currentTarget) {
                this.close();
            }
        });

    this.shadowRoot.querySelector(".confirm")
        .addEventListener("click", () => this.confirm());

    this.shadowRoot.querySelectorAll(".kind-button").forEach(button => {
        button.addEventListener("click", () => {
            this.shadowRoot.querySelectorAll(".kind-button")
                .forEach(el => el.classList.toggle("selected", el === button));
        });
    });

    if (this.selectedPath) {
        const row = [...this.shadowRoot.querySelectorAll(".tree-row")]
            .find(el => el.dataset.path === this.selectedPath);

        if (row) row.click();
    }

    this.shadowRoot.querySelectorAll(".kind-button").forEach(button => {
        button.addEventListener("click", () => {
            this.shadowRoot.querySelectorAll(".kind-button")
                .forEach(el => el.classList.toggle("selected", el === button));

            this.updateConfirmButton();
        });
    });

    const nameField = this.shadowRoot.querySelector("#entity-name");
    if (nameField) {
        nameField.addEventListener("input", () => this.updateConfirmButton());
    }

    this.updateConfirmButton();


}

confirmLabel() {
    return {
        open: "Open",
        download: "Download",
        new: "Create",
        rename: "Rename",
        delete: "Delete",
        revert: "Revert",
    }[this.operation] || "OK";
}

renderNodes(nodes, container, parentPath) {
    for (const node of nodes) {
        const path = parentPath
            ? `${parentPath}/${node.name}`
            : node.name;

        const item = document.createElement("div");
        item.className = "tree-item";

        const row = document.createElement("button");
        row.type = "button";
        row.className = "tree-row";
        row.dataset.path = path;
        row.dataset.type = node.type;
        row.setAttribute("role", "treeitem");

        const isFolder = node.type === "folder";

        const twisty = document.createElement("span");
        twisty.className = "twisty";
        twisty.textContent = isFolder ? "▸" : "";

        const icon = document.createElement("span");
        icon.className = "node-icon";
        icon.textContent = isFolder ? "▱" : "▤";

        const name = document.createElement("span");
        name.className = "node-name";
        name.textContent = node.name;

        row.append(twisty, icon, name);
        item.append(row);

        row.addEventListener("click", () => {
            this.shadowRoot
                .querySelectorAll(".tree-row.selected")
                .forEach(el => el.classList.remove("selected"));

            row.classList.add("selected");

            this.selected = {
                node,
                path,
                type: node.type,
            };

            if (isFolder) {
                const expanded = item.classList.toggle("expanded");
                twisty.textContent = expanded ? "▾" : "▸";
            }
            this.updateConfirmButton();
        });

        if (isFolder && node.children?.length) {
            const children = document.createElement("div");
            children.className = "tree-children";

            this.renderNodes(
                node.children,
                children,
                path
            );

            item.append(children);
        }

        container.append(item);
    }
}

confirm() {
    const detail = {
        operation: this.operation,
        selection: this.selected,
    };

    if (this.operation === "new") {
        detail.kind =
            this.shadowRoot
                .querySelector(".kind-button.selected")
                ?.dataset.kind || "file";

        detail.name =
            this.shadowRoot
                .querySelector("#entity-name")
                .value
                .trim();

        if (!detail.name) {
            this.shadowRoot
                .querySelector("#entity-name")
                .focus();

            return;
        }

        if (this.selected && this.selected.type !== "folder") {
            return;
        }
    } else if (
        ["open", "download", "rename", "delete", "revert"]
            .includes(this.operation)
        && !this.selected
    ) {
        return;
    }

    if (this.operation === "rename") {
        detail.name =
            this.shadowRoot
                .querySelector("#entity-name")
                .value
                .trim();

        if (!detail.name) {
            this.shadowRoot
                .querySelector("#entity-name")
                .focus();

            return;
        }
    }

    this.dispatchEvent(new CustomEvent("files-action", {
        detail,
        bubbles: true,
        composed: true,
    }));

    this.close();
}

updateConfirmButton() {
    const button = this.shadowRoot.querySelector(".confirm");
    if (!button) return;

    let enabled = true;

    if (["open", "download", "rename", "delete", "revert"].includes(this.operation)) {
        enabled = !!this.selected;
    }

    if (this.operation === "new") {
        const name = this.shadowRoot.querySelector("#entity-name")?.value.trim();
        enabled = !!name && (!this.selected || this.selected.type === "folder");
    }

    if (this.operation === "rename") {
        enabled = !!this.selected &&
            !!this.shadowRoot.querySelector("#entity-name")?.value.trim();
    }

    button.disabled = !enabled;
}

}

customElements.define("pf-files", PFFiles);
