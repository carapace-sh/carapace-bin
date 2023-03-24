package styles

import "github.com/rsteube/carapace/pkg/style"

var Git = struct {
	Branch     string `desc:"git branches"`
	Commit     string `desc:"git commits"`
	HeadCommit string `desc:"git HEAD~ commits"`
	Note       string `desc:"git notes"`
	Stash      string `desc:"git stashes"`
	Tag        string `desc:"git tags"`
}{
	Branch:     style.Blue,
	Commit:     style.Default,
	HeadCommit: style.Bold,
	Note:       style.Cyan,
	Stash:      style.Green,
	Tag:        style.Yellow,
}

func init() {
	style.Register("git", &Git)
}
