//go:build !js && !wasm

package text

func Red(s string) string {
	return RED + s + RESET
}

func Yellow(s string) string {
	return YELLOW + s + RESET
}

func Green(s string) string {
	return GREEN + s + RESET
}

func Cyan(s string) string {
	return CYAN + s + RESET
}

func Blue(s string) string {
	return BLUE + s + RESET
}

func Purple(s string) string {
	return PURPLE + s + RESET
}