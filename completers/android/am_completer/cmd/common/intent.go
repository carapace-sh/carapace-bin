package common

import (
	"github.com/carapace-sh/carapace"
	android "github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

// AddIntentFlags adds the intent specification flags of am
func AddIntentFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("action", "a", "", "intent `action`")
	cmd.Flags().StringP("category", "c", "", "intent `category` (repeatable)")
	cmd.Flags().StringP("component", "n", "", "component name as `package/class`")
	cmd.Flags().StringP("data", "d", "", "data `URI`")
	cmd.Flags().StringP("mime-type", "t", "", "MIME `type`")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"action":    android.ActionIntentActions(),
		"category":  android.ActionIntentCategories(),
		"component": android.ActionComponents(),
	})
}
