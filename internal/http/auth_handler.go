package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	service "github.com/rarrazaan/be-player-performance-app/internal/services"
	"github.com/rarrazaan/be-player-performance-app/internal/utils"
	"github.com/rarrazaan/be-player-performance-app/internal/utils/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authHandler struct {
	authService service.IAuthService
	config      config.Config
}

func NewAuthHandler(
	authService service.IAuthService,
	config config.Config,
) *authHandler {
	return &authHandler{
		authService: authService,
		config:      config,
	}
}

func (h *authHandler) GoogleLogin(c *gin.Context) {
	URL, err := url.Parse(google.Endpoint.AuthURL)
	if err != nil {
		_ = c.Error(err)
		return
	}
	parameters := url.Values{}
	parameters.Add("client_id", h.config.GOauth.ClientID)
	parameters.Add("scope", "https://www.googleapis.com/auth/userinfo.email")
	parameters.Add("redirect_uri", h.config.GOauth.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", "orenlite-state")
	URL.RawQuery = parameters.Encode()
	url := URL.String()

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *authHandler) GoogleLoginCallback(c *gin.Context) {
	oauthConfGl := &oauth2.Config{
		ClientID:     h.config.GOauth.ClientID,
		ClientSecret: h.config.GOauth.ClientSecret,
		RedirectURL:  h.config.GOauth.RedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	state := c.Query("state")
	if state != "orenlite-state" {
		_ = c.Error(errors.ErrInternalServer)
		return
	}

	code := c.Query("code")
	if code == "" {
		reason := c.Query("error_reason")
		if reason == "user_denied" {
			_ = c.Error(errors.ErrInternalServer)
			return
		}
	} else {
		token, err := oauthConfGl.Exchange(c, code)
		if err != nil {
			_ = c.Error(err)
			return
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			_ = c.Error(err)
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			_ = c.Error(err)
			return
		}
		respStruct := new(dto.GoogleResponse)
		err = json.Unmarshal(response, respStruct)
		if err != nil {
			_ = c.Error(err)
			return
		}
		ctx := c.Request.Context()
		loginResponse, err := h.authService.LoginWithGoogle(ctx, respStruct)
		if err != nil {
			_ = c.Error(err)
			return
		}
		utils.SetCookieAfterLogin(c, h.config, loginResponse.AccessToken)
		c.Redirect(http.StatusTemporaryRedirect, h.config.GOauth.Redirect)
	}
}
