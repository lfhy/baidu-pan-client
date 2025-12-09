package types

var (
	ClientId     string
	ClientSecret string
	RedirectUri  = "oob"
	AccessToken  string
	RefreshToken string
	AuthBaseURL  = "https://openapi.baidu.com"
	PanBaseURL   = "https://pan.baidu.com"
)

const (
	AuthRoute       = "/oauth/2.0/authorize"
	TokenRoute      = "/oauth/2.0/token"
	FileRoute       = "/rest/2.0/xpan/file"
	MultimediaRoute = "/rest/2.0/xpan/multimedia"
	NasRoute        = "/rest/2.0/xpan/nas"
	QuotaRoute      = "/api/quota"
)

func getENV(key string) string {
	switch key {
	case "$CLIENT_ID":
		return ClientId
	case "$CLIENT_SECRET":
		return ClientSecret
	case "$REDIRECT_URI":
		return RedirectUri
	default:
		return key
	}
}
