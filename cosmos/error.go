package cosmos

import "fmt"

type CosmosError struct {
	message    string `json:"message"`
	statusCode int    `json:"statusCode"`
}

func NewCosmosError(message string, statusCode int) *CosmosError {
	return &CosmosError{
		message:    message,
		statusCode: statusCode,
	}
}

// Error implements the error interface
func (e *CosmosError) Error() string {
	return fmt.Sprintf("%v, %v", e.statusCode, e.message)
}

func (e *CosmosError) StatusCode() int {
	return e.statusCode
}

func (e *CosmosError) StatusOK() bool {
	return e.statusCode == 200
}

func (e *CosmosError) Created() bool {
	return e.statusCode == 201
}

func (e *CosmosError) BadRequest() bool {
	return e.statusCode == 400
}

func (e *CosmosError) Unauthorized() bool {
	return e.statusCode == 401
}

func (e *CosmosError) Forbidden() bool {
	return e.statusCode == 403
}

func (e *CosmosError) NotFound() bool {
	return e.statusCode == 404
}

func (e *CosmosError) RequestTimeout() bool {
	return e.statusCode == 408
}

func (e *CosmosError) Conflict() bool {
	return e.statusCode == 409
}

func (e *CosmosError) EntityTooLarge() bool {
	return e.statusCode == 413
}

func (e *CosmosError) RetryWith() bool {
	return e.statusCode == 449
}
func (e *CosmosError) ServiceUnavailable() bool {
	return e.statusCode == 503
}
