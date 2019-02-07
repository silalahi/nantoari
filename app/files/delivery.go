package files

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	// HTTP Request
	request struct {
		URL string `json:"url" form:"url" query:"url"`
	}

	// HTTP Response
	response struct {
		URL  string `json:"url,omitempty"`
		UUID string `json:"uuid,omitempty" form:"uuid" query:"uuid"`
	}
)

// HTTPHandler contain HTTP handler for File
type HTTPHandler struct {
	u Usecase
}

// NewHTTPHandler returns an instance of HTTPHandler
func NewHTTPHandler(e *echo.Echo, u Usecase) {
	h := HTTPHandler{
		u: u,
	}

	e.GET("/:uuid", h.Get)
	e.POST("/", h.Set)
}

// Get handles GET request
func (h *HTTPHandler) Get(c echo.Context) error {
	param := c.Param("uuid")
	id, err := uuid.Parse(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	file, err := h.u.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// TODO: converting CSV to JSON

	res := response{
		URL: file.URL(),
	}

	return c.JSON(http.StatusOK, res)
}

// Set handles POST request
func (h *HTTPHandler) Set(c echo.Context) (err error) {
	req := new(request)
	if err = c.Bind(req); err != nil {
		return
	}

	file := File{
		url:  req.URL,
		uuid: uuid.New(),
	}

	// TODO: some validation

	if err := h.u.Set(file); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	res := response{
		UUID: file.UUID().String(),
	}

	return c.JSON(http.StatusOK, res)
}
