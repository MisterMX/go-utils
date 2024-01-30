package maps

import (
	"reflect"
	"testing"
)

func TestMergeMapsString(t *testing.T) {
	a := map[string]interface{}{
		"string": "string",
		"number": 0.0,
		"slice": []interface{}{
			"a",
			"b",
		},
		"map": map[string]interface{}{
			"a": "a",
			"b": "b",
		},
	}
	b := map[string]interface{}{
		"slice":  []interface{}{},
		"number": 100.0,
	}
	c := map[string]interface{}{
		"map": map[string]interface{}{
			"c": "c",
		},
	}
	expected := map[string]interface{}{
		"string": "string",
		"number": 100.0,
		"slice":  []interface{}{},
		"map": map[string]interface{}{
			"a": "a",
			"b": "b",
			"c": "c",
		},
	}

	merged := MergeMany(a, b, c)
	if !reflect.DeepEqual(expected, merged) {
		t.Error("merged != expected")
	}
}
