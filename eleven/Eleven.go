package eleven

import "net/http"

func Eleven(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("11"))
}
