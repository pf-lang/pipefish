import { load, save } from "./author-storage.js";

const KEYS = {
    destinations: "publishDestinations",
    publications: "bookPublications",
    lastDestination: "lastPublishDestination"
};

const GITHUB_API = "https://api.github.com";

class PfPublish extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });

        this.destinations = [];
        this.publications = [];
        this.selectedDestinationId = null;

        this.bookRoot = null;
        this.bookTitle = "";
        this.directory = "";

        const stylesheet = document.createElement("link");
        stylesheet.rel = "stylesheet";
        stylesheet.href = new URL("./publish.css", import.meta.url).href;

        this.shadowRoot.innerHTML = `
            <div class="backdrop setup-backdrop" hidden>
                <section class="dialog" role="dialog"
                         aria-modal="true"
                         aria-labelledby="dialog-title">

                    <header class="dialog-header">
                        <h2 id="dialog-title">Publish book</h2>
                        <button class="close" data-action="close"
                                aria-label="Close">&times;</button>
                    </header>

                    <div class="dialog-body">
                        <table>
                            <thead>
                                <tr>
                                    <th></th>
                                    <th>Owner</th>
                                    <th>Repository</th>
                                    <th>Authentication token</th>
                                    <th></th>
                                </tr>
                            </thead>
                            <tbody class="destination-rows"></tbody>
                            <tbody>
                                <tr class="add-row">
                                    <td></td>
                                    <td>
                                        <input class="add-owner"
                                            placeholder="Owner"
                                            autocomplete="off">
                                    </td>
                                    <td>
                                        <input class="add-repository"
                                            placeholder="Repository"
                                            autocomplete="off">
                                    </td>
                                    <td>
                                        <input class="add-token"
                                            type="password"
                                            placeholder="GitHub token"
                                            autocomplete="off">
                                    </td>
                                    <td>
                                        <button class="add-button"
                                                data-action="add"
                                                title="Add destination"
                                                aria-label="Add destination">+</button>
                                    </td>
                                </tr>
                                <tr class="directory-row">
                                    <td></td>
                                    <td colspan="4">
                                        <div class="field">
                                            <label for="publish-directory">
                                                Publish to directory
                                            </label>
                                            <input id="publish-directory"
                                                class="directory-input"
                                                readonly>
                                        </div>
                                    </td>
                                </tr>
                            </tbody>
                        </table>

                        <div class="status"
                             role="status"
                             aria-live="polite"></div>
                    </div>

                    <footer class="dialog-footer">
                        <button class="publish-button"
                                data-action="publish"
                                disabled>
                            Publish
                        </button>
                        <button class="cancel" data-action="close">
                            Cancel
                        </button>
                    </footer>
                </section>
            </div>

            <div class="backdrop success-backdrop" hidden>
                <section class="dialog success-dialog"
                         role="dialog"
                         aria-modal="true"
                         aria-labelledby="success-dialog-title">

                    <header class="dialog-header">
                        <h2 id="success-dialog-title">
                            Book published successfully
                        </h2>
                    </header>

                    <div class="dialog-body success-body">
                        <a class="published-url"
                           target="_blank"
                           rel="noopener"></a>
                    </div>

                    <footer class="dialog-footer">
                        <button data-action="success-ok">OK</button>
                        <button data-action="copy-url">Copy link</button>
                        <button data-action="visit-url">Visit book</button>
                    </footer>
                </section>
            </div>
        `;

        this.shadowRoot.prepend(stylesheet);

        this.backdrop =
            this.shadowRoot.querySelector(".setup-backdrop");

        this.successBackdrop =
            this.shadowRoot.querySelector(".success-backdrop");

        this.rows =
            this.shadowRoot.querySelector(".destination-rows");

        this.directoryInput =
            this.shadowRoot.querySelector(".directory-input");

        this.status =
            this.shadowRoot.querySelector(".status");

        this.publishedURL =
            this.shadowRoot.querySelector(".published-url");

        this.publishButton =
            this.shadowRoot.querySelector(".publish-button");

        this.shadowRoot.addEventListener("click", event => {
            const button = event.target.closest("[data-action]");
            if (!button) return;

            switch (button.dataset.action) {
                case "close":
                    this.close();
                    break;

                case "add":
                    this.addDestination();
                    break;

                case "delete":
                    this.deleteDestination(button.dataset.id);
                    break;

                case "publish":
                    this.publish();
                    break;

                case "success-ok":
                    this.closeSuccess();
                    break;

                case "copy-url":
                    this.copyPublishedURL();
                    break;

                case "visit-url":
                    window.open(
                        this.publishedURL.href,
                        "_blank",
                        "noopener"
                    );
                    break;
            }
        });

        this.shadowRoot.addEventListener("change", event => {
            const input = event.target;

            if (input.matches('input[type="radio"][name="destination"]')) {
                this.selectedDestinationId = input.value;
                this.updatePublishButton();
                return;
            }

            const row = input.closest("[data-destination-id]");
            if (!row || !input.dataset.field) return;

            const destination = this.destinations.find(
                item => item.id === row.dataset.destinationId
            );

            if (!destination) return;

            destination[input.dataset.field] = input.value.trim();

            save(KEYS.destinations, this.destinations)
                .catch(error => this.showStatus(error.message, true));

            this.updatePublishButton();
        });

        this.shadowRoot.addEventListener("focusin", event => {
            const row = event.target.closest("[data-destination-id]");
            if (!row) return;

            const radio = row.querySelector('input[type="radio"]');

            if (radio && !radio.checked) {
                radio.checked = true;
                this.selectedDestinationId = radio.value;
                this.updatePublishButton();
            }
        });
    }

    async open({ bookRoot, bookTitle }) {
        this.bookRoot = bookRoot;
        this.bookTitle = bookTitle || "Book";

        this.clearStatus();
        this.backdrop.hidden = true;
        this.successBackdrop.hidden = true;

        this.publishedURL.textContent = "";
        this.publishedURL.removeAttribute("href");

        try {
            const [destinations, publications, lastDestination] =
                await Promise.all([
                    load(KEYS.destinations),
                    load(KEYS.publications),
                    load(KEYS.lastDestination)
                ]);

            this.destinations = Array.isArray(destinations)
                ? destinations
                : [];

            this.publications = Array.isArray(publications)
                ? publications
                : [];

            const publication = this.publications.find(
                item => item.bookRoot === this.bookRoot
            );

            // Automatically republish an already-associated book.
            if (publication) {
                const destination = this.destinations.find(
                    item => item.id === publication.destinationId
                );

                if (destination?.token) {
                    this.selectedDestinationId = destination.id;
                    this.directory = publication.directory;
                    this.directoryInput.value = this.directory;

                    this.renderRows();
                    this.backdrop.hidden = false;
                    this.setBusy(true);
                    this.showStatus("Publishing book…");

                    try {
                        const result = await this.dispatchPublish(
                            destination,
                            publication.directory
                        );

                        this.showSuccess(result.url);
                    } catch (error) {
                        this.showStatus(error.message, true);
                    } finally {
                        this.setBusy(false);
                    }

                    return;
                }
            }

            this.selectedDestinationId =
                publication?.destinationId ??
                lastDestination ??
                this.destinations[0]?.id ??
                null;

            this.directory = publication?.directory ??
                this.slugify(this.bookTitle);

            this.directoryInput.value = this.directory;

            this.renderRows();
            this.backdrop.hidden = false;
            this.updatePublishButton();

        } catch (error) {
            console.error("Could not open publisher:", error);
            this.showStatus(error.message, true);
            this.backdrop.hidden = false;
        }
    }

    close() {
        this.backdrop.hidden = true;
    }

    closeSuccess() {
        this.successBackdrop.hidden = true;
    }

    slugify(value) {
        return value
            .toLowerCase()
            .trim()
            .replace(/[^a-z0-9]+/g, "-")
            .replace(/^-+|-+$/g, "") || "book";
    }

    renderRows() {
        this.rows.replaceChildren();

        for (const destination of this.destinations) {
            const row = document.createElement("tr");
            row.dataset.destinationId = destination.id;

            row.innerHTML = `
                <td>
                    <input type="radio"
                           name="destination"
                           value="${this.escape(destination.id)}"
                           aria-label="Select destination">
                </td>
                <td>
                    <input data-field="owner"
                           aria-label="Owner">
                </td>
                <td>
                    <input data-field="repository"
                           aria-label="Repository">
                </td>
                <td>
                    <input data-field="token"
                           type="password"
                           aria-label="Authentication token">
                </td>
                <td>
                    <button class="delete-button"
                            data-action="delete"
                            data-id="${this.escape(destination.id)}"
                            title="Delete destination"
                            aria-label="Delete destination">🗑</button>
                </td>
            `;

            row.querySelector('[type="radio"]').checked =
                destination.id === this.selectedDestinationId;

            row.querySelector('[data-field="owner"]').value =
                destination.owner || "";

            row.querySelector('[data-field="repository"]').value =
                destination.repository || "";

            row.querySelector('[data-field="token"]').value =
                destination.token || "";

            this.rows.append(row);
        }
    }

    escape(value) {
        return String(value).replace(/[&<>"']/g, character => ({
            "&": "&amp;",
            "<": "&lt;",
            ">": "&gt;",
            '"': "&quot;",
            "'": "&#39;"
        })[character]);
    }

    updatePublishButton() {
        const destination = this.destinations.find(
            item => item.id === this.selectedDestinationId
        );

        this.publishButton.disabled = !(
            destination &&
            destination.owner?.trim() &&
            destination.repository?.trim() &&
            destination.token?.trim() &&
            this.directory
        );
    }

    async addDestination() {
        const ownerInput =
            this.shadowRoot.querySelector(".add-owner");

        const repositoryInput =
            this.shadowRoot.querySelector(".add-repository");

        const tokenInput =
            this.shadowRoot.querySelector(".add-token");

        const owner = ownerInput.value.trim();
        const repository = repositoryInput.value.trim();
        const token = tokenInput.value.trim();

        if (!owner || !repository || !token) {
            this.showStatus("Please fill in all three fields.", true);
            return;
        }

        this.setBusy(true);
        this.showStatus("Checking GitHub destination…");

        try {
            const destination = await this.checkDestination({
                owner,
                repository,
                token
            });

            destination.id = crypto.randomUUID();

            this.destinations.push(destination);
            this.selectedDestinationId = destination.id;

            await save(KEYS.destinations, this.destinations);

            ownerInput.value = "";
            repositoryInput.value = "";
            tokenInput.value = "";

            this.renderRows();
            this.clearStatus();
            this.updatePublishButton();

        } catch (error) {
            // Leave the entered values in place for correction.
            this.showStatus(error.message, true);
        } finally {
            this.setBusy(false);
        }
    }

    async checkDestination({ owner, repository, token }) {
        if (owner.includes("/") || /\s/.test(owner)) {
            throw new Error("Please enter a valid GitHub owner.");
        }

        if (repository.includes("/") || /\s/.test(repository)) {
            throw new Error("Please enter a valid repository name.");
        }

        const headers = {
            "Accept": "application/vnd.github+json",
            "Authorization": `Bearer ${token}`,
            "X-GitHub-Api-Version": "2022-11-28"
        };

        const repoPath =
            `${GITHUB_API}/repos/${encodeURIComponent(owner)}/${encodeURIComponent(repository)}`;

        const repoResponse = await fetch(repoPath, { headers });

        if (repoResponse.status === 401) {
            throw new Error("GitHub rejected the authentication token.");
        }

        if (repoResponse.status === 404) {
            throw new Error(
                "Repository not found, or the token does not have access to it."
            );
        }

        if (repoResponse.status === 403) {
            throw new Error(
                "GitHub denied access. Check the token's repository permissions."
            );
        }

        if (!repoResponse.ok) {
            throw new Error(
                `Could not check repository (${repoResponse.status}).`
            );
        }

        const repo = await repoResponse.json();

        const pagesResponse = await fetch(
            `${repoPath}/pages`,
            { headers }
        );

        if (pagesResponse.status === 404) {
            throw new Error(
                "GitHub Pages does not appear to be enabled for this repository."
            );
        }

        if (pagesResponse.status === 403) {
            throw new Error(
                "The token can access the repository but cannot read its Pages settings. Check its Pages permissions."
            );
        }

        if (!pagesResponse.ok) {
            throw new Error(
                `Could not check GitHub Pages (${pagesResponse.status}).`
            );
        }

        const pages = await pagesResponse.json();

        return {
            owner,
            repository,
            token,
            pagesUrl: pages.html_url,
            sourceBranch: pages.source?.branch ?? null,
            sourcePath: pages.source?.path ?? "/",
            defaultBranch: repo.default_branch
        };
    }

    async deleteDestination(id) {
        const destination = this.destinations.find(item => item.id === id);
        if (!destination) return;

        if (!window.confirm(
            `Delete the destination "${destination.owner}/${destination.repository}"?`
        )) {
            return;
        }

        this.destinations = this.destinations.filter(
            item => item.id !== id
        );

        this.publications = this.publications.filter(
            item => item.destinationId !== id
        );

        if (this.selectedDestinationId === id) {
            this.selectedDestinationId =
                this.destinations[0]?.id ?? null;
        }

        await Promise.all([
            save(KEYS.destinations, this.destinations),
            save(KEYS.publications, this.publications)
        ]);

        this.renderRows();
        this.updatePublishButton();
    }

    async publish() {
        const destination = this.destinations.find(
            item => item.id === this.selectedDestinationId
        );

        if (!destination) return;

        this.setBusy(true);
        this.showStatus("Checking destination…");

        try {
            const checked = await this.checkDestination(destination);
            Object.assign(destination, checked);

            this.showStatus("Publishing book…");

            const result = await this.dispatchPublish(
                destination,
                this.directory
            );

            const publication = {
                bookRoot: this.bookRoot,
                destinationId: destination.id,
                directory: this.directory
            };

            const index = this.publications.findIndex(
                item => item.bookRoot === this.bookRoot
            );

            if (index === -1) {
                this.publications.push(publication);
            } else {
                this.publications[index] = publication;
            }

            await Promise.all([
                save(KEYS.destinations, this.destinations),
                save(KEYS.publications, this.publications),
                save(KEYS.lastDestination, destination.id)
            ]);

            this.showSuccess(result.url);

        } catch (error) {
            this.showStatus(error.message, true);
        } finally {
            this.setBusy(false);
        }
    }

    dispatchPublish(destination, directory) {
        const detail = {
            bookRoot: this.bookRoot,
            destination: { ...destination },
            directory,
            promise: null
        };

        this.dispatchEvent(new CustomEvent("publish-action", {
            bubbles: true,
            composed: true,
            detail
        }));

        if (!detail.promise) {
            return Promise.reject(
                new Error("No publication handler is available.")
            );
        }

        return detail.promise;
    }

    showSuccess(url) {
        this.backdrop.hidden = true;

        this.publishedURL.href = url;
        this.publishedURL.textContent = url;

        this.successBackdrop.hidden = false;

        this.shadowRoot
            .querySelector('[data-action="success-ok"]')
            .focus();
    }

    async copyPublishedURL() {
        try {
            await navigator.clipboard.writeText(this.publishedURL.href);
        } catch (error) {
            alert(
                "Could not copy the link automatically. " +
                "Please select and copy the URL."
            );
        }
    }

    setBusy(busy) {
        this.shadowRoot
            .querySelectorAll(".setup-backdrop button, .setup-backdrop input")
            .forEach(element => {
                element.disabled = busy;
            });

        if (!busy) {
            this.updatePublishButton();
        }
    }

    showStatus(message, isError = false) {
        this.status.textContent = message;
        this.status.classList.toggle("error", isError);
    }

    clearStatus() {
        this.status.textContent = "";
        this.status.classList.remove("error");
    }
}

if (!customElements.get("pf-publish")) {
    customElements.define("pf-publish", PfPublish);
}