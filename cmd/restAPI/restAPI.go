package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aaronsuns/assignment/pkg/numrange"
)

type RangeRequest struct {
	Includes []string `json:"includes"`
	Excludes []string `json:"excludes"`
}

type RangeResponse struct {
	Result string `json:"result"`
}

func processRanges(request RangeRequest) RangeResponse {
	var includesRanges, excludesRanges []numrange.Range

	for _, includeStr := range request.Includes {
		r, err := numrange.ParseRange(includeStr)
		if err != nil {
			return RangeResponse{Result: "Failed to parse include ranges: " + err.Error()}
		}
		includesRanges = append(includesRanges, r)
	}

	for _, excludeStr := range request.Excludes {
		r, err := numrange.ParseRange(excludeStr)
		if err != nil {
			return RangeResponse{Result: "Failed to parse exclude ranges: " + err.Error()}
		}
		excludesRanges = append(excludesRanges, r)
	}

	resultRanges := numrange.ProcessNumberRanges(includesRanges, excludesRanges)
	result := numrange.FormatRanges(resultRanges)

	return RangeResponse{Result: result}
}

func main() {
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		var request RangeRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := processRanges(request)

		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
