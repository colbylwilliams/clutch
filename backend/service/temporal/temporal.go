package temporal

// <!-- START clutchdoc -->
// description: Workflow client for temporal.io.
// <!-- END clutchdoc -->

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/uber-go/tally/v4"
	"go.temporal.io/sdk/client"
	temporaltally "go.temporal.io/sdk/contrib/tally"
	"go.temporal.io/sdk/log"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	temporalv1 "github.com/lyft/clutch/backend/api/config/service/temporal/v1"
	"github.com/lyft/clutch/backend/service"
)

const Name = "clutch.service.temporal"

func New(cfg *anypb.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {
	config := &temporalv1.Config{}
	if err := cfg.UnmarshalTo(config); err != nil {
		return nil, err
	}
	return newClient(config, logger, scope)
}

type ClientManager interface {
	GetNamespaceClient(namespace string) (client.Client, error)
}

func newClient(cfg *temporalv1.Config, logger *zap.Logger, scope tally.Scope) (ClientManager, error) {
	ret := &clientImpl{
		hostPort:       fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		metricsHandler: temporaltally.NewMetricsHandler(scope),
		logger:         newTemporalLogger(logger),
		copts:          client.ConnectionOptions{},
	}

	if cfg.ConnectionOptions != nil {
		if cfg.ConnectionOptions.UseSystemCaBundle {
			certs, err := x509.SystemCertPool()
			if err != nil {
				return nil, err
			}
			ret.copts.TLS = &tls.Config{
				RootCAs:    certs,
				MinVersion: tls.VersionTLS12,
			}
		}
	}
	return ret, nil
}

type clientImpl struct {
	hostPort       string
	logger         log.Logger
	metricsHandler client.MetricsHandler
	copts          client.ConnectionOptions
}

type lazyClient struct {
}

func (c *clientImpl) GetNamespaceClient(namespace string) (client.Client, error) {
	tc, err := client.NewClient(client.Options{
		HostPort:          c.hostPort,
		Logger:            c.logger,
		MetricsHandler:    c.metricsHandler,
		Namespace:         namespace,
		ConnectionOptions: c.copts,
	})
	if err != nil {
		return nil, err
	}

	// TODO: cache clients? is there any benefit?

	tc.CompleteActivity()

	return tc, err
}
