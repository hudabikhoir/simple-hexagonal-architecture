package content

import (
	"clean-hexa/api/v2/content/request"
	"clean-hexa/api/v2/content/response"
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

func (controller *Controller) GetAll(c echo.Context) error {
	contents, err := controller.service.GetContents()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := response.NewGetContents(contents)

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Create(c echo.Context) error {
	createContentRequest := new(request.CreateContentRequestV2)
	if err := c.Bind(createContentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("createContentRequest v2: ", createContentRequest)

	req := *createContentRequest.ToSpec()

	err := controller.service.CreateContent(req)
	if err != nil {
		fmt.Println("masuk err sini")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
