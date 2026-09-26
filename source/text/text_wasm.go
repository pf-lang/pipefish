//go:build js && wasm

package text

func Red(s string) string {
	return color(s, "red")
}

func Yellow(s string) string {
	return color(s, "yellow")
}

func Green(s string) string {
	return color(s, "green")
}

func Cyan(s string) string {
	return color(s, "cyan")
}

func Blue(s string) string {
	return color(s, "blue")
}

func Purple(s string) string {
	return color(s, "purple")
}

func color(s, c string) string {
	return("<span class = \"" + c + "\">" + s + "</span>")
}