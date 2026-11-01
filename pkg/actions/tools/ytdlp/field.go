package ytdlp

import "github.com/carapace-sh/carapace"

// ActionFields completes fields
//
//	id
//	title
func ActionFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"age_limit", "(numeric): Age restriction for the video (years)",
		"alt_title", "(string): A secondary title of the video",
		"autonumber", "(numeric): Number that will be increased with each download, starting at --autonumber-start",
		"availability", "(string): Whether the video is \"private\", \"premium_only\", \"subscriber_only\", \"needs_auth\", \"unlisted\" or \"public\"",
		"average_rating", "(numeric): Average rating give by users, the scale used depends on the webpage",
		"channel", "(string): Full name of the channel the video is uploaded on",
		"channel_follower_count", "(numeric): Number of followers of the channel",
		"channel_id", "(string): Id of the channel",
		"comment_count", "(numeric): Number of comments on the video (For some extractors, comments are only downloaded at the end, and so this field cannot be used)",
		"concurrent_view_count", "(numeric): How many users are currently watching the video on the platform.",
		"creator", "(string): The creator of the video",
		"description", "(string): The description of the video",
		"dislike_count", "(numeric): Number of negative ratings of the video",
		"display_id", "(string): An alternative identifier for the video",
		"duration", "(numeric): Length of the video in seconds",
		"duration_string", "(string): Length of the video (HH:mm:ss)",
		"end_time", "(numeric): Time in seconds where the reproduction should end, as specified in the URL",
		"epoch", "(numeric): Unix epoch of when the information extraction was completed",
		"ext", "(string): Video filename extension",
		"extractor", "(string): Name of the extractor",
		"extractor_key", "(string): Key name of the extractor",
		"fulltitle", "(string): Video title ignoring live timestamp and generic title",
		"id", "(string): Video identifier",
		"is_live", "(boolean): Whether this video is a live stream or a fixed-length video",
		"license", "(string): License name the video is licensed under",
		"like_count", "(numeric): Number of positive ratings of the video",
		"live_status", "(string): One of \"not_live\", \"is_live\", \"is_upcoming\", \"was_live\", \"post_live\" (was live, but VOD is not yet processed)",
		"location", "(string): Physical location where the video was filmed",
		"modified_date", "(string): The date (YYYYMMDD) when the video was last modified in UTC",
		"modified_timestamp", "(numeric): UNIX timestamp of the moment the video was last modified",
		"n_entries", "(numeric): Total number of extracted items in the playlist",
		"original_url", "(string): The URL given by the user (or same as webpage_url for playlist entries)",
		"playable_in_embed", "(string): Whether this video is allowed to play in embedded players on other sites",
		"playlist", "(string): playlist_id or playlist_title",
		"playlist_autonumber", "(numeric): Position of the video in the playlist download queue padded with leading zeros according to the total length of the playlist",
		"playlist_count", "(numeric): Total number of items in the playlist. May not be known if entire playlist is not extracted",
		"playlist_id", "(string): Identifier of the playlist that contains the video",
		"playlist_index", "(numeric): Index of the video in the playlist padded with leading zeros according the final index",
		"playlist_title", "(string): Name of the playlist that contains the video",
		"playlist_uploader", "(string): Full name of the playlist uploader",
		"playlist_uploader_id", "(string): Nickname or id of the playlist uploader",
		"release_date", "(string): The date (YYYYMMDD) when the video was released in UTC",
		"release_timestamp", "(numeric): UNIX timestamp of the moment the video was released",
		"repost_count", "(numeric): Number of reposts of the video",
		"start_time", "(numeric): Time in seconds where the reproduction should start, as specified in the URL",
		"timestamp", "(numeric): UNIX timestamp of the moment the video became available",
		"title", "(string): Video title",
		"upload_date", "(string): Video upload date in UTC (YYYYMMDD)",
		"uploader", "(string): Full name of the video uploader",
		"uploader_id", "(string): Nickname or id of the video uploader",
		"video_autonumber", "(numeric): Number that will be increased with each video",
		"view_count", "(numeric): How many users have watched the video on the platform",
		"was_live", "(boolean): Whether this video was originally a live stream",
		"webpage_url", "(string): A URL to the video webpage which if given to yt-dlp should allow to get the same result again",
		"webpage_url_basename", "(string): The basename of the webpage URL",
		"webpage_url_domain", "(string): The domain of the webpage URL",
	).Tag("fields")
}

