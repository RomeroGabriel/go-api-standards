package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RomeroGabriel/go-api-standards/internal/dto"
	"github.com/RomeroGabriel/go-api-standards/internal/entity"
	"github.com/RomeroGabriel/go-api-standards/internal/infra/db"
	pkgEntity "github.com/RomeroGabriel/go-api-standards/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB db.ProductInterface
}

func NewProductHandler(db db.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary      Create product
// @Description  Create products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		handlerBadResquest(w, err)
		return
	}

	p, err := entity.NewProduct(productDto.Name, productDto.Price)
	if err != nil {
		handlerBadResquest(w, err)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary      Get a product
// @Description  Get a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(uuid)
// @Success      200  {object}  entity.Product
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetByIDProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "product ID" Format(uuid)
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.UpdateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		handlerBadResquest(w, err)
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idDomain, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	product.ID = idDomain
	product.Name = productDto.Name
	product.Price = productDto.Price
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        path      string                  true  "product ID" Format(uuid)
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

// ListAccounts godoc
// @Summary      List products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "page number"
// @Param        limit     query     string  false  "limit"
// @Success      200       {array}   entity.Product
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	sort := chi.URLParam(r, "sort")
	if sort != "" {
		sort = "asc"
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
