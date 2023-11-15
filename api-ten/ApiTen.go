package apiten

import "net/http"

func ApiTen(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TEN"))
}
