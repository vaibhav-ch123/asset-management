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

func SerialExist(serialNumber string) (bool, error) {

	SQL := `SELECT id FROM assets WHERE serial_number = $1 AND archived_at IS NULL`

	var id string
	err := database.AssetDB.Get(&id, SQL, serialNumber)

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

func GetAssetByID(assetID string) (models.Asset, error) {

	SQL := `SELECT * from assets
	        WHERE id = $1 AND archived_at IS NULL`

	var asset models.Asset

	if err := database.AssetDB.Get(&asset, SQL, assetID); err != nil {
		return asset, err
	}

	return asset, nil
}

func GetLaptopSpecByAssetID(assetID string) (models.LaptopDetail, error) {

	SQL := `SELECT *
			FROM laptop_specs
			WHERE asset_id = $1`

	var laptop models.LaptopDetail

	if err := database.AssetDB.Get(&laptop, SQL, assetID); err != nil {
		return laptop, err
	}

	return laptop, nil
}

func GetMonitorSpecByAssetID(assetID string) (models.MonitorDetail, error) {

	SQL := `SELECT *
			FROM monitor_specs
			WHERE asset_id = $1`

	var monitor models.MonitorDetail

	if err := database.AssetDB.Get(&monitor, SQL, assetID); err != nil {
		return monitor, err
	}

	return monitor, nil
}

func GetMouseSpecByAssetID(assetID string) (models.MouseDetail, error) {

	SQL := `SELECT *
			FROM mouse_specs
			WHERE asset_id = $1`

	var mouse models.MouseDetail

	if err := database.AssetDB.Get(&mouse, SQL, assetID); err != nil {
		return mouse, err
	}

	return mouse, nil
}

func GetKeyboardSpecByAssetID(assetID string) (models.KeyboardDetail, error) {

	SQL := `SELECT *
			FROM keyboard_specs
			WHERE asset_id = $1`

	var keyboard models.KeyboardDetail

	if err := database.AssetDB.Get(&keyboard, SQL, assetID); err != nil {
		return keyboard, err
	}

	return keyboard, nil
}

func GetPhoneSpecByAssetID(assetID string) (models.PhoneDetail, error) {

	SQL := `SELECT *
			FROM phone_specs
			WHERE asset_id = $1`

	var phone models.PhoneDetail

	if err := database.AssetDB.Get(&phone, SQL, assetID); err != nil {
		return phone, err
	}

	return phone, nil
}

func GetAssetsWithSpec() ([]models.AssetWithSpecs, error) {

	SQL := `SELECT
                a.id,
                a.asset_name,
                a.asset_type,
                a.asset_brand,
                a.serial_number,
                a.purchase_date,
                a.warranty_expiry,
                a.asset_status,
                a.charger_available,

                ls.id AS laptop_id,
                ls.ram_gb,
                ls.storage_gb,
                ls.operating_system AS laptop_os,
                ls.screen_resolution AS laptop_screen_resolution,
                ls.processor,

                ms.id AS monitor_id,
                ms.screen_size,
                ms.screen_resolution AS monitor_screen_resolution,

                mos.id AS mouse_id,
                mos.wireless AS mouse_wireless,

                ks.id AS keyboard_id,
                ks.wireless AS keyboard_wireless,

                ps.id AS phone_id,
                ps.ram_gb AS phone_ram_gb,
                ps.storage_gb AS phone_storage_gb,
                ps.operating_system AS phone_os
            FROM assets a
            LEFT JOIN laptop_specs ls ON ls.asset_id = a.id
            LEFT JOIN monitor_specs ms ON ms.asset_id = a.id
            LEFT JOIN mouse_specs mos ON mos.asset_id = a.id
            LEFT JOIN keyboard_specs ks ON ks.asset_id = a.id
            LEFT JOIN phone_specs ps ON ps.asset_id = a.id
			WHERE a.archived_at IS NULL`

	var assetSpec []models.AssetWithSpecs

	err := database.AssetDB.Select(&assetSpec, SQL)

	return assetSpec, err
}

func UpdateAssetByAssetID(asset models.UpdateAssetRequest, assestID string) error {

	SQL := `UPDATE assets
	        SET asset_name = COALESCE($1, asset_name),
			    asset_type = COALESCE($2, asset_type),
			    asset_brand = COALESCE($3, asset_brand),
			    serial_number = COALESCE($4, serial_number),
			    purchase_date = COALESCE($5, purchase_date),
			    warranty_expiry = COALESCE($6, warranty_expiry),
			    asset_status = COALESCE($7, asset_status),
			    charger_available = COALESCE($8, charger_available),
			WHERE id = $9 AND archived_at IS NULL`

	_, err := database.AssetDB.Exec(SQL, asset.AssetName, asset.AssetType, asset.AssetBrand, asset.SerialNumber, asset.PurchaseDate, asset.WarrantyExpiry, asset.AssetStatus, asset.ChargerAvailable, assestID)

	return err
}

func UpdateLaptopSpecByAssetID(laptop *models.UpdateLaptopDetail, assetID string) error {

	SQL := `UPDATE laptop_specs
            SET ram_gb = COALESCE($1, ram_gb),
                storage_gb = COALESCE($2, storage_gb),
                operating_system = COALESCE($3, operating_system)
                screen_resolution = COALESCE($4, screen_resolution),
                processor = COALESCE($5, processor)
            WHERE asset_id = $6`

	_, err := database.AssetDB.Exec(SQL, laptop.RAM, laptop.Storage, laptop.OperatingSystem, laptop.ScreenResolution, laptop.Processor, assetID)

	return err
}

func UpdateMonitorSpecByAssetID(monitor *models.UpdateMonitorDetail, assetID string) error {

	SQL := `UPDATE monitor_specs
            SET screen_size = COALESCE($1, screen_size),
                screen_resolution = COALESCE($2, screen_resolution)
            WHERE asset_id = $3`

	_, err := database.AssetDB.Exec(SQL, monitor.ScreenSize, monitor.ScreenResolution, assetID)

	return err
}

func UpdateMouseSpecByAssetID(mouse *models.UpdateMouseDetail, assetID string) error {

	SQL := `UPDATE mouse_specs
            SET wireless = COALESCE($1, wireless)
            WHERE asset_id = $2`

	_, err := database.AssetDB.Exec(SQL, mouse.Wireless, assetID)

	return err
}

func UpdateKeyboardSpecByAssetID(keyboard *models.UpdateKeyboardDetail, assetID string) error {

	SQL := `UPDATE keyboard_specs
            SET wireless = COALESCE($1, wireless)
            WHERE asset_id = $2`

	_, err := database.AssetDB.Exec(SQL, keyboard.Wireless, assetID)

	return err
}

func UpdatePhoneSpecByAssetID(phone *models.UpdatePhoneDetail, assetID string) error {

	SQL := `UPDATE phone_specs
            SET ram_gb = COALESCE($1, ram_gb),
                storage_gb = COALESCE($2, storage_gb),
                operating_system = COALESCE($3, operating_system)
            WHERE asset_id = $4`

	_, err := database.AssetDB.Exec(SQL, phone.RAM, phone.Storage, phone.OperatingSystem)

	return err
}
