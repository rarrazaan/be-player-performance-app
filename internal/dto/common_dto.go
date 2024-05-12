package dto

type (
	JSONResponse struct {
		Data    any    `json:"data"`
		Message string `json:"message"`
	}

	PerformanceRequest struct {
		N int    `form:"n" validate:"required,gt=0"`
		M int    `form:"m" validate:"required"`
		A string `form:"a" validate:"required"`
		B string `form:"b" validate:"required"`
	}

	PerformanceResponse struct {
		Result int `json:"result"`
	}

	IdentityRequest struct {
		Name string `form:"name" validate:"required"`
	}

	IdentityResponse struct {
		FullName    string `json:"full_name"`
		Age         int    `json:"age"`
		Gender      string `json:"gender"`
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number"`
	}
)
