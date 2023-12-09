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

type ServiceProcessor struct {
	Object       *corev1.Service
	Action       Action
	Dependencies []processors.Processor
}

func (serviceProcessor *ServiceProcessor) getObject() client.Object {
	return serviceProcessor.Object
}

func (serviceProcessor *ServiceProcessor) getDependencies() []processors.Processor {
	return serviceProcessor.Dependencies
}

func (serviceProcessor *ServiceProcessor) getEmptyObject() client.Object {
	return &corev1.Service{}
}

func (serviceProcessor *ServiceProcessor) compareDiff(compareObject interface{}) string {
	apiVersionDiff := cmp.Diff(compareObject.(*corev1.Service).TypeMeta.APIVersion, serviceProcessor.Object.TypeMeta.APIVersion, cmpopts.IgnoreUnexported())
	annotationsDiff := cmp.Diff(compareObject.(*corev1.Service).Annotations, serviceProcessor.Object.Annotations,
		cmpopts.IgnoreUnexported(),
		cmpopts.IgnoreMapEntries(func(k, v string) bool {
			return k == "kubectl.kubernetes.io/last-applied-configuration"
		}),
	)
	specDiff := cmp.Diff(compareObject.(*corev1.Service).Spec, serviceProcessor.Object.Spec,
		cmpopts.IgnoreUnexported(),
		cmpopts.IgnoreFields(corev1.ServiceSpec{}, "ClusterIP"),
		cmpopts.IgnoreFields(corev1.ServiceSpec{}, "ClusterIPs"),
		cmpopts.IgnoreFields(corev1.ServiceSpec{}, "SessionAffinity"),
		cmpopts.IgnoreFields(corev1.ServiceSpec{}, "IPFamilies"),
		cmpopts.IgnoreFields(corev1.ServiceSpec{}, "IPFamilyPolicy"),
		cmpopts.IgnoreFields(corev1.ServiceSpec{}, "InternalTrafficPolicy"),
	)
	return fmt.Sprintf("%s%s%s", apiVersionDiff, annotationsDiff, specDiff)
}

func (serviceProcessor *ServiceProcessor) updateObject(object client.Object) {
	object.(*corev1.Service).TypeMeta.APIVersion = serviceProcessor.Object.TypeMeta.APIVersion
	object.(*corev1.Service).Annotations = serviceProcessor.Object.Annotations
	object.(*corev1.Service).Spec = serviceProcessor.Object.Spec
}

func (serviceProcessor *ServiceProcessor) IsReady(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return true, nil
}

func (serviceProcessor *ServiceProcessor) Process(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return processObject(r, log, ctx, application, serviceProcessor)
}
