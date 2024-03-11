package ytdlp

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type dump struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Formats []struct {
		FormatID   string  `json:"format_id"`
		FormatNote string  `json:"format_note"`
		Ext        string  `json:"ext"`
		Protocol   string  `json:"protocol"`
		Acodec     string  `json:"acodec"`
		Vcodec     string  `json:"vcodec"`
		URL        string  `json:"url"`
		Width      int     `json:"width"`
		Height     int     `json:"height"`
		Fps        float64 `json:"fps"`
		Rows       int     `json:"rows,omitempty"`
		Columns    int     `json:"columns,omitempty"`
		Fragments  []struct {
			URL      string  `json:"url"`
			Duration float64 `json:"duration"`
		} `json:"fragments,omitempty"`
		Resolution  string  `json:"resolution"`
		AspectRatio float64 `json:"aspect_ratio"`
		HTTPHeaders struct {
			UserAgent      string `json:"User-Agent"`
			Accept         string `json:"Accept"`
			AcceptLanguage string `json:"Accept-Language"`
			SecFetchMode   string `json:"Sec-Fetch-Mode"`
		} `json:"http_headers"`
		AudioExt           string  `json:"audio_ext"`
		VideoExt           string  `json:"video_ext"`
		Format             string  `json:"format"`
		Asr                int     `json:"asr,omitempty"`
		Filesize           int     `json:"filesize,omitempty"`
		SourcePreference   int     `json:"source_preference,omitempty"`
		AudioChannels      int     `json:"audio_channels,omitempty"`
		Quality            float64 `json:"quality,omitempty"`
		HasDrm             bool    `json:"has_drm,omitempty"`
		Tbr                float64 `json:"tbr,omitempty"`
		Language           any     `json:"language,omitempty"`
		LanguagePreference int     `json:"language_preference,omitempty"`
		Preference         any     `json:"preference,omitempty"`
		DynamicRange       any     `json:"dynamic_range,omitempty"`
		Abr                float64 `json:"abr,omitempty"`
		Container          string  `json:"container,omitempty"`
		Vbr                float64 `json:"vbr,omitempty"`
		FilesizeApprox     int     `json:"filesize_approx,omitempty"`
	} `json:"formats"`
	Thumbnails []struct {
		URL        string `json:"url"`
		Preference int    `json:"preference"`
		ID         string `json:"id"`
		Height     int    `json:"height,omitempty"`
		Width      int    `json:"width,omitempty"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"thumbnails"`
	Subtitles map[string][]struct {
		Ext  string `json:"ext"`
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"subtitles"`
}

func actionDump(url string, f func(d dump) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("yt-dlp", "--dump-json", url)(func(output []byte) carapace.Action {
		var d dump
		if err := json.Unmarshal(output, &d); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(d)
	})
}
