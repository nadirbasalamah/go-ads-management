package categories

import (
	"go-ads-management/businesses/categories"
	"go-ads-management/controllers"
	"go-ads-management/controllers/categories/request"
	"go-ads-management/controllers/categories/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.UseCase
}

func NewCategoryController(categoryUC categories.UseCase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	categoryRecords, err := cc.categoryUseCase.GetAll(c.Request().Context())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch categories", "")
	}

	categories := []response.Category{}

	for _, category := range categoryRecords {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	categoryID := c.Param("id")

	cID, err := strconv.Atoi(categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	category, err := cc.categoryUseCase.GetByID(c.Request().Context(), cID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category found", response.FromDomain(category))
}

func (cc *CategoryController) Create(c echo.Context) error {
	categoryReq := request.Category{}

	if err := c.Bind(&categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := c.Validate(categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	category, err := cc.categoryUseCase.Create(c.Request().Context(), categoryReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a category", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (cc *CategoryController) Update(c echo.Context) error {
	categoryReq := request.Category{}

	if err := c.Bind(&categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	categoryID := c.Param("id")

	cID, err := strconv.Atoi(categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	if err := c.Validate(categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	category, err := cc.categoryUseCase.Update(c.Request().Context(), categoryReq.ToDomain(), cID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(category))
}

func (cc *CategoryController) Delete(c echo.Context) error {
	categoryID := c.Param("id")

	cID, err := strconv.Atoi(categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = cc.categoryUseCase.Delete(c.Request().Context(), cID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
