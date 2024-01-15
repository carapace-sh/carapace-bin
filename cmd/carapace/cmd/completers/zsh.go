package completers

import (
	"os"
	"sort"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/execlog"
)

func ZshCompleters() []string {
	out, err := execlog.Command("zsh", "--no-rcs", "-c", "printf '%s\n' $fpath").Output()
	if err != nil {
		carapace.LOG.Println(err.Error())
		return []string{} // TODO
	}
	lines := strings.Split(string(out), "\n")

	unique := make(map[string]bool)
	for _, line := range lines {
		entries, err := os.ReadDir(line)
		if err != nil {
			carapace.LOG.Println(err.Error())
			continue // TODO
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasPrefix(entry.Name(), "_") {
				unique[strings.TrimPrefix(entry.Name(), "_")] = true
			}
		}
	}

	completers := make([]string, 0)
	for name := range unique {
		completers = append(completers, name)
	}
	sort.Strings(completers)
	return completers
}
