package yaml

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"unicode"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// UnmarshalFile reads a single object from a YAML file.
func UnmarshalFile(fsys fs.FS, filename string, target any) error {
	raw, err := fs.ReadFile(fsys, filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(raw, target)
}

// UnmarshalObjects parses all objects from a multi-document YAML stream.
// Documents that are empty (including no comments) or contain only white space
// are ignored.
func UnmarshalObjects[T any](rawYAML []byte) ([]T, error) {
	buf := bytes.NewBuffer(rawYAML)
	return UnmarshalObjectsReader[T](buf)
}

// UnmarshalObjectsFile call [UnmarshalObjectsReader] with a file stream.
func UnmarshalObjectsFile[T any](fsys fs.FS, filename string) ([]T, error) {
	file, err := fsys.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UnmarshalObjectsReader[T](file)
}

// UnmarshalObjectsReader parses all objects from a multi-document YAML stream.
// Documents that are empty (including no comments) or contain only white space
// are ignored.
func UnmarshalObjectsReader[T any](in io.Reader) ([]T, error) {
	objects := []T{}
	reader := yaml.NewYAMLReader(bufio.NewReader(in))
	for {
		data, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		if len(data) == 0 {
			continue
		}
		if isWhiteSpace(data) {
			continue
		}
		var o T
		if err := yaml.Unmarshal(data, &o); err != nil {
			return nil, err
		}
		objects = append(objects, o)
	}
	return objects, nil
}

// isWhiteSpace determines whether the passed in bytes are all unicode white
// space.
func isWhiteSpace(bytes []byte) bool {
	empty := true
	for _, b := range bytes {
		if !unicode.IsSpace(rune(b)) {
			empty = false
			break
		}
	}
	return empty
}
