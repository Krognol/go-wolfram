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
	//The pods are what hold the majority of the information
	Pods            []Pod         `xml:"pod"`

	//Warnings hold information about for example spelling errors
	Warnings        Warnings         `xml:"warnings"`

	//Assumptions show info if some assumption was made while parsing the query
	Assumptions     Assumptions         `xml:"assumptions"`

	//Each Source contains a link to a web page with the source information
	Sources         Sources         `xml:"sources"`

	//Generalizes the query to display more information
	Generalizations []Generalization `xml:"generalization"`

	//true or false depending on whether the input could be successfully
	//understood. If false there will be no <pod> subelements
	Success         string         `xml:"success,attr"`

	//true or false depending on whether a serious processing error occurred,
	//such as a missing required parameter. If true there will be no pod
	//content, just an <error> sub-element.
	Error           string         `xml:"error,attr"`

	//The number of pod elements
	NumPods         int         `xml:"numpods,attr"`

	//Categories and types of data represented in the results
	DataTypes       string         `xml:"datatypes,attr"`

	//The number of pods that are missing because they timed out (see the
	//scantimeout query parameter).
	TimedOut        string         `xml:"timedout,attr"`

	//The wall-clock time in seconds required to generate the output.
	Timing          int         `xml:"timing,attr"`

	//The time in seconds required by the parsing phase.
	ParseTiming     int         `xml:"parsetiming,attr"`

	//Whether the parsing stage timed out (try a longer parsetimeout parameter
	//if true)
	ParseTimedOut   string         `xml:"parsetimedout,attr"`

	//A URL to use to recalculate the query and get more pods.
	ReCalculate     string         `xml:"recalculate,attr"`

	//These elements are not documented currently
	Id              string         `xml:"id,attr"`
	Host            string         `xml:"host,attr"`
	Server          int         `xml:"server,attr"`
	Related         string         `xml:"related,attr"`

	//The version specification of the API on the server that produced this result.
	Version         string           `xml:"version,attr"`
}

type Generalization struct {
	Topic       string `xml:"topic,attr"`
	Description string `xml:"desc,attr"`
	Url         string `xml:"url,attr"`
}

type Warnings struct {
	//How many warnings were issued
	Count             int `xml:"count,attr"`

	//Suggestions for spelling corrections
	Spellchecks       []Spellcheck `xml:"spellcheck"`

	//"If you enter a query with mismatched delimiters like "sin(x", Wolfram|Alpha attempts to fix the problem and reports
	//this as a warning."
	Delimiters        []Delimiters `xml:"delimiters"`

	//"[The API] will translate some queries from non-English languages into English. In some cases when it does
	//this, you will get a <translation> element in the API result."
	Translations      []Translation `xml:"translation"`

	//"[The API] can automatically try to reinterpret a query that it does not understand but that seems close to one
	//that it can."
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
	Text         string 	   `xml:"text,attr"`
	New          string 	   `xml:"new,attr"`
}

type Alternative struct {
	Text string `xml:",innerxml"`
}

type Assumptions struct {
	Assumption []Assumption `xml:"assumption"`
	Count      int 		`xml:"count,attr"`
}

type Assumption struct {
	Values   []Value `xml:"value"`
	Type     string  `xml:"type,attr"`
	Word     string  `xml:"word,attr"`
	Template string  `xml:"template,attr"`
	Count    int 	 `xml:"count,attr"`
}

//Usually contains info about an assumption
type Value struct {
	Name        string `xml:"name,attr"`
	Word        string `xml:"word,attr"`
	Description string `xml:"desc,attr"`
	Input       string `xml:"input,attr"`
}

//<pod> elements are subelements of <queryresult>. Each contains the results for a single pod
type Pod struct {
	//The subpod elements of the pod
	SubPods    []SubPod `xml:"subpod"`

	//sub elements of the pod
	Infos      Infos    `xml:"infos"`
	States     States  `xml:"states"`

	//The pod title, used to identify the pod.
	Title      string   `xml:"title,attr"`

	//The name of the scanner that produced this pod. A guide to the type of
	//data it holds.
	Scanner    string   `xml:"scanner,attr"`

	//Not documented currently
	Id         string   `xml:"id,attr"`
	Position   int      `xml:"position,attr"`
	Error      string   `xml:"error,attr"`
	NumSubPods int      `xml:"numsubpods,attr"`
	Sounds     Sounds `xml:"sounds"`
}

//If there was a sound related to the query, if you for example query a musical note
//You will get a <sound> element which contains a link to the sound
type Sounds struct {
	Count int `xml:"count,attr"`
	Sound []Sound `xml:"sound"`
}

type Sound struct {
	Url  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

//If there's extra information for the pod, the pod will have a <infos> element
//which contains <info> elements with text, and/or images/links to that information
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

//Each Source contains a link to a web page with the source information
type Sources struct {
	Count  int `xml:"count,attr"`
	Source []Source `xml:"source"`
}

type Source struct {
	Url  string `xml:"url,attr"`
	Text string `xml:"text,attr"`
}


//"Many pods on the Wolfram|Alpha website have text buttons in their upper-right corners that substitute the
//contents of that pod with a modified version. In Figure 1, the Result pod has buttons titled "More days", "Sun and
//Moon", CDT", "GMT", and "Show metric". Clicking any of these buttons will recompute just that one pod to display
//different information."
type States struct {
	Count int `xml:"count"`
	State []State `xml:"state"`
}

type State struct {
	Name  string `xml:"name,attr"`
	Input string `xml:"input,attr"`
}

type SubPod struct {
	//HTML <img> element
	Image     Img      `xml:"img"`

	//Textual representation of the subpod
	Plaintext string   `xml:"plaintext"`

	//Usually an empty string because most subpod elements don't have a title
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
type AdditionalParameters struct {
	Parameters []string
}

//Gets the query result from the API and returns it
func (c *Client) GetQueryResult(query string, extra *AdditionalParameters) *QueryResult {
	query = strings.Replace(query, " ", "%20", -1)
	url := "https://api.wolframalpha.com/v2/query?input=" + query + "&appid=" + c.AppID

	if extra != nil {
		for i := range extra.Parameters {
			url += ("&" + extra.Parameters[i])
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