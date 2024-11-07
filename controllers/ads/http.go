package ads

import (
	"go-ads-management/businesses/ads"
	"go-ads-management/controllers"
	"go-ads-management/controllers/ads/request"
	"go-ads-management/controllers/ads/response"
	adsRecord "go-ads-management/drivers/mysql/ads"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"
)

type AdsController struct {
	adsUseCase ads.UseCase
}

func NewAdsController(adsUC ads.UseCase) *AdsController {
	return &AdsController{
		adsUseCase: adsUC,
	}
}

func (ac *AdsController) GetAll(c echo.Context) error {
	ads, err := ac.adsUseCase.GetAll(c.Request().Context())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch ads", "")
	}

	pg := paginate.New()

	result := pg.With(ads).Request(c.Request()).Response(&[]adsRecord.Ads{})

	return controllers.NewResponse(c, http.StatusOK, "success", "all ads", result)
}

func (ac *AdsController) GetByID(c echo.Context) error {
	adsID := c.Param("id")

	ads, err := ac.adsUseCase.GetByID(c.Request().Context(), adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "ad not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad found", response.FromDomain(ads))
}

func (ac *AdsController) Create(c echo.Context) error {
	adsReq := request.Ads{}

	if err := c.Bind(&adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := c.Validate(adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	ads, err := ac.adsUseCase.Create(c.Request().Context(), adsReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create an ad", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "ad created", response.FromDomain(ads))
}

func (ac *AdsController) Update(c echo.Context) error {
	adsID := c.Param("id")

	adsReq := request.Ads{}

	if err := c.Bind(&adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := c.Validate(adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	ads, err := ac.adsUseCase.Update(c.Request().Context(), adsReq.ToDomain(), adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad updated", response.FromDomain(ads))
}

func (ac *AdsController) Delete(c echo.Context) error {
	adsID := c.Param("id")

	err := ac.adsUseCase.Delete(c.Request().Context(), adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad deleted", "")
}

func (ac *AdsController) Restore(c echo.Context) error {
	adsID := c.Param("id")

	ads, err := ac.adsUseCase.Restore(c.Request().Context(), adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to restore an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad restored", response.FromDomain(ads))
}

func (ac *AdsController) ForceDelete(c echo.Context) error {
	adsID := c.Param("id")

	err := ac.adsUseCase.ForceDelete(c.Request().Context(), adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad deleted", "")
}
