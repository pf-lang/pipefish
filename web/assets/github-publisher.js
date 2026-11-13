const GITHUB_API = "https://api.github.com";

function apiHeaders(token) {
    return {
        "Accept": "application/vnd.github+json",
        "Authorization": `Bearer ${token}`,
        "X-GitHub-Api-Version": "2022-11-28"
    };
}

async function githubRequest(url, token, options = {}) {
    const response = await fetch(url, {
        ...options,
        headers: {
            ...apiHeaders(token),
            ...(options.headers ?? {})
        }
    });

    const text = await response.text();

    let result;
    try {
        result = text ? JSON.parse(text) : {};
    } catch {
        result = {};
    }

    if (!response.ok) {
        const message =
            result.message || response.statusText || "GitHub API request failed";

        throw new Error(
            `GitHub API error (${response.status}): ${message}`
        );
    }

    return result;
}

function jsonBody(value) {
    return {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(value)
    };
}

function encodeBase64(bytes) {
    let binary = "";
    const chunkSize = 0x8000;

    for (let offset = 0; offset < bytes.length; offset += chunkSize) {
        const chunk = bytes.subarray(
            offset,
            Math.min(offset + chunkSize, bytes.length)
        );

        binary += String.fromCharCode(...chunk);
    }

    return btoa(binary);
}

function normalizeSourcePath(path) {
    return (path ?? "/")
        .replace(/^\/+|\/+$/g, "")
        .replace(/\/+/g, "/");
}

function normalizeDirectory(directory) {
    const normalized = directory
        .replaceAll("\\", "/")
        .split("/")
        .filter(part => part && part !== "." && part !== "..")
        .join("/");

    if (!normalized || normalized.includes("/")) {
        throw new Error("The publication directory must be a single directory name.");
    }

    return normalized;
}

function makePublicURL(pagesUrl, directory) {
    return `${pagesUrl.replace(/\/+$/, "")}/${directory}/`;
}

/**
 * Publish a previously generated book export.
 *
 * destination:
 * {
 *     owner,
 *     repository,
 *     token,
 *     pagesUrl,
 *     sourceBranch,
 *     sourcePath
 * }
 *
 * files:
 * [
 *     { path: "index.html", data: Uint8Array },
 *     { path: "book/Contents.md", data: Uint8Array },
 *     ...
 * ]
 */
export async function publishBookToGitHub({
    destination,
    directory,
    files
}) {
    const {
        owner,
        repository,
        token
    } = destination;

    if (!owner || !repository || !token) {
        throw new Error("The GitHub destination is incomplete.");
    }

    directory = normalizeDirectory(directory);

    const repoURL =
        `${GITHUB_API}/repos/${encodeURIComponent(owner)}/${encodeURIComponent(repository)}`;

    // Resolve the Pages branch. If Pages is configured through GitHub
    // Actions, its source branch may be null; use the repository default.
    const repo = await githubRequest(repoURL, token);

    const branch =
        destination.sourceBranch ||
        repo.default_branch;

    if (!branch) {
        throw new Error("Could not determine the GitHub Pages source branch.");
    }

    const branchURL =
        `${repoURL}/branches/${encodeURIComponent(branch)}`;

    const branchInfo = await githubRequest(branchURL, token);
    const parentCommit = branchInfo.commit.sha;
    const baseTree = branchInfo.commit.commit.tree.sha;

    const existingTree = await githubRequest(
        `${repoURL}/git/trees/${baseTree}?recursive=1`,
        token
    );

    if (existingTree.truncated) {
        throw new Error(
            "The repository tree is too large to update safely in one operation."
        );
    }

    const sourcePath = normalizeSourcePath(destination.sourcePath);
    const directoryPath = [
        sourcePath,
        directory
    ].filter(Boolean).join("/");

    const prefix = `${directoryPath}/`;

    const newPaths = new Set(
        files.map(file => `${prefix}${file.path}`)
    );

    const entries = [];

    // Remove files previously published beneath this book directory but
    // absent from the new export. Do not touch other repository paths.
    for (const item of existingTree.tree) {
        if (
            item.type === "blob" &&
            item.path.startsWith(prefix) &&
            !newPaths.has(item.path)
        ) {
            entries.push({
                path: item.path,
                mode: "100644",
                type: "blob",
                sha: null
            });
        }
    }

    // Create blobs for every exported file.
    for (const file of files) {
        const path = `${prefix}${file.path}`;
        const bytes = file.data instanceof Uint8Array
            ? file.data
            : new Uint8Array(file.data);

        const blob = await githubRequest(
            `${repoURL}/git/blobs`,
            token,
            jsonBody({
                content: encodeBase64(bytes),
                encoding: "base64"
            })
        );

        entries.push({
            path,
            mode: "100644",
            type: "blob",
            sha: blob.sha
        });
    }

    // Create a new tree based on the complete existing repository tree.
    const newTree = await githubRequest(
        `${repoURL}/git/trees`,
        token,
        jsonBody({
            base_tree: baseTree,
            tree: entries
        })
    );

    const commit = await githubRequest(
        `${repoURL}/git/commits`,
        token,
        jsonBody({
            message: `Publish ${directory}`,
            tree: newTree.sha,
            parents: [parentCommit]
        })
    );

    // Update the branch without force-pushing.
    await githubRequest(
        `${repoURL}/git/refs/heads/${encodeURIComponent(branch)}`,
        token,
        {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                sha: commit.sha,
                force: false
            })
        }
    );

    if (!destination.pagesUrl) {
        throw new Error(
            "The commit succeeded, but the GitHub Pages URL is unavailable."
        );
    }

    return {
        commit: commit.sha,
        url: makePublicURL(destination.pagesUrl, directory)
    };
}