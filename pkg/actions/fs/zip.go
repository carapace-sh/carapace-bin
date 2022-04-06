package fs

import (
	"archive/zip"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionZipFileContents completes contents of given zip file
//   fileA
//   dirA/fileB
func ActionZipFileContents(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if reader, err := zip.OpenReader(file); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			defer reader.Close()
			vals := make([]string, len(reader.File))
			for index, f := range reader.File {
				vals[index] = f.Name
			}
			return carapace.ActionValues(vals...).Invoke(c).ToMultiPartsA("/").StyleF(style.ForPathExt)
		}
	})
}
