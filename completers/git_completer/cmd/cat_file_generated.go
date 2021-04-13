package cmd

import (
	"github.com/spf13/cobra"
)

var cat_fileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Provide content or type and size information for repository objects",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cat_fileCmd.Flags().Bool("allow-unknown-type", false, "allow -s and -t to work with broken/corrupt objects")
	cat_fileCmd.Flags().String("batch", "", "show info and content of objects fed from the standard input")
	cat_fileCmd.Flags().Bool("batch-all-objects", false, "show all objects with --batch or --batch-check")
	cat_fileCmd.Flags().String("batch-check", "", "show info about objects fed from the standard input")
	cat_fileCmd.Flags().Bool("buffer", false, "buffer --batch output")
	cat_fileCmd.Flags().BoolS("e", "e", false, "exit with zero when there's no error")
	cat_fileCmd.Flags().Bool("filters", false, "for blob objects, run filters on object's content")
	cat_fileCmd.Flags().Bool("follow-symlinks", false, "follow in-tree symlinks (used with --batch or --batch-check)")
	cat_fileCmd.Flags().BoolS("p", "p", false, "pretty-print object's content")
	cat_fileCmd.Flags().String("path", "", "use a specific path for --textconv/--filters")
	cat_fileCmd.Flags().BoolS("s", "s", false, "show object size")
	cat_fileCmd.Flags().BoolS("t", "t", false, "show object type")
	cat_fileCmd.Flags().Bool("textconv", false, "for blob objects, run textconv on object's content")
	cat_fileCmd.Flags().Bool("unordered", false, "do not order --batch-all-objects output")
	rootCmd.AddCommand(cat_fileCmd)
}
