package ads

import (
	"go-ads-management/businesses/ads"
	"go-ads-management/controllers"
	"go-ads-management/controllers/ads/request"
	"go-ads-management/controllers/ads/response"
	adsRecord "go-ads-management/drivers/mysql/ads"
	"go-ads-management/drivers/pinata"
	"go-ads-management/utils"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

	aID, err := strconv.Atoi(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	ads, err := ac.adsUseCase.GetByID(c.Request().Context(), aID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "ad not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad found", response.FromDomain(ads))
}

func (ac *AdsController) GetByCategory(c echo.Context) error {
	categoryID := c.Param("category_id")

	cID, err := strconv.Atoi(categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	ads, err := ac.adsUseCase.GetByCategory(c.Request().Context(), cID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch ads", "")
	}

	pg := paginate.New()

	result := pg.With(ads).Request(c.Request()).Response(&[]adsRecord.Ads{})

	return controllers.NewResponse(c, http.StatusOK, "success", "all ads", result)
}

func (ac *AdsController) GetByUser(c echo.Context) error {
	ads, err := ac.adsUseCase.GetByUser(c.Request().Context())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch ads", "")
	}

	pg := paginate.New()

	result := pg.With(ads).Request(c.Request()).Response(&[]adsRecord.Ads{})

	return controllers.NewResponse(c, http.StatusOK, "success", "all ads", result)
}

func (ac *AdsController) GetTrashed(c echo.Context) error {
	ads, err := ac.adsUseCase.GetTrashed(c.Request().Context())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch ads", "")
	}

	pg := paginate.New()

	result := pg.With(ads).Request(c.Request()).Response(&[]adsRecord.Ads{})

	return controllers.NewResponse(c, http.StatusOK, "success", "all trashed ads", result)
}

func (ac *AdsController) Create(c echo.Context) error {
	adsReq := request.Ads{}

	if err := c.Bind(&adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := c.Validate(adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "file not found", "")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))

	isFileValid := utils.ValidateFile(ext)

	if !isFileValid {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "invalid file format", "")
	}

	res, err := pinata.UploadFile(file)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "upload failed", "")
	}

	adsReq.MediaURL = res.SignedURL
	adsReq.MediaCID = res.FileCID
	adsReq.MediaID = res.FileID

	ads, err := ac.adsUseCase.Create(c.Request().Context(), adsReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create an ad", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "ad created", response.FromDomain(ads))
}

func (ac *AdsController) Update(c echo.Context) error {
	adsID := c.Param("id")

	aID, err := strconv.Atoi(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	adsReq := request.Ads{}

	adsData, err := ac.adsUseCase.GetByID(c.Request().Context(), aID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "ad not found", "")
	}

	if err := c.Bind(&adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := c.Validate(adsReq); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	file, err := c.FormFile("file")
	if err == nil { // file is provided
		ext := strings.ToLower(filepath.Ext(file.Filename))

		isFileValid := utils.ValidateFile(ext)

		if !isFileValid {
			return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "invalid file format", "")
		}

		res, err := pinata.UploadFile(file)
		if err != nil {
			return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "file upload failed", "")
		}

		if adsData.MediaID != "" {
			err = pinata.DeleteFile(adsData.MediaID)
			if err != nil {
				return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete old file", "")
			}
		}

		adsReq.MediaURL = res.SignedURL
		adsReq.MediaCID = res.FileCID
		adsReq.MediaID = res.FileID
	} else {
		// No new file provided, retain the old file data
		adsReq.MediaURL = adsData.MediaURL
		adsReq.MediaCID = adsData.MediaCID
		adsReq.MediaID = adsData.MediaID
	}

	ads, err := ac.adsUseCase.Update(c.Request().Context(), adsReq.ToDomain(), aID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad updated", response.FromDomain(ads))
}

func (ac *AdsController) Delete(c echo.Context) error {
	adsID := c.Param("id")

	aID, err := strconv.Atoi(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = ac.adsUseCase.Delete(c.Request().Context(), aID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad deleted", "")
}

func (ac *AdsController) Restore(c echo.Context) error {
	adsID := c.Param("id")

	aID, err := strconv.Atoi(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	ads, err := ac.adsUseCase.Restore(c.Request().Context(), aID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to restore an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad restored", response.FromDomain(ads))
}

func (ac *AdsController) ForceDelete(c echo.Context) error {
	adsID := c.Param("id")

	aID, err := strconv.Atoi(adsID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	ads, err := ac.adsUseCase.GetByID(c.Request().Context(), aID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "content not found", "")
	}

	err = pinata.DeleteFile(ads.MediaID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to remove the media", "")
	}

	err = ac.adsUseCase.ForceDelete(c.Request().Context(), aID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete an ad", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "ad deleted", "")
}
