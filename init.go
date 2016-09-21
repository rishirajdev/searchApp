package app

import (
	"github.com/revel/revel"
	"workype-app/workype-api/app/database"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		HeaderFilter,                  // Add some security based headers.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

/*
InitDB to connection to database
*/
func InitDB() {
	uri := revel.Config.StringDefault("database.uri", "mongodb://localhost:27017")
	name := revel.Config.StringDefault("database.name", "workype")
	if err := database.Init(uri, name); err != nil {
		revel.INFO.Println("DB Error", err)
	}
}

 var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")
	c.Response.Out.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	if c.Request.Header.Get("Access-Control-Request-Method") != "" && ("OPTIONS" == c.Request.Method) {
		c.Response.Out.Header().Set("Access-Control-Allow-Origin", "http://52.11.130.33")
		c.Response.Out.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Response.Out.Header().Set("Access-Control-Allow-Headers",c.Request.Header.Get("Access-Control-Request-Headers"))
		c.Response.Out.Header().Set("Access-Control-Allow-Credentials", "true")
	//	revel.FilterAction(users.List).Remove(RouterFilter)
	//	revel.FilterAction(users.List).Remove(ParamsFilter)
		fc[0](c,fc[2:])
	}else{
	
		fc[0](c, fc[1:]) // Execute the next filter stage.
	}
}
