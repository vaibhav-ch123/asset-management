package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/vaibhav-ch123/asset-management/database"
	"github.com/vaibhav-ch123/asset-management/models"
	"github.com/vaibhav-ch123/asset-management/repository"
)

func CreateAsset(asset models.CreateAssetRequest) (string, error) {

	isSerialExist, isSerialExistErr := repository.SerialExist(asset.SerialNumber)
	if isSerialExistErr != nil || isSerialExist {
		return "", isSerialExistErr
	}
    
	var assetID string
	txErr := database.Tx(func(tx *sqlx.Tx) error {
		assetID1, assestErr := repository.CreateAsset(asset)
		if assestErr != nil {
			return assestErr
		}
		assetID = assetID1

		var specErr error
		switch asset.AssetType {
		case "laptop":       specErr = repository.CreateLaptopSpec(asset.Laptop) 
		case "monitor":      specErr = repository.CreateMonitorSpec(asset.Monitor)
		case "mouse":        specErr = repository.CreateMouseSpec(asset.Mouse)
		case "keyboard"	:    specErr = repository.CreateKeyboardSpec(asset.Keyboard)
		case "phone":        specErr = repository.CreatePhoneSpec(asset.Phone)
		}
		if specErr != nil {
			return specErr
		}

		return nil
	})

	if txErr != nil {
		return "", txErr
	}

	return assetID, nil
}
