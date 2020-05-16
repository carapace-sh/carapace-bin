package cmd

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	statusCmd.Flags().Bool("ahead-behind", false, "compute full ahead/behind values")
	statusCmd.Flags().BoolP("branch", "b", false, "show branch information")
	statusCmd.Flags().String("column", "", "list untracked files in columns")
	statusCmd.Flags().String("ignored", "", "show ignored files, optional modes: traditional, matching, no. (Default: traditional)")
	statusCmd.Flags().String("ignore-submodules", "", "ignore changes to submodules, optional when: all, dirty, untracked. (Default: all)")
	statusCmd.Flags().Bool("long", false, "show status in long format (default)")
	statusCmd.Flags().StringP("find-renames", "M", "", "detect renames, optionally set similarity index")
	statusCmd.Flags().Bool("no-renames", false, "do not detect renames")
	statusCmd.Flags().String("porcelain", "", "machine-readable output")
	statusCmd.Flags().Bool("show-stash", false, "show stash information")
	statusCmd.Flags().BoolP("short", "s", false, "show status concisely")
	statusCmd.Flags().StringP("untracked-files", "u", "", "show untracked files, optional modes: all, normal, no. (Default: all)")
	statusCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	statusCmd.Flags().BoolP("null", "z", false, "terminate entries with NUL")
	rootCmd.AddCommand(statusCmd)
}
