package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/dtypes"
	"github.com/tim-hardcastle/pipefish/source/markdown"
	"github.com/tim-hardcastle/pipefish/source/settings"
	"github.com/tim-hardcastle/pipefish/source/text"
)

var templates map[string]string

type sub struct{ tag, text string }

func create(template string, subs ...sub) string {
	result := templates[template]
	for _, sub := range subs {
		result = strings.Replace(result, "{{"+sub.tag+"}}", sub.text, 1)
	}
	return result
}

func init() {
	templates = map[string]string{}
	files, _ := os.ReadDir(filepath.Join(settings.PipefishHomeDirectory, "website/templates"))
	for _, file := range files {
		path := filepath.Join(settings.PipefishHomeDirectory, "website/templates", file.Name())
		templateName := file.Name()[:len(file.Name())-5]
		text, _ := os.ReadFile(path)
		templates[templateName] = string(text)
	}
}

func main() {
	// We create a temporary output folder, `website-build`, ignored by git.

	os.RemoveAll(filepath.Join(settings.PipefishHomeDirectory, "website-build"))
	os.Mkdir(filepath.Join(settings.PipefishHomeDirectory, "website-build"), 0o755)
	os.Mkdir(filepath.Join(settings.PipefishHomeDirectory, "website-build/docs"), 0o755)
	os.Mkdir(filepath.Join(settings.PipefishHomeDirectory, "website-build/articles"), 0o755)
	copyDir(filepath.Join(settings.PipefishHomeDirectory, "website/assets"),
		filepath.Join(settings.PipefishHomeDirectory, "website-build/assets"))

	mdR := markdown.NewHtmlRenderer()

	// Put the header and the footer into the all-purpose template.

	templates["all"] = create("main",
		sub{"header", templates["header"]},
		sub{"footer", templates["footer"]},
	)

	// We assemble the landing page.

	// We make the cards.
	var builder strings.Builder
	sb := &builder
	files, _ := os.ReadDir(filepath.Join(settings.PipefishHomeDirectory, "website/content/cards"))
	for _, file := range files {
		path := filepath.Join(settings.PipefishHomeDirectory, "website/content/cards", file.Name())
		text, _ := os.ReadFile(path)
		sb.WriteString(create("card", sub{"content", mdR.Render(string(text))}))
		sb.WriteString("\n\n")
	}
	landingPage := create("all",
		sub{"title", "Pipefish"},
		sub{"description", "Landing page for Pipefish."},
		sub{"content", create("landing", sub{"cards", sb.String()})},
	)
	indexPage := filepath.Join(settings.PipefishHomeDirectory, "website-build/index.html")
	os.WriteFile(indexPage, []byte(landingPage), 0755)

	// Before generating the docs pages, we need to convert the docstring descriptions of the
	// libraries into markdown, which we can then convert into HTML. We'll put the generated
	// markdown into `website/content/docs` so that it'll be converted to HTML like everything
	// else.

	libraries := ""
	filepath.WalkDir(stdlib, func(path string, d fs.DirEntry, err error) error {
		if path == filepath.Join(settings.PipefishHomeDirectory, "source/initializer/libraries") {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		name := outputFile(path)
		libraries = libraries + "- The " + name + " library\n"
		return nil
	})

	// Having got the list of libraries in the previous step, we can construct the sidebars.
	libraries = strings.TrimSpace(libraries)
	docsBytes, _ := os.ReadFile(filepath.Join(settings.PipefishHomeDirectory, "website/content/indices/docs.md"))
	essaysBytes, _ := os.ReadFile(filepath.Join(settings.PipefishHomeDirectory, "website/content/indices/articles.md"))
	sidebars := make(map[string]string)
	sidebars["articles"] = string(essaysBytes)
	sidebars["docs"] = strings.Replace(string(docsBytes), "{{libraries}}", libraries, 1)
	rawArticleList := strings.Split(string(essaysBytes), "/n")
	articleSet := make(dtypes.Set[string])
	for _, raw := range rawArticleList {
		if raw[:2] == "- " {
			articleSet = articleSet.Add(raw[2:])
		}
	}

	for _, flavor := range []string{"articles", "docs"} {
		sidebars[flavor] = makeSidebar(flavor, sidebars[flavor])
	}

	// We convert everything in the docs folder from markdown to HTML and yeet the
	// results into the appropriate output folder.
	docsDir := filepath.Join(settings.PipefishHomeDirectory, "website/content/docs")
	files, _ = os.ReadDir(docsDir)
	for _, file := range files {
		path := filepath.Join(docsDir, file.Name())
		text, _ := os.ReadFile(path)
		name := file.Name()[0 : len(file.Name())-3]
		flavor := "docs"
		if articleSet.Contains(name) {
			flavor = "articles"
		}
		title := strings.ReplaceAll(name, "-", " ")
		ast := mdR.Parse(string(text))
		headInfo := mdR.ExtractHeadings(ast)
		tocContent := makeToc(headInfo)
		toc := create("toc", sub{"toc", tocContent})
		if tocContent == "" {
			toc = ""
		}
		article := create("article",
			sub{"title", title},
			sub{"maybe-toc", toc},
			sub{"content", mdR.RenderAst(ast)},
		)
		articlePlusSidebar := create(flavor,
			sub{"sidebar", sidebars[flavor]},
			sub{"content", article},
		)
		var description string
		if flavor == "docs" {
			description = "Description of " + strings.ToLower(title) + " in Pipefish."
		} else {
			description = "Article on '" + title + "'."
		}
		target := filepath.Join(settings.PipefishHomeDirectory, "website-build/docs", name+".html")
		page := create("all",
			sub{"title", title},
			sub{"description", description},
			sub{"content", articlePlusSidebar},
		)
		os.WriteFile(target, []byte(page), 0755)
	}
}

func makeSidebar(flavor, raw string) string {
	atStart := true
	result := ""
	for _, line := range strings.Split(raw, "\n") {
		switch {
		case line[0:3] == "## ":
			if !atStart {
				result = result + "</ul>\n</section>\n\n"
			}
			atStart = false
			result = result + "<section>\n<h3>" + line[3:] + "</h3>\n<ul>\n"
		case line[0:2] == "- ":
			name := line[2:]
			result = result + "<li><a href=\"/" + flavor + "/" + text.Hyphenate(name) + ".html\">" + name + "</a></li>\n"
		}
	}
	result = result + "</ul>\n</section>\n\n"
	return result
}

// Makes the table of contents. Each `ContentsItem` contains one <h2> heading and the following
// <h3> headings.
func makeToc(info []markdown.ContentsItem) string {
	if len(info) == 0 {
		return ""
	}
	var builder strings.Builder
	sb := &builder
	sb.WriteString("<ol>\n")
	for _, item := range info {
		sb.WriteString("<li><a href=\"#")
		sb.WriteString(text.Hyphenate(item.Heading))
		sb.WriteString("\">")
		sb.WriteString(item.Heading)
		sb.WriteString("</a>")
		if len(item.Subheading) > 0 {
			sb.WriteString("\n<ul>\n")
			for _, subhead := range item.Subheading {
				sb.WriteString("<li><a href=\"#")
				sb.WriteString(text.Hyphenate(subhead))
				sb.WriteString("\">")
				sb.WriteString(subhead)
				sb.WriteString("</a></li>\n")
			}
		}
		sb.WriteString("</ul>\n")
	}
	sb.WriteString("</ol>\n")
	return sb.String()
}

func copyDir(src string, dst string) {
	srcInfo, _ := os.Stat(src)
	os.MkdirAll(dst, srcInfo.Mode())
	entries, _ := os.ReadDir(src)
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			copyDir(srcPath, dstPath)
		} else {
			copyFile(srcPath, dst)
		}
	}
}

func copyFile(srcFile, destDir string) {
	text, _ := os.ReadFile(srcFile)
	os.MkdirAll(destDir, 0755)
	destFile := filepath.Join(destDir, filepath.Base(srcFile))
	os.WriteFile(destFile, text, 0o775)
}

var stdlib = filepath.Join(settings.PipefishHomeDirectory, "source", "initializer", "libraries")

func outputFile(path string) string {
	// We use Pipefish's CLI `wiki` command to generate the output.
	cmd := exec.Command(filepath.Join(settings.PipefishHomeDirectory, "pipefish"), "wiki", path)
	content, _ := cmd.CombinedOutput()
	outputFolder := filepath.Join(settings.PipefishHomeDirectory, "website", "content", "docs")
	rel, _ := filepath.Rel(stdlib, path)
	name := strings.TrimSuffix(rel, ".pf")
	hyphenatedName := text.Hyphenate(name)
	outputFile := filepath.Join(
		outputFolder,
		"The-"+hyphenatedName+"-library.md",
	)
	os.WriteFile(outputFile, []byte(content), 0755)
	return name
}
