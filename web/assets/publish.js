class PFPUBLISH extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });

        this.sites = [];
        this.selectedSite = null;
        this.credentials = [];

        this.editing = false;
    }

    open({ sites = [], credentials = [], bookTitle = "" } = {}) {
        this.sites = sites;
        this.credentials = credentials;
        this.bookTitle = bookTitle;

        this.selectedSite = sites[0] || null;

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
        this.shadowRoot.innerHTML = `
            <link rel="stylesheet"
                href="${new URL("./publish.css", import.meta.url)}">

            <div class="backdrop" hidden>
                <section class="dialog"
                    role="dialog"
                    aria-modal="true"
                    aria-labelledby="dialog-title">

                    <header class="dialog-header">
                        <span id="dialog-title">Publish book</span>

                        <button class="icon-button close"
                            type="button"
                            aria-label="Close">×</button>
                    </header>

                    <div class="dialog-body">

                        <section>
                            <div class="section-header">
                                <h3>GitHub Pages sites</h3>

                                <button class="add-site"
                                    type="button">+ Add site</button>
                            </div>

                            <div class="site-list"></div>
                        </section>

                        <section class="site-fields">

                            <div class="field">
                                <label for="repository">
                                    GitHub repository
                                </label>

                                <input id="repository"
                                    class="text-field"
                                    placeholder="username/repository">
                            </div>

                            <div class="field-row">
                                <div class="field">
                                    <label for="branch">Branch</label>

                                    <input id="branch"
                                        class="text-field"
                                        value="gh-pages">
                                </div>

                                <div class="field">
                                    <label for="directory">
                                        Book directory
                                    </label>

                                    <input id="directory"
                                        class="text-field"
                                        value="">
                                </div>
                            </div>

                            <div class="field">
                                <label for="pages-url">
                                    GitHub Pages URL
                                </label>

                                <input id="pages-url"
                                    class="text-field"
                                    placeholder="https://username.github.io/repository/">
                            </div>

                            <div class="field">
                                <label for="credential">
                                    Authentication
                                </label>

                                <select id="credential"
                                    class="select-field">
                                </select>
                            </div>

                            <button class="add-credential"
                                type="button">
                                + Add authentication
                            </button>

                        </section>

                    </div>

                    <footer class="dialog-footer">
                        <button class="confirm"
                            type="button">Publish</button>

                        <button class="cancel"
                            type="button">Cancel</button>
                    </footer>

                </section>
            </div>
        `;

        this.renderSites();
        this.renderCredentials();

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

        this.shadowRoot.querySelector(".add-site")
            .addEventListener("click", () => this.addSite());

        this.shadowRoot.querySelector(".add-credential")
            .addEventListener("click", () => {
                this.dispatchEvent(new CustomEvent(
                    "publish-add-credential",
                    { bubbles: true, composed: true }
                ));
            });

        this.shadowRoot.querySelector(".confirm")
            .addEventListener("click", () => this.confirm());

        this.populateFields();
    }

    renderSites() {
        const list = this.shadowRoot.querySelector(".site-list");

        list.innerHTML = "";

        for (const site of this.sites) {
            const button = document.createElement("button");

            button.type = "button";
            button.className = "site-option";

            if (site === this.selectedSite) {
                button.classList.add("selected");
            }

            const label = document.createElement("span");

            label.textContent = site.name || site.repository;

            button.append(label);

            button.addEventListener("click", () => {
                this.selectedSite = site;
                this.renderSites();
                this.populateFields();
            });

            list.append(button);
        }
    }

    renderCredentials() {
        const select = this.shadowRoot.querySelector("#credential");

        select.innerHTML = "";

        for (const credential of this.credentials) {
            const option = document.createElement("option");

            option.value = credential.id;
            option.textContent = credential.name;

            select.append(option);
        }
    }

    populateFields() {
        const site = this.selectedSite || {};

        this.shadowRoot.querySelector("#repository").value =
            site.repository || "";

        this.shadowRoot.querySelector("#branch").value =
            site.branch || "gh-pages";

        this.shadowRoot.querySelector("#pages-url").value =
            site.url || "";

        this.shadowRoot.querySelector("#directory").value =
            site.directory ||
            this.bookTitle.toLowerCase()
                .trim()
                .replace(/[^a-z0-9]+/g, "-")
                .replace(/^-|-$/g, "");

        this.shadowRoot.querySelector("#credential").value =
            site.credentialId || "";
    }

    addSite() {
        const site = {
            name: "",
            repository: "",
            branch: "gh-pages",
            url: "",
            directory: "",
            credentialId: ""
        };

        this.sites.push(site);
        this.selectedSite = site;

        this.renderSites();
        this.populateFields();
    }

    confirm() {
        const repository =
            this.shadowRoot.querySelector("#repository").value.trim();

        const branch =
            this.shadowRoot.querySelector("#branch").value.trim();

        const directory =
            this.shadowRoot.querySelector("#directory").value.trim();

        const url =
            this.shadowRoot.querySelector("#pages-url").value.trim();

        const credentialId =
            this.shadowRoot.querySelector("#credential").value;

        if (!repository || !branch || !url || !credentialId) {
            alert("Please complete the publishing destination.");
            return;
        }

        const site = {
            ...this.selectedSite,
            name: this.selectedSite?.name || repository,
            repository,
            branch,
            directory,
            url,
            credentialId
        };

        this.dispatchEvent(new CustomEvent("publish-action", {
            detail: { site },
            bubbles: true,
            composed: true
        }));

        this.close();
    }
}

customElements.define("pf-publish", PFPUBLISH);