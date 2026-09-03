package markdown

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/dtypes"
	"github.com/tim-hardcastle/pipefish/source/lexer"
)

type highlightWrapper struct {
	wrap     func(string, string) string
	brackets func(int, rune) string
	space    string
}

func BlockHighlighter(code string) string {
	lines := strings.Split(code, "\n")
	result := ""
	sep := ""
	for _, line := range lines {
		result = result + sep + htmlHighlighter(line)
		sep = "\n"
	}
	return result
}

func htmlHighlighter(code string) string {
	return highlightGivenWrapper([]rune(code), htmlWrapper)
}

// This provides HTML and terminal-code syntax highlighting.

func highlightGivenWrapper(code []rune, wrapper highlightWrapper) string {
	var out bytes.Buffer
	brackets := []rune{}
	runes := lexer.NewRuneSupplier(code, "Highlighter")
	for runes.CurrentRune() != '\n' && runes.CurrentRune() != 0 {
		switch {
		// First we deal with the brackets, which have their own rules.
		case leftBrackets.Contains(runes.CurrentRune()):
			out.WriteString(wrapper.brackets(len(brackets), runes.CurrentRune()))
			brackets = append(brackets, runes.CurrentRune())
		case rightBrackets.Contains(runes.CurrentRune()):
			if len(brackets) > 0 && bracketMatch[brackets[len(brackets)-1]] == runes.CurrentRune() {
				brackets = brackets[:len(brackets)-1]
			}
			out.WriteString(wrapper.brackets(len(brackets), runes.CurrentRune()))
		// We could be looking at protected punctuation.
		case IsProtectedPunctuation(runes.CurrentRune()) ||
			(runes.CurrentRune() == '!' && runes.PeekRune() == '='):
			out.WriteString(wrapper.wrap(string(runes.CurrentRune()), "reserved"))
		// A formatted string literal.
		case runes.CurrentRune() == '"':
			result, ok := runes.ReadPlaintextString('"')
			result = `"` + result
			if ok {
				result = result + `"`
			}
			out.WriteString(wrapper.wrap(result, "string"))
		// A plaintext string literal.
		case runes.CurrentRune() == '`':
			result, ok := runes.ReadPlaintextString('`')
			result = "`" + result
			if ok {
				result = result + "`"
			}
			out.WriteString(wrapper.wrap(result, "string"))
		// A rune literal.
		case runes.CurrentRune() == '\'':
			result, hasSingleQuote, _ := runes.ReadRuneLiteral()
			result = "'" + result
			if hasSingleQuote {
				result = result + "'"
			}
			out.WriteString(wrapper.wrap(result, "string"))
		// A comment.
		case runes.CurrentRune() == '/' && runes.PeekRune() == '/':
			result := "/" + runes.ReadComment()
			out.WriteString(wrapper.wrap(result, "comment"))
		// A comment.
		case runes.CurrentRune() == '~' && runes.PeekRune() == '~':
			result := "~" + runes.ReadComment()
			out.WriteString(wrapper.wrap(result, "docstring"))
		// Or a snippet.
		case runes.CurrentRune() == '-' && runes.PeekRune() == '-':
			out.WriteString(wrapper.wrap("--", "reserved"))
			runes.Next()
			runes.Next()
			for runes.CurrentRune() != '\n' && runes.CurrentRune() != 0 {
				out.WriteRune(runes.CurrentRune())
				runes.Next()
			}
		// A nondecimal integer literal.
		case runes.CurrentRune() == '0' && nondecimalIndicators.Contains(runes.PeekRune()):
			result := ""
			indicator := runes.PeekRune()
			switch indicator {
			case 'b', 'B':
				result = runes.ReadBinaryNumber()
			case 'o', 'O':
				result = runes.ReadOctalNumber()
			case 'x', 'X':
				result = runes.ReadHexNumber()
			}
			out.WriteString(wrapper.wrap("0"+string(indicator)+result, "number"))
		// It could be a perfectly ordinary number.
		case lexer.IsDigit(runes.CurrentRune()):
			result := runes.ReadNumber()
			out.WriteString(wrapper.wrap(result, "number"))
		// It could be an identifier (as the lexer defines that term, not the parser, e.g. it includes `else`.)
		case lexer.IsLegalStart(runes.CurrentRune()):
			result := runes.ReadIdentifier()
			switch {
			case control.Contains(result):
				out.WriteString(wrapper.wrap(result, "control"))
			case reserved.Contains(result):
				out.WriteString(wrapper.wrap(result, "reserved"))
			case enumlike.Match([]byte(result)):
				out.WriteString(wrapper.wrap(result, "constant"))
			case nativeTypes.Contains(result) || typelike.Match([]byte(result)):
				for runes.PeekRune() == '?' || runes.PeekRune() == '!' {
					runes.Next()
					result = result + string(runes.CurrentRune())
				}
				out.WriteString(wrapper.wrap(result, "type"))
			default:
				out.WriteString(wrapper.wrap(result, "identifier"))
			}
		case runes.CurrentRune() == ' ':
			out.WriteString(wrapper.space)
		default:
			out.WriteRune(runes.CurrentRune())
		}
		runes.Next()
	}
	return out.String()
}

func IsProtectedPunctuation(ch rune) bool {
	return ch == ',' || ch == ':' || ch == ';' || ch == '.' || ch == '='
}

var htmlWrapper = highlightWrapper{
	wrap: func(s, flavor string) string {
		return "<span class=\"" + flavor + "\">" + s + "</span>"
	},

	brackets: func(depth int, r rune) string {
		depth = depth % 3
		return "<span class=\"bracket-" + strconv.Itoa(depth) + "\">" + string(r) +
			"</span>"
	},

	space: "<span class=\"ws\"> </span>",
}

var (
	nondecimalIndicators = dtypes.SetOf('b', 'B', 'o', 'O', 'x', 'X')
	control              = dtypes.SetOf("break", "continue", "else", "test", "try", "cmd", "const", "def", "external", "global", "golang", "import", "include", "newtype", "private", "var")
	reserved             = dtypes.SetOf("and", "false", "given", "not", "or", "true", "->", ">>", "?>", "--")
	// Used by the syntax highlighter; should not be used by anything else without much forethought.
	// TODO --- there must be some principled way to generate this from something else.
	nativeTypes   = dtypes.SetOf("ok", "int", "string", "rune", "bool", "float", "error", "type", "pair", "list", "map", "set", "label", "func", "null", "snippet", "clone", "clones", "enum", "struct", "any", "ref", "tuple")
	enumlike, _   = regexp.Compile(`^[A-Z][A-Z_]+$`)
	typelike, _   = regexp.Compile(`^[A-Z][A-Z]*[a-z]+[A-Za-z]*$`)
	bracketMatch  = map[rune]rune{'(': ')', '[': ']', '{': '}'}
	leftBrackets  = dtypes.SetOf('(', '[', '{')
	rightBrackets = dtypes.SetOf(')', ']', '}')
)
