package oauth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	domain "github.com/miyayamamasaru/go-idp/pkg/domain/interfaces"
)

type AuthorizeRequest struct {
	ResponseType string `query:"response_type"`
	ClientID     string `query:"client_id"`
	RedirectURI  string `query:"redirect_uri"`
	Scope        string `query:"scope"`
	State        string `query:"state"`
}

type AuthorizeHandlerInterface interface {
	Execute(c *echo.Context) error
}

type AuthorizeHandler struct {
	ClientRepository domain.OauthClientRepositoryInterface
}

func NewAuthorizeHandler(clientRepository domain.OauthClientRepositoryInterface) AuthorizeHandlerInterface {
	return &AuthorizeHandler{ClientRepository: clientRepository}
}

func (h AuthorizeHandler) Execute(c *echo.Context) error {
	var req AuthorizeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid_request",
		})
	}

	if req.ResponseType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":             "invalid_request",
			"error_description": "response_type is required",
		})
	}

	if req.ResponseType != "code" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":             "unsupported_response_type",
			"error_description": "only 'code' response type is supported",
		})
	}

	if req.ClientID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":             "invalid_request",
			"error_description": "client_id is required",
		})
	}

	client, err := h.ClientRepository.FindByClientId(req.ClientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "server_error",
		})
	}
	if client == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":             "invalid_request",
			"error_description": "unknown client_id",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"response_type": req.ResponseType,
		"client_id":     req.ClientID,
		"redirect_uri":  req.RedirectURI,
		"scope":         req.Scope,
		"state":         req.State,
	})
}
