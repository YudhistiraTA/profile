package lib

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"go.abhg.dev/goldmark/toc"
)

func MdParse(dbResponse string) (body string, tableOfContent string) {
	content := []byte(dbResponse)
	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	)

	if err := markdown.Convert(content, &buf); err != nil {
		return "", ""
	}

	raw := markdown.Parser().Parse(text.NewReader(content))
	tree, err := toc.Inspect(raw, content)
	if err != nil {
		return "", ""
	}

	renderTree := toc.RenderList(tree)
	var list bytes.Buffer
	markdown.Renderer().Render(&list, content, renderTree)

	return buf.String(), list.String()
}
