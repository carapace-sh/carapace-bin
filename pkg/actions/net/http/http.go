// package http contains http related actions
package http

import "github.com/rsteube/carapace"

// ActionHttpRequestHeaderNames completes http reqest header names
//   Accept-Charset (Character sets that are acceptable.)
//   Accept-Datetime (Acceptable version in time.)

func ActionHttpRequestHeaderNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"A-IM", "Acceptable instance-manipulations for the request.[10]",
		"Accept", "Media type(s) that is/are acceptable for the response. See Content negotiation.",
		"Accept-Charset", "Character sets that are acceptable.",
		"Accept-Datetime", "Acceptable version in time.",
		"Accept-Encoding", "List of acceptable encodings. See HTTP compression.",
		"Accept-Language", "List of acceptable human languages for response. See Content negotiation.",
		"Access-Control-Request-Method", "Initiates a request for cross-origin resource sharing with Origin.",
		"Access-Control-Request-Headers", "Initiates a request for cross-origin resource sharing with Origin.",
		"Authorization", "Authentication credentials for HTTP authentication.",
		"Cache-Control", "Used to specify directives that must be obeyed by all caching mechanisms along the request-response chain.",
		"Connection", "Control options for the current connection and list of hop-by-hop request fields.[12] Must not be used with HTTP/2.[13]",
		"Content-Encoding", "The type of encoding used on the data. See HTTP compression.",
		"Content-Length", "The length of the request body in octets (8-bit bytes).",
		"Content-MD5", "A Base64-encoded binary MD5 sum of the content of the request body.",
		"Content-Type", "The Media type of the body of the request (used with POST and PUT requests).",
		"Cookie", "An HTTP cookie previously sent by the server with Set-Cookie (below).",
		"Date", "The date and time at which the message was originated (in \"HTTP-date\" format as defined by RFC 7231 Date/Time Formats).",
		"Expect", "Indicates that particular server behaviors are required by the client.",
		"Forwarded", "Disclose original information of a client connecting to a web server through an HTTP proxy.[15]",
		"From", "The email address of the user making the request.",
		"Host", "The domain name of the server (for virtual hosting), and the TCP port number on which the server is listening.",
		"HTTP2-Settings", "A request that upgrades from HTTP/1.1 to HTTP/2 MUST include exactly one HTTP2-Setting header field.",
		"If-Match", "Only perform the action if the client supplied entity matches the same entity on the server.",
		"If-Modified-Since", "Allows a 304 Not Modified to be returned if content is unchanged.",
		"If-None-Match", "Allows a 304 Not Modified to be returned if content is unchanged, see HTTP ETag.",
		"If-Range", "If the entity is unchanged, send me the part(s) that I am missing; otherwise, send me the entire new entity.",
		"If-Unmodified-Since", "Only send the response if the entity has not been modified since a specific time.",
		"Max-Forwards", "Limit the number of times the message can be forwarded through proxies or gateways.",
		"Origin", "Initiates a request for cross-origin resource sharing (asks server for Access-Control-* response fields).",
		"Pragma", "Implementation-specific fields that may have various effects anywhere along the request-response chain.",
		"Prefer", "Allows client to request that certain behaviors be employed by a server while processing a request.",
		"Proxy-Authorization", "Authorization credentials for connecting to a proxy.",
		"Range", "Request only part of an entity.  Bytes are numbered from 0.  See Byte serving.",
		"Referer", "This is the address of the previous web page from which a link to the currently requested page was followed.",
		"TE", "The transfer encodings the user agent is willing to accept",
		"Trailer", "The Trailer general field value indicates that the given set of header fields is present in the trailer of a message",
		"Transfer-Encoding", "The form of encoding used to safely transfer the entity to the user.",
		"User-Agent", "The user agent string of the user agent.",
		"Upgrade", "Ask the server to upgrade to another protocol. Must not be used in HTTP/2.[13]",
		"Via", "Informs the server of proxies through which the request was sent.",
		"Warning", "A general warning about possible problems with the entity body.",
	)
}

// ActionMediaTypes completes media types
//   "audio/ogg",
//   "image/gif",
func ActionMediaTypes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionValues(
			"application/x-executable",
			"application/graphql",
			"application/javascript",
			"application/json",
			"application/ld+json",
			"application/msword",
			"application/pdf",
			"application/sql",
			"application/vnd.api+json",
			"application/vnd.ms-excel",
			"application/vnd.ms-powerpoint",
			"application/vnd.oasis.opendocument.text",
			"application/vnd.openxmlformats-officedocument.presentationml.presentation",
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			"application/x-www-form-urlencoded",
			"application/xml",
			"application/zip",
			"application/zstd",
			"audio/mpeg",
			"audio/ogg",
			"image/gif",
			"image/apng",
			"image/flif",
			"image/webp",
			"image/x-mng",
			"image/jpeg",
			"image/png",
			"multipart/form-data",
			"text/css",
			"text/csv",
			"text/html",
			"text/php",
			"text/plain",
			"text/xml",
		).Invoke(c).ToMultiPartsA("/")
	})
}

// ActionContentEncodingTokens completes content encoding tokens
//   br (Brotli)
//   gzip (GNU zip format)
func ActionContentEncodingTokens() carapace.Action {
	return carapace.ActionValuesDescribed(
		"br", "Brotli",
		"compress", " UNIX \"compress\" program method",
		"deflate", " – compression based on the deflate algorithm",
		"exi", "W3C Efficient XML Interchange",
		"gzip", "GNU zip format",
		"identity", "No transformation is used.",
		"pack200-gzip", "Network Transfer Format for Java Archives",
		"zstd", "Zstandard compression",
	)
}

// ActionRequestMethods completes request methods
func ActionRequestMethods() carapace.Action {
	return carapace.ActionValuesDescribed(
		"GET", "The GET method requests a representation of the specified resource.",
		"HEAD", "The HEAD method asks for a response identical to that of a GET request, but without the response body.",
		"POST", "The POST method is used to submit an entity to the specified resource.",
		"PUT", "The PUT method replaces all current representations of the target resource with the request payload.",
		"DELETE", "The DELETE method deletes the specified resource.",
		"CONNECT", "The CONNECT method establishes a tunnel to the server identified by the target resource.",
		"OPTIONS", "The OPTIONS method is used to describe the communication options for the target resource.",
		"TRACE", "The TRACE method performs a message loop-back test along the path to the target resource.",
		"PATCH", "The PATCH method is used to apply partial modifications to a resource.",
	)
}
