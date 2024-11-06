package ads

import (
	"go-ads-management/app/middlewares"
	"go-ads-management/businesses/ads"
	"go-ads-management/controllers"
	"go-ads-management/controllers/ads/request"
	"go-ads-management/controllers/ads/response"
	"go-ads-management/models"
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
	ads, err := ac.adsUseCase.GetAll()

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch ads", "")
	}

	pg := paginate.New()

	result := pg.With(ads).Request(c.Request()).Response(&[]models.Ads{})

	return controllers.NewResponse(c, http.StatusOK, "success", "all ads", result)
}

func (ac *AdsController) GetByID(c echo.Context) error {
	adsID := c.Param("id")

	ads, err := ac.adsUseCase.GetByID(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "ad not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad found", response.FromDomain(ads))
}

func (ac *AdsController) Create(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", "")
	}

	adsReq := request.Ads{}

	if err := c.Bind(&adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	err = adsReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	adsReq.UserID = uint(claim.ID)

	ads, err := ac.adsUseCase.Create(adsReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create an ad", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "ad created", response.FromDomain(ads))
}

func (ac *AdsController) Update(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", "")
	}

	adsID := c.Param("id")

	adsReq := request.Ads{}

	isVerified := ac.verifyAdsOwner(adsID, uint(claim.ID))

	if !isVerified {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "operation not permitted", "")
	}

	if err := c.Bind(&adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	err = adsReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	ads, err := ac.adsUseCase.Update(adsReq.ToDomain(), adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad updated", response.FromDomain(ads))
}

func (ac *AdsController) Delete(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", "")
	}

	adsID := c.Param("id")

	isVerified := ac.verifyAdsOwner(adsID, uint(claim.ID))

	if !isVerified {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "operation not permitted", "")
	}

	err = ac.adsUseCase.Delete(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad deleted", "")
}

func (ac *AdsController) Restore(c echo.Context) error {
	adsID := c.Param("id")

	ads, err := ac.adsUseCase.Restore(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to restore an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad restored", response.FromDomain(ads))
}

func (ac *AdsController) ForceDelete(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", "")
	}

	adsID := c.Param("id")

	isVerified := ac.verifyAdsOwner(adsID, uint(claim.ID))

	if !isVerified {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "operation not permitted", "")
	}

	err = ac.adsUseCase.ForceDelete(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad deleted", "")
}

func (ac *AdsController) verifyAdsOwner(adsID string, userID uint) bool {
	ads, err := ac.adsUseCase.GetByID(adsID)

	if err != nil {
		return false
	}

	return ads.UserID == userID
}
