package wolfram

import (
	"encoding/xml"
	"net/http"
	"strings"
)

type WolfClient struct {
	AppID string
}

type QueryResult struct {
	XMLName       xml.Name `xml:"queryresult"`
	Pods          []Pod    `xml:"pod"`
	Success       string   `xml:"success,attr"`
	Error         string   `xml:"error,attr"`
	NumPods       int      `xml:"numpods,attr"`
	DataTypes     string   `xml:"datatypes,attr"`
	TimedOut      string   `xml:"timedout,attr"`
	Timing        int      `xml:"timing,attr"`
	ParseTiming   int      `xml:"parsetiming,attr"`
	ParseTimedOut string   `xml:"parsetimedout,attr"`
	ReCalculate   string   `xml:"recalculate,attr"`
	Id            string   `xml:"id,attr"`
	Host          string   `xml:"host,attr"`
	Server        int      `xml:"server,attr"`
	Related       string   `xml:"related,attr"`
	Version       string   `xml:"version,attr"`
}

type Pod struct {
	XMLName    xml.Name `xml:"pod"`
	SubPods    []SubPod `xml:"subpod"`
	Title      string   `xml:"title,attr"`
	Scanner    string   `xml:"scanner,attr"`
	Id         string   `xml:"id,attr"`
	Position   int      `xml:"position,attr"`
	Error      string   `xml:"error,attr"`
	NumSubPods int      `xml:"numsubpods,attr"`
}

type SubPod struct {
	XMLName   xml.Name `xml:"subpod"`
	Image     Img      `xml:"img"`
	Plaintext string   `xml:"plaintext"`
	Title     string   `xml:"title,attr"`
}

type Img struct {
	XMLName xml.Name `xml:"img"`
	Src     string   `xml:"src,attr"`
	Alt     string   `xml:"alt,attr"`
	Title   string   `xml:"title,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

func (c *WolfClient) GetKnowledge(query string) *QueryResult {
	query = strings.Replace(query, " ", "", -1)
	url := "http://api.wolframalpha.com/v2/query?input=" + query + "&appid=" + c.AppID

	data := &QueryResult{}

	err := GetXML(url, data)

	if err != nil {
		return nil
	}
	return data
}


func GetXML(url string, target interface{}) error {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	get, err := client.Do(req)

	if err != nil {
		return err
	}

	defer get.Body.Close()

	return xml.NewDecoder(get.Body).Decode(target)
}