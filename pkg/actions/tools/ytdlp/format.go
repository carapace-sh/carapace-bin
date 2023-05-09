package ytdlp

import (
	"fmt"
	"math"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionOutputFormats completes output formats
//
//	avi
//	flv
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues(
		"avi",
		"flv",
		"mkv",
		"mov",
		"mp4",
		"webm",
	).StyleF(func(s string, sc style.Context) string {
		return style.ForPathExt("."+s, sc)
	}).Tag("output formats")
}

// ActionSubtitleFormats completes subtitle formats
//
//	ass
//	lrc
func ActionSubtitleFormats() carapace.Action {
	return carapace.ActionValues(
		"ass",
		"lrc",
		"srt",
		"vtt",
	).StyleF(func(s string, sc style.Context) string {
		return style.ForPathExt("."+s, sc)
	}).Tag("subtitle formats")
}

// ActionThumbnailFormats completes thumbnail formats
//
//	jpg
//	png
func ActionThumbnailFormats() carapace.Action {
	return carapace.ActionValues(
		"jpg",
		"png",
		"webp",
	).StyleF(func(s string, sc style.Context) string {
		return style.ForPathExt("."+s, sc)
	}).Tag("thumbnail formats")
}

// ActionVideoFormats completes video formats
//
//	avi
//	flv
func ActionVideoFormats() carapace.Action {
	return carapace.ActionValues(
		"avi",
		"flv",
		"gif",
		"mkv",
		"mov",
		"mp4",
		"webm",
		"aac",
		"aiff",
		"alac",
		"flac",
		"m4a",
		"mka",
		"mp3",
		"ogg",
		"opus",
		"vorbis",
		"wav",
	).StyleF(func(s string, sc style.Context) string {
		return style.ForPathExt("."+s, sc)
	}).Tag("video formats")
}

// ActionFormats completes formats
//
//	133 (426x240 - avc1 - none - 0)
//	134 (640x360 - avc1 - none - 0)
func ActionFormats(url string) carapace.Action {
	return actionDump(url, func(d dump) carapace.Action {
		vals := make([]string, 0)
		for _, format := range d.Formats {
			description := fmt.Sprintf("%v - %v - %v - %v",
				format.Resolution,
				strings.Split(format.Vcodec, ".")[0],
				strings.Split(format.Acodec, ".")[0],
				math.Round(format.Abr),
			)

			s := style.Default
			switch {
			case format.Vcodec == "none":
				s = style.Yellow
			}

			vals = append(vals, format.FormatID, description, s) // TODO enrich description
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("formats")
}

// ActionSubtitles completes subtitles
//
//	en (English)
//	fr (French)
func ActionSubtitles(url string) carapace.Action {
	return actionDump(url, func(d dump) carapace.Action {
		vals := make([]string, 0)
		for id, s := range d.Subtitles {
			vals = append(vals, id, s[0].Name)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("subtitles")
}
