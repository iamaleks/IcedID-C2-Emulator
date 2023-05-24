package main

import (
	"net/http"
	"log"
	"bufio"
	"os"
)

func GetFile(path string) ([]byte, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil,statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_,err = bufr.Read(bytes)

	return bytes,err


}

func HandleRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		log.Printf("[+] Got %s request from %s", req.Method, req.RemoteAddr)

		file,err := GetFile("FileToServe.dat")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("Error opening local file")
			return
		}

		w.Header().Set("Content-Type", "application/gzip")
		w.WriteHeader(http.StatusOK)
		w.Write(file)
	}
}

func main() {

	log.Printf("[+] Server is Running")

	http.HandleFunc("/", HandleRequest)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
