package k8sprocessors

import (
	"context"
	"fmt"

	"github.com/gastonsalgado/k8s-operator-lib/applications"
	"github.com/gastonsalgado/k8s-operator-lib/processors"
	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	rbacv1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RoleBindingProcessor struct {
	Object       *rbacv1.RoleBinding
	Action       Action
	Dependencies []processors.Processor
}

func (roleBindingProcessor *RoleBindingProcessor) getObject() client.Object {
	return roleBindingProcessor.Object
}

func (roleBindingProcessor *RoleBindingProcessor) getDependencies() []processors.Processor {
	return roleBindingProcessor.Dependencies
}

func (roleBindingProcessor *RoleBindingProcessor) getEmptyObject() client.Object {
	return &rbacv1.RoleBinding{}
}

func (roleBindingProcessor *RoleBindingProcessor) compareDiff(compareObject interface{}) string {
	apiVersionDiff := cmp.Diff(compareObject.(*rbacv1.RoleBinding).TypeMeta.APIVersion, roleBindingProcessor.Object.TypeMeta.APIVersion, cmpopts.IgnoreUnexported())
	annotationsDiff := cmp.Diff(compareObject.(*rbacv1.RoleBinding).Annotations, roleBindingProcessor.Object.Annotations, cmpopts.IgnoreUnexported())
	roleRefDiff := cmp.Diff(compareObject.(*rbacv1.RoleBinding).RoleRef, roleBindingProcessor.Object.RoleRef, cmpopts.IgnoreUnexported())
	subjectsDiff := cmp.Diff(compareObject.(*rbacv1.RoleBinding).Subjects, roleBindingProcessor.Object.Subjects, cmpopts.IgnoreUnexported())
	return fmt.Sprintf("%s%s%s%s", apiVersionDiff, annotationsDiff, roleRefDiff, subjectsDiff)
}

func (roleBindingProcessor *RoleBindingProcessor) updateObject(object client.Object) {
	object.(*rbacv1.RoleBinding).TypeMeta.APIVersion = roleBindingProcessor.Object.TypeMeta.APIVersion
	object.(*rbacv1.RoleBinding).Annotations = roleBindingProcessor.Object.Annotations
	object.(*rbacv1.RoleBinding).RoleRef = roleBindingProcessor.Object.RoleRef
	object.(*rbacv1.RoleBinding).Subjects = roleBindingProcessor.Object.Subjects
}

func (roleBindingProcessor *RoleBindingProcessor) IsReady(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return true, nil
}

func (roleBindingProcessor *RoleBindingProcessor) Process(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return processObject(r, log, ctx, application, roleBindingProcessor)
}
