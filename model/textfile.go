package model

type TextFile struct {
	name    string
	content []byte
}

func NewTextFile(name string, content []byte) *TextFile {
	return &TextFile{name: name, content: content}
}

func (t *TextFile) Name() string {
	return t.name
}

func (t *TextFile) Content() []byte {
	return t.content
}
