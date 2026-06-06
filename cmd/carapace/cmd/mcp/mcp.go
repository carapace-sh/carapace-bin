package mcp

import (
	"encoding/json"
	"fmt"
	"io"
)

type MCPServer struct {
	version string
	stdin   io.Reader
	stdout  io.Writer
}

func NewMCPServer(version string, stdin io.Reader, stdout io.Writer) *MCPServer {
	return &MCPServer{
		version: version,
		stdin:   stdin,
		stdout:  stdout,
	}
}

func (s *MCPServer) Run() error {
	decoder := json.NewDecoder(s.stdin)
	encoder := json.NewEncoder(s.stdout)

	for {
		var message json.RawMessage
		if err := decoder.Decode(&message); err != nil {
			if isEOF(err) {
				return nil
			}
			return err
		}

		result, err := s.processMessage(message)
		if err != nil {
			return err
		}
		if result != nil {
			if err := encoder.Encode(result); err != nil {
				return err
			}
		}
	}
}

func (s *MCPServer) processMessage(message json.RawMessage) (any, error) {
	switch firstNonSpace(message) {
	case '[':
		var requests []mcpRequest
		if err := json.Unmarshal(message, &requests); err != nil {
			return nil, err
		}
		responses := make([]mcpResponse, 0, len(requests))
		for _, request := range requests {
			response, ok := s.handleRequest(request)
			if ok {
				responses = append(responses, response)
			}
		}
		if len(responses) == 0 {
			return nil, nil
		}
		return responses, nil
	default:
		var request mcpRequest
		if err := json.Unmarshal(message, &request); err != nil {
			return nil, err
		}
		response, ok := s.handleRequest(request)
		if !ok {
			return nil, nil
		}
		return response, nil
	}
}

func (s *MCPServer) handleRequest(request mcpRequest) (mcpResponse, bool) {
	if len(request.ID) == 0 {
		return mcpResponse{}, false
	}

	response := mcpResponse{
		JSONRPC: "2.0",
		ID:      request.ID,
	}

	switch request.Method {
	case "initialize":
		response.Result = map[string]any{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]any{
				"tools": map[string]any{},
			},
			"serverInfo": map[string]any{
				"name":    "carapace",
				"version": s.version,
			},
		}
	case "tools/list":
		response.Result = map[string]any{
			"tools": tools,
		}
	case "tools/call":
		result, err := s.handleToolCall(request.Params)
		if err != nil {
			response.Error = &mcpError{Code: -32602, Message: err.Error()}
			return response, true
		}
		response.Result = result
	default:
		response.Error = &mcpError{Code: -32601, Message: "method not found"}
	}

	return response, true
}

func (s *MCPServer) handleToolCall(params json.RawMessage) (map[string]any, error) {
	var call mcpToolCallParams
	if err := json.Unmarshal(params, &call); err != nil {
		return nil, fmt.Errorf("invalid tool call parameters: %w", err)
	}

	switch call.Name {
	case "complete_command":
		return s.handleCompleteCommand(call.Arguments)
	// TODO: re-enable complete_macro once macro invocation is properly implemented
	// case "complete_macro":
	// 	return s.handleCompleteMacro(call.Arguments)
	case "list_macros":
		return s.handleListMacros(call.Arguments)
	case "codegen":
		return s.handleCodegen(call.Arguments)
	default:
		return nil, fmt.Errorf("unknown tool %q", call.Name)
	}
}
