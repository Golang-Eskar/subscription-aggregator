package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Golang-Eskar/subscription-aggregator/internal/database"
	"github.com/Golang-Eskar/subscription-aggregator/internal/models"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var sub models.Subscription
	json.NewDecoder(r.Body).Decode(&sub)

	sub.StartDate = time.Now()

	id, err := database.Create(sub)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"id": id})
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	subs, _ := database.GetAll()
	json.NewEncoder(w).Encode(subs)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	s, err := database.GetByID(id)

	if err != nil {
		http.Error(w, "not found", 404)
		return
	}

	json.NewEncoder(w).Encode(s)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var sub models.Subscription
	json.NewDecoder(r.Body).Decode(&sub)

	err := database.Update(id, sub)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("updated"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := database.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("deleted"))
}

func Filter(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user_id")
	service := r.URL.Query().Get("service_name")

	result, _ := database.Filter(user, service)
	json.NewEncoder(w).Encode(result)
}

func Sum(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	total, _ := database.Sum(from, to)
	json.NewEncoder(w).Encode(map[string]any{"total": total})
}
