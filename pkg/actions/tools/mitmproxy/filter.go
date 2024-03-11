package mitmproxy

import "github.com/carapace-sh/carapace"

func ActionFlowFilters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"a", "Match asset in response: CSS, JavaScript, images, fonts",
		"all", "Match all flows",
		"b", "Body",
		"bq", "Request body",
		"bs", "Response body",
		"c", "HTTP response code",
		"comment", "Flow comment",
		"d", "Domain",
		"dns", "Match DNS flows",
		"dst", "Match destination address",
		"e", "Match error",
		"h", "Header",
		"hq", "Request header",
		"hs", "Response header",
		"http", "Match HTTP flows",
		"m", "Method",
		"marked", "Match marked flows",
		"marker", "Match marked flows with specified marker",
		"meta", "Flow metadata",
		"q", "Match request with no response",
		"replay", "Match replayed flows",
		"replayq", "Match replayed client request",
		"replays", "Match replayed server response",
		"s", "Match response",
		"src", "Match source address",
		"t", "Content-type header",
		"tcp", "Match TCP flows",
		"tq", "Request Content-Type header",
		"ts", "Response Content-Type header",
		"u", "URL",
		"websocket", "Match WebSocket flows",
	)
}
