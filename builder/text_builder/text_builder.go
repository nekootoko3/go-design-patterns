package text_builder

import (
	"bytes"
	"fmt"
)

type textBuilder struct {
	buf      *bytes.Buffer
	hasTitle bool
}

func NewTextBuilder() *textBuilder {
	return &textBuilder{buf: bytes.NewBuffer([]byte{})}
}

func (tb *textBuilder) MakeTitle(title string) {
	tb.buf.WriteString("=========================\n")
	tb.buf.WriteString("\n")
	tb.buf.WriteString(fmt.Sprintf("# %s\n", title))
	tb.buf.WriteString("\n")
	tb.hasTitle = true
}

func (tb *textBuilder) MakeString(str string) {
	if !tb.hasTitle {
		return
	}

	tb.buf.WriteString(fmt.Sprintf("## %s\n", str))
	tb.buf.WriteString("\n")
}

func (tb *textBuilder) MakeItems(items []string) {
	if !tb.hasTitle {
		return
	}

	for _, item := range items {
		tb.buf.WriteString(fmt.Sprintf("- %s\n", item))
	}
	tb.buf.WriteString("\n")
}

func (tb *textBuilder) Close() {
	if !tb.hasTitle {
		return
	}

	tb.buf.WriteString("=========================\n")
}

func (tb *textBuilder) GetResult() string {
	if !tb.hasTitle {
		return "MakeTitle should be first"
	}

	return tb.buf.String()
}
