package ytdlp

import "github.com/carapace-sh/carapace"

// ActionFields completes fields
func ActionFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"id", "(string): Video identifier",
		"title", "(string): Video title",
		"fulltitle", "(string): Video title ignoring live timestamp and generic title",
		"ext", "(string): Video filename extension",
		"alt_title", "(string): A secondary title of the video",
		"description", "(string): The description of the video",
		"display_id", "(string): An alternative identifier for the video",
		"uploader", "(string): Full name of the video uploader",
		"license", "(string): License name the video is licensed under",
		"creator", "(string): The creator of the video",
		"timestamp", "(numeric): UNIX timestamp of the moment the video became available",
		"upload_date", "(string): Video upload date in UTC (YYYYMMDD)",
		"release_timestamp", "(numeric): UNIX timestamp of the moment the video was released",
		"release_date", "(string): The date (YYYYMMDD) when the video was released in UTC",
		"modified_timestamp", "(numeric): UNIX timestamp of the moment the video was last modified",
		"modified_date", "(string): The date (YYYYMMDD) when the video was last modified in UTC",
		"uploader_id", "(string): Nickname or id of the video uploader",
		"channel", "(string): Full name of the channel the video is uploaded on",
		"channel_id", "(string): Id of the channel",
		"channel_follower_count", "(numeric): Number of followers of the channel",
		"location", "(string): Physical location where the video was filmed",
		"duration", "(numeric): Length of the video in seconds",
		"duration_string", "(string): Length of the video (HH:mm:ss)",
		"view_count", "(numeric): How many users have watched the video on the platform",
		"concurrent_view_count", "(numeric): How many users are currently watching the video on the platform.",
		"like_count", "(numeric): Number of positive ratings of the video",
		"dislike_count", "(numeric): Number of negative ratings of the video",
		"repost_count", "(numeric): Number of reposts of the video",
		"average_rating", "(numeric): Average rating give by users, the scale used depends on the webpage",
		"comment_count", "(numeric): Number of comments on the video (For some extractors, comments are only downloaded at the end, and so this field cannot be used)",
		"age_limit", "(numeric): Age restriction for the video (years)",
		"live_status", "(string): One of \"not_live\", \"is_live\", \"is_upcoming\", \"was_live\", \"post_live\" (was live, but VOD is not yet processed)",
		"is_live", "(boolean): Whether this video is a live stream or a fixed-length video",
		"was_live", "(boolean): Whether this video was originally a live stream",
		"playable_in_embed", "(string): Whether this video is allowed to play in embedded players on other sites",
		"availability", "(string): Whether the video is \"private\", \"premium_only\", \"subscriber_only\", \"needs_auth\", \"unlisted\" or \"public\"",
		"start_time", "(numeric): Time in seconds where the reproduction should start, as specified in the URL",
		"end_time", "(numeric): Time in seconds where the reproduction should end, as specified in the URL",
		"extractor", "(string): Name of the extractor",
		"extractor_key", "(string): Key name of the extractor",
		"epoch", "(numeric): Unix epoch of when the information extraction was completed",
		"autonumber", "(numeric): Number that will be increased with each download, starting at --autonumber-start",
		"video_autonumber", "(numeric): Number that will be increased with each video",
		"n_entries", "(numeric): Total number of extracted items in the playlist",
		"playlist_id", "(string): Identifier of the playlist that contains the video",
		"playlist_title", "(string): Name of the playlist that contains the video",
		"playlist", "(string): playlist_id or playlist_title",
		"playlist_count", "(numeric): Total number of items in the playlist. May not be known if entire playlist is not extracted",
		"playlist_index", "(numeric): Index of the video in the playlist padded with leading zeros according the final index",
		"playlist_autonumber", "(numeric): Position of the video in the playlist download queue padded with leading zeros according to the total length of the playlist",
		"playlist_uploader", "(string): Full name of the playlist uploader",
		"playlist_uploader_id", "(string): Nickname or id of the playlist uploader",
		"webpage_url", "(string): A URL to the video webpage which if given to yt-dlp should allow to get the same result again",
		"webpage_url_basename", "(string): The basename of the webpage URL",
		"webpage_url_domain", "(string): The domain of the webpage URL",
		"original_url", "(string): The URL given by the user (or same as webpage_url for playlist entries)",
	).Tag("fields")
}

