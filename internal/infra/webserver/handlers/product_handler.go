package handlers

import (
	"encoding/json"
	"net/http"

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
