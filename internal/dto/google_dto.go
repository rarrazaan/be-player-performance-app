package dto

type (
	GoogleResponse struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Picture       string `json:"picture"`
	}

	LoginResponsePayload struct {
		AccessToken string `json:"access_token"`
	}
)
