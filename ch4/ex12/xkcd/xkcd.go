package xkcd

const XkcdUrl = "https://xkcd.com/"
const JsonPath = "/info.0.json"

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}
