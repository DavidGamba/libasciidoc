package html5_test

import (
	"bytes"
	"context"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(t GinkgoTInterface, expectedResult, content string, rendererOpts ...renderer.Option) {
	// t.Logf("processing '%s'", content)
	reader := strings.NewReader(content)
	actualDocument, err := parser.ParseDocument("", reader)
	require.NoError(t, err, "Error found while parsing the document")
	t.Logf("actual document: `%s`", spew.Sdump(actualDocument))
	buff := bytes.NewBuffer(nil)
	rendererCtx := renderer.Wrap(context.Background(), actualDocument, rendererOpts...)
	// insert tables of contents, preamble and process file inclusions
	err = renderer.Prerender(rendererCtx)
	require.NoError(t, err)
	_, err = html5.Render(rendererCtx, buff)
	require.NoError(t, err)
	if strings.Contains(expectedResult, "{{.LastUpdated}}") {
		expectedResult = strings.Replace(expectedResult, "{{.LastUpdated}}", rendererCtx.LastUpdated(), 1)
	}
	t.Log("* Done processing document:")
	result := buff.String()
	// expectedResult = strings.Replace(expectedResult, "\t", "", -1)
	t.Logf("** Actual output:\n`%s`\n", result)
	t.Logf("** expectedResult output:\n`%s`\n", expectedResult) // remove tabs that can be inserted by VSCode while formatting the tests code
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(result, expectedResult, true)
	assert.Equal(t, expectedResult, result, dmp.DiffPrettyText(diffs))
}
