package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var jsonFile, yamlFile string
	flag.StringVar(&jsonFile, "json", "", "path to json file.")
	flag.StringVar(&yamlFile, "yaml", "", "path to yaml file.")

	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	if jsonFile != "" {
		jsonData, err := ioutil.ReadFile(jsonFile)
		if err != nil {
			panic(err)
		}
		jsonHandler, err := JSONHandler([]byte(jsonData), mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Jamming on port 8080")
		http.ListenAndServe(":8080", jsonHandler)
	} else if yamlFile != "" {
		yamlData, err := ioutil.ReadFile(yamlFile)
		if err != nil {
			panic(err)
		}

		yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Mixing it up on port 8080")
		http.ListenAndServe(":8080", yamlHandler)
	} else {
		fmt.Println("Riffing on port 8080")
		http.ListenAndServe(":8080", mapHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello!")
}
