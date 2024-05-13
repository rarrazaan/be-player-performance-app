package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	service "github.com/rarrazaan/be-player-performance-app/internal/services"
)

type identityPerformanceHandler struct {
	identityPerformanceService service.IIdentityPerformanceService
	validator                  *validator.Validate
}

func NewIdentityPerformanceHandler(
	identityPerformanceService service.IIdentityPerformanceService,
	validator *validator.Validate,
) *identityPerformanceHandler {
	return &identityPerformanceHandler{
		identityPerformanceService: identityPerformanceService,
		validator:                  validator,
	}
}

//	@Summary		CalculatePerformance Function
//	@Description	Get Performance result of Diki based on input (N, M, A array, B array)
//	@ID				CalculatePerformance
//	@Tags			CalculatePerformance V1
//	@Produce		json
//	@Param			Cookie	header		string	true	"access_token={jwt_token}"
//	@Param			n		query		int		true	"length of array A and B"
//	@Param			m		query		int		true	"proficiency level of Diki at first"
//	@Param			a		query		string	true	"N number of Diki's opponent's proficiency level"
//	@Param			b		query		string	true	"N number of proficiency level that Diki will get if Diki can beat his opponent"
//	@Success		200		{object}	dto.JSONResponse{data=dto.PerformanceResponse}
//	@Failure		401		{object}	dto.JSONResponse
//	@Failure		400		{object}	dto.JSONResponse
//	@Failure		500		{object}	dto.JSONResponse
//	@Router			/api/calculate [get]
func (h *identityPerformanceHandler) CalculatePerformance(c *gin.Context) {
	params := new(dto.PerformanceRequest)
	if err := c.ShouldBindQuery(&params); err != nil {
		_ = c.Error(err)
		return
	}

	if err := h.validator.Struct(*params); err != nil {
		_ = c.Error(err.(validator.ValidationErrors))
		return
	}

	ctx := c.Request.Context()
	res := h.identityPerformanceService.CalculatePerformance(ctx, params)

	c.JSON(http.StatusOK, dto.JSONResponse{
		Data: res,
	})
}

//	@Summary		Identity Function
//	@Description	Get Identity of user based on firs name from query paramaters
//	@ID				Identity
//	@Tags			Identity V1
//	@Produce		json
//	@Param			Cookie	header		string	true	"access_token={jwt_token}"
//	@Param			name	query		string	true	"keyword for searched user identity"
//	@Success		200		{object}	dto.JSONResponse{data=[]dto.IdentityResponse}
//	@Failure		401		{object}	dto.JSONResponse
//	@Failure		400		{object}	dto.JSONResponse
//	@Failure		500		{object}	dto.JSONResponse
//	@Router			/api/identity [get]
func (h *identityPerformanceHandler) Identity(c *gin.Context) {
	params := new(dto.IdentityRequest)
	if err := c.ShouldBindQuery(&params); err != nil {
		_ = c.Error(err)
		return
	}

	if err := h.validator.Struct(*params); err != nil {
		_ = c.Error(err.(validator.ValidationErrors))
		return
	}

	ctx := c.Request.Context()
	res, err := h.identityPerformanceService.Identity(ctx, params)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JSONResponse{
		Data: res,
	})
}
