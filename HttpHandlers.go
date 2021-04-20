package main


import (
	"encoding/json"
	"net/http"
)


type HttpHandler struct{
	 Repo Repository
}


func (h *HttpHandler) CreateTraining(w http.ResponseWriter, r *http.Request){
	var training Training
	json.NewDecoder(r.Body).Decode(&training)
	h.Repo.InsertTranining(&training)
	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) GetTrainings(w http.ResponseWriter, r *http.Request){
	var training []Training
	training = h.Repo.GetAllTrainings()
	json.NewEncoder(w).Encode(&training)
	w.WriteHeader(http.StatusOK)
}