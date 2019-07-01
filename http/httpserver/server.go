package httpserver

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/larry-dev/plugins/eglog"
	"github.com/larry-dev/plugins/ssl"
	"net/http"
	"os"
)

func NewServer(c Config) error {
	server := &http.Server{
		Addr:    c.ListenAddr,
		Handler: c.Handler,
	}
	// 开启TLS访问
	if c.EnableHttps {
		tlsC, err := ssl.ServerTslConfVerity(c.CertFile, c.KeyFile, c.Password)
		if err != nil {
			return fmt.Errorf("generate tls config failed. err: %v", err)
		}
		server.TLSConfig = tlsC
	}
	eglog.Info().Msgf("Serving %s with pid %d .", server.Addr, os.Getpid())
	if err := gracehttp.Serve(server); err != nil {
		eglog.Error().Msgf(err.Error())
		return err
	}
	return nil
}
