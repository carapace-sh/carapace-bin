package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_toolStubCmd = &cobra.Command{
	Use:   "tool-stub",
	Short: "Generate or update a tool stub",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_toolStubCmd).Standalone()

	generate_toolStubCmd.Flags().StringP("bin", "b", "", "Binary path within the extracted archive")
	generate_toolStubCmd.Flags().Bool("bootstrap", false, "Make the stub install mise when it is missing")
	generate_toolStubCmd.Flags().String("bootstrap-version", "", "mise version the --bootstrap script installs; defaults to the latest release")
	generate_toolStubCmd.Flags().String("checksum-algorithm", "blake3", "Checksum algorithm for downloaded artifacts")
	generate_toolStubCmd.Flags().Bool("fetch", false, "Fetch checksums and sizes for an existing tool stub file")
	generate_toolStubCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_toolStubCmd.Flags().String("http", "http", "Has no effect; kept so existing scripts that pass it keep working")
	generate_toolStubCmd.Flags().Bool("lock", false, "Resolve and record lock data (exact version, platform URLs and checksums) for an existing stub")
	generate_toolStubCmd.Flags().StringSlice("platform-bin", nil, "Binary path inside the archive for one platform, such as `linux-x64:bin/tool`; repeatable")
	generate_toolStubCmd.Flags().StringSlice("platform-url", nil, "Download URL for one platform, as `PLATFORM:URL` or a bare URL; repeatable")
	generate_toolStubCmd.Flags().Bool("skip-download", false, "Do not download artifacts, so checksums, sizes, and the detected binary path are left out")
	generate_toolStubCmd.Flags().StringP("url", "u", "", "URL to download the tool from")
	generate_toolStubCmd.Flags().String("version", "latest", "Version of the tool")
	generate_toolStubCmd.Flag("http").Hidden = true
	generateCmd.AddCommand(generate_toolStubCmd)

	carapace.Gen(generate_toolStubCmd).FlagCompletion(carapace.ActionMap{
		"checksum-algorithm": carapace.ActionValues("blake3", "sha256"),
	})

	carapace.Gen(generate_toolStubCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
