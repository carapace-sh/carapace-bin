package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty Git repository or reinitialize an existing one",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	initCmd.Flags().Bool("bare", false, "create a bare repository")
	initCmd.Flags().BoolP("quiet", "q", false, "be quiet")
	initCmd.Flags().String("separate-git-dir", "", "separate git dir from working tree")
	initCmd.Flags().String("shared", "", "specify that the git repository is to be shared amongst several users")
	initCmd.Flags().String("template", "", "directory from which templates will be used")
	rootCmd.AddCommand(initCmd)
}
