package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":4545", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func readFile(fileTitle string) string {
	file, err := os.Open("./upload/" + fileTitle + ".json")
	if err != nil {
		fmt.Println("Error with reading the file:" + fileTitle)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	var json string
	for {
		count, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if count < 64 {
			endpart := data[:count]
			json += string(endpart)
		} else {
			json += string(data)
		}
	}
	return json
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param 'key' is missing")
		w.Write([]byte("{\"success\":\"false\"}"))
		return
	}
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	w.Header().Add("status", "success")
	w.Write([]byte("{'success':'true'"))
	var returnValue string = readFile(key)
	fmt.Println(returnValue + "end")
	fmt.Println(len(returnValue))
	w.Write([]byte(",'body':'" + returnValue + "'}"))
}
