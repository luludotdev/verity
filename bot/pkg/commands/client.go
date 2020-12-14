package commands

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lolPants/edsm-go"
	"github.com/lolPants/verity/bot/pkg/version"
)

var (
	api *edsm.Client
)

func init() {
	client := http.Client{
		Timeout:   time.Second * 10,
		Transport: &userAgentTransport{},
	}

	api = edsm.New(&client)
}

type userAgentTransport struct {
	rt http.RoundTripper
}

func (t *userAgentTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	var v string
	if version.IsDev() {
		v = fmt.Sprintf("git+%s", version.CommitHash())
	} else {
		v = version.AppVersion()
	}

	r.Header.Set("User-Agent", fmt.Sprintf("VerityBot/%s", v))
	return t.transport().RoundTrip(r)
}

func (t *userAgentTransport) transport() http.RoundTripper {
	if t.rt != nil {
		return t.rt
	}

	return http.DefaultTransport
}
