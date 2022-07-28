package templates

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Annotate(meta metav1.ObjectMeta, key, val string) metav1.ObjectMeta {
	if meta.Annotations == nil {
		meta.Annotations = make(map[string]string)
	}
	meta.Annotations[key] = val
	return meta
}
