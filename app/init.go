package app

import (
	"path"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"gopkg.in/gorp.v2"

	"database/sql"
	"fmt"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	// Database map
	DbMap *gorp.DbMap
)

func init() {
	revel.ConfPaths = append(
		revel.ConfPaths,
		path.Join(revel.BasePath, "conf", "dev"),
		path.Join(revel.BasePath, "conf", "prod"),
	)

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// Init database
	revel.OnAppStart(func() {
		cred := fmt.Sprintf(
			"user=%s password='%s' dbname=%s",
			revel.Config.StringDefault("db.user", ""),
			revel.Config.StringDefault("db.password", ""),
			revel.Config.StringDefault("db.name", ""),
		)
		db, err := sql.Open("postgres", cred)
		if err != nil {
			revel.INFO.Println("Database connection error:", err)
		}
		DbMap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	})
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
