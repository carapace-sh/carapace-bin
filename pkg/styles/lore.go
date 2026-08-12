package styles

import "github.com/carapace-sh/carapace/pkg/style"

var Lore = struct {
	Branch       string `description:"lore branches"`
	CurrentBranch string `description:"lore current branch"`
	Revision     string `description:"lore revisions"`
	Identity     string `description:"lore identities"`
}{
	Branch:        style.Blue,
	CurrentBranch: style.Bold,
	Revision:      style.Default,
	Identity:      style.Yellow,
}

func init() {
	style.Register("lore", &Lore)
}