// ActionChapterFields completes chapter fields
//
//	title
//	start_time
func ActionChapterFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"chapter", "(string): Name or title of the chapter the video belongs to",
		"chapter_id", "(string): Id of the chapter the video belongs to",
		"chapter_number", "(numeric): Number of the chapter the video belongs to",
	).Tag("chapter fields")
}

// ActionEpisodeFields completes episode fields
//
//	title
//	id
func ActionEpisodeFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"episode", "(string): Title of the video episode",
		"episode_id", "(string): Id of the video episode",
		"episode_number", "(numeric): Number of the video episode within a season",
		"season", "(string): Title of the season the video episode belongs to",
		"season_id", "(string): Id of the season the video episode belongs to",
		"season_number", "(numeric): Number of the season the video episode belongs to",
		"series", "(string): Title of the series or programme the video episode belongs to",
	).Tag("episode fields")
}

// ActionTrackFields completes track fields
//
//	title
//	number
func ActionTrackFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"album", "(string): Title of the album the track belongs to",
		"album_artist", "(string): List of all artists appeared on the album",
		"album_type", "(string): Type of the album",
		"artist", "(string): Artist(s) of the track",
		"disc_number", "(numeric): Number of the disc or other physical medium the track belongs to",
		"genre", "(string): Genre(s) of the track",
		"release_year", "(numeric): Year (YYYY) when the album was released",
		"track", "(string): Title of the track",
		"track_id", "(string): Id of the track",
		"track_number", "(numeric): Number of the track within an album or a disc",
	).Tag("track fields")
}

// ActionSectionFields completes section fields
//
//	title
//	category
func ActionSectionFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"section_end", "(numeric): End time of the chapter in seconds",
		"section_number", "(numeric): Number of the chapter within the file",
		"section_start", "(numeric): Start time of the chapter in seconds",
		"section_title", "(string): Title of the chapter",
	).Tag("section fields")
}

// ActionPrintFields completes print fields
//
//	id
//	title
func ActionPrintFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"automatic_captions_table", "(table): The automatic subtitle format table as printed by --list-subs",
		"filename", "(string): Name of the video file. Note that the actual filename may differ",
		"formats_table", "(table): The video format table as printed by --list-formats",
		"subtitles_table", "(table): The subtitle format table as printed by --list-subs",
		"thumbnails_table", "(table): The thumbnail format table as printed by --list-thumbnails",
		"urls", "(string): The URLs of all requested formats, one in each line",
	).Tag("print fields")
}

// ActionSponsorblockFields completes sponsorblock fields
//
//	category
//	start_time
func ActionSponsorblockFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"categories", "(list): The SponsorBlock categories the chapter belongs to",
		"category", "(string): The smallest SponsorBlock category the chapter belongs to",
		"category_names", "(list): Friendly names of the categories",
		"end_time", "(numeric): End time of the chapter in seconds",
		"name", "(string): Friendly name of the smallest category",
		"start_time", "(numeric): Start time of the chapter in seconds",
		"type", "(string): The SponsorBlock action type of the chapter",
	).Tag("sponsorblock fields")
}

// ActionNumericMetaFields completes numeric meta fields
//
//	like_count
//	view_count
func ActionNumericMetaFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"abr", "Average audio bitrate in KBit/s",
		"aspect_ratio", "Aspect ratio of the video, if known",
		"asr", "Audio sampling rate in Hertz",
		"audio_channels", "The number of audio channels",
		"filesize", "The number of bytes, if known in advance",
		"filesize_approx", "An estimate for the number of bytes",
		"fps", "Frame rate",
		"height", "Height of the video, if known",
		"stretched_ratio", "width:height of the video's pixels, if not square)",
		"tbr", "Average bitrate of audio and video in KBit/s",
		"vbr", "Average video bitrate in KBit/s",
		"width", "Width of the video, if known",
	).Tag("numeric meta fields")
}

// ActionStringMetaFields completes string meta fields
//
//	title
//	uploader
func ActionStringMetaFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"acodec", "Name of the audio codec in use",
		"container", "Name of the container format",
		"dynamic_range", "The dynamic range of the video",
		"ext", "File extension",
		"format", "A human-readable description of the format",
		"format_id", "A short description of the format",
		"format_note", "Additional info about the format",
		"language", "Language code",
		"protocol", "The protocol that will be used for the actual download, lower-case (http, https, rtsp, rtmp, rtmpe, mms, f4m, ism, http_dash_segments, m3u8, or m3u8_native)",
		"resolution", "Textual description of width and height",
		"url", "Video URL",
		"vcodec", "Name of the video codec in use",
	).Tag("string meta fields")
}
