package pass

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace/pkg/style"
)

func passwordStore() (string, error) {
	if val, exists := os.LookupEnv("PASSWORD_STORE_DIR"); !exists {
		if home, err := os.UserHomeDir(); err != nil {
			return "", err
		} else {
			return home + "/.password-store", nil
		}
	} else {
		return val, nil
	}
}

// ActionDirectories completes password directories
//
//	dir1/
//	dir2/
func ActionDirectories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path, err := passwordStore(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return fs.ActionSubDirectories(path).StyleF(style.ForPathExt)
		}
	})
}

// ActionPasswords completes password names
//
//	mypass
//	dir/mypass2
func ActionPasswords() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path, err := passwordStore(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			names := make([]string, 0)
			filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
				if !info.IsDir() && strings.HasSuffix(info.Name(), ".gpg") {
					names = append(names, strings.TrimSuffix(strings.TrimPrefix(walkPath, path+"/"), ".gpg"))
				}
				return nil
			})
			return carapace.ActionValues(names...).
				MultiParts("/").
				StyleF(func(s string, sc style.Context) string {
					if strings.HasSuffix(s, "/") {
						return style.ForPath(fmt.Sprintf("%v/%v", path, s), c)
					}
					return style.ForPath(fmt.Sprintf("%v/%v.gpg", path, s), c)
				})
		}
	})
}
