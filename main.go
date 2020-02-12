package main

import (
    "bytes"
	"bufio"
    "encoding/json"
    "fmt"
	"flag"
    "log"
    "net/http"
    "net/url"
    "os"
	"io/ioutil"
)

var (
	verbose bool
	to string
	from string
	txt string
	apiKey string
	filename string
)

const endpoint = "https://api.cognitive.microsofttranslator.com"

// Trans is the type use to unmarshal service response
type Trans struct {
	Translations []TransData `json:"translations"`
}

// Data unmarshals the content of the reponse
type TransData struct {
	Text string `json:"text"`
	To string `json:"to"`
}

type ErrorData struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Data ErrorData `json:"error"`
}

func init() {
	flag.Usage = func() {
		h := "-------------------------------------------------------\n"
		h += "gotrans - Translate text using Azure Cognitive Services\n"
		h += "-------------------------------------------------------\n"

		h += "\nUsage:\n"
		h += "    ./gotrans -t nl \"<text-to-translate>\"\n"
		h += "    ./gotrans -t es -file <file>\n"
		h += "    echo \"text\" | gotrans [OPTIONS]\n"
		h += "    cat <file> | gotrans [OPTIONS]\n"
		h += "\nOptions:\n"
		h += "  -h, --help\n        Show help\n"

		fmt.Fprintf(os.Stderr, h)

		flag.PrintDefaults()
	}
}

func parseFlags(){

	flag.StringVar(&txt, "", "", "Text to translate")
	flag.BoolVar(&verbose, "v", false, "Run in verbose mode")
	flag.StringVar(&to, "t", "", "Target language. See list here: https://bit.ly/37o3PFX")
	flag.StringVar(&from, "f", "", "Source Language (optional). Set empty to auto-detect")
	flag.StringVar(&filename, "file", "", "File name")

	flag.Parse()
	txt = flag.Arg(0)

	d("Parsing flags...")
	d(len(os.Args), "args:", os.Args)
	d("Source language:", from)
	d("Target language:", to)
	d("Verbose:", verbose)
	d("File name:", filename)

    apiKey = os.Getenv("TRKEY")
	d("ApiKey:", apiKey)

    if apiKey == "" {
		log.Fatal("Please set/export the environment variable TRKEY.")
    }

	// read source file if specified
	if filename != "" {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		txt = string(content)
	}

	// if nothing else, load it from stdin
	if txt == "" {
		txt = readStdin()
	}

}

func err(msg string){
	fmt.Printf("\nError: %v\n\n", msg)
	flag.Usage()
	os.Exit(1)
}

func main() {

	parseFlags()


	if txt == ""{
		err("No text to translate provided. Please specify it via command line, file or stdin")
	} else if to == "" {
		err("No target language provided. Please specify a target language.\nExample: gotrans -t fr \"text to translate\"")
	}
	
    uri := endpoint + "/translate?api-version=3.0"
    translate(apiKey, uri, txt, from, to);

	os.Exit(0)
}

func readStdin() string {
	var in string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan(){
		in += s.Text() + "\n"
	} 

	if l := len(in); l > 0 {
		in = in[:l-1]
	}

	d("Stdin:", in)

	return in
}

// translate Translates the string `txt` using `from` as source
// language, `to` as target language. Set `from` empty to let
// Cognitive Services to auto-determine source language
func translate(apiKey, uri, txt, from, to string) {
    u, _ := url.Parse(uri)
    q := u.Query()

	if from != "" {
		q.Add("from", from);
	}

	if to != "" {
		q.Add("to", to);
	}

	d("Text to translate:", txt)
    u.RawQuery = q.Encode()

    // Create an anonymous struct for your request body and encode it to JSON
    body := []struct {
        Text string
    }{
        {Text: txt},
    }
    b, _ := json.Marshal(body)

	d("Request Url:", u)
	d("Payload:", string(b))

    // Build the HTTP POST request
    req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
    if err != nil {
        log.Fatal(err)
    }

    // Add required headers to the request
    req.Header.Add("Ocp-Apim-Subscription-Key", apiKey)
    req.Header.Add("Content-Type", "application/json")

    // Call the Translator Text API
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
	defer res.Body.Close()

	d("Request Status:", res.Status)
	d("Headers:", res.Header)
    resp, _ := ioutil.ReadAll(res.Body)
    d("Response:", string(resp))

    // unmarshall the response
	var t []Trans
	err = json.Unmarshal([]byte(resp), &t)
	if err != nil {
		// try to cast as error if response was not a valid request
		var e Error
		err = json.Unmarshal([]byte(resp), &e)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal("Error: ", e.Data.Message)
	}

	fmt.Println(t[0].Translations[0].Text)
}

func d(v ...interface{}) {
	if !verbose {
		return
	}

	log.Println("[debug]", v)
}
