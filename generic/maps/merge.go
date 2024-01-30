package maps

import "maps"

// Merge recursivley merges b into a and returns the result as a new map.
// If the same entries in a and b contain maps, they are merged recursively.
// Other values are only shallowly copied into the result map.
//
// Credits go to https://stackoverflow.com/a/70291996
func Merge[K comparable](a, b map[K]interface{}) map[K]interface{} {
	out := make(map[K]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[K]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[K]interface{}); ok {
					out[k] = Merge(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

// MergeMany recursively merges a series maps using [Merge].
// Maps are merged in the order they are passed to the function.
func MergeMany[K comparable](a map[K]interface{}, others ...map[K]interface{}) map[K]interface{} {
	res := maps.Clone(a)
	for _, m := range others {
		res = Merge(res, m)
	}
	return res
}
