package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugFileIndexCmd = &cobra.Command{
	Use:    "debug::file-index",
	Short:  "inspect or manipulate the file index",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugFileIndexCmd).Standalone()

	debugFileIndexCmd.Flags().Bool("docket", false, "dump docket file")
	debugFileIndexCmd.Flags().Bool("gc", false, "run garbage collection")
	debugFileIndexCmd.Flags().StringP("path", "p", "", "look up path")
	debugFileIndexCmd.Flags().StringP("template", "T", "", "template for --docket")
	debugFileIndexCmd.Flags().StringP("token", "t", "", "look up token")
	debugFileIndexCmd.Flags().Bool("tree", false, "dump tree file")
	debugFileIndexCmd.Flags().Bool("vacuum", false, "vacuum the tree file")
	rootCmd.AddCommand(debugFileIndexCmd)
}
