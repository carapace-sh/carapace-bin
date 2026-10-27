package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:     "tag",
	Short:   "List tags, or create / delete one",
	GroupID: "branch",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()

	tagCmd.Flags().String("contains", "", "Only tags whose peeled commit contains this revision (`HEAD` when omitted)")
	tagCmd.Flags().BoolP("delete", "d", false, "Delete the named tag instead of creating it")
	tagCmd.Flags().Bool("dry-run", false, "Preview create or delete without changing refs")
	tagCmd.Flags().StringSlice("exclude-submodule", nil, "Submodule path globs excluded from the digest (repeatable)")
	tagCmd.Flags().StringP("file", "F", "", "Read the annotated tag message from a file (`-` reads stdin)")
	tagCmd.Flags().BoolP("force", "f", false, "Replace an existing tag with the same name")
	tagCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tagCmd.Flags().StringP("limit", "n", "", "Maximum number of rows to list")
	tagCmd.Flags().StringSliceP("list", "l", nil, "List tags matching `PATTERN` (repeatable). Omit `PATTERN` to list all tags")
	tagCmd.Flags().String("max-count", "", "Maximum number of rows to list")
	tagCmd.Flags().StringP("message", "m", "", "Annotated tag message (`-m` creates an annotated tag)")
	tagCmd.Flags().Bool("no-sign", false, "Do not sign even when `tag.gpgSign` is true")
	tagCmd.Flags().Bool("no-signature", false, "Skip signature verification (still checks `tree-sha256` when present)")
	tagCmd.Flags().String("object", "", "Object to tag (revision). Defaults to `HEAD`")
	tagCmd.Flags().StringSlice("points-at", nil, "Only tags pointing at this object (repeatable)")
	tagCmd.Flags().Bool("reverse", false, "Reverse the sort order")
	tagCmd.Flags().BoolP("sign", "s", false, "Sign the tag object (implies annotated). Honors `tag.gpgSign` when omitted")
	tagCmd.Flags().String("signing-key", "", "Signing key id or path (overrides `user.signingkey` for this tag)")
	tagCmd.Flags().String("sort", "version", "Sort listed tags by version, date, or name")
	tagCmd.Flags().Bool("tree-sha256", false, "Attach a tree-sha256 content digest to the tag (implies `--sign`)")
	tagCmd.Flags().BoolP("verify", "v", false, "Verify annotated tag signatures and optional `tree-sha256` claims")
	tagCmd.Flag("object").Hidden = true
	rootCmd.AddCommand(tagCmd)

	carapace.Gen(tagCmd).FlagCompletion(carapace.ActionMap{
		"contains":  git.ActionRefs(git.RefOption{}.Default()),
		"points-at": git.ActionRefs(git.RefOption{}.Default()),
		"sort":      carapace.ActionValues("date", "name", "version"),
	})

	carapace.Gen(tagCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{Tags: true}),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
