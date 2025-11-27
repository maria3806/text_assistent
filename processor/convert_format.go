package processor

import (
	"io/ioutil"
	"os"
	"strings"
)

func SaveAsMD(text, path string) error {
	if !strings.HasSuffix(path, ".md") {
		path += ".md"
	}
	return ioutil.WriteFile(path, []byte(text), os.ModePerm)
}

func SaveAsHTML(text, path string) error {
	if !strings.HasSuffix(path, ".html") {
		path += ".html"
	}
	html := "<html><body><p>" + text + "</p></body></html>"
	return ioutil.WriteFile(path, []byte(html), os.ModePerm)
}
