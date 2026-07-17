package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/vaibhav-ch123/asset-management/database"
	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
	"github.com/vaibhav-ch123/asset-management/repository"
)

func CreateAssetAssigned(assetAssign models.AssignRequest) error {

	assetStatus, err := repository.GetAssetStatus(assetAssign.AssetID)
	if err != nil {
		return err
	}
	if assetStatus != "available" {
		return errors.ErrAssetAlreadyAssigned
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		err = repository.CreateAssetAssign(tx, assetAssign)
		if err != nil {
			return err
		}

		err = repository.UpdateAssetStatusAssigned(tx, assetAssign.AssetID)
		if err != nil {
			return err
		}

		return nil
	})

	if txErr != nil {
		return txErr
	}

	return nil
}

func DeleteAssetAssign(assetAssign models.AssignRequest, assignedID string) error {

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		err := repository.DeleteAssetAssign(tx, assetAssign, assignedID)
		if err != nil {
			return err
		}

		err = repository.UpdateAssetStatusUnassigned(tx, assetAssign.AssetID)
		if err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		return txErr
	}

	return nil
}

func GetAssetAssigned() ([]models.AssetAssignedResponse, error) {

	assetAssigned, err := repository.GetAssignedAssets()

	return assetAssigned, err
}

func GetEmployeeAssignedAssets(employeeID string) ([]models.EmployeeAssetAssignedResponse, error) {

	assetAssigned, err := repository.GetEmployeeAssignedAssets(employeeID)

	return assetAssigned, err
}
