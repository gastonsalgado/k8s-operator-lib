package k8sprocessors

import (
	"context"
	"fmt"

	"github.com/gastonsalgado/k8s-operator-lib/applications"
	"github.com/gastonsalgado/k8s-operator-lib/processors"
	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ConfigMapProcessor struct {
	Object       *corev1.ConfigMap
	Action       Action
	Dependencies []processors.Processor
}

func (configMapProcessor *ConfigMapProcessor) getObject() client.Object {
	return configMapProcessor.Object
}

func (configMapProcessor *ConfigMapProcessor) getDependencies() []processors.Processor {
	return configMapProcessor.Dependencies
}

func (configMapProcessor *ConfigMapProcessor) getEmptyObject() client.Object {
	return &corev1.ConfigMap{}
}

func (configMapProcessor *ConfigMapProcessor) compareDiff(compareObject interface{}) string {
	apiVersionDiff := cmp.Diff(compareObject.(*corev1.ConfigMap).TypeMeta.APIVersion, configMapProcessor.Object.TypeMeta.APIVersion, cmpopts.IgnoreUnexported())
	annotationsDiff := cmp.Diff(compareObject.(*corev1.ConfigMap).Annotations, configMapProcessor.Object.Annotations, cmpopts.IgnoreUnexported())
	dataDiff := cmp.Diff(compareObject.(*corev1.ConfigMap).Data, configMapProcessor.Object.Data, cmpopts.IgnoreUnexported())
	return fmt.Sprintf("%s%s%s", apiVersionDiff, annotationsDiff, dataDiff)
}

func (configMapProcessor *ConfigMapProcessor) updateObject(object client.Object) {
	object.(*corev1.ConfigMap).TypeMeta.APIVersion = configMapProcessor.Object.TypeMeta.APIVersion
	object.(*corev1.ConfigMap).TypeMeta.Kind = configMapProcessor.Object.TypeMeta.Kind
	object.(*corev1.ConfigMap).Annotations = configMapProcessor.Object.Annotations
	object.(*corev1.ConfigMap).Data = configMapProcessor.Object.Data
}

func (configMapProcessor *ConfigMapProcessor) IsReady(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return true, nil
}

func (configMapProcessor *ConfigMapProcessor) Process(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return processObject(r, log, ctx, application, configMapProcessor)
}
