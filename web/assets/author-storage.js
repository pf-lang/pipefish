const DATABASE_NAME = "pipefish-author";
const DATABASE_VERSION = 1;
const STORE_NAME = "state";

let databasePromise = null;

function openDatabase() {
    if (databasePromise) {
        return databasePromise;
    }

    databasePromise = new Promise((resolve, reject) => {
        const request = indexedDB.open(
            DATABASE_NAME,
            DATABASE_VERSION
        );

        request.onupgradeneeded = () => {
            const database = request.result;

            if (!database.objectStoreNames.contains(STORE_NAME)) {
                database.createObjectStore(STORE_NAME);
            }
        };

        request.onsuccess = () => {
            const database = request.result;

            database.onversionchange = () => {
                database.close();
                databasePromise = null;
            };

            resolve(database);
        };

        request.onerror = () => {
            databasePromise = null;
            reject(request.error);
        };

        request.onblocked = () => {
            console.warn(
                "Opening Pipefish author storage is blocked by another database connection."
            );
        };
    });

    return databasePromise;
}

async function runTransaction(mode, operation) {
    const database = await openDatabase();

    return new Promise((resolve, reject) => {
        const transaction = database.transaction(
            STORE_NAME,
            mode
        );

        let result;

        transaction.oncomplete = () => {
            resolve(result);
        };

        transaction.onerror = () => {
            reject(
                transaction.error ??
                new Error("Author storage transaction failed.")
            );
        };

        transaction.onabort = () => {
            reject(
                transaction.error ??
                new Error("Author storage transaction was aborted.")
            );
        };

        try {
            result = operation(
                transaction.objectStore(STORE_NAME)
            );
        } catch (error) {
            transaction.abort();
            reject(error);
        }
    });
}

export async function load(key) {
    const database = await openDatabase();

    return new Promise((resolve, reject) => {
        const transaction = database.transaction(
            STORE_NAME,
            "readonly"
        );

        const request = transaction
            .objectStore(STORE_NAME)
            .get(key);

        request.onsuccess = () => {
            resolve(request.result);
        };

        request.onerror = () => {
            reject(request.error);
        };
    });
}

export async function save(key, value) {
    await runTransaction("readwrite", store => {
        store.put(value, key);
    });
}

export async function remove(key) {
    await runTransaction("readwrite", store => {
        store.delete(key);
    });
}