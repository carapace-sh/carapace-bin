package rex

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionClientActions completes actions of a client or the server
//
//	com.superlogical.terminal.clear (Clear Terminal)
//	pane.split (Split a pane)
func ActionClientActions(client string) carapace.Action {
	args := []string{"actions", "--json"}
	if client != "" {
		args = append(args, "--client", client)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		return actionValues(output, client)
	}).Cache(24*time.Hour, key.String(client)).Tag("actions")
}

func actionValues(output []byte, client string) carapace.Action {
	actions, err := unmarshalActions(output)
	if err != nil {
		return carapace.ActionMessage(err.Error())
	}
	vals := make([]string, 0, len(actions)*2)
	for _, action := range actions {
		description := action.Title
		if description == "" {
			description = action.Description
		}
		vals = append(vals, action.Name, description)
	}
	return carapace.ActionValuesDescribed(vals...)
}

type action struct {
	Name        string           `json:"name"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Args        actionArgsSchema `json:"args"`
}

type actionArgsSchema struct {
	Properties map[string]struct {
		Type json.RawMessage `json:"type"`
	} `json:"properties"`
	Required []string `json:"required"`
}

func schemaType(raw json.RawMessage) string {
	switch {
	case len(raw) == 0:
		return ""
	case raw[0] == '"':
		var value string
		if err := json.Unmarshal(raw, &value); err == nil {
			return value
		}
	case raw[0] == '[':
		var values []string
		if err := json.Unmarshal(raw, &values); err == nil {
			return strings.Join(values, "|")
		}
	}
	return ""
}

func unmarshalActions(output []byte) ([]action, error) {
	var actions struct {
		Actions []action `json:"actions"`
	}
	if err := json.Unmarshal(output, &actions); err != nil {
		return nil, err
	}
	return actions.Actions, nil
}
