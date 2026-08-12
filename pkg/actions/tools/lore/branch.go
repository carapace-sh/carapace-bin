package lore

import (
	"bufio"
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
)

type branchListEntry struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Latest    string `json:"latest"`
	Location  string `json:"location"`
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
			local := carapace.Batch()
			remote := carapace.Batch()
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
				s := styles.Lore.Branch
				if entry.IsCurrent {
					desc += " (current)"
					s = styles.Lore.CurrentBranch
				}
				if entry.Archived {
					desc += " (archived)"
					s = style.Of(s, style.Dim)
				}
				a := carapace.ActionValuesDescribed(entry.Name, desc).Style(s)
				if entry.Location == "remote" {
					remote = append(remote, a.Tag("remote branches"))
				} else {
					local = append(local, a.Tag("local branches"))
				}
			}
			remote = append(remote, local...)
			return remote.ToA()
		})
	})
}
