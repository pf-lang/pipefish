package compiler

import (
	"github.com/tim-hardcastle/pipefish/source/markdown"
	"github.com/tim-hardcastle/pipefish/source/text"
	"github.com/tim-hardcastle/pipefish/source/values"
)

func (cp *Compiler) Api(name string, path []string, fonts values.Map, width int) string {
	markdown := markdown.GetTuiRenderer(width)
	return cp.RenderApi(name, path, fonts, markdown, false)
}

func (cp *Compiler) Wiki(path []string) string {
	return cp.RenderApi("", path, values.Map{}, markdown.Iota, true)
}

func (cp *Compiler) RenderApi(name string, path []string, fonts values.Map, render func(string)string, wiki bool) string {
	if len(path) > 0 {
		newCp, ok := cp.Modules[path[0]]
		if !ok {
			return render("The module `" + path[0] + "` does not exist.")
		}
		if newCp.P.Private {
			return render("The module `" + path[0] + "` is private.")
		}
		return newCp.RenderApi(name, path[1:], fonts, render, wiki)
	}
	hasContents := false
	result := ""
	if name != "" {
		result = render("# " + name)
		result = result + "\n"
	}
	if cp.DocString != "" && !wiki {
		result = result + "\n" + render("## Overview")
		result = result + "\n\n"
		result = result + render(cp.DocString)
		result = result + "\n"
	}
	for i, items := range cp.ApiDescription {
		if len(items) == 0 {
			continue
		}
		hasContents = true
		result = result + "\n" + render("## " + headings[i]) + "\n"
		for _, item := range items {
			heading := item.Declaration
			if item.DocString != "" {
				heading = append(heading, ' ', ':')
			}
			if !wiki {
				result = result + "\n" + text.Cyan("•") + " " + cp.Highlight(heading, fonts) + "\n"
			} else {
				result = result + "\n### `" + string(heading) + "`\n"
			}
			if item.DocString != "" {
				result = result + "\n" + render(item.DocString) + "\n"
			}
		}
	}
	if !hasContents {
		result = result + "Nothing has been declared.\n"
	}
	return result + "\n"
}

type ApiItem struct {
	Declaration []rune
	DocString   string
}

var headings = []string{"Modules", "Types", "Constants", "Variables", "Commands", "Functions"}

