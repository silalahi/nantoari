package files

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPHandler contain HTTP handler for File
type HTTPHandler struct {
	u Usecase
}

func NewHTTPHandler(e *echo.Echo, u Usecase) {
	h := HTTPHandler{
		u: u,
	}

	e.GET("/:uuid", h.Get)
	e.POST("/", h.Set)
}

// Get ...
func (h *HTTPHandler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from HTTPHandler@GET")
}

func (h *HTTPHandler) Set(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from HTTPHandler@POST")
}
