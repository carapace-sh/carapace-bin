// package http contains http related actions
package http

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
)

// ActionRequestHeaders completes http request headers
//   Accept:application/json
//   Accept-Encoding:exi,br
func ActionRequestHeaders() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionRequestHeaderNames().Invoke(c).Suffix(":").ToA()
		case 1:
			return ActionRequestHeaderValues(c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionRequestHeaderNames completes http request header names
//   Accept-Charset (Character sets that are acceptable.)
//   Accept-Datetime (Acceptable version in time.)
func ActionRequestHeaderNames() carapace.Action {
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

// ActionRequestHeaderValues completes values for given request header
//   ActionRequestHeaderValues("Accept")
//   ActionRequestHeaderValues("Accept-Encoding")
func ActionRequestHeaderValues(header string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch header {
		// TODO complete more headers
		case "Accept":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return ActionMediaTypes().Invoke(c).Filter(c.Parts).ToMultiPartsA("/")
			})
		case "Accept-Encoding":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return ActionContentEncodingTokens().Invoke(c).Filter(c.Parts).ToA()
			})
		case "Accept-Language":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return os.ActionLanguages().Invoke(c).Filter(c.Parts).ToA()
			})
		case "Cache-Control":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return ActionCacheControlRequestDirectives()
			})
		case "Content-Encoding":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return ActionContentEncodingTokens().Invoke(c).Filter(c.Parts).ToA()
			})
		case "Content-Type":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return ActionMediaTypes().Invoke(c).Filter(c.Parts).ToMultiPartsA("/")
			})
		case "Transfer-Encoding":
			return ActionTransferEncodingTokens()
		case "User-Agent":
			return ActionUserAgents()
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionMediaTypes completes media types
//   "audio/ogg",
//   "image/gif",
func ActionMediaTypes() carapace.Action {
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
	)
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

// ActionTransferEncodingTokens completes transfer encoding tokens
//   chunked (Transfer in a series of chunks)
//   gzip (GZIP file format)
func ActionTransferEncodingTokens() carapace.Action {
	return carapace.ActionValuesDescribed(
		"chunked", "Transfer in a series of chunks",
		"compress", "UNIX \"compress\" data format",
		"deflate", "\"deflate\" compressed data",
		"gzip", "GZIP file format",
	)
}

// ActionCacheControlRequestDirectives completes Cache-Control directives for a request
//   no-store (The response may not be stored in any cache.)
//   no-transform (An intermediate cache or proxy cannot edit the response body.)
func ActionCacheControlRequestDirectives() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"max-age=", "The maximum amount of time a resource is considered fresh.",
				"max-stale", "Indicates the client will accept a stale response.",
				"min-fresh=", "Indicates the client wants a response that will still be fresh for at least the specified number of seconds.",
				"no-cache", "The response may be stored by any cache, even if the response is normally non-cacheable.",
				"no-store", "The response may not be stored in any cache.",
				"no-transform", "An intermediate cache or proxy cannot edit the response body.",
				"only-if-cached", "Set by the client to indicate \"do not use the network\" for the response.",
			)
		default:
			return carapace.ActionValues("")
		}
	})
}

// ActionUserAgents completes common user agents
//   curl/7.35.0 (Curl)
//   Lynx/2.8.8pre.4 libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/2.12.23 (Lynx)
func ActionUserAgents() carapace.Action {
	// https://www.networkinghowtos.com/howto/common-user-agent-list/
	return carapace.ActionValuesDescribed(
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36", "Google Chrome",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:53.0) Gecko/20100101 Firefox/53.0", "Mozilla Firefox",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393", "Microsoft Edge",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)", "Microsoft Internet Explorer 6 / IE 6",
		"Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)", "Microsoft Internet Explorer 7 / IE 7",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)", "Microsoft Internet Explorer 8 / IE 8",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0;  Trident/5.0)", "Microsoft Internet Explorer 9 / IE 9",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; MDDCJS)", "Microsoft Internet Explorer 10 / IE 10",
		"Mozilla/5.0 (compatible, MSIE 11, Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko", "Microsoft Internet Explorer 11 / IE 11",
		"Mozilla/5.0 (iPad; CPU OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H321 Safari/600.1.4", "Apple iPad",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1", "Apple iPhone",
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Googlebot (Google Search Engine Bot)",
		"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", "Bing Bot (Bing Search Engine Bot)",
		"Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG SM-G570Y Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36", "Samsung Phone",
		"Mozilla/5.0 (Linux; Android 5.0; SAMSUNG SM-N900 Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/2.1 Chrome/34.0.1847.76 Mobile Safari/537.36", "Samsung Galaxy Note 3",
		"Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG SM-N910F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36", "Samsung Galaxy Note 4",
		"Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7", "Google Nexus",
		"Mozilla/5.0 (Linux; Android 7.0; HTC 10 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.83 Mobile Safari/537.36", "HTC",
		"curl/7.35.0", "Curl",
		"Wget/1.15 (linux-gnu)", "Wget",
		"Lynx/2.8.8pre.4 libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/2.12.23", "Lynx",
	)
}
