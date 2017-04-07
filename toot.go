package toot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

type AppRequest struct {
	ClientName   string `json:"client_name"`
	Scopes       string `json:"scopes"`
	RedirectUris string `json:"redirect_uris"`
}

type App struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func NewApp(name string, scope Scope) App {
	url := "https://cmpwn.com/api/v1/apps"

	r := AppRequest{
		ClientName:   name,
		Scopes:       scope.String(),
		RedirectUris: "urn:ietf:wg:oauth:2.0:oob",
	}

	jsonStr, _ := json.Marshal(r)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	app := App{}
	json.Unmarshal(body, &app)

	return app
}
