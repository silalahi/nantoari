package files

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

var v = validator.New()

type (
	// HTTP Request
	request struct {
		URL string `json:"url" form:"url" query:"url" validate:"required"`
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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": ErrInvalidUUID.Error(),
		})
	}

	file, err := h.u.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	result, err := file.Parse()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

// Set handles POST request
func (h *HTTPHandler) Set(c echo.Context) (err error) {
	req := new(request)
	if err = c.Bind(req); err != nil {
		return
	}
	if err = v.Struct(req); err != nil {
		return
	}

	file := File{
		url:  req.URL,
		uuid: uuid.New(),
	}

	if !file.IsCSV() {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": ErrInvalidFileFormat.Error(),
		})
	}

	if err := h.u.Set(file); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res := response{
		UUID: file.UUID().String(),
	}

	return c.JSON(http.StatusOK, res)
}
