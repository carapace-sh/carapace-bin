package lore

import (
	"bufio"
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

type authIdentity struct {
	UserID string `json:"userId"`
}

// ActionIdentities completes user IDs from `lore auth list --json`
//
//	user1
//	user2
func ActionIdentities(opts GlobalOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := append(opts.args(), "--json", "auth", "list")
		return carapace.ActionExecCommand("lore", args...)(func(output []byte) carapace.Action {
			var vals []string
			seen := make(map[string]bool)
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
				if event.TagName != "authIdentity" {
					continue
				}
				var identity authIdentity
				if err := json.Unmarshal(event.Data, &identity); err != nil {
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
			return carapace.ActionValuesDescribed(vals...).Style(styles.Lore.Identity)
		})
	}).Tag("identities")
}
