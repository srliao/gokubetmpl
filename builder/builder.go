package builder

import "path/filepath"

var registry = make(map[string]Resource)

type Resource interface {
	Write(path string) error
}

func Register(path string, r Resource) {
	_, ok := registry[path]
	if ok {
		panic("resource already registered: " + path)
	}
	registry[path] = r
}

func Build(root string) error {
	for path, r := range registry {
		if err := r.Write(filepath.Join(root, path)); err != nil {
			return err
		}
	}
	return nil
}
