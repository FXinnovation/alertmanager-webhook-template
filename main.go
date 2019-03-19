package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/alertmanager/template"
	"log"
	"net/http"
	"os"
)

type responseJSON struct {
	Status  int
	Message string
}

func webhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data := template.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		asJson(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("Alerts: Status=%v, GroupLabels=%v, CommonLabels=%v", data.Status, data.GroupLabels, data.CommonLabels)
	for _, alert := range data.Alerts {
		log.Printf("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)
	}

	asJson(w, http.StatusOK, "success")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ok!")
	asJson(w, http.StatusOK, "success")
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/webhook", webhook)

	listenAddress := ":9876"
	if os.Getenv("PORT") != "" {
		listenAddress = ":" + os.Getenv("PORT")
	}

	log.Printf("listening on: %v", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func asJson(w http.ResponseWriter, status int, message string) {
	data := responseJSON{
		Status:  status,
		Message: message,
	}
	bytes, _ := json.Marshal(data)
	json := string(bytes[:])

	w.WriteHeader(status)
	fmt.Fprint(w, json)
}