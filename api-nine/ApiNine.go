package apinine

func ApiNine(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
	w.writer([]byte("9"))
}