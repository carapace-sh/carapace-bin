// package lore contains lore related actions
package lore

import "encoding/json"

// GlobalOpts contains global lore flags that affect which repository or data source a command reads.
type GlobalOpts struct {
	// Repository is the path to the repository to operate on
	Repository string

	// Identity is the identity to use
	Identity string

	// Offline forces offline mode
	Offline bool

	// Remote uses remote data
	Remote bool

	// Local uses local data
	Local bool
}

func (o GlobalOpts) args() []string {
	args := make([]string, 0, 8)
	if o.Repository != "" {
		args = append(args, "--repository", o.Repository)
	}
	if o.Identity != "" {
		args = append(args, "--identity", o.Identity)
	}
	if o.Offline {
		args = append(args, "--offline")
	}
	if o.Remote {
		args = append(args, "--remote")
	}
	if o.Local {
		args = append(args, "--local")
	}
	return args
}

// loreEvent is the top-level JSON event emitted by lore --json commands.
type loreEvent struct {
	TagName string          `json:"tagName"`
	Data    json.RawMessage `json:"data"`
}
