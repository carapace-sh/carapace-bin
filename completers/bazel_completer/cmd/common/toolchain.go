package common

import (
	"github.com/spf13/cobra"
)

func AddToolchainFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("incompatible_enable_proto_toolchain_resolution", false, "If true, proto lang rules define toolchains from rules_proto, rules_java, rules_cc repositories")
	cmd.Flags().Bool("no-incompatible_enable_proto_toolchain_resolution", false, "If true, proto lang rules define toolchains from rules_proto, rules_java, rules_cc repositories")
}
