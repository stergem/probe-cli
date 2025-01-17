package geolocate

import (
	"context"
	"net/http"
	"strings"

	"github.com/ooni/probe-cli/v3/internal/engine/httpheader"
	"github.com/ooni/probe-cli/v3/internal/httpx"
	"github.com/ooni/probe-cli/v3/internal/model"
)

func ipConfigIPLookup(
	ctx context.Context,
	httpClient *http.Client,
	logger model.Logger,
	userAgent string,
) (string, error) {
	data, err := (&httpx.APIClientTemplate{
		BaseURL:    "https://ipconfig.io",
		HTTPClient: httpClient,
		Logger:     logger,
		UserAgent:  httpheader.CLIUserAgent(),
	}).WithBodyLogging().Build().FetchResource(ctx, "/")
	if err != nil {
		return DefaultProbeIP, err
	}
	ip := strings.Trim(string(data), "\r\n\t ")
	logger.Debugf("ipconfig: body: %s", ip)
	return ip, nil
}
