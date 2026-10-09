package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_pool_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create subnet pool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_pool_createCmd).Standalone()

	subnet_pool_createCmd.Flags().String("address-scope", "", "Set address scope associated with the subnet pool (name or ID), prefixes must be unique across address scopes")
	subnet_pool_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	subnet_pool_createCmd.Flags().Bool("default", false, "Set this as a default subnet pool")
	subnet_pool_createCmd.Flags().String("default-prefix-length", "", "Set subnet pool default prefix length")
	subnet_pool_createCmd.Flags().String("default-quota", "", "Set default per-project quota for this subnet pool as the number of IP addresses that can be allocated from the subnet pool")
	subnet_pool_createCmd.Flags().String("description", "", "Set subnet pool description")
	subnet_pool_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	subnet_pool_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	subnet_pool_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	subnet_pool_createCmd.Flags().String("max-prefix-length", "", "Set subnet pool maximum prefix length")
	subnet_pool_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	subnet_pool_createCmd.Flags().String("min-prefix-length", "", "Set subnet pool minimum prefix length")
	subnet_pool_createCmd.Flags().Bool("no-default", false, "Set this as a non-default subnet pool")
	subnet_pool_createCmd.Flags().Bool("no-share", false, "Set this subnet pool as not shared")
	subnet_pool_createCmd.Flags().Bool("no-tag", false, "No tags associated with the subnet pool")
	subnet_pool_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	subnet_pool_createCmd.Flags().String("pool-prefix", "", "Set subnet pool prefixes (in CIDR notation) (repeat option to set multiple prefixes)")
	subnet_pool_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	subnet_pool_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	subnet_pool_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	subnet_pool_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	subnet_pool_createCmd.Flags().Bool("share", false, "Set this subnet pool as shared")
	subnet_pool_createCmd.Flags().String("tag", "", "Tag to be added to the subnet pool (repeat option to set multiple tags)")
	subnet_pool_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	subnet_pool_createCmd.MarkFlagRequired("pool-prefix")
	subnet_poolCmd.AddCommand(subnet_pool_createCmd)
}
