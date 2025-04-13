package handler

import (
	"encoding/json"
	"net/http"
	"product-management/models"
	"product-management/service"
	"strconv"

	"product-management/helper"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		response := helper.ErrorResponse("Gagal mengambil data produk", []string{err.Error()})
		helper.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}
	
	response := helper.SuccessResponse("Berhasil mengambil data produk", products)
	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response := helper.ErrorResponse("ID tidak valid", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		response := helper.ErrorResponse("Produk tidak ditemukan", []string{err.Error()})
		helper.WriteJSON(w, http.StatusNotFound, response)
		return
	}

	response := helper.SuccessResponse("Berhasil mengambil data produk", product)
	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response := helper.ErrorResponse("Data produk tidak valid", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	newProduct, err := h.service.CreateProduct(product)
	if err != nil {
		response := helper.ErrorResponse("Gagal membuat produk", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessResponse("Berhasil membuat produk", newProduct)
	helper.WriteJSON(w, http.StatusCreated, response)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response := helper.ErrorResponse("ID tidak valid", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response := helper.ErrorResponse("Data produk tidak valid", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	updatedProduct, err := h.service.UpdateProduct(uint(id), product)
	if err != nil {
		response := helper.ErrorResponse("Gagal mengupdate produk", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessResponse("Berhasil mengupdate produk", updatedProduct)
	helper.WriteJSON(w, http.StatusOK, response)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response := helper.ErrorResponse("ID tidak valid", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		response := helper.ErrorResponse("Gagal menghapus produk", []string{err.Error()})
		helper.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessResponse("Berhasil menghapus produk", nil)
	helper.WriteJSON(w, http.StatusOK, response)
}
