package main


import (
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)


type HttpHandler struct{
	 Repo Repository
	 UserRepo UserIRepository
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

func(h *HttpHandler) CreateTrainRelatTrainGroup (w http.ResponseWriter, r *http.Request){
	var TrainRelateTG TrainingRelationTrainingGroup
	json.NewDecoder(r.Body).Decode(&TrainRelateTG)
	res, err := h.Repo.InsertTrainRelatTG(&TrainRelateTG)

	if(!res){
		panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) GetTrainingsFromGroup (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
	res, err := h.Repo.GetTrainingsFromGroup(id)
	if(err!=nil){
		panic(err)
	}
	json.NewEncoder(w).Encode(&res)
	w.WriteHeader(http.StatusOK)
}

func(h *HttpHandler) CreateUser (w http.ResponseWriter, r *http.Request){
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	res, err := h.UserRepo.Insert(&user)

	if(!res){
		panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}
