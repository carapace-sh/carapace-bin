package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "avconvert",
	Short: "AV Foundation media conversion",
	Long:  "https://keith.github.io/xcode-manpages/avconvert.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("audioCodec", "", "Audio codec")
	rootCmd.Flags().String("dest", "", "Destination file")
	rootCmd.Flags().String("fileType", "", "File type")
	rootCmd.Flags().Bool("help", false, "Show help")
	rootCmd.Flags().String("preset", "", "Preset")
	rootCmd.Flags().Bool("progress", false, "Show progress")
	rootCmd.Flags().Bool("showMetadata", false, "Show metadata")
	rootCmd.Flags().String("source", "", "Source file")
	rootCmd.Flags().Bool("version", false, "Show version")
	rootCmd.Flags().String("videoCodec", "", "Video codec")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"audioCodec": carapace.ActionValues("aac", "alac", "ilbc", "ima4", "MAC3", "MAC6", "mp1", "mp2", "mp3", "msbc", "pcm", "qcelp", "qclp"),
		"dest":       carapace.ActionFiles(),
		"fileType":   carapace.ActionValues("aac", "aiff", "caf", "m4a", "m4v", "mov", "mp3", "mp4", "wav", "sd2"),
		"preset":     carapace.ActionValues("AVPresetPreset640x480", "AVPresetPreset960x540", "AVPresetPreset1280x720", "AVPresetPreset1920x1080", "AVPresetPreset3840x2160"),
		"source":     carapace.ActionFiles(),
		"videoCodec": carapace.ActionValues("h264", "hevc", "prores422", "prores4444", "jpeg"),
	})
}
