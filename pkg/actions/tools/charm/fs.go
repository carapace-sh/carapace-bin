package charm

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionFiles completes charm cloud files
//
//	charm:/dir/
//	charm:/one.txt
func ActionFiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if !strings.HasPrefix(c.Value, "charm:") {
			return carapace.ActionValues("charm:").NoSpace(':')
		}

		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			dir := filepath.Dir(c.Value)
			switch dir {
			case "/":
			case ".":
				dir = ""
			default:
				dir += "/"
			}

			return carapace.ActionExecCommand("charm", "fs", "ls", dir)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				vals := make([]string, 0)
				for _, line := range lines {
					if len(line) > 28 {
						switch line[1] {
						case 'd':
							vals = append(vals, line[29:]+"/")
						default:
							vals = append(vals, line[29:])
						}
					}
				}
				return carapace.ActionValues(vals...).StyleF(style.ForPathExt).NoSpace('/').Prefix(dir)
			})
		}).Prefix("charm:")
	})
}
