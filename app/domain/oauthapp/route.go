package oauthapp

import (
	"net/http"

	"github.com/angrieralien/seeother/app/sdk/auth"
	"github.com/angrieralien/seeother/foundation/logger"
	"github.com/angrieralien/seeother/foundation/web"
)

// Config contains all the configuration for the auth app.
type Config struct {
	Auth              *auth.Auth
	Log               *logger.Logger
	TokenKey          string
	GoogleKey         string
	GoogleSecret      string
	GoogleUIURL       string
	GoogleCallBackURL string
	APIHost           string
}

// Routes adds the routes for the auth app.
func Routes(app *web.App, cfg Config) {
	api := newApp(cfg)

	app.HandlerFunc(http.MethodGet, "", "/api/auth/{provider}", api.authenticate)
	app.HandlerFunc(http.MethodGet, "", "/api/logout/{provider}", api.logout)
	app.HandlerFunc(http.MethodGet, "", "/api/auth/{provider}/callback", api.authCallback)
}
