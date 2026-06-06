package mcp

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type mcpCodegenRequest struct {
	Path string `json:"path"`
}

func (s *MCPServer) handleCodegen(args json.RawMessage) (map[string]any, error) {
	var request mcpCodegenRequest
	if err := json.Unmarshal(args, &request); err != nil {
		return nil, fmt.Errorf("invalid codegen arguments: %w", err)
	}
	if request.Path == "" {
		return nil, errors.New("path is required")
	}

	absPath := request.Path
	if !filepath.IsAbs(absPath) {
		abs, err := filepath.Abs(absPath)
		if err != nil {
			return mcpTextResult(err.Error(), true), nil
		}
		absPath = abs
	}

	executable, err := selfExecutable()
	if err != nil {
		return mcpTextResult(err.Error(), true), nil
	}

	cmdOutput, err := runCommandRaw(executable, []string{"--codegen", absPath}, os.Environ())
	if err != nil {
		return mcpTextResult(fmt.Sprintf("codegen failed: %v\n%s", err, cmdOutput), true), nil
	}

	var filenames []string
	scanner := bufio.NewScanner(strings.NewReader(cmdOutput))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasSuffix(line, ".go") {
			filenames = append(filenames, line)
		}
	}

	if len(filenames) == 0 {
		return mcpTextResult("codegen completed but no files were generated", false), nil
	}

	sort.Strings(filenames)
	return mcpJSONResult(filenames), nil
}
