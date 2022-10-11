package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "Choose an option from a list of choices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseCmd).Standalone()

	chooseCmd.Flags().String("cursor", "", "Prefix to show on item that corresponds to the cursor position")
	chooseCmd.Flags().String("cursor-prefix", "", "Prefix to show on the cursor item")
	chooseCmd.Flags().String("cursor.foreground", "", "Foreground Color")
	chooseCmd.Flags().String("height", "", "Height of the list")
	chooseCmd.Flags().String("item.foreground", "", "Foreground Color")
	chooseCmd.Flags().String("limit", "", "Maximum number of options to pick")
	chooseCmd.Flags().Bool("no-limit", false, "Pick unlimited number of options (ignores limit)")
	chooseCmd.Flags().String("selected", "", "Options that should start as selected")
	chooseCmd.Flags().String("selected-prefix", "", "Prefix to show on selected items")
	chooseCmd.Flags().String("selected.foreground", "", "Foreground Color")
	chooseCmd.Flags().String("unselected-prefix", "", "Prefix to show on selected items")
	rootCmd.AddCommand(chooseCmd)

	carapace.Gen(chooseCmd).FlagCompletion(carapace.ActionMap{
		"cursor.foreground": gum.ActionColors(),
		"item.foreground":   gum.ActionColors(),
		"selected": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(c.Args...).Invoke(c).Filter(c.Parts).ToA()
		}),
		"selected.foreground": gum.ActionColors(),
	})
}
