package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kinit",
	Short: "obtain and cache Kerberos ticket-granting ticket",
	Long:  "https://man.freebsd.org/cgi/man.cgi?kinit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "No addresses")
	rootCmd.Flags().BoolS("R", "R", false, "Renew")
	rootCmd.Flags().StringS("S", "S", "", "Server principal")
	rootCmd.Flags().BoolS("a", "a", false, "Extra addresses")
	rootCmd.Flags().StringS("c", "c", "", "Cache name")
	rootCmd.Flags().Bool("canonicalize", false, "Canonicalize")
	rootCmd.Flags().StringS("e", "e", "", "Encryption types")
	rootCmd.Flags().BoolS("f", "f", false, "Forwardable")
	rootCmd.Flags().String("fcache-version", "", "File cache version")
	rootCmd.Flags().BoolS("k", "k", false, "Use keytab")
	rootCmd.Flags().StringS("l", "l", "", "Lifetime")
	rootCmd.Flags().Bool("no-forwardable", false, "Do not forwardable")
	rootCmd.Flags().BoolS("p", "p", false, "Proxiable")
	rootCmd.Flags().String("password-file", "", "Password file")
	rootCmd.Flags().StringS("r", "r", "", "Renewable life")
	rootCmd.Flags().Bool("renewable", false, "Renewable")
	rootCmd.Flags().StringS("s", "s", "", "Start time")
	rootCmd.Flags().StringS("t", "t", "", "Keytab name")
	rootCmd.Flags().BoolS("v", "v", false, "Validate")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c":             carapace.ActionFiles(),
		"password-file": carapace.ActionFiles(),
		"t":             carapace.ActionFiles(),
	})
}
