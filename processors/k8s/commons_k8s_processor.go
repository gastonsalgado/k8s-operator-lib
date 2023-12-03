package k8sprocessor

import (
	"context"

	"github.com/gastonsalgado/k8s-operator-lib/processors"
	"github.com/go-logr/logr"
)

type K8sAction string

const (
	CreateOrUpdate K8sAction = "CreateOrUpdate"
	Delete         K8sAction = "Delete"
	Ignore         K8sAction = "Ignore"
)

func processObject(r processors.Reconcile, log logr.Logger, ctx context.Context, application processors.Application, processor processors.Processor) (bool, error) {
	return true, nil
}
