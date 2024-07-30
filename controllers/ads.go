package controllers

import (
	"go-ads-management/middlewares"
	"go-ads-management/models"
	"go-ads-management/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdsController struct {
	service services.AdsService
}

func InitAdsController() AdsController {
	return AdsController{
		service: services.InitAdsService(),
	}
}

func (cc *AdsController) GetAll(c echo.Context) error {
	ads, err := cc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch ad data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Ads]{
		Status:  "success",
		Message: "all ads",
		Data:    ads,
	})
}

func (cc *AdsController) GetByID(c echo.Context) error {
	adsID := c.Param("id")

	ads, err := cc.service.GetByID(adsID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "ads not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Ads]{
		Status:  "success",
		Message: "ads found",
		Data:    ads,
	})
}

func (cc *AdsController) Create(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	var adsInput models.AdsInput

	if err := c.Bind(&adsInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "request invalid",
		})
	}

	err = adsInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	adsInput.UserID = uint(claim.ID)

	ads, err := cc.service.Create(adsInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create an ad",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Ads]{
		Status:  "success",
		Message: "ad created",
		Data:    ads,
	})
}

func (cc *AdsController) Update(c echo.Context) error {
	adsID := c.Param("id")

	var adsInput models.AdsInput

	if err := c.Bind(&adsInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := adsInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	ads, err := cc.service.Update(adsInput, adsID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update ad",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Ads]{
		Status:  "success",
		Message: "ad updated",
		Data:    ads,
	})
}

func (cc *AdsController) Delete(c echo.Context) error {
	adsID := c.Param("id")

	err := cc.service.Delete(adsID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete ad",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "ad deleted",
	})
}
