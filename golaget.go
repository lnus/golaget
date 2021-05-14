package golaget

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// SystemGoLaget is what holds http.Client
// Might store more data in the future
type SystemGoLaget struct {
	Client *http.Client
}

// Position holds the stores lat and lon
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// OpeningHour holds the stores open hours
type OpeningHour struct {
	Date     string `json:"date"`
	OpenFrom string `json:"openFrom"`
	OpenTo   string `json:"openTo"`
	Reason   string `json:"reason"`
}

// HeaderRoundTripper holds settings for the RoundTrip middleware
type HeaderRoundTripper struct {
	APIKey  string
	Proxy   http.RoundTripper
	Verbose bool
}

// RoundTrip sets the API key as 'middleware' for the http.Client
func (hrt HeaderRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	// Will only log if verbose is true
	if hrt.Verbose {
		log.Printf("Sending request to %v\n", req.URL)
	}

	// Sets the API header
	req.Header.Add("Ocp-Apim-Subscription-Key", hrt.APIKey)

	res, err = hrt.Proxy.RoundTrip(req)
	if err != nil {
		return
	}

	if hrt.Verbose {
		log.Printf("Received %v response\n", res.Status)
	}

	return
}

// JsonDecode does a request and decodes it onto target structure
func (s SystemGoLaget) JsonDecode(url string, target interface{}) error {
	// Performs the request
	r, err := s.Client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// Tries to decode the body into the target interface
	return json.NewDecoder(r.Body).Decode(target)
}

// New creates a new API HTTP instance and returns it
func New(key string, verbose bool) (instance *SystemGoLaget, err error) {
	// Sets up a http.Client to use with requests
	// Elimanites the need for keys in every call.
	instance = &SystemGoLaget{
		Client: &http.Client{
			Timeout: time.Second * 15,
			Transport: HeaderRoundTripper{
				Proxy:   http.DefaultTransport,
				APIKey:  key,
				Verbose: verbose,
			},
		}}

	// Returns object with set options keys
	return
}
