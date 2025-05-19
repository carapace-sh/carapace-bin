package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showFaceCmd = &cobra.Command{
	Use:   "+show-face",
	Short: "show what font face Ghostty will use to render a specific codepoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showFaceCmd).Standalone()

	showFaceCmd.Flags().String("cp", "", "Find the face for a single codepoint")
	showFaceCmd.Flags().Bool("help", false, "show help")
	showFaceCmd.Flags().String("presentation", "", "force searching for a specific presentation style")
	showFaceCmd.Flags().String("string", "", "Find the face for all of the codepoints in a string")
	showFaceCmd.Flags().String("style", "", "Search for a specific style")
	rootCmd.AddCommand(showFaceCmd)

	carapace.Gen(showFaceCmd).FlagCompletion(carapace.ActionMap{
		"presentation": carapace.ActionValues("text", "emoji"),
		"style":        carapace.ActionValues("regular", "bold", "italic", "bold_italic"),
	})
}
