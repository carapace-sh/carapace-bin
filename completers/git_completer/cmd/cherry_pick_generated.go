package cmd

import (
	"github.com/spf13/cobra"
)

var cherry_pickCmd = &cobra.Command{
	Use:   "cherry-pick",
	Short: "Apply the changes introduced by some existing commits",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cherry_pickCmd.Flags().Bool("abort", false, "cancel revert or cherry-pick sequence")
	cherry_pickCmd.Flags().Bool("allow-empty-message", false, "allow commits with empty messages")
	cherry_pickCmd.Flags().Bool("allow-empty", false, "preserve initially empty commits")
	cherry_pickCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	cherry_pickCmd.Flags().Bool("continue", false, "resume revert or cherry-pick sequence")
	cherry_pickCmd.Flags().BoolP("edit", "e", false, "edit the commit message")
	cherry_pickCmd.Flags().Bool("ff", false, "allow fast-forward")
	cherry_pickCmd.Flags().Bool("keep-redundant-commits", false, "keep redundant, empty commits")
	cherry_pickCmd.Flags().BoolP("mainline", "m", false, "<parent-number>    select mainline parent")
	cherry_pickCmd.Flags().BoolP("no-commit", "n", false, "don't automatically commit")
	cherry_pickCmd.Flags().Bool("quit", false, "end revert or cherry-pick sequence")
	cherry_pickCmd.Flags().Bool("rerere-autoupdate", false, "update the index with reused conflict resolution if possible")
	cherry_pickCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	cherry_pickCmd.Flags().Bool("skip", false, "skip current commit and continue")
	cherry_pickCmd.Flags().BoolP("signoff", "s", false, "add Signed-off-by:")
	cherry_pickCmd.Flags().String("strategy", "", "merge strategy")
	cherry_pickCmd.Flags().BoolP("x", "x", false, "append commit name")
	cherry_pickCmd.Flags().BoolP("strategy-option", "X", false, "<option>    option for merge strategy")
	rootCmd.AddCommand(cherry_pickCmd)
}
