package products

import (
	"fmt"
	"net/http"

	"github.com/RadekKusiak71/goEcom/types"
	"github.com/RadekKusiak71/goEcom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	db types.ProductStore
}

func NewHandler(db types.ProductStore) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.handleGetProducts).Methods("GET")
	r.HandleFunc("/products", h.handleCreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", h.handleRetrieveProduct).Methods("GET")
	r.HandleFunc("/products/{id}", h.handleUpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", h.handleDeleteProduct).Methods("DELETE")

	r.HandleFunc("/categories/", h.handleCreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", h.handleRetrieveCategory).Methods("GET")
	r.HandleFunc("/categories", h.handleGetCategories).Methods("GET")
}
func (h *Handler) handleGetCategories(w http.ResponseWriter, r *http.Request) {
	products, err := h.db.GetCategories()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleRetrieveCategory(w http.ResponseWriter, r *http.Request) {
	productID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	cat, err := h.db.GetCategory(productID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, cat)
}

func (h *Handler) handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	type CategoryPayload struct {
		Name string
	}
	newCat := new(CategoryPayload)
	if err := utils.ParseJSON(r, newCat); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"))
		return
	}
	err := h.db.CreateCategory(types.Category{
		Name: newCat.Name,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}
	utils.WriteINFO(w, http.StatusCreated, "message", "category created")
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.db.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}
func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	productPayload := new(types.ProductPayload)
	if err := utils.ParseJSON(r, productPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"))
		return
	}

	res, err := h.db.CreateProduct(types.Product{
		Name:        productPayload.Name,
		CategoryID:  productPayload.CategoryID,
		Description: productPayload.Description,
		Price:       productPayload.Price,
		Quantity:    productPayload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("interal server error"))
		return
	}
	utils.WriteINFO(w, http.StatusCreated, "message", res)

}
func (h *Handler) handleRetrieveProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.db.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, product)
}
func (h *Handler) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	productPayload := new(types.ProductPayload)
	if err := utils.ParseJSON(r, productPayload); err != nil {
		utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("invalid payload"))
		return
	}

	account, err := h.db.UpdateProduct(types.Product{
		Name:        productPayload.Name,
		CategoryID:  productPayload.CategoryID,
		Description: productPayload.Description,
		Price:       productPayload.Price,
		Quantity:    productPayload.Quantity,
	}, productID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("product was not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, account)
}
func (h *Handler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := utils.ReadRequestID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.db.DeleteProduct(productID)
	if err != nil {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	}
	utils.WriteINFO(w, http.StatusOK, "message", res)
}
