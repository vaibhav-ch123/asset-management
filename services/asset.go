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
		assetID1, assestErr := repository.CreateAsset(tx, asset)
		if assestErr != nil {
			return assestErr
		}
		assetID = assetID1

		var specErr error
		switch asset.AssetType {
		case "laptop":
			asset.Laptop.AssetID = assetID
			specErr = repository.CreateLaptopSpec(tx, asset.Laptop)
		case "monitor":
			asset.Monitor.AssetID = assetID
			specErr = repository.CreateMonitorSpec(tx, asset.Monitor)
		case "mouse":
			asset.Mouse.AssetID = assetID
			specErr = repository.CreateMouseSpec(tx, asset.Mouse)
		case "keyboard":
			asset.Keyboard.AssetID = assetID
			specErr = repository.CreateKeyboardSpec(tx, asset.Keyboard)
		case "phone":
			asset.Phone.AssetID = assetID
			specErr = repository.CreatePhoneSpec(tx, asset.Phone)
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

func GetAsset(assetID string) (models.AssetResponse, error) {

	var assetWithSpec models.AssetResponse
	var laptop models.LaptopDetail
	var monitor models.MonitorDetail
	var keyboard models.KeyboardDetail
	var mouse models.MouseDetail
	var phone models.PhoneDetail
	var asset models.Asset

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		var assetErr error
		asset, assetErr = repository.GetAssetByID(assetID)
		if assetErr != nil {
			return assetErr
		}

		var specErr error
		switch asset.AssetType {
		case "laptop":
			laptop, specErr = repository.GetLaptopSpecByAssetID(assetID)
		case "monitor":
			monitor, specErr = repository.GetMonitorSpecByAssetID(assetID)
		case "mouse":
			mouse, specErr = repository.GetMouseSpecByAssetID(assetID)
		case "keyboard":
			keyboard, specErr = repository.GetKeyboardSpecByAssetID(assetID)
		case "phone":
			phone, specErr = repository.GetPhoneSpecByAssetID(assetID)
		}

		if specErr != nil {
			return specErr
		}

		return nil
	})

	if txErr != nil {
		return assetWithSpec, txErr
	}

	if asset.ID != "" {
		assetWithSpec.ID = assetID
		assetWithSpec.AssetName = asset.AssetName
		assetWithSpec.AssetBrand = asset.AssetBrand
		assetWithSpec.AssetStatus = asset.AssetStatus
		assetWithSpec.AssetType = asset.AssetType
		assetWithSpec.SerialNumber = asset.SerialNumber
		assetWithSpec.PurchaseDate = asset.PurchaseDate.Format("2006-01-01")
		assetWithSpec.ChargerAvailable = asset.ChargerAvailable
		assetWithSpec.WarrantyExpiry = asset.WarrantyExpiry.Format("2006-01-01")
	}

	if laptop.ID != "" {
		assetWithSpec.Laptop = &laptop
	}

	if monitor.ID != "" {
		assetWithSpec.Monitor = &monitor
	}

	if mouse.ID != "" {
		assetWithSpec.Mouse = &mouse
	}

	if keyboard.ID != "" {
		assetWithSpec.Keyboard = &keyboard
	}

	if phone.ID != "" {
		assetWithSpec.Phone = &phone
	}

	return assetWithSpec, nil
}

