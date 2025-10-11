package http

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionStatusCodes completes http status codes
//
//	100 (Continue)
//	200 (OK)
func ActionStatusCodes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"100", "Continue",
		"101", "Switching Protocol",
		"102", "Processing",
		"103", "Early Hints",
		"200", "OK",
		"201", "Created",
		"202", "Accepted",
		"203", "Non-Authoritative Information",
		"204", "No Content",
		"205", "Reset Content",
		"206", "Partial Content",
		"207", "Multi-Status",
		"208", "Already Reported",
		"226", "IM Used",
		"300", "Multiple Choice",
		"301", "Moved Permanently",
		"302", "Found",
		"303", "See Other",
		"304", "Not Modified",
		"305", "Use Proxy",
		"306", "unused",
		"307", "Temporary Redirect",
		"308", "Permanent Redirect",
		"400", "Bad Request",
		"401", "Unauthorized",
		"402", "Payment Required",
		"403", "Forbidden",
		"404", "Not Found",
		"405", "Method Not Allowed",
		"406", "Not Acceptable",
		"407", "Proxy Authentication Required",
		"408", "Request Timeout",
		"409", "Conflict",
		"410", "Gone",
		"411", "Length Required",
		"412", "Precondition Failed",
		"413", "Payload Too Large",
		"414", "URI Too Long",
		"415", "Unsupported Media Type",
		"416", "Requested Range Not Satisfiable",
		"417", "Expectation Failed",
		"421", "Misdirected Request",
		"426", "Upgrade Required",
		"428", "Precondition Required",
		"429", "Too Many Requests",
		"431", "Request Header Fields Too Large",
		"451", "Unavailable For Legal Reasons",
		"500", "Internal Server Error",
		"501", "Not Implemented",
		"502", "Bad Gateway",
		"503", "Service Unavailable",
		"504", "Gateway Timeout",
		"505", "HTTP Version Not Supported",
		"506", "Variant Also Negotiates",
		"507", "Insufficient Storage",
		"508", "Loop Detected",
		"510", "Not Extended",
		"511", "Network Authentication Required",
	).StyleF(func(s string, _ style.Context) string {
		switch []rune(s)[0] {
		case '1':
			return styles.CarapaceBin.HttpStatusInformational
		case '2':
			return styles.CarapaceBin.HttpStatusSuccessful
		case '3':
			return styles.CarapaceBin.HttpStatusRedirection
		case '4':
			return styles.CarapaceBin.HttpStatusClientError
		case '5':
			return styles.CarapaceBin.HttpStatusServerError
		default:
			return ""
		}
	}).Uid("http", "status")
}
