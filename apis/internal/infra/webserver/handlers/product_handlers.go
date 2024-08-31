package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mwives/go-expert/apis/internal/dto"
	"github.com/mwives/go-expert/apis/internal/entity"
	"github.com/mwives/go-expert/apis/internal/infra/database"
	entityPkg "github.com/mwives/go-expert/apis/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// Create Product godoc
// @Summary      Create product
// @Description  Create a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request	body	dto.CreateProductInput	true	"Product request"
// @Success      201
// @Failure      400
// @Failure      500	{object}	Error
// @Router       /products [post]
// @Security     ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Products godoc
// @Summary      List products
// @Description  Get a paginated list of products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page   query     int     false  "Page number"
// @Param        limit  query     int     false  "Number of products per page"
// @Param        sort   query     string  false  "Sort order"
// @Success      200    {array}   entity.Product
// @Failure      500    {object}  Error
// @Router       /products [get]
// @Security     ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}

	products, err := h.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Get Product godoc
// @Summary      Get product
// @Description  Get a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Product ID"
// @Success      200   {object}  entity.Product
// @Failure      400   {object}  Error
// @Failure      404   {object}  Error
// @Router       /products/{id} [get]
// @Security     ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary      Update product
// @Description  Update a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id      path      string              true  "Product ID"
// @Param        request body      entity.Product      true  "Product data"
// @Success      200
// @Failure      400   {object}  Error
// @Failure      404   {object}  Error
// @Failure      500   {object}  Error
// @Router       /products/{id} [put]
// @Security     ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary      Delete product
// @Description  Delete a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Product ID"
// @Success      200
// @Failure      400   {object}  Error
// @Failure      404   {object}  Error
// @Failure      500   {object}  Error
// @Router       /products/{id} [delete]
// @Security     ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}
