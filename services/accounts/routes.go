package accounts

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RadekKusiak71/goEcom/services/authentication"
	"github.com/RadekKusiak71/goEcom/types"
	"github.com/RadekKusiak71/goEcom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	db types.AccountStore
}

func NewHandler(db types.AccountStore) *Handler {
	return &Handler{
		db: db,
	}
}

func (h Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", h.handleCreateAccount)
	r.HandleFunc("/login", h.handleLoginAccount)
}

func (h *Handler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	accountPayload := new(types.AccountPayload)
	if err := utils.ParseJSON(r, accountPayload); err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"))
		return
	}

	if accountPayload.Password != accountPayload.Password2 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("password's don't match"))
		return
	}

	hashedPassword, err := authentication.HashPassword(accountPayload.Password)
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}
	res, err := h.db.CreateAccount(types.Account{
		FirstName: accountPayload.FirstName,
		LastName:  accountPayload.LastName,
		Email:     accountPayload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusConflict, err)
		return
	}
	utils.WriteINFO(w, http.StatusCreated, "message", res)
}

func (h *Handler) handleLoginAccount(w http.ResponseWriter, r *http.Request) {
	loginPayload := new(types.LoginPayload)
	if err := utils.ParseJSON(r, loginPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"))
		return
	}
	account, err := h.db.GetAccountByEmail(loginPayload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid password or email"))
		return
	}
	if authentication.ComparePasswords([]byte(account.Password), []byte(loginPayload.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password or email"))
		return
	}

	token, err := authentication.CreateJWT(account.ID)
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	loginResponse := types.LoginResponse{
		Token:   token,
		Message: "login successfull",
	}

	utils.WriteJSON(w, http.StatusOK, loginResponse)
}
