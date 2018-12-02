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
	"os"
	"fmt"
	"bytes"
	"mime/multipart"
)

type Client struct {
	UserId string
	PersonalAccessToken string
	OrganizationId string
}

func NewClient(userId string, personalAccessToken string, organizationId string) *Client {
	return &Client{UserId: userId, PersonalAccessToken: personalAccessToken, OrganizationId:organizationId}
}

func (c *Client) CreateModel(param *CreateModelParam) (*ModelResult, error) {
	jsonBytes, err := json.Marshal(param)

	payload := strings.NewReader(string(jsonBytes))

	print(string(jsonBytes))

	body, err := c.DoRequest("POST", "models", payload)
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

func (c *Client) DeleteModel(id string) error {
	_, err := c.DoRequest("DELETE", path.Join("models", id), nil)
	return err
}

func (c *Client) CreateModelVersion(modelId string, uploadFilePath string, param *CreateModelVersionParam) (*ModelVersionResult, error) {
	jsonBytes, err := json.Marshal(param)
	payload := strings.NewReader(string(jsonBytes))

	body, err := c.DoRequest("POST", path.Join("models", modelId, "versions"), payload)
	if err != nil {
		return nil, err
	}
	result := &ModelVersionResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)

	file, err := os.Open(uploadFilePath)
	defer file.Close()

	if err != nil {
		return nil, err
	}
	fw, err := w.CreateFormFile("file", uploadFilePath)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return nil, err
	}

	//url, err := url.QueryUnescape(result.UploadUrl)
	url := result.UploadUrl
	req, err := http.NewRequest(http.MethodPut, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("error upload file error code:%d %s %s %s", res.StatusCode, url, body))
	}

	return result, nil
}

func (c *Client) DeleteModelVersion(modelId string, modelVersionId string) error {
	_, err := c.DoRequest("DELETE", path.Join("models", modelId, "versions", modelVersionId), nil)
	return err
}

func (c *Client) CreateDeployment(modelId string, param *CreateDeploymentParam) (*DeploymentResult, error) {
	// FIXME: 不定形なJSONをMarshalで扱う方法がわからなかったのでとりあえずSprintfで作る
	payload := strings.NewReader(fmt.Sprintf(`{"name": "%s", "default_environment": %s}`, param.Name, param.DefaultEnvironment))
	body, err := c.DoRequest("POST", path.Join("models", modelId, "deployments"), payload)
	if err != nil {
		return nil, err
	}
	result := &DeploymentResult{}
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
