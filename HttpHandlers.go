package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
	"github.com/komtriangle/Bodroe_ytro_GO/repositories"
)

type HttpHandler struct {
	UserRepo          repositories.UserIRepository
	ProgressRepo      repositories.ProgressIRepository
	TrainingRepo      repositories.TrainingIRepository
	TrainingGroupRepo repositories.TrainingGroupIRepository
	TrainRelatTG      repositories.TrainRelatTrainGroupIRepository
}

func (h *HttpHandler) CreateTraining(w http.ResponseWriter, r *http.Request) {
	var training models.Training
	json.NewDecoder(r.Body).Decode(&training)
	res, err := h.TrainingRepo.Insert(&training)
	w.Header().Set("Content-Type", "application/json")
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&training)
		w.WriteHeader(http.StatusCreated)
	}

}

func (h *HttpHandler) GetTrainings(w http.ResponseWriter, r *http.Request) {
	var training []models.Training
	training, err := h.TrainingRepo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	} else {
		json.NewEncoder(w).Encode(&training)
		w.WriteHeader(http.StatusOK)
	}

}

func (h *HttpHandler) CreateTrainingGroup(w http.ResponseWriter, r *http.Request) {
	var trainingGroup models.TrainingGroup
	json.NewDecoder(r.Body).Decode(&trainingGroup)
	res, err := h.TrainingGroupRepo.Insert(&trainingGroup)
	w.Header().Set("Content-Type", "application/json")
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	} else {
		json.NewEncoder(w).Encode(&trainingGroup)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *HttpHandler) GetAllTrainingGroups(w http.ResponseWriter, r *http.Request) {
	var trainingGroups []models.TrainingGroup
	trainingGroups, err := h.TrainingGroupRepo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	} else {
		json.NewEncoder(w).Encode(&trainingGroups)
		w.WriteHeader(http.StatusOK)
	}

}

func (h *HttpHandler) CreateTrainRelatTrainGroup(w http.ResponseWriter, r *http.Request) {
	var TrainRelateTG models.TrainingRelationTrainingGroup
	json.NewDecoder(r.Body).Decode(&TrainRelateTG)
	res, err := h.TrainRelatTG.Insert(&TrainRelateTG)
	w.Header().Set("Content-Type", "application/json")
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&TrainRelateTG)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) GetTrainingsFromGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	res, err := h.TrainingRepo.GetByGroupId(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&res)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	res, err := h.UserRepo.Insert(&user)
	w.Header().Set("Content-Type", "application/json")
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&user)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	users, err := h.UserRepo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&users)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) CreateProgress(w http.ResponseWriter, r *http.Request) {
	var progress struct {
		Id        int
		UserToken string
		DateTime  string
	}
	json.NewDecoder(r.Body).Decode(&progress)
	fmt.Println(progress.DateTime)
	DateTime, err := time.Parse("2006-01-02", progress.DateTime)
	if err != nil {
		panic(err)
	}
	progress_With_Date := models.Progress{Id: progress.Id, UserToken: progress.UserToken, DateTime: DateTime}
	res, err := h.ProgressRepo.Insert(&progress_With_Date)
	w.Header().Set("Content-Type", "application/json")
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&progress)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) GetAllProgresses(w http.ResponseWriter, r *http.Request) {
	var progresses []models.Progress
	progresses, err := h.ProgressRepo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&progresses)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *HttpHandler) GetProgressByUser(w http.ResponseWriter, r *http.Request) {
	var progressbyUser models.ProgressbyUser
	var IdForProgress struct {
		UserId string
	}
	json.NewDecoder(r.Body).Decode(&IdForProgress)
	progressbyUser, err := h.ProgressRepo.GetProgressByUser(IdForProgress.UserId)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&progressbyUser)
		w.WriteHeader(http.StatusOK)
	}
}
