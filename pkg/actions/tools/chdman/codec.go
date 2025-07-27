package chdman

import "github.com/carapace-sh/carapace"

// ActionCodecs completes codecs
//
//	avhu (Huffman coding for audio-visual data)
//	cdfl (Free Lossless Audio Codec/zlib deflate for CD-ROM data)
func ActionCodecs() carapace.Action {
	return carapace.ActionValuesDescribed(
		"avhu", "Huffman coding for audio-visual data",
		"cdfl", "Free Lossless Audio Codec/zlib deflate for CD-ROM data",
		"cdlz", "Lempel-Ziv-Markov chain algorithm/zlib deflate for CD-ROM data",
		"cdzl", "zlib deflate for CD-ROM data",
		"cdzs", "Zstandard for CD-ROM data",
		"flac", "Free Lossless Audio Codec",
		"huff", "Huffman coding",
		"lzma", "Lempel-Ziv-Markov chain algorithm",
		"zlib", "zlib deflate",
		"zstd", "Zstandard",
	).Tag("codecs")
}
