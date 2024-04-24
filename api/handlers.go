package api

import (
	"encoding/json"
	"net/http"
	"user-service/db"
	"user-service/service"
)

type Handler struct {
	usersService service.UsersHttpService
}

func NewHandler(httpService service.UsersHttpService) *Handler {
	return &Handler{
		usersService: httpService,
	}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим данные из запроса
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.usersService.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := h.usersService.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) FilterUsersHandler(w http.ResponseWriter, r *http.Request) {
	dateStart := r.URL.Query().Get("dateStart")
	dateEnd := r.URL.Query().Get("dateEnd")
	ageStart := r.URL.Query().Get("ageStart")
	ageEnd := r.URL.Query().Get("ageEnd")

	filteredUsers, _ := h.usersService.FilterUser(
		dateStart, dateEnd, ageStart, ageEnd)

	response := struct {
		Users []db.User `json:"users"`
		Total int       `json:"total"`
	}{
		Users: filteredUsers,
		Total: len(filteredUsers),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
