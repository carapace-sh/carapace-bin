package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Create a new user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createUserCmd).Standalone()

	createUserCmd.Flags().Bool("ephemeral", false, "create an ephemeral user")
	createUserCmd.Flags().Bool("for-testing", false, "create a user for testing")
	createUserCmd.Flags().Bool("guest", false, "shorthand for --user-type android.os.usertype.full.GUEST")
	createUserCmd.Flags().Bool("managed", false, "shorthand for --user-type android.os.usertype.profile.MANAGED")
	createUserCmd.Flags().Bool("pre-create-only", false, "pre-create the user without initializing it")
	createUserCmd.Flags().String("profileOf", "", "create a profile for the given `USER_ID`")
	createUserCmd.Flags().Bool("restricted", false, "shorthand for --user-type android.os.usertype.full.RESTRICTED")
	createUserCmd.Flags().String("user-type", "", "the name of a `USER_TYPE`")

	rootCmd.AddCommand(createUserCmd)

	carapace.Gen(createUserCmd).FlagCompletion(carapace.ActionMap{
		"user-type": carapace.ActionValues(
			"android.os.usertype.full.SYSTEM",
			"android.os.usertype.full.SECONDARY",
			"android.os.usertype.full.GUEST",
			"android.os.usertype.profile.MANAGED",
			"android.os.usertype.system.HEADLESS",
		),
	})

	carapace.Gen(createUserCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
