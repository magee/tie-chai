package main;

import (
	"net/http"
	"io/ioutil"
)

func handleYelp(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		client_yelp := &http.Client{}

		keyword  := req.Header.Get("Keyword")
		location := req.Header.Get("Location")
		// TODO: use URL.Query to generate query string
		r, _ := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search?term=" + keyword + "&location=" + location , nil)
		// TODO: make a seperate post call to get token, and add to header
		r.Header.Add("Authorization", YELP_TOKEN)

		resp, _ := client_yelp.Do(r)

		data, _ := ioutil.ReadAll(resp.Body)

		w.Write(data)	

	} else if req.Method != http.MethodOptions {
		badRequest(w, "need get request" , 400)	
	}
}
