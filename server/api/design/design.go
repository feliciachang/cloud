package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var cors = func() {
	Headers("Authorization")
	Expose("Authorization")
	Methods("GET", "OPTIONS", "POST", "DELETE")
}

var _ = API("fieldkit", func() {
	Title("Fieldkit API")
	Description("A one-click open platform for field researchers and explorers.")
	Host("api.fieldkit.org")
	Scheme("https")
	Origin("https://fieldkit.org", cors)
	Origin("https://*.fieldkit.org", cors)
	Origin("https://fieldkit.team", cors)
	Origin("https://*.fieldkit.team", cors)
	Origin("http://localhost:3000", cors)
	Consumes("application/json")
	Produces("application/json")
})

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

var Location = MediaType("application/vnd.app.location+json", func() {
	TypeName("Location")
	Attributes(func() {
		Attribute("location", func() {
			Format("uri")
		})
		Required("location")
	})
	View("default", func() {
		Attribute("location")
	})
})
