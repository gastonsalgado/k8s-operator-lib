package k8sprocessors

import (
	"context"
	"fmt"

	"github.com/gastonsalgado/k8s-operator-lib/applications"
	"github.com/gastonsalgado/k8s-operator-lib/processors"
	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeploymentProcessor struct {
	Object       *appsv1.Deployment
	Action       Action
	Dependencies []processors.Processor
}

func (deploymentProcessor *DeploymentProcessor) getObject() client.Object {
	return deploymentProcessor.Object
}

func (deploymentProcessor *DeploymentProcessor) getDependencies() []processors.Processor {
	return deploymentProcessor.Dependencies
}

func (deploymentProcessor *DeploymentProcessor) getEmptyObject() client.Object {
	return &appsv1.Deployment{}
}

func (deploymentProcessor *DeploymentProcessor) compareDiff(compareObject interface{}) string {
	apiVersionDiff := cmp.Diff(compareObject.(*appsv1.Deployment).TypeMeta.APIVersion, deploymentProcessor.Object.TypeMeta.APIVersion, cmpopts.IgnoreUnexported())
	annotationsDiff := cmp.Diff(compareObject.(*appsv1.Deployment).Annotations, deploymentProcessor.Object.Annotations, cmpopts.IgnoreUnexported())
	specDiff := cmp.Diff(compareObject.(*appsv1.Deployment).Spec, deploymentProcessor.Object.Spec,
		cmpopts.IgnoreUnexported(),
		cmpopts.IgnoreFields(appsv1.DeploymentSpec{}, "Replicas"),
		cmpopts.IgnoreFields(appsv1.DeploymentSpec{}, "ProgressDeadlineSeconds"),
		cmpopts.IgnoreFields(corev1.PodSpec{}, "RestartPolicy"),
		cmpopts.IgnoreFields(corev1.PodSpec{}, "DNSPolicy"),
		cmpopts.IgnoreFields(corev1.PodSpec{}, "DeprecatedServiceAccount"),
		cmpopts.IgnoreFields(corev1.PodSpec{}, "SecurityContext"),
		cmpopts.IgnoreFields(corev1.PodSpec{}, "SchedulerName"),
		cmpopts.IgnoreFields(corev1.SecretVolumeSource{}, "DefaultMode"),
		cmpopts.IgnoreFields(corev1.Container{}, "TerminantionMessagePath"),
		cmpopts.IgnoreFields(corev1.Container{}, "TerminantionMessagePolicy"),
	)
	return fmt.Sprintf("%s%s%s", apiVersionDiff, annotationsDiff, specDiff)
}

func (deploymentProcessor *DeploymentProcessor) updateObject(object client.Object) {
	object.(*appsv1.Deployment).TypeMeta.APIVersion = deploymentProcessor.Object.TypeMeta.APIVersion
	object.(*appsv1.Deployment).Annotations = deploymentProcessor.Object.Annotations
	object.(*appsv1.Deployment).Spec = deploymentProcessor.Object.Spec
}

func (deploymentProcessor *DeploymentProcessor) IsReady(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return true, nil
}

func (deploymentProcessor *DeploymentProcessor) Process(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return processObject(r, log, ctx, application, deploymentProcessor)
}
