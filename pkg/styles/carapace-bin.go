package styles

import "github.com/rsteube/carapace/pkg/style"

var CarapaceBin = struct {
	HttpStatusInformational string `desc:"http informational" tag:"http status styles"`
	HttpStatusSuccessful    string `desc:"http successful" tag:"http status styles"`
	HttpStatusRedirection   string `desc:"http redirection" tag:"http status styles"`
	HttpStatusClientError   string `desc:"http client error" tag:"http status styles"`
	HttpStatusServerError   string `desc:"http server error" tag:"http status styles"`

	HttpMethodCONNECT string `desc:"http CONNECT method" tag:"http method styles"`
	HttpMethodDELETE  string `desc:"http DELETE method" tag:"http method styles"`
	HttpMethodGET     string `desc:"http GET method" tag:"http method styles"`
	HttpMethodHEAD    string `desc:"http HEAD method" tag:"http method styles"`
	HttpMethodOPTIONS string `desc:"http OPTIONS method" tag:"http method styles"`
	HttpMethodPATCH   string `desc:"http PATCH method" tag:"http method styles"`
	HttpMethodPOST    string `desc:"http POST method" tag:"http method styles"`
	HttpMethodPUT     string `desc:"http PUT method" tag:"http method styles"`
	HttpMethodTRACE   string `desc:"http TRACE method" tag:"http method styles"`

	KillSignalTerm string `desc:"Default action is to terminate the process" tag:"kill signal styles"`
	KillSignalIgn  string `desc:"Default action is to ignore the signal" tag:"kill signal styles"`
	KillSignalCore string `desc:"Default action is to terminate the process and dump core" tag:"kill signal styles"`
	KillSignalStop string `desc:"Default action is to stop the proces" tag:"kill signal styles"`
	KillSignalCont string `desc:"Default action is to continue the process if it is currently stopped" tag:"kill signal styles"`
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

	KillSignalTerm: style.Red,
	KillSignalIgn:  style.Blue,
	KillSignalCore: style.Magenta,
	KillSignalStop: style.Yellow,
	KillSignalCont: style.Green,
}

func init() {
	style.Register("carapace-bin", &CarapaceBin)
}
