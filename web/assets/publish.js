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

        this.shadowRoot.append(stylesheet);

        this.shadowRoot.insertAdjacentHTML("beforeend", `
            <div class="backdrop" hidden>
                <section class="dialog" role="dialog"
                         aria-modal="true"
                         aria-labelledby="dialog-title">

                    <header>
                        <h2 id="dialog-title">Publish book</h2>
                        <button class="close" data-action="close"
                                aria-label="Close">&times;</button>
                    </header>

                    <div class="dialog-body">
                        <table>
                            <thead>
                                <tr>
                                    <th></th>
                                    <th>Destination</th>
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
                                        <input class="add-name"
                                               placeholder="Site name"
                                               autocomplete="off">
                                    </td>
                                    <td>
                                        <input class="add-repository"
                                               placeholder="owner/repository"
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
                            </tbody>
                        </table>

                        <div class="directory">
                            <label for="publish-directory">
                                Publish to directory
                            </label>
                            <input id="publish-directory"
                                   class="directory-input"
                                   readonly>
                        </div>

                        <div class="status" role="status" aria-live="polite"></div>
                    </div>

                    <footer>
                        <button class="cancel" data-action="close">
                            Cancel
                        </button>
                        <button class="publish-button"
                                data-action="publish"
                                disabled>
                            Publish
                        </button>
                    </footer>
                </section>
            </div>
        `);

        this.backdrop = this.shadowRoot.querySelector(".backdrop");
        this.rows = this.shadowRoot.querySelector(".destination-rows");
        this.directoryInput =
            this.shadowRoot.querySelector(".directory-input");
        this.status = this.shadowRoot.querySelector(".status");
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

            // Republishing an already-associated book needs no dialog.
            if (publication) {
                const destination = this.destinations.find(
                    item => item.id === publication.destinationId
                );

                if (destination?.token) {
                    this.dispatchPublish(
                        destination,
                        publication.directory
                    );
                    return;
                }
            }

            this.selectedDestinationId =
                publication?.destinationId ??
                lastDestination ??
                this.destinations[0]?.id ??
                null;

            // Keep the existing directory when this book has been published
            // before. Otherwise derive it from the title.
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
                    <input data-field="name"
                           aria-label="Destination name">
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

            row.querySelector('[data-field="name"]').value =
                destination.name || "";

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
            destination.name?.trim() &&
            destination.repository?.trim() &&
            destination.token?.trim() &&
            this.directory
        );
    }

    async addDestination() {
        const nameInput = this.shadowRoot.querySelector(".add-name");
        const repositoryInput =
            this.shadowRoot.querySelector(".add-repository");
        const tokenInput = this.shadowRoot.querySelector(".add-token");

        const name = nameInput.value.trim();
        const repository = repositoryInput.value.trim();
        const token = tokenInput.value.trim();

        if (!name || !repository || !token) {
            this.showStatus("Please fill in all three fields.", true);
            return;
        }

        this.setBusy(true);
        this.showStatus("Checking GitHub destination…");

        try {
            const destination = await this.checkDestination({
                name,
                repository,
                token
            });

            destination.id = crypto.randomUUID();

            this.destinations.push(destination);
            this.selectedDestinationId = destination.id;

            await save(KEYS.destinations, this.destinations);

            // Clear the add row only after validation and persistence succeed.
            nameInput.value = "";
            repositoryInput.value = "";
            tokenInput.value = "";

            this.renderRows();
            this.showStatus("Destination added.");
            this.updatePublishButton();

        } catch (error) {
            // Leave the entered values in place for correction.
            this.showStatus(error.message, true);
        } finally {
            this.setBusy(false);
        }
    }

    async checkDestination({ name, repository, token }) {
        const match = repository.match(/^([^/\s]+)\/([^/\s]+)$/);

        if (!match) {
            throw new Error(
                "Repository must be in the form owner/repository."
            );
        }

        const [, owner, repo] = match;

        const headers = {
            "Accept": "application/vnd.github+json",
            "Authorization": `Bearer ${token}`,
            "X-GitHub-Api-Version": "2022-11-28"
        };

        const repoResponse = await fetch(
            `${GITHUB_API}/repos/${encodeURIComponent(owner)}/${encodeURIComponent(repo)}`,
            { headers }
        );

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

        const pagesResponse = await fetch(
            `${GITHUB_API}/repos/${encodeURIComponent(owner)}/${encodeURIComponent(repo)}/pages`,
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
            name,
            repository,
            token,
            pagesUrl: pages.html_url,
            sourceBranch: pages.source?.branch ?? null,
            sourcePath: pages.source?.path ?? "/"
        };
    }

    async deleteDestination(id) {
        const destination = this.destinations.find(item => item.id === id);
        if (!destination) return;

        if (!window.confirm(
            `Delete the destination "${destination.name}"?`
        )) {
            return;
        }

        this.destinations = this.destinations.filter(
            item => item.id !== id
        );

        // Remove publication associations which refer to this destination.
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
            // Recheck editable destination fields before publishing.
            const checked = await this.checkDestination(destination);

            Object.assign(destination, checked);

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

            this.close();
            this.dispatchPublish(destination, this.directory);

        } catch (error) {
            this.showStatus(error.message, true);
        } finally {
            this.setBusy(false);
        }
    }

    dispatchPublish(destination, directory) {
        this.dispatchEvent(new CustomEvent("publish-action", {
            bubbles: true,
            composed: true,
            detail: {
                bookRoot: this.bookRoot,
                destination: { ...destination },
                directory
            }
        }));
    }

    setBusy(busy) {
        this.shadowRoot.querySelectorAll("button, input").forEach(element => {
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