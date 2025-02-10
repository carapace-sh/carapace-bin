package common

import (
	"github.com/spf13/cobra"
)

func AddOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("bep_maximum_open_remote_upload_files", "", "Maximum number of open files allowed during BEP artifact upload")
	cmd.Flags().String("remote_download_outputs", "", "If set to 'minimal' doesn't download any remote build outputs to the local machine, except the ones required by local actions")
	cmd.Flags().String("remote_download_symlink_template", "", "Instead of downloading remote build outputs to the local machine, create symbolic links")
	cmd.Flags().StringSlice("repo_env", []string{}, "Specifies additional environment variables to be available only for repository rules")

	// TODO add flag completion
}
