package router

import (
	"encoding/json"
	"net/http"
	"strings"
	"web-service/pkg/log"
)

// ResSuccess Struct
type ResSuccess struct {
	Status  bool   `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResWithData struct {
	Status  bool   `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// ResError Struct
type ResError struct {
	Status  bool   `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// ResponseWrite Function
func ResponseWrite(w http.ResponseWriter, responseCode int, responseData any) {
	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	// Write JSON to Response
	json.NewEncoder(w).Encode(responseData)
}

// ResponseSuccess Function
func ResponseSuccess(w http.ResponseWriter, statusCode, message string) {
	var response ResSuccess

	// Set Default Message
	if len(message) == 0 {
		message = "Success"
	}

	// Set Response Data
	response.Status = true
	response.Code = statusCode
	response.Message = message

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusOK, response)
}

// ResponseSuccess Function With any data type
func ResponseSuccessWithData(w http.ResponseWriter, statusCode, message string, data ...any) {
	var responseData any
	if len(data) == 1 {
		responseData = data[0]
	} else {
		responseData = data
	}
	var response ResWithData

	// Set Default Message
	if len(message) == 0 {
		message = "Success"
	}

	if len(statusCode) == 0 {
		statusCode = "200"
	}

	// Set Response Data
	response.Status = true
	response.Code = statusCode
	response.Message = message
	response.Data = responseData

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusOK, response)
}

func ResponseCreatedWithData(w http.ResponseWriter, statusCode, message string, data ...any) {
	var responseData any
	if len(data) == 1 {
		responseData = data[0]
	} else {
		responseData = data
	}
	var response ResWithData

	// Set Default Message
	if len(message) == 0 {
		message = "Created successfully"
	}

	// Set Response Data
	response.Status = true
	response.Code = statusCode
	response.Message = message
	response.Data = responseData

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusCreated, response)
}

// ResponseCreated Function
func ResponseCreated(w http.ResponseWriter, statusCode string) {
	var response ResSuccess

	// Set Response Data
	response.Status = true
	response.Code = statusCode
	response.Message = "Created"

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusCreated, response)
}

// ResponseUpdated Function
func ResponseUpdated(w http.ResponseWriter, statusCode string) {
	var response ResSuccess

	// Set Response Data
	response.Status = true
	response.Code = statusCode
	response.Message = "Updated"

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusOK, response)
}

// ResponseNoContent Function
func ResponseNoContent(w http.ResponseWriter) {
	w.WriteHeader(204)
}

// ResponseNotFound Function
func ResponseNotFound(w http.ResponseWriter, message string) {
	var response ResError

	// Set Default Message
	if len(message) == 0 {
		message = "Not Found"
	}

	// Set Response Data
	response.Status = false
	response.Code = "B.ALL.404.C1"
	response.Message = "Not Found"
	response.Error = message

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusNotFound, response)
}

// ResponseMethodNotAllowed Function
func ResponseMethodNotAllowed(w http.ResponseWriter, message string) {
	var response ResError

	// Set Default Message
	if len(message) == 0 {
		message = "Method Not Allowed"
	}

	// Set Response Data
	response.Status = false
	response.Code = "B.ALL.405.C3"
	response.Message = "Method Not Allowed"
	response.Error = message

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusMethodNotAllowed, response)
}

// ResponseBadRequest Function
func ResponseBadRequest(w http.ResponseWriter, statusCode, message string) {
	var response ResError

	// Set Default Message
	if len(message) == 0 {
		message = "Bad Request"
	}

	// Set Response Data
	response.Status = false
	response.Code = statusCode
	response.Message = "Bad Request"
	response.Error = message

	// Logging Error
	log.Println(log.LogLevelError, "http-access", strings.ToLower(message))

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusBadRequest, response)
}

// ResponseInternalError Function
func ResponseInternalError(w http.ResponseWriter, message string) {
	var response ResError

	// Set Default Message
	if len(message) == 0 {
		message = "Internal Server Error"
	}

	// Set Response Data
	response.Status = false
	response.Code = "B.ALL.500.C5"
	response.Message = "Internal Server Error"
	response.Error = message

	// Logging Error
	log.Println(log.LogLevelError, "http-access", strings.ToLower(message))

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusInternalServerError, response)
}

// ResponseBadGateway Function
func ResponseBadGateway(w http.ResponseWriter, message string) {
	var response ResError

	// Set Default Message
	if len(message) == 0 {
		message = "Bad Gateway"
	}

	// Set Response Data
	response.Status = false
	response.Code = "B.ALL.502.C2"
	response.Message = "Bad Gateway"
	response.Error = message

	// Logging Error
	log.Println(log.LogLevelError, "http-access", strings.ToLower(message))

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusBadGateway, response)
}

// ResponseUnauthorized Function
func ResponseUnauthorized(w http.ResponseWriter, message string) {
	var response ResError

	// Set Default Message
	if len(message) == 0 {
		message = "Unauthorized"
	}

	// Set Response Data
	response.Status = false
	response.Code = "B.ALL.401.C4"
	response.Message = "Unauthorized"
	response.Error = message

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusUnauthorized, response)
}
