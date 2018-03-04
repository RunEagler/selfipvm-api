package design


import (
	"github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)


var _ = apidsl.Resource("activity", func() {
	apidsl.BasePath("/activity")
	apidsl.Response(design.InternalServerError)
	apidsl.Action("entry", func() {
		apidsl.Description("day activity entry")
		apidsl.Routing(apidsl.POST(""))
		apidsl.Payload(ActivityPayload)
		apidsl.Response(design.OK)
		apidsl.Response(design.BadRequest)
	})
})

//ActivityPayload :POST /activity payload
var ActivityPayload = apidsl.Type("ActivityPayload",func(){
	apidsl.Member("date",design.String)
	apidsl.Member("type",design.Integer)
	apidsl.Member("minutes",design.Integer,func(){
		apidsl.Minimum(0)
		apidsl.Maximum(1440)
	})
	apidsl.Member("content",design.String)
	apidsl.Required("date","type","minutes")
})