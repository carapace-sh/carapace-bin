package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var passwdCmd = &cobra.Command{
	Use:     "passwd",
	Short:   "Generation of hashed passwords",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(passwdCmd).Standalone()

	passwdCmd.Flags().BoolS("1", "1", false, "MD5-based password algorithm")
	passwdCmd.Flags().BoolS("5", "5", false, "SHA256-based password algorithm")
	passwdCmd.Flags().BoolS("6", "6", false, "SHA512-based password algorithm")
	passwdCmd.Flags().BoolS("aixmd5", "aixmd5", false, "AIX MD5-based password algorithm")
	passwdCmd.Flags().BoolS("apr1", "apr1", false, "MD5-based password algorithm, Apache variant")
	passwdCmd.Flags().StringS("in", "in", "", "Read passwords from file")
	passwdCmd.Flags().BoolS("noverify", "noverify", false, "Never verify when reading password from terminal")
	passwdCmd.Flags().BoolS("quiet", "quiet", false, "No warnings")
	passwdCmd.Flags().BoolS("reverse", "reverse", false, "Switch table columns")
	passwdCmd.Flags().StringS("salt", "salt", "", "Use provided salt")
	passwdCmd.Flags().BoolS("stdin", "stdin", false, "Read passwords from stdin")
	passwdCmd.Flags().BoolS("table", "table", false, "Format output as table")
	common.AddProviderFlags(passwdCmd)
	common.AddRandomStateFlags(passwdCmd)
	rootCmd.AddCommand(passwdCmd)

	carapace.Gen(passwdCmd).FlagCompletion(carapace.ActionMap{
		"in": carapace.ActionFiles(),
	})
}
