package markdown

import (
	"strings"

	"github.com/tim-hardcastle/pipefish/source/text"
)

type mdMode int

const (
	mdUnassigned mdMode = iota
	mdGettingParagraph
	mdGettingCodeBlock
	mdGettingCliBlock
	mdGettingQuote
	mdGettingList
	mdGettingHeading
	mdEnd
)

// The block parser.
func (rnd Renderer) Parse(raw string) mdDocument {
	lines := strings.Split(raw, "\n")
	docNodes := []mdNode{}
	accumulator := []string{}
	mode := mdUnassigned
mainloop:
	for i := range len(lines) + 1 {
		line := ""
		if i < len(lines) {
			line = lines[i]
		}
		newMode := mdGettingParagraph
		if mode == mdGettingCodeBlock || mode == mdGettingCliBlock {
			newMode = mode
		}
		switch {
		case line == "" && !(mode == mdGettingCodeBlock || mode == mdGettingCliBlock):
			newMode = mdUnassigned
		case text.Head(line, "#"):
			newMode = mdGettingHeading
		case text.Head(line, "* "):
			newMode = mdGettingList
		case text.Head(line, "> "):
			newMode = mdGettingQuote
		case text.Head(line, "```tui"):
			newMode = mdGettingCliBlock
		case text.Head(line, "```"):
			if mode == mdGettingCodeBlock || mode == mdGettingCliBlock {
				newMode = mdUnassigned
			} else {
				newMode = mdGettingCodeBlock
			}
		}
		if newMode == mode {
			accumulator = append(accumulator, line)
			continue mainloop
		}
		// So if we're here, we've switched from one sort of block to another.
		if len(accumulator) > 0 {
			// Then we can emit the appropriate old block.
			switch mode {
			case mdGettingHeading:
				docNodes = append(docNodes, makeHeading(accumulator))
			case mdGettingParagraph:
				docNodes = append(docNodes, makeParagraph(accumulator))
			case mdGettingCodeBlock:
				docNodes = append(docNodes, makeCodeBlock(accumulator))
			case mdGettingCliBlock:
				docNodes = append(docNodes, makeCliBlock(accumulator))
			case mdGettingQuote:
				docNodes = append(docNodes, makeQuote(accumulator))
			case mdGettingList:
				docNodes = append(docNodes, makeList(accumulator))
			}
		}
		// And we start a new block
		if text.Head(line, "```") { // We discard code block fences.
			accumulator = []string{}
		} else { // Otherwise the line that marked the end of the old block is the start of the new one.
			accumulator = []string{line}
		}
		mode = newMode
	}
	return mdDocument{docNodes}
}

func makeParagraph(lines []string) mdParagraph {
	ip := newInlineParser(strings.Join(lines, " "))
	return mdParagraph{ip.parseAll()}
}

func makeHeading(lines []string) mdHeading {
	heading := strings.Join(lines, " ")
	level := 0
	var ch rune
	for _, ch = range heading {
		if ch != '#' {
			break
		}
		level = level + 1
	}
	i := level
	for ; i < len(heading) && heading[i] == ' '; i++ {
	}
	ip := newInlineParser(heading[i:])
	return mdHeading{level, ip.parseAll()}
}

func makeCodeBlock(lines []string) mdCodeBlock {
	return mdCodeBlock{lines}
}

func makeCliBlock(lines []string) mdCliBlock {
	return mdCliBlock{lines}
}

func makeQuote(lines []string) mdQuote {
	quote := ""
	sep := ""
	for _, line := range lines {
		quote = quote + sep + line[2:]
		sep = " "
	}
	ip := newInlineParser(quote)
	return mdQuote{ip.parseAll()}
}

func makeList(lines []string) mdList {
	items := []mdNode{}
	for _, line := range lines {
		ip := newInlineParser(line[2:])
		lineItems := ip.parseAll()
		items = append(items, mdListItem{lineItems})
	}
	return mdList{items}
}

// The inline parser.

type inlineParser struct {
	line string
	pos  int
}

type parserMode int

const (
	pmNone parserMode = iota
	pmText
	pmBold
	pmItalic
)

var stopAt = map[parserMode][]string{
	pmText:   {"**", "*", "`"},
	pmBold:   {"**"},
	pmItalic: {"*"},
}

func (ip *inlineParser) parseAll() []mdNode {
	result := []mdNode{}
	for !ip.done() {
		result = append(result, ip.parse(pmNone)...)
	}
	return result
}

func (ip *inlineParser) parse(pM parserMode) []mdNode {
	txt := []byte{}
	for {
		if ip.done() {
			return []mdNode{mdText{string(txt)}}
		}
		// Inline code is handled differently since it doesn't get bold or italics or colors
		// in it.
		if pM == pmNone && ip.char() == '`' {
			return append([]mdNode{ip.parseInlineCode()}, ip.parse(pM)...)
		}
		if (pM == pmNone || pM == pmItalic) && ip.headIs("**") {
			ip.skip(2)
			bolded := ip.parse(pmBold)
			ip.skip(2)
			return append([]mdNode{mdFormat{stBold, bolded}})
		}
		if (pM == pmNone && ip.headIs("*")) ||
			(pM == pmBold && ip.headIs("*") && !ip.headIs("**")) {
			ip.next()
			italicized := ip.parse(pmItalic)
			ip.next()
			return append([]mdNode{mdFormat{stItalic, italicized}})
		}
		if pM == pmNone {
			pM = pmText
		}
		// If we've hit some closing item we were looking for, then we return the text we
		// accumulated as `mdText`, which may then be wrapped in something else by the caller.
		if upTo, ok := stopAt[pM]; ok {
			for _, limit := range upTo {
				if ip.headIs(limit) && !(pM == pmItalic && ip.headIs("**")) {
					return []mdNode{mdText{string(txt)}}
				}
			}
		}
		txt = append(txt, ip.char())
		ip.next()
	}
}

func (ip *inlineParser) parseInlineCode() mdInlineCode {
	code := []byte{}
	ip.next()
	for !ip.done() {
		if ip.char() == '`' {
			ip.next()
			break
		}
		code = append(code, ip.char())
		ip.next()
	}
	return mdInlineCode{string(code)}
}

func (ip *inlineParser) done() bool {
	return ip.pos >= len(ip.line)
}

func (ip *inlineParser) char() byte {
	return ip.line[ip.pos]
}

func (ip *inlineParser) next() {
	ip.pos = ip.pos + 1
}

func (ip *inlineParser) skip(i int) {
	ip.pos = ip.pos + i
}

func (ip *inlineParser) headIs(s string) bool {
	upperBound := ip.pos + len(s)
	return upperBound < len(ip.line) && ip.line[ip.pos:upperBound] == s
}

func (ip *inlineParser) passHead(s string) bool {
	upperBound := ip.pos + len(s)
	if upperBound < len(ip.line) && ip.line[ip.pos:upperBound] == s {
		ip.skip(len(s))
		return true
	}
	return false
}

func newInlineParser(raw string) inlineParser {
	return inlineParser{raw, 0}
}
