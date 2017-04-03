package mastodon

type Account struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Acct           string `json:"acct"`
	DisplayName    string `json:"display_name"`
	Note           string `json:"note"`
	Url            string `json:"url"`
	Avatar         string `json:"avatar"`
	Header         string `json:"header"`
	Locked         bool   `json:"locked"`
	CreatedAt      string `json:"created_at"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	StatusesCount  int    `json:"statuses_count"`
}
