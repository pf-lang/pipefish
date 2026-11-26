export function parseFlags(source) {
    const flags = {};

    if (!source) return flags;

    const tokens = source.match(/--[^\s]+/g) || [];

    for (const token of tokens) {
        const flag = token.slice(2);
        const equals = flag.indexOf("=");

        if (equals !== -1) {
            const name = flag.slice(0, equals);
            const value = flag.slice(equals + 1);
            flags[name] = value;
        } else if (flag.startsWith("no-")) {
            flags[flag.slice(3)] = false;
        } else {
            flags[flag] = true;
        }
    }

    return flags;
}

export function setFlags(component, flags) {
    component.flags = { ...flags };
}