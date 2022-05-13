package content

import (
	"clean-hexa/api/v1/content/request"
	contentBusiness "clean-hexa/business/content"
	"fmt"
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

// func (controller *Controller) GetByID(c echo.Context) error {
// return Controller.service.GetContentByID(id)
// }

func (controller *Controller) GetAll(c echo.Context) error {
	contents, err := controller.service.GetContents()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, contents)
}

// Create godoc
// @Summary Create content
// @description create content with data
// @tags content
// @Accept json
// @Produce json
// @Success 200 {object} content.Content
// @Router /v1/content [post]
func (controller *Controller) Create(c echo.Context) error {
	createContentRequest := new(request.CreateContentRequest)
	if err := c.Bind(createContentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("createContentRequest v1: ", createContentRequest)

	req := *createContentRequest.ToSpec()

	err := controller.service.CreateContent(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// controller.service.CreateContent
	return c.JSON(http.StatusOK, "")
}
