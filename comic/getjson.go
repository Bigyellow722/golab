package comic

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type comicResponse struct {
	Month      string `json:"month"`
	Num        string `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safetitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type Comic struct {
	Alt         string
	PublishDate time.Time
	ImageURL    string `json:"img"`
	URL         string
	News        string
	Number      int
	SafeTitle   string
	Title       string
	Transcript  string
}

func (c *Client) doReq(rawurl, path string) (Comic, error) {
	var comic Comic

	req, err := http.NewRequest("GET", c.rawURL(rawurl)+path, nil)
	if err != nil {
		return comic, err
	}

	body, err := c.do(req)
	if err != nil {
		return comic, err
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&comic); err != nil {
		return comic, err
	}

	return comic, nil
}

func (c *Client) Get(rawurl string, number int) (Comic, error) {
	numStr := strconv.Itoa(number)
	return c.doReq(rawurl, "/"+numStr+"/info.0.json")
}

func (comic *Comic) Download(number int) error {
	res, err := http.Get(comic.ImageURL)
	if err != nil {
		panic(err)

	}
	tmp := strconv.Itoa(number)
	f, err := os.Create(tmp + ".png")
	if err != nil {
		panic(err)

	}

	io.Copy(f, res.Body)

	return err

}
