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
	training, err = h.Repo.GetAllTrainings()
	if(err!=nil){
		panic(err)
	}
	json.NewEncoder(w).Encode(&training)
	w.WriteHeader(http.StatusOK)
}

func(h *HttpHandler) CreateTrainingGroup(w http.ResponseWriter, r *http.Request){
	var trainingGroup TrainingGroup
	json.NewDecoder(r.Body).Decode(&trainingGroup)
	res, _ :=h.Repo.InsertTrainingGroup(&trainingGroup)
	if(!res){
		w.WriteHeader(http.StatusBadRequest)

	}else{
		w.WriteHeader(http.StatusCreated)
	}
}

func(h *HttpHandler) GetAllTrainingGroups(w http.ResponseWriter, r *http.Request){
	var trainingGroups []TrainingGroup
	trainingGroups, err = h.Repo.GetAllTrainingGroups()
	if(err!=nil){
		panic(err)
	}
	json.NewEncoder(w).Encode(&trainingGroups)
	w.WriteHeader(http.StatusOK)
}
