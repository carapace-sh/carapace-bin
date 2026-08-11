package action

import (
	"bufio"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// loreArgs builds the global flag prefix for lore invocations, forwarding
// flags that affect which repository or data source the command reads.
func loreArgs(cmd *cobra.Command) []string {
	args := make([]string, 0, 8)

	if flag := cmd.Flag("repository"); flag != nil && flag.Changed {
		args = append(args, "--repository", flag.Value.String())
	}
	if flag := cmd.Flag("identity"); flag != nil && flag.Changed {
		args = append(args, "--identity", flag.Value.String())
	}
	if flag := cmd.Flag("offline"); flag != nil && flag.Changed {
		args = append(args, "--offline")
	}
	if flag := cmd.Flag("remote"); flag != nil && flag.Changed {
		args = append(args, "--remote")
	}
	if flag := cmd.Flag("local"); flag != nil && flag.Changed {
		args = append(args, "--local")
	}
	return args
}

type branchListEntry struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Latest    string `json:"latest"`
	IsCurrent bool   `json:"isCurrent"`
	Archived  bool   `json:"archived"`
}

// ActionBranches completes branch names from `lore branch list --json`.
func ActionBranches(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append(loreArgs(cmd), "--json", "branch", "list")
		return carapace.ActionExecCommand("lore", args...)(func(output []byte) carapace.Action {
			var vals []string
			scanner := bufio.NewScanner(strings.NewReader(string(output)))
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line == "" {
					continue
				}
				var event map[string]json.RawMessage
				if err := json.Unmarshal([]byte(line), &event); err != nil {
					continue
				}
				raw, ok := event["BranchListEntry"]
				if !ok {
					continue
				}
				var entry branchListEntry
				if err := json.Unmarshal(raw, &entry); err != nil {
					continue
				}
				desc := entry.ID
				if entry.IsCurrent {
					desc += " (current)"
				}
				if entry.Archived {
					desc += " (archived)"
				}
				vals = append(vals, entry.Name, desc)
			}
			if len(vals) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type revisionHistoryEntry struct {
	Revision       string   `json:"revision"`
	RevisionNumber uint64   `json:"revisionNumber"`
	Parent         []string `json:"parent"`
}

// ActionRevisions completes revision hashes from `lore revision history --json`.
func ActionRevisions(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append(loreArgs(cmd), "--json", "revision", "history")
		if flag := cmd.Flag("branch"); flag != nil && flag.Changed {
			args = append(args, "--branch", flag.Value.String())
		}
		if flag := cmd.Flag("revision"); flag != nil && flag.Changed {
			args = append(args, "--revision", flag.Value.String())
		}
		return carapace.ActionExecCommand("lore", args...)(func(output []byte) carapace.Action {
			var vals []string
			scanner := bufio.NewScanner(strings.NewReader(string(output)))
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line == "" {
					continue
				}
				var event map[string]json.RawMessage
				if err := json.Unmarshal([]byte(line), &event); err != nil {
					continue
				}
				raw, ok := event["RevisionHistoryEntry"]
				if !ok {
					continue
				}
				var entry revisionHistoryEntry
				if err := json.Unmarshal(raw, &entry); err != nil {
					continue
				}
				vals = append(vals, entry.Revision, "r"+strconv.FormatUint(entry.RevisionNumber, 10))
			}
			if len(vals) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type authIdentity struct {
	UserID string `json:"userId"`
}

// ActionIdentities completes user IDs from `lore auth list --json`.
func ActionIdentities(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append(loreArgs(cmd), "--json", "auth", "list")
		return carapace.ActionExecCommand("lore", args...)(func(output []byte) carapace.Action {
			var vals []string
			seen := make(map[string]bool)
			scanner := bufio.NewScanner(strings.NewReader(string(output)))
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line == "" {
					continue
				}
				var event map[string]json.RawMessage
				if err := json.Unmarshal([]byte(line), &event); err != nil {
					continue
				}
				raw, ok := event["AuthIdentity"]
				if !ok {
					continue
				}
				var identity authIdentity
				if err := json.Unmarshal(raw, &identity); err != nil {
					continue
				}
				if identity.UserID == "" || seen[identity.UserID] {
					continue
				}
				seen[identity.UserID] = true
				vals = append(vals, identity.UserID, identity.UserID)
			}
			if len(vals) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
