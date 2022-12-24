package reverse

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

// /go:generate packr2
func write() (err error) {
	// TODO

	return
}

func doWriteFile(path, tmpl string) (err error) {
	data, err := parseTmpl(tmpl)
	if err != nil {
		return
	}

	fmt.Println(yellow("File generated ->"), cyan("$pwd"+path[len(argz.Pwd):]))

	content := string(data)
	content = replaceSpace(content)

	return os.WriteFile(path, []byte(content), 0o755)
}

func replaceSpace(content string) string {
	// TODO

	return content
}

func parseTmpl(tmpl string) ([]byte, error) {
	tmp, err := template.New("").Parse(tmpl)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	if err = tmp.Execute(&buf, ctx); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
