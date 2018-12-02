package apiclient

import (
	"net/http"
	"encoding/json"
	"strings"
	"io"
	"net/url"
	"path"
	"io/ioutil"
	"github.com/pkg/errors"
)

type Client struct {
	UserId string
	PersonalAccessToken string
	OrganizationId string
}

func NewClient(userId string, personalAccessToken string, organizationId string) *Client {
	return &Client{UserId: userId, PersonalAccessToken: personalAccessToken, OrganizationId:organizationId}
}

func (c *Client) CreateModel(name string, description string) (*ModelResult, error) {
	jsonBytes, err := json.Marshal(struct {
		Name string 		`json:"name"`
		Description string	`json:"description"`
	} {Name:name, Description:description})

	payload := strings.NewReader(string(jsonBytes))

	body, err := c.DoRequest("POST", "models", payload,)
	if err != nil {
		return nil, err
	}

	result := &ModelResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) RetrieveModel(id string) (*ModelResult, error) {
	body, err := c.DoRequest("GET", path.Join("models", id), nil)
	if err != nil {
		return nil, err
	}

	result := &ModelResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) NewRequest(method string, endpoint string, payload io.Reader) (*http.Request, error) {
	u, err := c.GetUrl()
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, endpoint)

	println(u.String())

	req, err := http.NewRequest(method, u.String(), payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	username := "user-" + c.UserId
	req.SetBasicAuth(username, c.PersonalAccessToken)

	return req, nil
}

func (c *Client) DoRequest(method string, endpoint string, payload io.Reader) ([]byte, error) {
	req, _ := c.NewRequest(
		method,
		endpoint,
		payload,)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) GetUrl() (*url.URL, error) {
	u, err := url.Parse("https://api.abeja.io")
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "organizations", string(c.OrganizationId))
	return u, nil
}
