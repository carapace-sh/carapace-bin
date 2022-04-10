package styles

import "github.com/rsteube/carapace/pkg/style"

var CarapaceBin = struct {
	HttpStatusInformational string `desc:"http informational"`
	HttpStatusSuccessful    string `desc:"http successful"`
	HttpStatusRedirection   string `desc:"http redirection"`
	HttpStatusClientError   string `desc:"http client error"`
	HttpStatusServerError   string `desc:"http server error"`

	HttpMethodCONNECT string `desc:"http CONNECT method"`
	HttpMethodDELETE  string `desc:"http DELETE method"`
	HttpMethodGET     string `desc:"http GET method"`
	HttpMethodHEAD    string `desc:"http HEAD method"`
	HttpMethodOPTIONS string `desc:"http OPTIONS method"`
	HttpMethodPATCH   string `desc:"http PATCH method"`
	HttpMethodPOST    string `desc:"http POST method"`
	HttpMethodPUT     string `desc:"http PUT method"`
	HttpMethodTRACE   string `desc:"http TRACE method"`
}{
	HttpStatusInformational: style.Blue,
	HttpStatusSuccessful:    style.Green,
	HttpStatusRedirection:   style.Yellow,
	HttpStatusClientError:   style.Red,
	HttpStatusServerError:   style.Magenta,

	HttpMethodCONNECT: style.Default,
	HttpMethodDELETE:  style.Red,
	HttpMethodGET:     style.Blue,
	HttpMethodHEAD:    style.Magenta,
	HttpMethodOPTIONS: style.Of(style.Dim, style.Blue),
	HttpMethodPATCH:   style.Green,
	HttpMethodPOST:    style.Of(style.Dim, style.Green),
	HttpMethodPUT:     style.Yellow,
	HttpMethodTRACE:   style.Default,
}

func init() {
	style.Register("carapace-bin", &CarapaceBin)
}
