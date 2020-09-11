// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ardanlabs/service/business/auth"
	"github.com/ardanlabs/service/business/mid"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, db *sqlx.DB, a *auth.Auth) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := check{
		db: db,
	}
	app.Handle(http.MethodGet, "/readiness", check.readiness)

	// Register user management and authentication endpoints.
	u := userHandlers{
		db:   db,
		auth: a,
	}
	app.Handle(http.MethodGet, "/users", u.query, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))
	app.Handle(http.MethodPost, "/users", u.create, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))
	app.Handle(http.MethodGet, "/users/:id", u.queryByID, mid.Authenticate(a))
	app.Handle(http.MethodPut, "/users/:id", u.update, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))
	app.Handle(http.MethodDelete, "/users/:id", u.delete, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))

	return app
}
