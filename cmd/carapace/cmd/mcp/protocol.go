package mcp

import (
	"encoding/json"
	"io"
)

type mcpRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type mcpResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *mcpError       `json:"error,omitempty"`
}

type mcpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type mcpToolCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
}

func firstNonSpace(message []byte) byte {
	for _, b := range message {
		switch b {
		case ' ', '\n', '\r', '\t':
			continue
		default:
			return b
		}
	}
	return 0
}

func mcpTextResult(text string, isError bool) map[string]any {
	result := map[string]any{
		"content": []map[string]string{
			{
				"type": "text",
				"text": text,
			},
		},
	}
	if isError {
		result["isError"] = true
	}
	return result
}

func mcpJSONResult(v any) map[string]any {
	return map[string]any{
		"content": []map[string]any{
			{
				"type": "text",
				"text": toJSON(v),
			},
		},
	}
}

func isEOF(err error) bool {
	return err == io.EOF
}
