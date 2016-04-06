package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Route struct {
	Name       string
	Pattern    string
	HandleFunc http.HandlerFunc
}

var routes = []Route{
	Route{
		"Index",
		"/",
		Index,
	},
	Route{
		"Tracking Url Log",
		"/tracking",
		RecordTrackingUrl,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}

func main() {
	/*
		for i := 0; i < len(os.Args); i++ {
			fmt.Println(os.Args[i])
		}*/
	fmt.Println(os.Args[1])
	logfilename := os.Args[1] + ".log"
	logfile, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I am a mock server!")
}

func RecordTrackingUrl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("%s\n", r.URL.RawQuery)
	//	fmt.Println(r.Form)

	if r.Method == "GET" {
		fmt.Printf("This is a GET\n")
		log.Println("This is a GET")
		//log r.URL.RawQuery
	} else if r.Method == "POST" {
		log.Println("This is a POST")
		//	fmt.Printf("This is a POST\n")
		result, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Printf("Read Post body failed\n")
			panic("Read Post body failed")
		}
		// log result
		fmt.Printf("%s\n", string(result))
	}
}
