package api

import (
	"context"
	"github.com/gitcodeporter/okex-sdk"
	"github.com/gitcodeporter/okex-sdk/api/rest"
	"github.com/gitcodeporter/okex-sdk/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okex.Destination) (*Client, error) {
	restURL := okex.RestURL
	wsPubURL := okex.PublicWsURL
	wsPriURL := okex.PrivateWsURL
	wsBusURL := okex.BusinessWsURL
	switch destination {
	case okex.AwsServer:
		restURL = okex.AwsRestURL
		wsPubURL = okex.AwsPublicWsURL
		wsPriURL = okex.AwsPrivateWsURL
	case okex.DemoServer:
		restURL = okex.DemoRestURL
		wsPubURL = okex.DemoPublicWsURL
		wsPriURL = okex.DemoPrivateWsURL
	}
	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[okex.URLType]okex.BaseURL{
		okex.PublicURL:   wsPubURL,
		okex.PrivateURL:  wsPriURL,
		okex.BusinessURL: wsBusURL,
	})
	return &Client{r, c, ctx}, nil
}
