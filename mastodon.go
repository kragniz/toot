package toot

import (
	"strings"
)

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

type Scope struct {
	Read   bool
	Write  bool
	Follow bool
}

func (s Scope) String() string {
	scopes := []string{}

	if s.Read {
		scopes = append(scopes, "read")
	}

	if s.Write {
		scopes = append(scopes, "write")
	}

	if s.Follow {
		scopes = append(scopes, "follow")
	}

	return strings.Join(scopes, " ")
}
