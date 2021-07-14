package action

import (
	"archive/zip"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionJarFileClasses(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if reader, err := zip.OpenReader(file); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			defer reader.Close()
			vals := make([]string, 0)
			for _, f := range reader.File {
				if strings.HasSuffix(f.Name, ".class") {
					name := strings.TrimPrefix(f.Name, "BOOT-INF/classes/")
					name = strings.Replace(name, "/", ".", -1)
					name = strings.TrimSuffix(name, ".class")
					vals = append(vals, name)
				}
			}
			return carapace.ActionValues(vals...)
		}
	})
}
