package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imageCmd).Standalone()
	rootCmd.AddCommand(imageCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"image", "Image commands"},
		{"schema", "Schema commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		imageCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