// ActionChapterFields completes chapter fields
func ActionChapterFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"chapter", "(string): Name or title of the chapter the video belongs to",
		"chapter_number", "(numeric): Number of the chapter the video belongs to",
		"chapter_id", "(string): Id of the chapter the video belongs to",
	).Tag("chapter fields")
}

// ActionEpisodeFields completes episode fields
func ActionEpisodeFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"series", "(string): Title of the series or programme the video episode belongs to",
		"season", "(string): Title of the season the video episode belongs to",
		"season_number", "(numeric): Number of the season the video episode belongs to",
		"season_id", "(string): Id of the season the video episode belongs to",
		"episode", "(string): Title of the video episode",
		"episode_number", "(numeric): Number of the video episode within a season",
		"episode_id", "(string): Id of the video episode",
	).Tag("episode fields")
}

// ActionTrackFields completes track fields
func ActionTrackFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"track", "(string): Title of the track",
		"track_number", "(numeric): Number of the track within an album or a disc",
		"track_id", "(string): Id of the track",
		"artist", "(string): Artist(s) of the track",
		"genre", "(string): Genre(s) of the track",
		"album", "(string): Title of the album the track belongs to",
		"album_type", "(string): Type of the album",
		"album_artist", "(string): List of all artists appeared on the album",
		"disc_number", "(numeric): Number of the disc or other physical medium the track belongs to",
		"release_year", "(numeric): Year (YYYY) when the album was released",
	).Tag("track fields")
}

// ActionSectionFields completes section fields
func ActionSectionFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"section_title", "(string): Title of the chapter",
		"section_number", "(numeric): Number of the chapter within the file",
		"section_start", "(numeric): Start time of the chapter in seconds",
		"section_end", "(numeric): End time of the chapter in seconds",
	).Tag("section fields")
}

// ActionPrintFields completes print fields
func ActionPrintFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"urls", "(string): The URLs of all requested formats, one in each line",
		"filename", "(string): Name of the video file. Note that the actual filename may differ",
		"formats_table", "(table): The video format table as printed by --list-formats",
		"thumbnails_table", "(table): The thumbnail format table as printed by --list-thumbnails",
		"subtitles_table", "(table): The subtitle format table as printed by --list-subs",
		"automatic_captions_table", "(table): The automatic subtitle format table as printed by --list-subs",
	).Tag("print fields")
}

// ActionSponsorblockFields completes sponsorblock fields
func ActionSponsorblockFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"start_time", "(numeric): Start time of the chapter in seconds",
		"end_time", "(numeric): End time of the chapter in seconds",
		"categories", "(list): The SponsorBlock categories the chapter belongs to",
		"category", "(string): The smallest SponsorBlock category the chapter belongs to",
		"category_names", "(list): Friendly names of the categories",
		"name", "(string): Friendly name of the smallest category",
		"type", "(string): The SponsorBlock action type of the chapter",
	).Tag("sponsorblock fields")
}

// ActionNumericMetaFields completes numeric meta fields
func ActionNumericMetaFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"filesize", "The number of bytes, if known in advance",
		"filesize_approx", "An estimate for the number of bytes",
		"width", "Width of the video, if known",
		"height", "Height of the video, if known",
		"aspect_ratio", "Aspect ratio of the video, if known",
		"tbr", "Average bitrate of audio and video in KBit/s",
		"abr", "Average audio bitrate in KBit/s",
		"vbr", "Average video bitrate in KBit/s",
		"asr", "Audio sampling rate in Hertz",
		"fps", "Frame rate",
		"audio_channels", "The number of audio channels",
		"stretched_ratio", "width:height of the video's pixels, if not square)",
	).Tag("numeric meta fields")
}

// ActionStringMetaFields completes string meta fields
func ActionStringMetaFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"url", "Video URL",
		"ext", "File extension",
		"acodec", "Name of the audio codec in use",
		"vcodec", "Name of the video codec in use",
		"container", "Name of the container format",
		"protocol", "The protocol that will be used for the actual download, lower-case (http, https, rtsp, rtmp, rtmpe, mms, f4m, ism, http_dash_segments, m3u8, or m3u8_native)",
		"language", "Language code",
		"dynamic_range", "The dynamic range of the video",
		"format_id", "A short description of the format",
		"format", "A human-readable description of the format",
		"format_note", "Additional info about the format",
		"resolution", "Textual description of width and height",
	).Tag("string meta fields")
}
