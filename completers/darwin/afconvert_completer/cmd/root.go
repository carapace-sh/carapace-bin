package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "afconvert",
	Short: "Audio File Convert",
	Long:  "https://man.freebsd.org/cgi/man.cgi?afconvert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print help text")
	rootCmd.Flags().Bool("help-formats", false, "Print list of supported file/data formats")
	rootCmd.Flags().BoolP("tag", "t", false, "Tag")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print progress verbosely")

	rootCmd.Flags().StringP("bitrate", "b", "", "Total bit rate in bps")
	rootCmd.Flags().StringP("channellayout", "l", "", "Channel layout tag")
	rootCmd.Flags().StringP("channelmap", "m", "", "Channel map")
	rootCmd.Flags().StringP("channels", "c", "", "Number of channels")
	rootCmd.Flags().StringP("data", "d", "", "Data format")
	rootCmd.Flags().StringP("file", "f", "", "File format")
	rootCmd.Flags().StringS("o", "o", "", "Output file")
	rootCmd.Flags().StringP("quality", "q", "", "Codec quality (0-127)")
	rootCmd.Flags().StringP("src-quality", "r", "", "Sample rate converter quality (0-127)")
	rootCmd.Flags().StringP("strategy", "s", "", "Bitrate allocation strategy (0=CBR, 1=ABR, 2=VBR_constrained, 3=VBR)")
	rootCmd.Flags().StringP("userproperty", "u", "", "Set an AudioConverter property")

	rootCmd.Flags().Bool("generate-hash", false, "Generate SHA-1 hash of input audio data")
	rootCmd.Flags().Bool("leaks", false, "Run leaks at the end")
	rootCmd.Flags().Bool("mix", false, "Enable channel downmixing")
	rootCmd.Flags().Bool("no-filler", false, "Don't page-align audio data")
	rootCmd.Flags().Bool("profile", false, "Collect and print performance information")
	rootCmd.Flags().Bool("soundcheck-generate", false, "Analyze audio, add SoundCheck data")
	rootCmd.Flags().Bool("soundcheck-read", false, "Read SoundCheck data from source file")

	rootCmd.Flags().String("codec-manuf", "", "Codec manufacturer code")
	rootCmd.Flags().String("dither", "", "Dither algorithm (1-2)")
	rootCmd.Flags().String("gapless-after", "", "File coming after this one in a gapless album")
	rootCmd.Flags().String("gapless-before", "", "File coming before this one in a gapless album")
	rootCmd.Flags().String("media-kind", "", "Media kind string")
	rootCmd.Flags().String("offset", "", "Starting offset in frames")
	rootCmd.Flags().String("prime-method", "", "Decode priming method")
	rootCmd.Flags().String("prime-override", "", "Override packet table info")
	rootCmd.Flags().String("read-track", "", "Track index to read")
	rootCmd.Flags().String("src-complexity", "", "Sample rate converter complexity")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"channels":       carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8"),
		"dither":         carapace.ActionValues("1", "2"),
		"gapless-after":  carapace.ActionFiles(),
		"gapless-before": carapace.ActionFiles(),
		"media-kind":     carapace.ActionValues("Music", "Podcast", "iTunes U", "Audiobook", "Voice Memo", "Movie", "TV Show", "Music Video", "Home Video", "Audio Ad", "Video Ad"),
		"o":              carapace.ActionFiles(),
		"strategy":       carapace.ActionValuesDescribed("0", "CBR", "1", "ABR", "2", "VBR_constrained", "3", "VBR"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
