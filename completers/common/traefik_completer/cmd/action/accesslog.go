package action

import "github.com/carapace-sh/carapace"

func ActionAccessLogFieldNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ClientAddr", "The remote address in its original form (usually IP:port).",
		"ClientHost", "The remote IP address from which the client request was received.",
		"ClientPort", "The remote TCP port from which the client request was received.",
		"ClientUsername", "The username provided in the URL, if present.",
		"DownstreamContentSize", "The number of bytes in the response entity returned to the client.",
		"DownstreamStatus", "The HTTP status code returned to the client.",
		"DownstreamStatusLine", "DownstreamStatus + Status code explanation",
		"Duration", "The total time taken (in nanoseconds) by processing the response, including the origin server's time but not the log writing time.",
		"GzipRatio", "The response body compression ratio achieved.",
		"OriginContentSize", "The content length specified by the origin server, or 0 if unspecified.",
		"OriginDuration", "The time taken (in nanoseconds) by the origin server ('upstream') to return its response.",
		"OriginStatus", "The HTTP status code returned by the origin server.",
		"OriginStatusLine", "OriginStatus + Status code explanation",
		"Overhead", "The processing time overhead (in nanoseconds) caused by Traefik.",
		"RequestAddr", "The HTTP Host header (usually IP:port). This is treated as not a header by the Go API.",
		"RequestContentSize", "The number of bytes in the request entity (a.k.a. body) sent by the client.",
		"RequestCount", "The number of requests received since the Traefik instance started.",
		"RequestHost", "The HTTP Host server name (not including port).",
		"RequestLine", "RequestMethod + RequestPath + RequestProtocol",
		"RequestMethod", "The HTTP method.",
		"RequestPath", "The HTTP request URI, not including the scheme, host or port.",
		"RequestPort", "The TCP port from the HTTP Host.",
		"RequestProtocol", "The version of HTTP requested.",
		"RequestScheme", "The HTTP scheme requested http or https.",
		"RetryAttempts", "The amount of attempts the request was retried.",
		"RouterName", "The name of the Traefik router.",
		"ServiceAddr", "The IP:port of the Traefik backend (extracted from ServiceURL)",
		"ServiceName", "The name of the Traefik backend.",
		"ServiceURL", "The URL of the Traefik backend.",
		"StartLocal", "The local time at which request processing started.",
		"StartUTC", "The time at which request processing started.",
		"TLSCipher", "The TLS cipher used by the connection",
		"TLSVersion", "The TLS version used by the connection (e.g. 1.2) (if connection is TLS).",
	)
}
