package charm

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDatabases completes databases
//
//	db1
//	db2
func ActionDatabases() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		out, err := c.Command("charm", "where").Output()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		dir := strings.SplitN(string(out), "\n", 2)[0]
		entries, err := os.ReadDir(dir + "/kv/")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if entry.IsDir() && entry.Name() != "charm.sh.kv.user.default" {
				vals = append(vals, entry.Name())
			}
		}
		return carapace.ActionValues(vals...)
	})
}
