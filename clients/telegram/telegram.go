package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"random-music-bot/lib/errs"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) GetUpdates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendAudio(chat_id int, file_id string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chat_id))
	q.Add("audio", file_id)

	_, err := c.doRequest("sendAudio", q)
	if err != nil {
		return errs.Wrap(err, "err while sending audio")
	}

	return nil
}

func (c *Client) doRequest(method string, q url.Values) ([]byte, error) {
	const errMsg = "err while doing request"
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, errs.Wrap(err, errMsg)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errs.Wrap(err, errMsg)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
