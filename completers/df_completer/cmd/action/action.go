package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"source", "The source of the mount point, usually a device",
		"fstype", "File system type",
		"itotal", "Total number of inodes",
		"iused", "Number of used inodes",
		"iavail", "Number of available inodes",
		"ipcent", "Percentage of IUSED divided by ITOTAL",
		"size", "Total number of blocks",
		"used", "Number of used blocks",
		"avail", "Number of available blocks",
		"pcent", "Percentage of USED divided by SIZE",
		"file", "The file name if specified on the command line",
		"target", "The mount point",
	)
}

func ActionFilesystemTypes() carapace.Action {
	return carapace.ActionExecCommand("df", "--output=fstype")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1:]...)
	})
}
