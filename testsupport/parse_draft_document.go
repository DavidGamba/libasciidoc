package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ParseDraftDocument parses the actual source with the options
func ParseDraftDocument(actual string) (types.DraftDocument, error) {
	r := strings.NewReader(actual)
	c := &draftDocumentParserConfig{
		filename: "test.adoc",
	}
	parserOptions := []parser.Option{}
	config := configuration.NewConfiguration(configuration.WithFilename(c.filename))
	rawDoc, err := parser.ParseRawDocument(r, config, parserOptions...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	return parser.ApplySubstitutions(rawDoc, config)
}

type draftDocumentParserConfig struct {
	filename string
}

func (c *draftDocumentParserConfig) setFilename(f string) {
	c.filename = f
}
