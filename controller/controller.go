package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type inputRequest struct {
	requestIdMap map[int]int
	mu           sync.Mutex
}

var req *inputRequest

func sendRequest(endpoint string) {
	req.mu.Lock()

	// count of unique requests in the current minute
	count := len(req.requestIdMap)

	req.mu.Unlock()

	// Create the request URL with the unique request count as a query parameter
	requestURL := fmt.Sprintf("%s?count=%d", endpoint, count)

	resp, err := http.Get(requestURL)
	if err != nil {
		log.Printf("Failed to send GET request to %s: %v", requestURL, err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Sent GET request to %s, got response status: %s", requestURL, resp.Status)
}

func LogRequest() {
	req.mu.Lock()
	log.Printf("Unique requests in the last minute: %d", len(req.requestIdMap))

	// re-initialise it for the next minute after logging
	req.requestIdMap = make(map[int]int)

	req.mu.Unlock()
}

func Accept(w http.ResponseWriter, r *http.Request) {
	idInput := r.URL.Query().Get("id")
	if idInput == "" {
		http.Error(w, "failed", http.StatusBadRequest)
		return
	} else {

	}

	// Convert the id to an integer
	id, err := strconv.Atoi(idInput)
	_ = id
	if err != nil {
		http.Error(w, "failed", http.StatusBadRequest)
		return
	}

	req.mu.Lock()

	idCount, ok := req.requestIdMap[id]
	if ok {
		req.requestIdMap[id] = idCount + 1
	} else {
		req.requestIdMap[id] = 1
	}

	req.mu.Unlock()

	endpoint := r.URL.Query().Get("endpoint")
	if endpoint != "" {
		go sendRequest(endpoint)

		// extension #1
		//go sendRequestExtension(endpoint)
	}

	_, writeErr := w.Write([]byte("ok"))
	if writeErr != nil {
		log.Println("Error", writeErr)
	}
}
