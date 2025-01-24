package res

import (
	"encoding/json"
	"net/http"
)



func Response(w http.ResponseWriter,data any , status int){
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(status)
}