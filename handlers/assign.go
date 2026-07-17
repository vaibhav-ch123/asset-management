package handlers

import (
	"net/http"
	"time"

	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
	"github.com/vaibhav-ch123/asset-management/repository"
	"github.com/vaibhav-ch123/asset-management/services"
	"github.com/vaibhav-ch123/asset-management/utils"
)

func AssetAssigned(w http.ResponseWriter, r *http.Request) {

	var body models.AssignRequest

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.ResponseError(w, http.StatusBadRequest, parseErr, "failed to parse request body!")
		return
	}

	if body.AssetID == "" && body.EmployeeID == "" {
		utils.ResponseError(w, http.StatusBadRequest, nil, "asseID and employeeID is required!")
		return
	}

	_, assignedDateErr := time.Parse("2006-01-02", body.AssignedDate)
	if assignedDateErr != nil {
		utils.ResponseError(w, http.StatusBadRequest, nil, "Assigned Date format is not valid!")
		return
	}

	err := services.CreateAssetAssigned(body)
	if err != nil {
		switch err {
		case errors.ErrAssetAlreadyAssigned:
			utils.ResponseError(w, http.StatusBadRequest, err, "asset already assigned!")
		default:
			utils.ResponseError(w, http.StatusInternalServerError, err, "failed to assigned asset!")
		}
		return
	}

	utils.ResponseJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "assigned successfully!",
	})

}

func AssetUnassigned(w http.ResponseWriter, r *http.Request) {

	var body models.AssignRequest

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.ResponseError(w, http.StatusBadRequest, parseErr, "failed to parse request body!")
		return
	}

	if body.AssetID == "" {
		utils.ResponseError(w, http.StatusBadRequest, nil, "asseID is required!")
		return
	}

	_, returnedDateErr := time.Parse("2006-01-02", body.ReturnedDate)
	if returnedDateErr != nil {
		utils.ResponseError(w, http.StatusBadRequest, nil, "returned Date format is not valid!")
		return
	}

	assignedID := r.PathValue("id")

	err := services.DeleteAssetAssign(body, assignedID)

	if err != nil {
		switch err {
		case errors.ErrAssetAssignedNotFound:
			utils.ResponseError(w, http.StatusBadRequest, err, "asset is not assigned!")
		default:
			utils.ResponseError(w, http.StatusInternalServerError, err, "failed to assigned asset!")
		}
		return
	}

	utils.ResponseJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "unassigned successfully!",
	})
}

func GetAssignedAssets(w http.ResponseWriter, r *http.Request) {

	assetAssigned, err := services.GetAssetAssigned()

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err, "failed to get assigned assets")
		return
	}

	utils.ResponseJSON(w, http.StatusOK, assetAssigned)
}

func GetEmployeeAssignedAssets(w http.ResponseWriter, r *http.Request) {

	employeeID := r.PathValue("id")

	assignedAsset, err := repository.GetEmployeeAssignedAssets(employeeID)

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err, "failed to get assigned assets of employee")
		return
	}

	utils.ResponseJSON(w, http.StatusOK, assignedAsset)
}
