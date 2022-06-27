package zap

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	scanv1 "github.com/kappsec-io/kappsec/api/v1"
	"github.com/kappsec-io/kappsec/scanners"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ZapScanner struct {
	scanners.Scanner
	*scanv1.Zap
}

func New(ctx context.Context, logger logr.Logger, client client.Client, zap *scanv1.Zap) *ZapScanner {
	return &ZapScanner{
		Scanner: scanners.Scanner{
			Client:  client,
			Logger:  logger,
			Context: ctx,
		},
		Zap: zap,
	}
}

func (s *ZapScanner) Reconcile() error {
	s.Scanner.Logger.Info(fmt.Sprintf("Status: %s", s.Zap.Status.Phase))

	if s.Zap.Status.Phase == "" {
		//TODO: deploy zap job
		s.Zap.Status.Phase = "Running"
		s.Client.Status().Update(s.Scanner.Context, s.Zap)
	}

	// if s.Zap.Status.Phase == "Running" || s.Zap.Status.Phase == "Done" {
	// 	return nil
	// }

	// if s.Zap.Status.Phase == "Complete" {
	// 	//TODO: get result
	// 	s.Zap.Status.Phase = "Running"
	// 	s.Client.Status().Update(s.Scanner.Context, s.Zap)
	// }

	return fmt.Errorf("not implementaed")
}
