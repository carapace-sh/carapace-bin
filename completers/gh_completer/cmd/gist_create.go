package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gist_createCmd = &cobra.Command{
	Use:   "create [<filename>... | -]",
	Short: "Create a new gist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	gist_createCmd.Flags().StringP("desc", "d", "", "A description for this gist")
	gist_createCmd.Flags().StringP("filename", "f", "", "Provide a filename to be used when reading from STDIN")
	gist_createCmd.Flags().BoolP("public", "p", false, "List the gist publicly (default: private)")
	gist_createCmd.Flags().BoolP("web", "w", false, "Open the web browser with created gist")
	gistCmd.AddCommand(gist_createCmd)

	carapace.Gen(gist_createCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
