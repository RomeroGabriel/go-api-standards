package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RomeroGabriel/go-api-standards/internal/dto"
	"github.com/RomeroGabriel/go-api-standards/internal/entity"
	"github.com/RomeroGabriel/go-api-standards/internal/infra/db"
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
