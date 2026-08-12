package lore

import (
	"bufio"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
)

type revisionHistoryEntry struct {
	Revision       string   `json:"revision"`
	RevisionNumber uint64   `json:"revisionNumber"`
	Parent         []string `json:"parent"`
}

// RevisionOpts contains global lore flags plus revision-specific filters.
type RevisionOpts struct {
	GlobalOpts

	// Branch shows revisions for the given branch
	Branch string

	// Revision starts listing from the specified revision
	Revision string
}

// ActionRevisions completes revision hashes from `lore revision history --json`
//
//	74fba5efe1201f0c1117b8655b25be878d2c9cf11cdd61e0128d0dc3155d642f (r1)
func ActionRevisions(opts RevisionOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append(opts.args(), "--json", "revision", "history")
		if opts.Branch != "" {
			args = append(args, "--branch", opts.Branch)
		}
		if opts.Revision != "" {
			args = append(args, "--revision", opts.Revision)
		}
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
				if event.TagName != "revisionHistoryEntry" {
					continue
				}
				var entry revisionHistoryEntry
				if err := json.Unmarshal(event.Data, &entry); err != nil {
					continue
				}
				vals = append(vals, entry.Revision, "r"+strconv.FormatUint(entry.RevisionNumber, 10))
			}
			if len(vals) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("revisions")
}
