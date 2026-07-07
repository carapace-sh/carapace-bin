package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Print a substring",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subCmd).Standalone()

	subCmd.Flags().StringP("end", "e", "", "1-based end index")
	subCmd.Flags().StringP("length", "l", "", "length of substring")
	subCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	subCmd.Flags().StringP("start", "s", "", "1-based start index")

	rootCmd.AddCommand(subCmd)
}
