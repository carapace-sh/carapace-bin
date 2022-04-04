package styles

import "github.com/rsteube/carapace/pkg/style"

var Git = struct {
	Branch     string `desc:"git branches"`
	Commit     string `desc:"git commits"`
	HeadCommit string `desc:"git HEAD~ commits"`
	Stash      string `desc:"git stashes"`
	Tag        string `desc:"git tags"`
}{
	Branch:     "blue",
	Commit:     "default",
	HeadCommit: "bold",
	Stash:      "green",
	Tag:        "yellow",
}

func init() {
	style.Register("git", &Git)
}
