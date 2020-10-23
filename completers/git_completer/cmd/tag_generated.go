package cmd

import (
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create, list, delete or verify a tag object signed with GPG",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	tagCmd.Flags().BoolP("annotate", "a", false, "annotated tag, needs a message")
	tagCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	tagCmd.Flags().String("color", "", "respect format colors")
	tagCmd.Flags().String("column", "", "show tag list in columns")
	tagCmd.Flags().String("contains", "", "print only tags that contain the commit")
	tagCmd.Flags().Bool("create-reflog", false, "create a reflog")
	tagCmd.Flags().BoolP("delete", "d", false, "delete tags")
	tagCmd.Flags().BoolP("edit", "e", false, "force edit of tag message")
	tagCmd.Flags().BoolP("file", "F", false, "<file>     read message from file")
	tagCmd.Flags().BoolP("force", "f", false, "replace the tag if exists")
	tagCmd.Flags().String("format", "", "format to use for the output")
	tagCmd.Flags().BoolP("ignore-case", "i", false, "sorting and filtering are case insensitive")
	tagCmd.Flags().BoolP("list", "l", false, "list tag names")
	tagCmd.Flags().String("merged", "", "print only tags that are merged")
	tagCmd.Flags().BoolP("message", "m", false, "<message>    tag message")
	tagCmd.Flags().StringS("n", "n", "", "print <n> lines of each tag message")
	tagCmd.Flags().String("no-contains", "", "print only tags that don't contain the commit")
	tagCmd.Flags().String("no-merged", "", "print only tags that are not merged")
	tagCmd.Flags().String("points-at", "", "print only tags of the object")
	tagCmd.Flags().String("sort", "", "field name to sort on")
	tagCmd.Flags().BoolP("sign", "s", false, "annotated and GPG-signed tag")
	tagCmd.Flags().BoolP("local-user", "u", false, "<key-id>    use another key to sign the tag")
	tagCmd.Flags().BoolP("verify", "v", false, "verify tags")
	rootCmd.AddCommand(tagCmd)
}
