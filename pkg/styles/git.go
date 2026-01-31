package styles

import "github.com/carapace-sh/carapace/pkg/style"

var Git = struct {
	Branch string `description:"git branches"`
	Commit string `description:"git commits"`
	Head   string `description:"git heads"`
	Note   string `description:"git notes"`
	Stash  string `description:"git stashes"`
	Tag    string `description:"git tags"`
}{
	Branch: style.Blue,
	Commit: style.Default,
	Head:   style.Bold,
	Note:   style.Cyan,
	Stash:  style.Green,
	Tag:    style.Yellow,
}

func init() {
	style.Register("git", &Git)
}
