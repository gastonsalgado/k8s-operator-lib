package processors

import (
	"context"

	"github.com/gastonsalgado/k8s-operator-lib/applications"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconcile interface {
	client.Reader
	client.Writer
	client.StatusClient

	GetSchema() *runtime.Scheme
	GetRecorder() record.EventRecorder
	GetFinalizer() string
	CheckFinalizer(logr.Logger, context.Context, ctrl.Request, applications.Application) (bool, error)
}

type Processor interface {
	IsReady(Reconcile, logr.Logger, context.Context, applications.Application) (bool, error)
	Process(Reconcile, logr.Logger, context.Context, applications.Application) (bool, error)
}
