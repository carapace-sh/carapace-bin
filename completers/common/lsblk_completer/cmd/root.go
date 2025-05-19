package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lsblk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsblk",
	Short: "list block devices",
	Long:  "https://man7.org/linux/man-pages/man8/lsblk.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print all devices")
	rootCmd.Flags().BoolP("ascii", "i", false, "use ascii characters only")
	rootCmd.Flags().BoolP("bytes", "b", false, "print SIZE in bytes rather than in human readable format")
	rootCmd.Flags().String("ct", "", "define a custom counter")
	rootCmd.Flags().String("ct-filter", "", "restrict the next counter")
	rootCmd.Flags().StringP("dedup", "E", "", "de-duplicate output by <column>")
	rootCmd.Flags().BoolP("discard", "D", false, "print discard capabilities")
	rootCmd.Flags().StringP("exclude", "e", "", "exclude devices by major number (default: RAM disks)")
	rootCmd.Flags().StringP("filter", "Q", "", "print only lines matching the expression")
	rootCmd.Flags().BoolP("fs", "f", false, "output info about filesystems")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().String("highlight", "", "colorize lines matching the expression")
	rootCmd.Flags().String("hyperlink", "", "Print mountpoint paths as terminal hyperlinks")
	rootCmd.Flags().StringP("include", "I", "", "show only devices with specified major numbers")
	rootCmd.Flags().BoolP("inverse", "s", false, "inverse dependencies")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("list", "l", false, "use list format output")
	rootCmd.Flags().BoolP("list-columns", "H", false, "list the available columns")
	rootCmd.Flags().BoolP("merge", "M", false, "group parents of sub-trees (usable for RAIDs, Multi-path)")
	rootCmd.Flags().BoolP("nodeps", "d", false, "don't print slaves or holders")
	rootCmd.Flags().BoolP("noempty", "A", false, "don't print empty devices")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().BoolP("nvme", "N", false, "output info about NVMe devices")
	rootCmd.Flags().StringP("output", "o", "", "output columns")
	rootCmd.Flags().BoolP("output-all", "O", false, "output all columns")
	rootCmd.Flags().BoolP("pairs", "P", false, "use key=\"value\" output format")
	rootCmd.Flags().BoolP("paths", "p", false, "print complete device path")
	rootCmd.Flags().BoolP("perms", "m", false, "output info about permissions")
	rootCmd.Flags().String("properties-by", "", "methods used to gather data")
	rootCmd.Flags().BoolP("raw", "r", false, "use raw output format")
	rootCmd.Flags().BoolP("scsi", "S", false, "output info about SCSI devices")
	rootCmd.Flags().BoolP("shell", "y", false, "use column names that can be used as shell variables")
	rootCmd.Flags().StringP("sort", "x", "", "sort output by <column>")
	rootCmd.Flags().String("sysroot", "", "use specified directory as system root")
	rootCmd.Flags().BoolP("topology", "t", false, "output info about topology")
	rootCmd.Flags().StringP("tree", "T", "", "use tree format output")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().BoolP("virtio", "v", false, "output info about virtio devices")
	rootCmd.Flags().StringP("width", "w", "", "specifies output width as number of characters")
	rootCmd.Flags().BoolP("zoned", "z", false, "print zone related information")

	rootCmd.Flag("hyperlink").NoOptDefVal = " "
	rootCmd.Flag("tree").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dedup":         action.ActionColumns().UniqueList(","),
		"hyperlink":     carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"output":        action.ActionColumns().UniqueList(","),
		"properties-by": carapace.ActionValues("udev", "blkid", "file", "none").UniqueList(","),
		"sort":          action.ActionColumns(),
		"sysroot":       carapace.ActionDirectories(),
		"tree":          action.ActionColumns(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		fs.ActionBlockDevices(),
	)
}
