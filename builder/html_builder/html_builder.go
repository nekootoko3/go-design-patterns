package html_builder

import (
	"fmt"
	"io/ioutil"
	"os"
)

type htmlBuilder struct {
	fileName string
	file     *os.File
	hasTitle bool
}

func NewHtmlBuilder() *htmlBuilder {
	return &htmlBuilder{hasTitle: false}
}

func (hb *htmlBuilder) MakeTitle(title string) {
	fileName := fmt.Sprintf("%s.html", title)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	hb.fileName = fileName
	hb.file = f
	hb.file.WriteString("<html>")
	hb.file.WriteString(fmt.Sprintf("<head>%s</head>\n", title))
	hb.hasTitle = true
}

func (hb *htmlBuilder) MakeString(str string) {
	if !hb.hasTitle {
		return
	}

	hb.file.WriteString(fmt.Sprintf("<p>%s</p>\n", str))
}

func (hb *htmlBuilder) MakeItems(items []string) {
	if !hb.hasTitle {
		return
	}
	hb.file.WriteString("<ul>\n")
	for _, item := range items {
		hb.file.WriteString(fmt.Sprintf("<li>%s</li>\n", item))
	}
	hb.file.WriteString("</ul>\n")
}

func (hb *htmlBuilder) Close() {
	if !hb.hasTitle {
		return
	}
	hb.file.WriteString("</html>")
	hb.file.Close()
}

func (hb *htmlBuilder) Print() {
	if !hb.hasTitle {
		return
	}
	content, err := ioutil.ReadFile(hb.fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}
