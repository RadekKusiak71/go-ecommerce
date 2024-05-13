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

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", h.handleCreateAccount).Methods("POST")
	r.HandleFunc("/login", h.handleLoginAccount).Methods("POST")

	r.HandleFunc("/accounts", h.handleGetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", h.handleRetrieveAccount).Methods("GET")
	r.HandleFunc("/accounts/{id}", h.handleUpdateAccount).Methods("PUT")
	r.HandleFunc("/accounts/{id}", h.handleDeleteAccount).Methods("DELETE")
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

func (h *Handler) handleGetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.db.GetAccounts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("accounts were not found"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, accounts)
}
func (h *Handler) handleRetrieveAccount(w http.ResponseWriter, r *http.Request) {

	accountID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	account, err := h.db.GetAccountByID(accountID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("account was not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, account)
}
func (h *Handler) handleUpdateAccount(w http.ResponseWriter, r *http.Request) {
	accountID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	accountPayload := new(types.AccountUpdate)
	if err := utils.ParseJSON(r, accountPayload); err != nil {
		utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("invalid payload"))
		return
	}

	account, err := h.db.UpdateAccount(types.Account{
		FirstName: accountPayload.FirstName,
		LastName:  accountPayload.LastName,
		Email:     accountPayload.Email,
	}, accountID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("accounts were not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, account)
}
func (h *Handler) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	accountID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if res, err := h.db.DeleteAccount(accountID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	} else {
		utils.WriteINFO(w, http.StatusOK, "message", res)
	}
}
