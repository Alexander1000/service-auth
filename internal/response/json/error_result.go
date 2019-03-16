package json

import "net/http"

type ErrorResult struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

var (
	// ErrorUnauthorized
	ErrorUnauthorized = ErrorResult{
		Error: ErrorData{
			Message: "Unauthorized",
			Code: http.StatusUnauthorized,
		},
	}
	// ErrorForbidden
	ErrorForbidden = ErrorResult{
		Error: ErrorData{
			Message: "Forbidden",
			Code: http.StatusForbidden,
		},
	}
	// ErrorNotAllowed
	ErrorNotAllowed = ErrorResult{
		Error: ErrorData{
			Message: "Method not allowed",
			Code: http.StatusMethodNotAllowed,
		},
	}
	// ErrorInternalServerError
	ErrorInternalServerError = ErrorResult{
		Error: ErrorData{
			Message: "Internal server error",
			Code: http.StatusInternalServerError,
		},
	}
	// ErrorNotImplemented
	ErrorNotImplemented = ErrorResult{
		Error: ErrorData{
			Message: "Not implemented",
			Code: http.StatusNotImplemented,
		},
	}
	ErrorNotFound = ErrorResult{
		Error: ErrorData{
			Message: "Not found",
			Code: http.StatusNotFound,
		},
	}
	ErrorBadRequest = ErrorResult{
		Error: ErrorData{
			Message: "Bad request",
			Code: http.StatusBadRequest,
		},
	}
)
