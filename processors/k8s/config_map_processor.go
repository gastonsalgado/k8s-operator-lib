package k8sprocessors

import (
	"context"
	"fmt"

	"github.com/gastonsalgado/k8s-operator-lib/processors"
	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ConfigMapsProcessor struct {
	Object       *corev1.ConfigMap
	Action       Action
	Dependencies []processors.Processor
}

func (configMapsProcessor *ConfigMapsProcessor) getObject() client.Object {
	return configMapsProcessor.Object
}

func (configMapsProcessor *ConfigMapsProcessor) getDependencies() []processors.Processor {
	return configMapsProcessor.Dependencies
}

func (configMapsProcessor *ConfigMapsProcessor) getEmptyObject() client.Object {
	return &corev1.ConfigMap{}
}

func (configMapsProcessor *ConfigMapsProcessor) compareDiff(compareObject interface{}) string {
	typeMetaDiff := cmp.Diff(compareObject.(*corev1.ConfigMap).TypeMeta, configMapsProcessor.Object.TypeMeta, cmpopts.IgnoreUnexported())
	annotationsDiff := cmp.Diff(compareObject.(*corev1.ConfigMap).Annotations, configMapsProcessor.Object.Annotations, cmpopts.IgnoreUnexported())
	dataDiff := cmp.Diff(compareObject.(*corev1.ConfigMap).Data, configMapsProcessor.Object.Data, cmpopts.IgnoreUnexported())
	return fmt.Sprintf("%s%s%s", typeMetaDiff, annotationsDiff, dataDiff)
}

func (configMapsProcessor *ConfigMapsProcessor) updateObject(object client.Object) {
	object.(*corev1.ConfigMap).TypeMeta.APIVersion = configMapsProcessor.Object.TypeMeta.APIVersion
	object.(*corev1.ConfigMap).TypeMeta.Kind = configMapsProcessor.Object.TypeMeta.Kind
	object.(*corev1.ConfigMap).Annotations = configMapsProcessor.Object.Annotations
	object.(*corev1.ConfigMap).Data = configMapsProcessor.Object.Data
}

func (configMapsProcessor *ConfigMapsProcessor) IsReady(r processors.Reconcile, log logr.Logger, ctx context.Context, application processors.Application) (bool, error) {
	return true, nil
}

func (configMapsProcessor *ConfigMapsProcessor) Process(r processors.Reconcile, log logr.Logger, ctx context.Context, application processors.Application) (bool, error) {
	return processObject(r, log, ctx, application, configMapsProcessor)
}
