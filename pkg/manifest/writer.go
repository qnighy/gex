package manifest

import (
	"bytes"
	"html/template"

	"github.com/spf13/afero"
)

// Writer creates a manifest file to manage tool dependencies.
type Writer interface {
	Write(path string, m *Manifest) error
}

// NewWriter creates a new Writer instance.
func NewWriter(fs afero.Fs) Writer {
	return &writerImpl{
		fs: fs,
	}
}

type writerImpl struct {
	fs afero.Fs
}

func (w *writerImpl) Write(path string, m *Manifest) error {
	buf := new(bytes.Buffer)
	err := toolsGoTemplate.Execute(buf, m)
	if err != nil {
		return err
	}
	err = afero.WriteFile(w.fs, path, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

var (
	toolsGoTemplate = template.Must(template.New("tools.go").Parse(`// Code generated by github.com/izumin5210/gex. DO NOT EDIT.

// +build tools

package tools

// tool dependencies
import (
{{- range $t := .Tools}}
	_ "{{$t}}"
{{- end}}
)
`))
)
