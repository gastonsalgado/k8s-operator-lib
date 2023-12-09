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

type RoleProcessor struct {
	Object       *rbacv1.Role
	Action       Action
	Dependencies []processors.Processor
}

func (roleProcessor *RoleProcessor) getObject() client.Object {
	return roleProcessor.Object
}

func (roleProcessor *RoleProcessor) getDependencies() []processors.Processor {
	return roleProcessor.Dependencies
}

func (roleProcessor *RoleProcessor) getEmptyObject() client.Object {
	return &rbacv1.Role{}
}

func (roleProcessor *RoleProcessor) compareDiff(compareObject interface{}) string {
	apiVersionDiff := cmp.Diff(compareObject.(*rbacv1.Role).TypeMeta.APIVersion, roleProcessor.Object.TypeMeta.APIVersion, cmpopts.IgnoreUnexported())
	annotationsDiff := cmp.Diff(compareObject.(*rbacv1.Role).Annotations, roleProcessor.Object.Annotations, cmpopts.IgnoreUnexported())
	rulesDiff := cmp.Diff(compareObject.(*rbacv1.Role).Rules, roleProcessor.Object.Rules, cmpopts.IgnoreUnexported())
	return fmt.Sprintf("%s%s%s", apiVersionDiff, annotationsDiff, rulesDiff)
}

func (roleProcessor *RoleProcessor) updateObject(object client.Object) {
	object.(*rbacv1.Role).TypeMeta.APIVersion = roleProcessor.Object.TypeMeta.APIVersion
	object.(*rbacv1.Role).Annotations = roleProcessor.Object.Annotations
	object.(*rbacv1.Role).Rules = roleProcessor.Object.Rules
}

func (roleProcessor *RoleProcessor) IsReady(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return true, nil
}

func (roleProcessor *RoleProcessor) Process(r processors.Reconcile, log logr.Logger, ctx context.Context, application applications.Application) (bool, error) {
	return processObject(r, log, ctx, application, roleProcessor)
}
