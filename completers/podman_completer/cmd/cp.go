package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp [options] [CONTAINER:]SRC_PATH [CONTAINER:]DEST_PATH",
	Short: "Copy files/folders between a container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()

	cpCmd.Flags().BoolP("archive", "a", false, "Chown copied files to the primary uid/gid of the destination container.")
	cpCmd.Flags().Bool("extract", false, "Deprecated...")
	cpCmd.Flags().Bool("overwrite", false, "Allow to overwrite directories with non-directories and vice versa")
	cpCmd.Flags().Bool("pause", false, "Deprecated")
	cpCmd.Flag("extract").Hidden = true
	cpCmd.Flag("pause").Hidden = true
	rootCmd.AddCommand(cpCmd)
}
