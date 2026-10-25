package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"avail", "Number of available blocks",
		"file", "The file name if specified on the command line",
		"fstype", "File system type",
		"iavail", "Number of available inodes",
		"ipcent", "Percentage of IUSED divided by ITOTAL",
		"itotal", "Total number of inodes",
		"iused", "Number of used inodes",
		"pcent", "Percentage of USED divided by SIZE",
		"size", "Total number of blocks",
		"source", "The source of the mount point, usually a device",
		"target", "The mount point",
		"used", "Number of used blocks",
	)
}

func ActionFilesystemTypes() carapace.Action {
	return carapace.ActionExecCommand("df", "--output=fstype")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1:]...)
	})
}
