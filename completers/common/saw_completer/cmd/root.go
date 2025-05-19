package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "saw <command>",
	Short: "A fast, multipurpose tool for AWS CloudWatch Logs",
	Long:  "https://github.com/TylerBrock/saw",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("endpoint-url", "", "override default endpoint URL")
	rootCmd.PersistentFlags().String("profile", "", "override default AWS profile")
	rootCmd.PersistentFlags().String("region", "", "override profile AWS region")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"profile": aws.ActionProfiles(),
		"region":  aws.ActionRegions(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := rootCmd.Flag("profile"); f.Changed {
				c.Setenv("AWS_PROFILE", f.Value.String())
			}
			if f := rootCmd.Flag("region"); f.Changed {
				c.Setenv("AWS_REGION", f.Value.String())
			}
			return action.Invoke(c).ToA()
		})
	})
}
