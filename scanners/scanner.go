package scanners

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ScannerInterface interface {
	Reconcile() error
}

type Scanner struct {
	client.Client
	logr.Logger
	context.Context
}
