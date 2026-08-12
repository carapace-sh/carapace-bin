package lore

import (
	"bufio"
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

type branchListEntry struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Latest    string `json:"latest"`
	IsCurrent bool   `json:"isCurrent"`
	Archived  bool   `json:"archived"`
}

// ActionBranches completes branch names from `lore branch list --json`
//
//	main (e726318bbc3fd75ac8733a7e030cc35b)
//	arrr (019ff72384157580b25676c743f9b516)
func ActionBranches(opts GlobalOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append(opts.args(), "--json", "branch", "list")
		return carapace.ActionExecCommand("lore", args...)(func(output []byte) carapace.Action {
			var vals []string
			scanner := bufio.NewScanner(strings.NewReader(string(output)))
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line == "" {
					continue
				}
				var event loreEvent
				if err := json.Unmarshal([]byte(line), &event); err != nil {
					continue
				}
				if event.TagName != "branchListEntry" {
					continue
				}
				var entry branchListEntry
				if err := json.Unmarshal(event.Data, &entry); err != nil {
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
	}).Tag("branches")
}
