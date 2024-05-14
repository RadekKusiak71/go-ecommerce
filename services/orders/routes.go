package orders

import (
	"fmt"
	"net/http"

	"github.com/RadekKusiak71/goEcom/types"
	"github.com/RadekKusiak71/goEcom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	db types.OrderStore
}

func NewHandler(db types.OrderStore) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/orders", h.handleGetOrders)
	r.HandleFunc("/orders/user/{id}", h.handleGetOrders)
	r.HandleFunc("/orders/{id}/details", h.handleGetOrderDetails)
}

func (h *Handler) handleGetOrderDetails(w http.ResponseWriter, r *http.Request) {
	accountID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	orderDetails, err := h.db.GetFullOrderDetails(accountID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orderDetails)
}

func (h *Handler) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.db.GetOrders()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error while fetching orders"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *Handler) handleGetUserOrders(w http.ResponseWriter, r *http.Request) {
	accountID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	orders, err := h.db.GetOrderByID(accountID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error while fetching orders"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, orders)
}
