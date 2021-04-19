
func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
  }

func create_training(w http.ResponseWriter, r *http.Request){
	var training Training
	json.NewDecoder(r.Body).Decode(&training)
  
  
	fmt.Println(training)
	db.Create(&training)
  }
  
  func get_trainings(w http.ResponseWriter, r *http.Request){
	  var training []Training
	
	  db.Find(&training)
	
	  json.NewEncoder(w).Encode(&training)
  }
  