package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WordCount struct {
	Key   string `json:"word"`
	Value int    `json:"count"`
}

func CallWordCountService(w http.ResponseWriter, r *http.Request) {

	// Read String from file
	byteData, errFileRead := ioutil.ReadFile("GoLang_Test.txt")
	if errFileRead != nil {
		fmt.Print(errFileRead)
	}
	str := string(byteData)

	// Create JSON of text file string
	var jsonData = []byte(`{"text":"` + str + `"}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/TopTenWordCount", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Fprintf(w, "Error in creating POST NewRequest")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, errApi := client.Do(req)
	if errApi != nil {
		fmt.Fprintf(w, "Error in calling API")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	outputData, _ := ioutil.ReadAll(res.Body)
	errClose := res.Body.Close()
	if errClose != nil {
		fmt.Fprintf(w, "Error in closing response body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//converting byte data into Struct
	var outputJSON []WordCount
	errJson := json.Unmarshal(outputData, &outputJSON)
	if errJson != nil {
		fmt.Fprintf(w, "Error in converting JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	errorOutput := json.NewEncoder(w).Encode(outputJSON)
	if errorOutput != nil {
		fmt.Fprintf(w, "Error encoding response")
	}
}
