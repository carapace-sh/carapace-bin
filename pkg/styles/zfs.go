package styles

import "github.com/carapace-sh/carapace/pkg/style"

var Zfs = struct {
	Pool       string `description:"zfs pools"`
	Filesystem string `description:"zfs filesystems"`
	Volume     string `description:"zfs volumes"`
	Snapshot   string `description:"zfs snapshots"`
	Bookmark   string `description:"zfs bookmarks"`
}{
	Pool:       style.Bold,
	Filesystem: style.Blue,
	Volume:     style.Yellow,
	Snapshot:   style.Magenta,
	Bookmark:   style.Cyan,
}

func init() {
	style.Register("zfs", &Zfs)
}
