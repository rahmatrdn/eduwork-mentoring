package helper

import (
	"encoding/json"
	"net/http"
)

// Response adalah struct untuk format response standar
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  []string   `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseWithPagination adalah struct untuk response dengan pagination
type ResponseWithPagination struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Errors     []string   `json:"errors,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination adalah struct untuk informasi pagination
type Pagination struct {
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
	PerPage      int   `json:"per_page"`
}

// WriteJSON adalah fungsi helper untuk menulis response JSON
func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Status:  false,
			Message: "Error marshalling JSON",
			Errors:  []string{err.Error()},
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// SuccessResponse mengembalikan response sukses
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse mengembalikan response error
func ErrorResponse(message string, errors []string) Response {
	return Response{
		Status:  false,
		Message: message,
		Errors:  errors,
	}
}

// PaginationResponse mengembalikan response dengan pagination
func PaginationResponse(message string, data interface{}, pagination *Pagination) ResponseWithPagination {
	return ResponseWithPagination{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}

// BuildPagination membuat struct Pagination
func BuildPagination(currentPage, perPage int, totalRecords int64) *Pagination {
	totalPages := int(totalRecords) / perPage
	if int(totalRecords)%perPage > 0 {
		totalPages++
	}

	return &Pagination{
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
		PerPage:      perPage,
	}
}
