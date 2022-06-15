package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["id"])

	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
