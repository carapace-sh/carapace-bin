package styles

import "github.com/rsteube/carapace/pkg/style"

var Git = struct {
	Branch     string `desc:"git branches"`
	Commit     string `desc:"git commits"`
	HeadCommit string `desc:"git HEAD~ commits"`
	Stash      string `desc:"git stashes"`
	Tag        string `desc:"git tags"`
}{
	Branch:     style.Blue,
	Commit:     style.Default,
	HeadCommit: style.Bold,
	Stash:      style.Green,
	Tag:        style.Yellow,
}

func init() {
	style.Register("git", &Git)
}
