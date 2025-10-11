package http

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionRequestMethods completes request methods
//
//	DELETE (The DELETE method deletes the specified resource.)
//	GET (The GET method requests a representation of the specified resource.)
func ActionRequestMethods() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionStyledValuesDescribed(
			"CONNECT", "Establish a tunnel to the server identified by the target resource", styles.CarapaceBin.HttpMethodCONNECT,
			"DELETE", "Delete the specified resource", styles.CarapaceBin.HttpMethodDELETE,
			"GET", "Request a representation of the specified resource", styles.CarapaceBin.HttpMethodGET,
			"HEAD", "Identical to a GET request, but without the response body", styles.CarapaceBin.HttpMethodHEAD,
			"OPTIONS", "Describe the communication options for the target resource", styles.CarapaceBin.HttpMethodOPTIONS,
			"PATCH", "Apply partial modifications to a resource", styles.CarapaceBin.HttpMethodPATCH,
			"POST", "Submit an entity to the specified resource", styles.CarapaceBin.HttpMethodPOST,
			"PUT", "Replace all current representations of the target resource", styles.CarapaceBin.HttpMethodPUT,
			"TRACE", "Perform a message loop-back test along the path to the target resource", styles.CarapaceBin.HttpMethodTRACE,
		)
	}).Tag("request methods").Uid("http", "method")
}
