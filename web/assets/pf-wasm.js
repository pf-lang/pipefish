let ready = null;

export function loadWasm() {
if (ready) {
return ready;
}

ready = new Promise(async (resolve, reject) => {
    try {
        if (!window.Go) {
            await loadScript("assets/wasm_exec.js");
        }

        const go = new window.Go();

        window.pipefishReady = WebAssembly.instantiateStreaming(
            fetch("assets/pipefish.wasm"),
            go.importObject
        ).then(result => {
            go.run(result.instance);
        });

        await waitForFunctions();

        resolve();
    } catch (error) {
        reject(error);
    }
});

return ready;

}

function loadScript(url) {
return new Promise((resolve, reject) => {
const script = document.createElement("script");
    script.src = url;
    script.onload = resolve;
    script.onerror = reject;

    document.head.appendChild(script);
});
}

function waitForFunctions() {
return new Promise((resolve, reject) => {
const deadline = Date.now() + 10000;

    function check() {
        if (
            window.pipefishCompile &&
            window.pipefishCompileMain &&
            window.pipefishCreateAuthorDirectory &&
            window.pipefishCreateDirectory &&
            window.pipefishDeleteAuthorDirectory &&
            window.pipefishDeleteAuthorFile &&
            window.pipefishDeleteDirectory &&
            window.pipefishDeleteFile &&
            window.pipefishDo &&
            window.pipefishGetAuthorFiles &&
            window.pipefishGetAuthorFileTree &&
            window.pipefishGetBookFileTree &&
            window.pipefishGetFilepath &&
            window.pipefishGetFileTree &&
            window.pipefishHighlight &&
            window.pipefishInitAuthorFiles &&
            window.pipefishOpenBook &&
            window.pipefishReadAuthorFile && 
            window.pipefishReadBookFile &&   
            window.pipefishReadFile &&    
            window.pipefishRenameAuthorPath &&   
            window.pipefishRenamePath &&
            window.pipefishRenderMdAsHtml &&
            window.pipefishRevertFile &&
            window.pipefishUpdateFile &&
            window.pipefishWasmReady &&
            window.pipefishWriteAuthorFile &&
            window.pipefishWriteBookFile &&
            window.pipefishZipDirectory
        ) {
            resolve();
            return;
        }

        if (Date.now() >= deadline) {
            reject(
                new Error(
                    "Pipefish WASM did not initialize its functions."
                )
            );
            return;
        }

        setTimeout(check, 10);
    }

    check();
});

}
