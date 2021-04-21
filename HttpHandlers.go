package main


import (
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)


type HttpHandler struct{
	 UserRepo UserIRepository
	 ProgressRepo ProgressIRepository
	 TrainingRepo TrainingIRepository
	 TrainingGroupRepo TrainingGroupIRepository
	 TrainRelatTG TrainRelatTrainGroupIRepository
}


func (h *HttpHandler) CreateTraining(w http.ResponseWriter, r *http.Request){
	var training Training
	json.NewDecoder(r.Body).Decode(&training)
	h.TrainingRepo.Insert(&training)
	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) GetTrainings(w http.ResponseWriter, r *http.Request){
	var training []Training
	training, err = h.TrainingRepo.GetAll()
	if(err!=nil){
		panic(err)
	}
	json.NewEncoder(w).Encode(&training)
	w.WriteHeader(http.StatusOK)
}

func(h *HttpHandler) CreateTrainingGroup(w http.ResponseWriter, r *http.Request){
	var trainingGroup TrainingGroup
	json.NewDecoder(r.Body).Decode(&trainingGroup)
	res, _ :=h.TrainingGroupRepo.Insert(&trainingGroup)
	if(!res){
		w.WriteHeader(http.StatusBadRequest)

	}else{
		w.WriteHeader(http.StatusCreated)
	}
}

func(h *HttpHandler) GetAllTrainingGroups(w http.ResponseWriter, r *http.Request){
	var trainingGroups []TrainingGroup
	trainingGroups, err = h.TrainingGroupRepo.GetAll()
	if(err!=nil){
		panic(err)
	}
	json.NewEncoder(w).Encode(&trainingGroups)
	w.WriteHeader(http.StatusOK)
}

func(h *HttpHandler) CreateTrainRelatTrainGroup (w http.ResponseWriter, r *http.Request){
	var TrainRelateTG TrainingRelationTrainingGroup
	json.NewDecoder(r.Body).Decode(&TrainRelateTG)
	res, err := h.TrainRelatTG.Insert(&TrainRelateTG)

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
	res, err := h.TrainingRepo.GetByGroupId(id)
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


func (h *HttpHandler) GetAllUsers(w http.ResponseWriter, r *http.Request){
	var users []User
	users, err = h.UserRepo.GetAll()
	if(err!=nil){
		panic(err)
	}
	json.NewEncoder(w).Encode(&users)
	w.WriteHeader(http.StatusOK)
}

func (h *HttpHandler) CreateProgress(w http.ResponseWriter, r *http.Request){
	var progress Progress
	json.NewDecoder(r.Body).Decode(&progress)
	res, err := h.ProgressRepo.Insert(&progress)
	if(!res){
		panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) GetAllProgresses(w http.ResponseWriter, r *http.Request){
	var progresses []Progress
	progresses, err = h.ProgressRepo.GetAll()
	if(err!=nil){
		panic(err)
	}
	json.NewEncoder(w).Encode(&progresses)
	w.WriteHeader(http.StatusOK)
}