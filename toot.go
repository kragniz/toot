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

type Client struct {
	Url         string
	App         App
	AccessToken string
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

type appRequest struct {
	ClientName   string `json:"client_name"`
	Scopes       string `json:"scopes"`
	RedirectUris string `json:"redirect_uris"`
}

type App struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (c *Client) NewApp(name string, scope Scope) App {
	url := c.Url + "/api/v1/apps"

	r := appRequest{
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

	c.App = app

	return app
}

type loginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

func (c *Client) Login(username, password string, scope Scope) *Client {
	url := c.Url + "/oauth/token"

	r := loginRequest{
		Username:     username,
		Password:     password,
		Scope:        scope.String(),
		ClientID:     c.App.ClientID,
		ClientSecret: c.App.ClientSecret,
		GrantType:    "password",
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

	login := loginResponse{}
	json.Unmarshal(body, &login)

	c.AccessToken = login.AccessToken

	return c
}

type statusRequest struct {
	Status       string `json:"status"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

func (c *Client) Toot(status string) {
	url := c.Url + "/api/v1/statuses"

	r := statusRequest{
		Status: status,
	}

	jsonStr, _ := json.Marshal(r)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
