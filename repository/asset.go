package repository

import (
	"database/sql"

	"github.com/vaibhav-ch123/asset-management/database"
	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
)

func CreateAsset(asset models.CreateAssetRequest) (string, error) {

	SQL := `INSERT INTO assets (asset_name, asset_type, asset_brand, serial_number, purchase_date, warranty_expiry, asset_status, charger_available)
	        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var assetID string

	if err := database.AssetDB.Get(&assetID, SQL, asset.AssetName, asset.AssetType, asset.AssetBrand, asset.SerialNumber, asset.PurchaseDate, asset.WarrantyExpiry, asset.AssetStatus, asset.ChargerAvailable); err != nil {
		return "", err
	}

	return assetID, nil
}

func SerialExist(serial_number string) (bool, error) {

	SQL := `SELECT id FROM assets WHERE serial_number = $1 AND archived_at IS NULL`

	var id string
	err := database.AssetDB.Get(&id, SQL, serial_number)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return true, errors.ErrAssetSerialNumberMatch
	}

	return false, nil
}

func CreateLaptopSpec(laptop models.LaptopDetail) error {

	SQL := `INSERT INTO laptop_specs (assest_id, ram_gb, storage_gb, operating_system, screen_resolution, processor)
	        VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := database.AssetDB.Exec(SQL, laptop.AssetID, laptop.RAM, laptop.Storage, laptop.OperatingSystem, laptop.ScreenResolution, laptop.Processor)

	return err
}

func CreateMonitorSpec(monitor models.MonitorDetail) error {

	SQL := `INSERT INTO monitor_specs
			(asset_id, screen_size, screen_resolution)
			VALUES ($1, $2, $3)`

	_, err := database.AssetDB.Exec(
		SQL,
		monitor.AssetID,
		monitor.ScreenSize,
		monitor.ScreenResolution,
	)

	return err
}

func CreateMouseSpec(mouse models.MouseDetail) error {

	SQL := `INSERT INTO mouse_specs
			(asset_id, wireless)
			VALUES ($1, $2)`

	_, err := database.AssetDB.Exec(
		SQL,
		mouse.AssetID,
		mouse.Wireless,
	)

	return err
}

func CreateKeyboardSpec(keyboard models.KeyboardDetail) error {

	SQL := `INSERT INTO keyboard_specs
			(asset_id, wireless)
			VALUES ($1, $2)`

	_, err := database.AssetDB.Exec(
		SQL,
		keyboard.AssetID,
		keyboard.Wireless,
	)

	return err
}

func CreatePhoneSpec(phone models.PhoneDetail) error {

	SQL := `INSERT INTO phone_specs
			(asset_id, ram_gb, storage_gb, operating_system)
			VALUES ($1, $2, $3, $4)`

	_, err := database.AssetDB.Exec(
		SQL,
		phone.AssetID,
		phone.RAM,
		phone.Storage,
		phone.OperatingSystem,
	)

	return err
}
