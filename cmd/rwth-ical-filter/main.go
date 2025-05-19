package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"rwth-ical-filter/pkg/ical"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		studid := r.URL.Query().Get("pStud")
		token := r.URL.Query().Get("pToken")
		lvs := r.URL.Query()["lv"]

		if studid == "" || token == "" || len(lvs) == 0 {
			http.Error(w, "Missing pStud, pToken or lv", http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf(
			"https://online.rwth-aachen.de/RWTHonlinej/ws/termin/ical?pStud=%s&pToken=%s",
			studid, token,
		)
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			http.Error(w, "Failed to fetch iCal", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read iCal content", http.StatusInternalServerError)
			return
		}

		modified := ical.RemoveAllMatchingEvents(string(body), lvs)
		w.Header().Set("Content-Type", "text/calendar")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(modified))
	})

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
