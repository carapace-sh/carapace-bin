package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugrevlogindexCmd = &cobra.Command{
	Use:    "debugrevlogindex",
	Short:  "dump the contents of a revlog index",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugrevlogindexCmd).Standalone()

	debugrevlogindexCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugrevlogindexCmd.Flags().String("dir", "", "open directory manifest")
	debugrevlogindexCmd.Flags().StringP("format", "f", "", "revlog format")
	debugrevlogindexCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	rootCmd.AddCommand(debugrevlogindexCmd)
}
