package file

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

// HttpGetHandler handles GET request
func HttpGetHandler(c echo.Context) error {
	param := c.Param("uuid")
	id, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": ErrInvalidUUID.Error(),
		})
	}

	file, err := Get(id)
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
func HttpSetHandler(c echo.Context) (err error) {
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

	if err := Set(file); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res := response{
		UUID: file.UUID().String(),
	}

	return c.JSON(http.StatusOK, res)
}
