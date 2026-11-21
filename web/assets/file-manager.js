import "./files.js";

export class FileManager {
    constructor(dialog, vfs, host = {}) {
        this.dialog = dialog;
        this.vfs = vfs;
        this.host = host;
        this.pendingUploads = [];

        dialog.addEventListener("files-action", event => {
            this.handleAction(event).catch(error => {
                console.error("File operation failed:", error);
            });
        });
    }

    open(operation = "open") {
        if (operation === "upload files" || operation === "upload") {
            this.selectUploadFiles();
            return;
        }

        this.refresh();
        this.dialog.open(operation);
    }

    selectUploadFiles() {
        const input = document.createElement("input");

        input.type = "file";
        input.multiple = true;
        input.hidden = true;

        input.addEventListener("change", () => {
            const files = Array.from(input.files);
            input.remove();

            if (!files.length) {
                return;
            }

            this.pendingUploads = files;
            this.refresh();
            this.dialog.open("upload");
        }, { once: true });

        input.addEventListener("cancel", () => {
            input.remove();
        }, { once: true });

        document.body.append(input);
        input.click();
    }

    refresh() {
        this.dialog.setVFS(this.vfs.getTree());
    }

    check(result) {
        if (typeof result === "string" && result) {
            throw new Error(result);
        }

        if (result?.ok === false) {
            throw new Error(result.error);
        }
    }

    async handleAction(event) {
        const { operation, selection, kind, name } = event.detail;
        const path = selection?.path ?? null;
        const type = selection?.type ?? selection?.node?.type;

        switch (operation) {
            case "upload": {
                if (type !== "folder") {
                    break;
                }

                const destination = path === "."
                    ? ""
                    : path;

                for (const file of this.pendingUploads) {
                    const filepath = destination
                        ? `${destination}/${file.name}`
                        : file.name;

                    const data = new Uint8Array(
                        await file.arrayBuffer()
                    );

                    await this.vfs.writeFile(filepath, data);
                }

                this.pendingUploads = [];
                break;
            }

            case "open-book": {
                if (type !== "folder") {
                    break;
                }
                await this.host.openBook?.(path);
                break;
            }

            case "open": {
                if (type === "file") {
                    await this.host.openPaths?.([path]);
                } else if (type === "folder") {
                    const paths = (selection.node.children || selection.children || [])
                        .filter(child => child.type === "file")
                        .map(child => path === "."
                            ? child.name
                            : `${path}/${child.name}`);

                    if (paths.length) {
                        await this.host.openPaths?.(paths);
                    }
                }

                break;
            }

            case "download": {
                if (type === "file") {
                    const result = this.vfs.readFile(path);
                    this.check(result);

                    this.download(
                        new Blob([result.data], {
                            type: "application/octet-stream"
                        }),
                        path.split("/").pop()
                    );
                } else if (type === "folder") {
                    const result = this.vfs.zipDirectory(path);
                    this.check(result);

                    this.download(
                        new Blob([result.data], {
                            type: "application/zip"
                        }),
                        `${path === "." ? "vfs" : selection.node.name}.zip`
                    );
                }

                break;
            }

            case "new": {
                if (!name || name.includes("/") || name.includes("\\")) {
                    throw new Error("Enter a valid file or folder name.");
                }

                const parent = type === "folder" ? path : ".";
                const newPath = parent === "."
                    ? name
                    : `${parent}/${name}`;

                if (kind === "folder") {
                    this.check(this.vfs.createDirectory(newPath));
                } else {
                    this.check(this.vfs.writeFile(newPath, ""));
                    await this.host.openPaths?.([newPath]);
                }

                break;
            }

            case "rename": {
                if (
                    !name ||
                    name.includes("/") ||
                    name.includes("\\") ||
                    path === "."
                ) {
                    throw new Error("Enter a valid new name.");
                }

                const normalizedPath = path.replace(/^\.\//, "");
                const parent = normalizedPath.includes("/")
                    ? normalizedPath.slice(0, normalizedPath.lastIndexOf("/"))
                    : "";

                const newPath = parent
                    ? `${parent}/${name}`
                    : name;

                this.check(this.vfs.renamePath(normalizedPath, newPath));

                await this.host.renamePaths?.(normalizedPath, newPath);

                break;
            }

            case "delete": {
                if (type === "folder") {
                    this.check(this.vfs.deleteDirectory(path));
                } else {
                    this.check(this.vfs.deleteFile(path));
                }

                const removedPaths = [];

                const collect = (node, nodePath) => {
                    if (node.type === "file") {
                        removedPaths.push(nodePath);
                        return;
                    }

                    for (const child of node.children || []) {
                        collect(
                            child,
                            nodePath === "."
                                ? child.name
                                : `${nodePath}/${child.name}`
                        );
                    }
                };

                collect(selection.node, path);

                await this.host.removePaths?.(removedPaths);

                break;
            }

            case "revert": {
                if (type !== "file" && type !== "folder") {
                    break;
                }

                const target = path.replace(/^\.\//, "");
                this.check(this.vfs.revertFile(target));

                await this.host.revertPaths?.(target, type);

                break;
            }
        }

        this.refresh();
    }

    download(blob, filename) {
        const url = URL.createObjectURL(blob);
        const anchor = document.createElement("a");

        anchor.href = url;
        anchor.download = filename;

        document.body.append(anchor);
        anchor.click();
        anchor.remove();

        setTimeout(() => URL.revokeObjectURL(url), 1000);
    }
}