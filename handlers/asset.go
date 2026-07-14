package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
	"github.com/vaibhav-ch123/asset-management/services"
	"github.com/vaibhav-ch123/asset-management/utils"
)

func CreateAsset(w http.ResponseWriter, r *http.Request) {

	var body models.CreateAssetRequest

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.ResponseError(w, http.StatusBadRequest, parseErr, "failed to parse request body!")
		return
	}

	body.AssetName = strings.Trim(body.AssetName, " ")
	body.AssetType = strings.Trim(body.AssetType, " ")
	body.AssetBrand = strings.Trim(body.AssetBrand, " ")
	body.SerialNumber = strings.Trim(body.SerialNumber, " ")
	body.PurchaseDate = strings.Trim(body.PurchaseDate, " ")
	body.WarrantyExpiry = strings.Trim(body.WarrantyExpiry, " ")
	body.AssetStatus = strings.Trim(body.AssetStatus, " ")

	if body.AssetName == "" && body.AssetType == "" && body.AssetBrand == "" && body.SerialNumber == "" && body.AssetStatus == "" {
		utils.ResponseError(w, http.StatusBadRequest, nil, "all field must required!")
		return
	}

	if len(body.AssetName) > 30 {
		utils.ResponseError(w, http.StatusBadRequest, nil, "asset name must less than 30")
		return
	}

	if len(body.SerialNumber) > 30 {
		utils.ResponseError(w, http.StatusBadRequest, nil, "serial number must less than 30")
		return
	}

	validAssetTypes := map[string]bool{
		"laptop":   true,
		"monitor":  true,
		"mouse":    true,
		"keyboard": true,
	}

	if !validAssetTypes[body.AssetType] {
		utils.ResponseError(w, http.StatusBadRequest, nil, "invalid assetType")
		return
	}

	validBrands := map[string]bool{
		"dell":     true,
		"hp":       true,
		"lenovo":   true,
		"apple":    true,
		"logitech": true,
	}

	if !validBrands[body.AssetBrand] {
		utils.ResponseError(w, http.StatusBadRequest, nil, "invaild assetbrand")
		return
	}

	validStatuses := map[string]bool{
		"available": true,
		"assigned":  true,
		"repair":    true,
		"service":   true,
		"damaged":   true,
	}

	if !validStatuses[body.AssetStatus] {
		utils.ResponseError(w, http.StatusBadRequest, nil, "invalid assetStatus")
		return
	}

	if body.PurchaseDate != "" {
		_, err := time.Parse("2006-01-02", body.PurchaseDate)
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, nil, "purchaseDate must be in YYYY-MM-DD format")
			return
		}
	}

	if body.WarrantyExpiry != "" {
		_, err := time.Parse("2006-01-02", body.WarrantyExpiry)
		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, nil, "warrantyExpiry must be in YYYY-MM-DD format")
			return
		}
	}

	assetID, err := services.CreateAsset(body)

	if err != nil {
		switch err {
		case errors.ErrAssetSerialNumberMatch:
			utils.ResponseError(w, http.StatusBadRequest, err, "Serial number already exists")
		default:
			utils.ResponseError(w, http.StatusInternalServerError, err, "failed to create asset!")
		}
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, assetID)
}

func GetAsset(w http.ResponseWriter, r *http.Request) {

	// assetID := r.PathValue("id")


}