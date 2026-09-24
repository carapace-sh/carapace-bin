package styles

import "github.com/carapace-sh/carapace/pkg/style"

var Hg = struct {
	Bookmark string `description:"hg bookmarks"`
	Branch   string `description:"hg branches"`
	Stash    string `description:"hg shelves"`
	Tag      string `description:"hg tags"`
}{
	Bookmark: style.Blue,
	Branch:   style.Blue,
	Stash:    style.Green,
	Tag:      style.Yellow,
}

func init() {
	style.Register("hg", &Hg)
}
