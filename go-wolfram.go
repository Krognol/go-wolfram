package wolfram

import (
	"encoding/xml"
	"net/http"
	"strings"
)

//The client, requires an App ID, which you can sign up for at https://developer.wolframalpha.com/
type Client struct {
	AppID string
}

//The QueryResult is what you get back after a request
type QueryResult struct {
	Pods            []Pod    `xml:"pod"`
	Success         string   `xml:"success,attr"`
	Error           string   `xml:"error,attr"`
	NumPods         int      `xml:"numpods,attr"`
	DataTypes       string   `xml:"datatypes,attr"`
	TimedOut        string   `xml:"timedout,attr"`
	Timing          int      `xml:"timing,attr"`
	ParseTiming     int      `xml:"parsetiming,attr"`
	ParseTimedOut   string   `xml:"parsetimedout,attr"`
	ReCalculate     string   `xml:"recalculate,attr"`
	Id              string   `xml:"id,attr"`
	Host            string   `xml:"host,attr"`
	Server          int      `xml:"server,attr"`
	Related         string      `xml:"related,attr"`
	Version         string          `xml:"version,attr"`
	Warnings        Warnings   `xml:"warnings"`
	Assumptions     Assumptions `xml:"assumptions"`
	Sources         Sources     `xml:"sources"`
	Generalizations []Generalization `xml:"generalization"`
}

type Generalization struct {
	Topic       string `xml:"topic,attr"`
	Description string `xml:"desc,attr"`
	Url         string `xml:"url,attr"`
}

type Warnings struct {
	Count             int `xml:"count,attr"`
	Spellchecks       []Spellcheck `xml:"spellcheck"`
	Delimiters        []Delimiters `xml:"delimiters"`
	Translations      []Translation `xml:"translation"`
	ReInterpretations []ReInterpretation `xml:"reinterpret"`
}

type Spellcheck struct {
	Word       string `xml:"word,attr"`
	Suggestion string `xml:"suggestion,attr"`
	Text       string `xml:"text,attr"`
}

type Delimiters struct {
	Text string `xml:"text,attr"`
}

type Translation struct {
	Phrase      string `xml:"phrase,attr"`
	Translation string `xml:"trans,attr"`
	Language    string `xml:"lang,attr"`
	Text        string `xml:"text,attr"`
}

type ReInterpretation struct {
	Alternatives []Alternative `xml:"alternative"`
	Text         string `xml:"text,attr"`
	New          string `xml:"new,attr"`
}

type Alternative struct{}

type Assumptions struct {
	Assumption []Assumption `xml:"assumption"`
	Count      int `xml:"count,attr"`
}

type Assumption struct {
	Values   []Value `xml:"value"`
	Type     string `xml:"type,attr"`
	Word     string `xml:"word,attr"`
	Template string `xml:"template,attr"`
	Count    int `xml:"count,attr"`
}

type Value struct {
	Name        string `xml:"name,attr"`
	Word        string `xml:"word,attr"`
	Description string `xml:"desc,attr"`
	Input       string `xml:"input,attr"`
}

type Pod struct {
	SubPods    []SubPod `xml:"subpod"`
	Infos      Infos    `xml:"infos"`
	States     States  `xml:"states"`
	Title      string   `xml:"title,attr"`
	Scanner    string   `xml:"scanner,attr"`
	Id         string   `xml:"id,attr"`
	Position   int      `xml:"position,attr"`
	Error      string   `xml:"error,attr"`
	NumSubPods int      `xml:"numsubpods,attr"`
	Sounds     Sounds `xml:"sounds"`
}

type Sounds struct {
	Count int `xml:"count,attr"`
	Sound []Sound `xml:"sound"`
}

type Sound struct {
	Url  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

type Infos struct {
	Count int `xml:"count,attr"`
	Info  []Info `xml:"info"`
}

type Info struct {
	Text string `xml:"text,attr"`
	Img  []Img  `xml:"img"`
	Link []Link  `xml:"link"`
}

type Link struct {
	Url   string `xml:"url,attr"`
	Text  string `xml:"text,attr"`
	Title string `xml:"title,attr"`
}

type Sources struct {
	Count  int `xml:"count,attr"`
	Source []Source `xml:"source"`
}

type Source struct {
	Url  string `xml:"url,attr"`
	Text string `xml:"text,attr"`
}

type States struct {
	Count int `xml:"count"`
	State []State `xml:"state"`
}

type State struct {
	Name  string `xml:"name,attr"`
	Input string `xml:"input,attr"`
}

type SubPod struct {
	Image     Img      `xml:"img"`
	Plaintext string   `xml:"plaintext"`
	Title     string   `xml:"title,attr"`
}

type Img struct {
	Src    string   `xml:"src,attr"`
	Alt    string   `xml:"alt,attr"`
	Title  string   `xml:"title,attr"`
	Width  int      `xml:"width,attr"`
	Height int      `xml:"height,attr"`
}

//Example: Add[0] = "format=image"
//Additional parameters can be found at http://products.wolframalpha.com/docs/WolframAlpha-API-Reference.pdf, page 42
type AdditionalUrl struct {
	Add []string
}

//Gets the query result from the API and returns it
func (c *Client) GetQueryResult(query string, extra *AdditionalUrl) *QueryResult {
	query = strings.Replace(query, " ", "", -1)
	url := "https://api.wolframalpha.com/v2/query?input=" + query + "&appid=" + c.AppID

	if extra != nil {
		for i := range extra.Add {
			url += ("&" + extra.Add[i])
		}
	}

	data := &QueryResult{}

	err := GetXML(url, data)

	if err != nil {
		panic(err)
	}
	return data
}

//Gets the XML from the API and assigns the data to the target
//The target being a QueryResult struct
func GetXML(url string, target interface{}) error {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	get, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	
	defer get.Body.Close()
	return xml.NewDecoder(get.Body).Decode(target)
}