func GetAssets() ([]models.AssetResponse, error) {

	assetSpec, err := repository.GetAssetsWithSpec()
	if err != nil {
		return nil, err
	}

	assets := make([]models.AssetResponse, 0, len(assetSpec))

	for _, a := range assetSpec {

		resp := models.AssetResponse{
			ID:               a.ID,
			AssetName:        a.AssetName,
			AssetBrand:       a.AssetBrand,
			AssetStatus:      a.AssetStatus,
			AssetType:        a.AssetType,
			SerialNumber:     a.SerialNumber,
			ChargerAvailable: a.ChargerAvailable,
		}

		if a.PurchaseDate != nil {
			resp.PurchaseDate = a.PurchaseDate.Format("2006-01-02")
		}

		if a.WarrantyExpiry != nil {
			resp.WarrantyExpiry = a.WarrantyExpiry.Format("2006-01-02")
		}

		// Laptop
		if a.LaptopID != nil {
			resp.Laptop = &models.LaptopDetail{
				ID:      *a.LaptopID,
				AssetID: a.ID,
			}

			if a.RAM != nil {
				resp.Laptop.RAM = *a.RAM
			}
			if a.Storage != nil {
				resp.Laptop.Storage = *a.Storage
			}
			if a.LaptopOS != nil {
				resp.Laptop.OperatingSystem = *a.LaptopOS
			}
			if a.ScreenResolution != nil {
				resp.Laptop.ScreenResolution = *a.ScreenResolution
			}
			if a.Processor != nil {
				resp.Laptop.Processor = *a.Processor
			}
		}

		// Monitor
		if a.MonitorID != nil {
			resp.Monitor = &models.MonitorDetail{
				ID:      *a.MonitorID,
				AssetID: a.ID,
			}

			if a.MonitorScreenSize != nil {
				resp.Monitor.ScreenSize = *a.MonitorScreenSize
			}
			if a.MonitorScreenResolution != nil {
				resp.Monitor.ScreenResolution = *a.MonitorScreenResolution
			}
		}

		// Mouse
		if a.MouseID != nil {
			resp.Mouse = &models.MouseDetail{
				ID:      *a.MouseID,
				AssetID: a.ID,
			}

			if a.MouseWireless != nil {
				resp.Mouse.Wireless = *a.MouseWireless
			}
		}

		// Keyboard
		if a.KeyboardID != nil {
			resp.Keyboard = &models.KeyboardDetail{
				ID:      *a.KeyboardID,
				AssetID: a.ID,
			}

			if a.KeyboardWireless != nil {
				resp.Keyboard.Wireless = *a.KeyboardWireless
			}
		}

		// Phone
		if a.PhoneID != nil {
			resp.Phone = &models.PhoneDetail{
				ID:      *a.PhoneID,
				AssetID: a.ID,
			}

			if a.PhoneRAM != nil {
				resp.Phone.RAM = *a.PhoneRAM
			}
			if a.PhoneStorage != nil {
				resp.Phone.Storage = *a.PhoneStorage
			}
			if a.PhoneOS != nil {
				resp.Phone.OperatingSystem = *a.PhoneOS
			}
		}

		assets = append(assets, resp)
	}

	return assets, nil
}

func UpdateAsset(assetSpec models.UpdateAssetRequest, assetID string) error {

	asset, err := repository.GetAssetByID(assetID)
	if err != nil {
		return err
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		err := repository.UpdateAssetByAssetID(tx, assetSpec, assetID)
		if err != nil {
			return err
		}

		var specErr error
		switch asset.AssetType {
		case "laptop":
			if assetSpec.Laptop != nil {
				specErr = repository.UpdateLaptopSpecByAssetID(tx, assetSpec.Laptop, assetID)
			}
		case "monitor":
			if assetSpec.Monitor != nil {
				specErr = repository.UpdateMonitorSpecByAssetID(tx, assetSpec.Monitor, assetID)
			}
		case "mouse":
			if assetSpec.Mouse != nil {
				specErr = repository.UpdateMouseSpecByAssetID(tx, assetSpec.Mouse, assetID)
			}
		case "keyboard":
			if assetSpec.Keyboard != nil {
				specErr = repository.UpdateKeyboardSpecByAssetID(tx, assetSpec.Keyboard, assetID)
			}
		case "phone":
			if assetSpec.Phone != nil {
				specErr = repository.UpdatePhoneSpecByAssetID(tx, assetSpec.Phone, assetID)
			}
		}

		if specErr != nil {
			return specErr
		}

		return nil
	})

	if txErr != nil {
		return txErr
	}

	return nil
}
