package main

import (
	"encoding/json"
	"github.com/prometheus/alertmanager/template"
	"log"
	"net/http"
	"os"
)

type JsonResponse struct {
	Status  int
	Message string
}

func webhook(w http.ResponseWriter, r *http.Request) {
	// Do not forget to close the body at the end
	defer r.Body.Close()

	// Extract data from the body in the Data template provided by AlertManager
	data := template.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Do stuff here
	log.Printf("Alerts: Status=%s, GroupLabels=%v, CommonLabels=%v", data.Status, data.GroupLabels, data.CommonLabels)
	for _, alert := range data.Alerts {
		log.Printf("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)
	}

	// Returns a 200 if everything went smoothly
	sendJsonResponse(w, http.StatusOK, "Success")
}

// Function used to give a status on the webhook receiver
func health(w http.ResponseWriter, r *http.Request) {
	sendJsonResponse(w, http.StatusOK, "Success")
}


// Starts 2 listeners
// - first one to give a status on the receiver itself
// - second one to actually process the data
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

func sendJsonResponse(w http.ResponseWriter, status int, message string) {
	data := JsonResponse{
		Status:  status,
		Message: message,
	}
	bytes, _ := json.Marshal(data)

	w.WriteHeader(status)
	w.Write(bytes)
}