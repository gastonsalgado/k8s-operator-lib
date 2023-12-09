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

type ServiceAccountProcessor struct {
	Object       *corev1.ServiceAccount
	Action       Action
	Dependencies []processors.Processor
}

func (serviceAccountProcessor *ServiceAccountProcessor) getObject() client.Object {
	return serviceAccountProcessor.Object
}

func (serviceAccountProcessor *ServiceAccountProcessor) getDependencies() []processors.Processor {
	return serviceAccountProcessor.Dependencies
}

func (serviceAccountProcessor *ServiceAccountProcessor) getEmptyObject() client.Object {
	return &corev1.ServiceAccount{}
}

func (serviceAccountProcessor *ServiceAccountProcessor) compareDiff(compareObject interface{}) string {
	apiVersionDiff := cmp.Diff(compareObject.(*corev1.ServiceAccount).TypeMeta.APIVersion, serviceAccountProcessor.Object.TypeMeta.APIVersion, cmpopts.IgnoreUnexported())
	annotationsDiff := cmp.Diff(compareObject.(*corev1.ServiceAccount).Annotations, serviceAccountProcessor.Object.Annotations, cmpopts.IgnoreUnexported())
	return fmt.Sprintf("%s%s", apiVersionDiff, annotationsDiff)
}

func (serviceAccountProcessor *ServiceAccountProcessor) updateObject(object client.Object) {
	object.(*corev1.ServiceAccount).TypeMeta.APIVersion = serviceAccountProcessor.Object.TypeMeta.APIVersion
	object.(*corev1.ServiceAccount).Annotations = serviceAccountProcessor.Object.Annotations
}

func (serviceAccountProcessor *ServiceAccountProcessor) IsReady(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return true, nil
}

func (serviceAccountProcessor *ServiceAccountProcessor) Process(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return processObject(r, log, ctx, application, serviceAccountProcessor)
}
