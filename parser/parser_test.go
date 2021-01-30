package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitize(t *testing.T) {
	d := []byte("// comment\n\n// comment comment\npush constant 10 // comment\n\npop local 0\n\n\npush constant 21\n")
	p := NewParser()
	p.Sanitize(d)
	assert.Equal(t, "push constant 10\npop local 0\npush constant 21", strings.Join(p.SourceLines, "\n"))
}
