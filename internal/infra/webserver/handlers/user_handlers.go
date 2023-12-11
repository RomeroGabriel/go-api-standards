package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RomeroGabriel/go-api-standards/internal/dto"
	"github.com/RomeroGabriel/go-api-standards/internal/entity"
	"github.com/RomeroGabriel/go-api-standards/internal/infra/db"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB        db.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHandler(db db.UserInterface, jwt *jwtauth.JWTAuth, jwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        db,
		Jwt:           jwt,
		JwtExperiesIn: jwtExperiesIn,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handlerBadResquest(w, err)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		handlerBadResquest(w, err)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var userJWT dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&userJWT)
	if err != nil {
		handlerBadResquest(w, err)
		return
	}
	u, err := h.UserDB.FindByEmail(userJWT.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(userJWT.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, err := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExperiesIn)).Unix(),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tokenString)
	fmt.Println(tokenString)
	fmt.Println(tokenString)
	fmt.Println(tokenString)
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
