package content

import (
	"clean-hexa/api/v2/content/response"
	contentBusiness "clean-hexa/business/content"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service contentBusiness.Service
}

func NewController(service contentBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	contents, err := controller.service.GetContents()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := response.NewGetContents(contents)

	return c.JSON(http.StatusOK, response)
}
