package k8sprocessors

import (
	"context"

	"github.com/gastonsalgado/k8s-operator-lib/processors"
	"github.com/go-logr/logr"
)

type Action string

const (
	CreateOrUpdate Action = "CreateOrUpdate"
	Delete         Action = "Delete"
	Ignore         Action = "Ignore"
)

func processObject(r processors.Reconcile, log logr.Logger, ctx context.Context, application processors.Application, processor processors.Processor) (bool, error) {
	return true, nil
}
