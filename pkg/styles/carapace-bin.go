package styles

import "github.com/rsteube/carapace/pkg/style"

var CarapaceBin = struct {
	HttpInformational string `desc:"http informational"`
	HttpSuccessful    string `desc:"http successful"`
	HttpRedirection   string `desc:"http redirection"`
	HttpClientError   string `desc:"http client error"`
	HttpServerError   string `desc:"http server error"`
}{
	HttpInformational: style.Blue,
	HttpSuccessful:    style.Green,
	HttpRedirection:   style.Yellow,
	HttpClientError:   style.Red,
	HttpServerError:   style.Magenta,
}

func init() {
	style.Register("carapace-bin", &CarapaceBin)
}
