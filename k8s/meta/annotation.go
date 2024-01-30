package meta

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HasAnnotation determines if o contains an annotation key.
func HasAnnotation(o metav1.Object, key string) bool {
	_, exists := o.GetAnnotations()[key]
	return exists
}

// GetAnnotation returns the value of the annotation key of the object o
// and a boolean that determines if the key exists in the annotation map.
// If the key does not exist, the returned values is an empty string.
func GetAnnotation(o metav1.Object, key string) (string, bool) {
	val, exists := o.GetAnnotations()[key]
	return val, exists
}

// GetAnnotationFallback returns the value of the annotation key of the
// object o or fallback if it does not exist.
func GetAnnotationFallback(o metav1.Object, key, fallback string) string {
	val, exists := GetAnnotation(o, key)
	if !exists {
		return fallback
	}
	return val
}
