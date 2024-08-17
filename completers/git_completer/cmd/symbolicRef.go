package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var symbolicRefCmd = &cobra.Command{
	Use:     "symbolic-ref",
	Short:   "Read, modify and delete symbolic refs",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(symbolicRefCmd).Standalone()

	symbolicRefCmd.Flags().BoolP("delete", "d", false, "delete the symbolicRef <name>")
	symbolicRefCmd.Flags().StringS("m", "m", "", "update the reflog for <name> with <reason>")
	symbolicRefCmd.Flags().Bool("no-recurse", false, "stop after dereferencing single level of symbolic ref")
	symbolicRefCmd.Flags().BoolP("quiet", "q", false, "do not issue an error message if the <name> is not a symbolicRef")
	symbolicRefCmd.Flags().Bool("recurse", false, "follow symbolicReferences")
	symbolicRefCmd.Flags().Bool("short", false, "try to shorten the symbolicRef")
	rootCmd.AddCommand(symbolicRefCmd)
}
