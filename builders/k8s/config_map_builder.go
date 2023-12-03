package k8sbuilders

import (
	"github.com/gastonsalgado/k8s-operator-lib/processors"
	k8sprocessors "github.com/gastonsalgado/k8s-operator-lib/processors/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const configMapAPIVersion = "v1"
const configMapKind = "ConfigMap"

type ConfigMapBuilder struct {
	object       *corev1.ConfigMap
	action       k8sprocessors.Action
	dependencies []processors.Processor
}

func newConfigMapBuilder(name string, namespace string, labels Labels) *ConfigMapBuilder {
	return &ConfigMapBuilder{
		object: &corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				APIVersion: configMapAPIVersion,
				Kind:       configMapKind,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
				Labels:    getCommonLabels(labels),
			},
		},
	}
}

func (configMapBuilder *ConfigMapBuilder) setApiVersion(apiVersion string) {
	configMapBuilder.object.TypeMeta.APIVersion = apiVersion
}

func (configMapBuilder *ConfigMapBuilder) setKind(kind string) {
	configMapBuilder.object.TypeMeta.Kind = kind
}

func (configMapBuilder *ConfigMapBuilder) setAnnotations(annotations map[string]string) {
	configMapBuilder.object.Annotations = annotations
}

func (configMapBuilder *ConfigMapBuilder) setData(data map[string]string) {
	configMapBuilder.object.Data = data
}

func (configMapBuilder *ConfigMapBuilder) setAction(action k8sprocessors.Action) {
	configMapBuilder.action = action
}

func (configMapBuilder *ConfigMapBuilder) setDependencies(dependencies []processors.Processor) {
	configMapBuilder.dependencies = dependencies
}

func (configMapBuilder *ConfigMapBuilder) getConfigMap() *k8sprocessors.ConfigMapsProcessor {
	return &k8sprocessors.ConfigMapsProcessor{
		Object:       configMapBuilder.object,
		Action:       configMapBuilder.action,
		Dependencies: configMapBuilder.dependencies,
	}
}
