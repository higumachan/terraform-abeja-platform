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
	"bufio"
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

	result := &ModelResult{}
	err = c.doRequest("POST", "models", payload, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) RetrieveModel(id string) (*ModelResult, error) {
	result := &ModelResult{}
	err := c.doRequest("GET", path.Join("models", id), nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) DeleteModel(id string) error {
	err := c.doRequest("DELETE", path.Join("models", id), nil, &struct{}{})
	return err
}

func (c *Client) CreateModelVersion(modelId string, uploadFilePath string, param *CreateModelVersionParam) (*ModelVersionResult, error) {
	jsonBytes, err := json.Marshal(param)
	payload := strings.NewReader(string(jsonBytes))

	result := &ModelVersionResult{}
	err = c.doRequest("POST", path.Join("models", modelId, "versions"), payload, result)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	file, err := os.Open(uploadFilePath)
	defer file.Close()

	if err != nil {
		c.DeleteModelVersion(modelId, result.VersionId)
		return nil, err
	}
	if _, err = io.Copy(w, file); err != nil {
		return nil, err
	}

	//url, err := url.QueryUnescape(result.UploadUrl)
	url := result.UploadUrl
	req, err := http.NewRequest(http.MethodPut, url, &buf)
	if err != nil {
		c.DeleteModelVersion(modelId, result.VersionId)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		c.DeleteModelVersion(modelId, result.VersionId)
		return nil, err
	}

	if res.StatusCode != 200 {
		c.DeleteModelVersion(modelId, result.VersionId)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("error upload file error code:%d %s %s %s", res.StatusCode, url, body))
	}

	return result, nil
}

func (c *Client) DeleteModelVersion(modelId string, modelVersionId string) error {
	err := c.doRequest("DELETE", path.Join("models", modelId, "versions", modelVersionId), nil, &struct {}{})
	return err
}

func (c *Client) CreateDeployment(modelId string, param *CreateDeploymentParam) (*DeploymentResult, error) {
	// FIXME: 不定形なJSONをMarshalで扱う方法がわからなかったのでとりあえずSprintfで作る
	payload := strings.NewReader(fmt.Sprintf(`{"name": "%s", "default_environment": %s}`, param.Name, param.DefaultEnvironment))
	result := &DeploymentResult{}
	err := c.doRequest("POST", path.Join("models", modelId, "deployments"), payload, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteDeployment(deploymentId string) error {
	err := c.doRequest("DELETE", path.Join("deployments", deploymentId), nil, &struct{}{})
	return err
}

func (c *Client) CreateDeploymentService(deploymentId string, param *CreateDeploymentServiceParam) (*DeploymentServiceResult, error) {
	// FIXME: 不定形なJSONをMarshalで扱う方法がわからなかったのでとりあえずSprintfで作る
	payload := strings.NewReader(fmt.Sprintf(
		`{"instance_number": %d, "instance_type": "%s", "environment": %s, "version_id": "%s"}`,
		param.InstanceNumber,
		param.InstanceType,
		param.Environment,
		param.VersionId,
	))
	result := &DeploymentServiceResult{}
	err := c.doRequest("POST", path.Join("deployments", deploymentId, "services"), payload, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteDeploymentService(deploymentId, serviceId string) error {
	return c.doRequest("DELETE", path.Join("deployments", deploymentId, "services", serviceId), nil, &struct{}{})
}

func (c *Client) newRequest(method string, endpoint string, payload io.Reader) (*http.Request, error) {
	u, err := c.getUrl()
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

func (c *Client) doRequest(method string, endpoint string, payload io.Reader, result interface{}) error {
	req, _ := c.newRequest(
		method,
		endpoint,
		payload,)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(string(body))
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, result); err != nil {
		return err
	}

	return nil
}

func (c *Client) getUrl() (*url.URL, error) {
	u, err := url.Parse("https://api.abeja.io")
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "organizations", string(c.OrganizationId))
	return u, nil
}
