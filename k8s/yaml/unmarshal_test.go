package yaml

import (
	_ "embed"
	"reflect"
	"testing"
)

var (
	//go:embed testdata/unmarshal.yaml
	objectsYAML []byte
)

func TestUnmarshalObjects(t *testing.T) {
	expected := []map[string]interface{}{
		{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "test-1",
			},
			"data": map[string]interface{}{
				"hello": "world",
			},
		},
		{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "test-2",
			},
			"data": map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "test-3",
			},
			"data": map[string]interface{}{
				"test": "string",
			},
		},
	}
	result, err := UnmarshalObjects[map[string]interface{}](objectsYAML)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if !reflect.DeepEqual(expected, result) {
		t.Error("result != expected")
	}
}
