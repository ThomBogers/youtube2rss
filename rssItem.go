package main

type YoutubeDlData_s struct {
	WebpageURLBasename string `json:"webpage_url_basename"`
	AutomaticCaptions  struct {
	} `json:"automatic_captions"`
	URL           string      `json:"url"`
	Tbr           float64     `json:"tbr"`
	ID            string      `json:"id"`
	EpisodeNumber interface{} `json:"episode_number"`
	Categories    []string    `json:"categories"`
	UploaderURL   string      `json:"uploader_url"`
	StartTime     interface{} `json:"start_time"`
	Acodec        string      `json:"acodec"`
	SeasonNumber  interface{} `json:"season_number"`
	Description   string      `json:"description"`
	Ext           string      `json:"ext"`
	AltTitle      interface{} `json:"alt_title"`
	Formats       []struct {
		Format            string `json:"format"`
		URL               string `json:"url"`
		FormatNote        string `json:"format_note"`
		Protocol          string `json:"protocol"`
		Abr               int    `json:"abr,omitempty"`
		PlayerURL         string `json:"player_url"`
		Filesize          int    `json:"filesize,omitempty"`
		DownloaderOptions struct {
			HTTPChunkSize int `json:"http_chunk_size"`
		} `json:"downloader_options,omitempty"`
		FormatID    string  `json:"format_id"`
		Ext         string  `json:"ext"`
		Acodec      string  `json:"acodec"`
		Vcodec      string  `json:"vcodec"`
		Tbr         float64 `json:"tbr,omitempty"`
		HTTPHeaders struct {
			Accept         string `json:"Accept"`
			AcceptCharset  string `json:"Accept-Charset"`
			UserAgent      string `json:"User-Agent"`
			AcceptEncoding string `json:"Accept-Encoding"`
			AcceptLanguage string `json:"Accept-Language"`
		} `json:"http_headers"`
		Container  string `json:"container,omitempty"`
		Height     int    `json:"height,omitempty"`
		Fps        int    `json:"fps,omitempty"`
		Width      int    `json:"width,omitempty"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"formats"`
	Playlist          interface{} `json:"playlist"`
	WebpageURL        string      `json:"webpage_url"`
	DownloaderOptions struct {
		HTTPChunkSize int `json:"http_chunk_size"`
	} `json:"downloader_options"`
	Duration   int `json:"duration"`
	Thumbnails []struct {
		URL string `json:"url"`
		ID  string `json:"id"`
	} `json:"thumbnails"`
	PlayerURL          string      `json:"player_url"`
	AgeLimit           int         `json:"age_limit"`
	FormatID           string      `json:"format_id"`
	RequestedSubtitles interface{} `json:"requested_subtitles"`
	Filesize           int         `json:"filesize"`
	AverageRating      float64     `json:"average_rating"`
	Series             interface{} `json:"series"`
	Creator            interface{} `json:"creator"`
	Chapters           []struct {
		EndTime   float64 `json:"end_time"`
		StartTime float64 `json:"start_time"`
		Title     string  `json:"title"`
	} `json:"chapters"`
	Protocol      string      `json:"protocol"`
	EndTime       interface{} `json:"end_time"`
	DisplayID     string      `json:"display_id"`
	Filename      string      `json:"_filename"`
	Tags          []string    `json:"tags"`
	Thumbnail     string      `json:"thumbnail"`
	Fulltitle     string      `json:"fulltitle"`
	IsLive        interface{} `json:"is_live"`
	ExtractorKey  string      `json:"extractor_key"`
	Vcodec        string      `json:"vcodec"`
	PlaylistIndex interface{} `json:"playlist_index"`
	HTTPHeaders   struct {
		Accept         string `json:"Accept"`
		AcceptCharset  string `json:"Accept-Charset"`
		UserAgent      string `json:"User-Agent"`
		AcceptEncoding string `json:"Accept-Encoding"`
		AcceptLanguage string `json:"Accept-Language"`
	} `json:"http_headers"`
	DislikeCount int `json:"dislike_count"`
	Subtitles    struct {
	} `json:"subtitles"`
	Extractor   string      `json:"extractor"`
	Format      string      `json:"format"`
	LikeCount   int         `json:"like_count"`
	FormatNote  string      `json:"format_note"`
	UploadDate  string      `json:"upload_date"`
	License     string      `json:"license"`
	Abr         int         `json:"abr"`
	Uploader    string      `json:"uploader"`
	Annotations interface{} `json:"annotations"`
	ViewCount   int         `json:"view_count"`
	UploaderID  string      `json:"uploader_id"`
	Title       string      `json:"title"`
}
