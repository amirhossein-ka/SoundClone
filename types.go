// Thanks to chatgpt for these (:
package main

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Subscription struct {
	Product   Product `json:"product"`
	Recurring bool    `json:"recurring"`
}

type QuotaInfo struct {
	UnlimitedUploadQuota bool `json:"unlimited_upload_quota"`
	UploadSecondsUsed    int  `json:"upload_seconds_used"`
	UploadSecondsLeft    int  `json:"upload_seconds_left,omitempty"`
}

type Me struct {
	AvatarURL             string         `json:"avatar_url"`
	City                  string         `json:"city"`
	CommentsCount         int            `json:"comments_count,omitempty"`
	Country               string         `json:"country"`
	CreatedAt             string         `json:"created_at"`
	Description           string         `json:"description"`
	DiscogsName           string         `json:"discogs_name"`
	FirstName             string         `json:"first_name"`
	FollowersCount        int            `json:"followers_count"`
	FollowingsCount       int            `json:"followings_count"`
	FullName              string         `json:"full_name"`
	ID                    int            `json:"id"`
	Kind                  string         `json:"kind"`
	LastModified          string         `json:"last_modified"`
	LastName              string         `json:"last_name"`
	LikesCount            int            `json:"likes_count"`
	Locale                string         `json:"locale,omitempty"`
	Online                bool           `json:"online"`
	MyspaceName           string         `json:"myspace_name,omitempty"`
	Permalink             string         `json:"permalink"`
	PermalinkURL          string         `json:"permalink_url"`
	Plan                  string         `json:"plan"`
	PlaylistCount         int            `json:"playlist_count"`
	PrimaryEmailConfirmed bool           `json:"primary_email_confirmed"`
	PrivatePlaylistsCount int            `json:"private_playlists_count"`
	PrivateTracksCount    int            `json:"private_tracks_count"`
	PublicFavoritesCount  int            `json:"public_favorites_count"`
	Quota                 QuotaInfo      `json:"quota"`
	RepostsCount          int            `json:"reposts_count"`
	Subscriptions         []Subscription `json:"subscriptions"`
	TrackCount            int            `json:"track_count"`
	UploadSecondsLeft     int            `json:"upload_seconds_left,omitempty"`
	URI                   string         `json:"uri"`
	Username              string         `json:"username"`
	Website               string         `json:"website"`
	WebsiteTitle          string         `json:"website_title"`
}

type Playlist struct {
	Title         string   `json:"title"`
	ID            int      `json:"id"`
	Kind          string   `json:"kind"`
	ArtworkURL    string   `json:"artwork_url"`
	CreatedAt     string   `json:"created_at"`
	Description   string   `json:"description"`
	Downloadable  bool     `json:"downloadable"`
	Duration      int      `json:"duration"`
	EAN           string   `json:"ean"`
	EmbeddableBy  string   `json:"embeddable_by"`
	Genre         string   `json:"genre"`
	LabelID       int      `json:"label_id"`
	LabelName     string   `json:"label_name"`
	LastModified  string   `json:"last_modified"`
	License       string   `json:"license"`
	Permalink     string   `json:"permalink"`
	PermalinkURL  string   `json:"permalink_url"`
	PlaylistType  string   `json:"playlist_type"`
	PurchaseTitle string   `json:"purchase_title"`
	PurchaseURL   string   `json:"purchase_url"`
	Release       string   `json:"release"`
	ReleaseDay    int      `json:"release_day"`
	ReleaseMonth  int      `json:"release_month"`
	ReleaseYear   int      `json:"release_year"`
	Sharing       string   `json:"sharing"`
	Streamable    bool     `json:"streamable"`
	TagList       string   `json:"tag_list"`
	TrackCount    int      `json:"track_count"`
	Tracks        []*Track `json:"tracks"`
	Type          string   `json:"type"`
	URI           string   `json:"uri"`
	User          User     `json:"user"`
	UserID        int      `json:"user_id"`
	LikesCount    int      `json:"likes_count"`
	Label         *User    `json:"label,omitempty"`
	TracksURI     string   `json:"tracks_uri,omitempty"`
	Tags          string   `json:"tags,omitempty"`
}

type User struct {
	AvatarURL            string         `json:"avatar_url"`
	City                 string         `json:"city"`
	Country              string         `json:"country"`
	Description          string         `json:"description"`
	DiscogsName          string         `json:"discogs_name"`
	FirstName            string         `json:"first_name"`
	FollowersCount       int            `json:"followers_count"`
	FollowingsCount      int            `json:"followings_count"`
	FullName             string         `json:"full_name"`
	ID                   int            `json:"id"`
	Kind                 string         `json:"kind"`
	CreatedAt            string         `json:"created_at"`
	LastModified         string         `json:"last_modified"`
	LastName             string         `json:"last_name"`
	MyspaceName          string         `json:"myspace_name,omitempty"`
	Permalink            string         `json:"permalink"`
	PermalinkURL         string         `json:"permalink_url"`
	Plan                 string         `json:"plan"`
	PlaylistCount        int            `json:"playlist_count"`
	PublicFavoritesCount int            `json:"public_favorites_count"`
	RepostsCount         int            `json:"reposts_count"`
	TrackCount           int            `json:"track_count"`
	URI                  string         `json:"uri"`
	Username             string         `json:"username"`
	Website              string         `json:"website"`
	WebsiteTitle         string         `json:"website_title"`
	Subscriptions        []Subscription `json:"subscriptions,omitempty"`
}

type Track struct {
	Title                 string `json:"title"`
	ArtworkURL            string `json:"artwork_url"`
	BPM                   int    `json:"bpm"`
	CommentCount          int    `json:"comment_count"`
	Commentable           bool   `json:"commentable"`
	CreatedAt             string `json:"created_at"`
	Description           string `json:"description"`
	DownloadCount         int    `json:"download_count"`
	Downloadable          bool   `json:"downloadable"`
	Duration              int    `json:"duration"`
	EmbeddableBy          string `json:"embeddable_by,omitempty"`
	FavoritingsCount      int    `json:"favoritings_count"`
	Genre                 string `json:"genre"`
	ID                    int    `json:"id"`
	ISRC                  string `json:"isrc"`
	KeySignature          string `json:"key_signature"`
	Kind                  string `json:"kind"`
	LabelName             string `json:"label_name"`
	License               string `json:"license"`
	PermalinkURL          string `json:"permalink_url"`
	PlaybackCount         int    `json:"playback_count"`
	PurchaseTitle         string `json:"purchase_title"`
	PurchaseURL           string `json:"purchase_url"`
	Release               string `json:"release"`
	ReleaseDay            int    `json:"release_day"`
	ReleaseMonth          int    `json:"release_month"`
	ReleaseYear           int    `json:"release_year"`
	Sharing               string `json:"sharing"`
	StreamURL             string `json:"stream_url"`
	Streamable            bool   `json:"streamable"`
	TagList               string `json:"tag_list"`
	URI                   string `json:"uri"`
	User                  User   `json:"user"`
	UserFavorite          bool   `json:"user_favorite"`
	UserPlaybackCount     int    `json:"user_playback_count"`
	WaveformURL           string `json:"waveform_url"`
	AvailableCountryCodes string `json:"available_country_codes"`
	Access                string `json:"access,omitempty"`
	DownloadURL           string `json:"download_url"`
	RepostsCount          int    `json:"reposts_count"`
	SecretURI             string `json:"secret_uri"`
}

type Playlists struct {
	Collection []Playlist `json:"collection"`
	NextHref   string     `json:"next_href"`
}
