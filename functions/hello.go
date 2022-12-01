package functions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pteich/cloudfunctions/hello/pkg/helloservice"
)

type NameRequest struct {
	Name string `json:"name"`
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var req NameRequest

	hs := helloservice.New()

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	greeting := hs.HelloName(req.Name)

	fmt.Fprintln(w, greeting)
}
