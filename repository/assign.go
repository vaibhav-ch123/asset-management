package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vaibhav-ch123/asset-management/database"
	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
)

func GetAssetStatus(assetID string) (string, error) {

	SQL := `SELECT asset_status
            FROM assets
            WHERE id = $1 AND archived_at IS NULL`

	var assetStatus string

	err := database.AssetDB.Get(&assetStatus, SQL, assetID)

	return assetStatus, err
}

func CreateAssetAssign(tx sqlx.Ext, assetAssign models.AssignRequest) error {

	SQL := `INSERT INTO asset_assignments (employee_id, asset_id, assigned_date, remark)
            VALUES ($1, $2, $3, $4)`

	_, err := tx.Exec(SQL, assetAssign.EmployeeID, assetAssign.AssetID, assetAssign.AssignedDate, assetAssign.Remark)

	return err
}

func UpdateAssetStatusAssigned(tx sqlx.Ext, assetID string) error {

	SQL := `UPDATE assets
            SET asset_status = $1
            WHERE id = $2 AND archived_at IS NULL`

	_, err := tx.Exec(SQL, "assigned", assetID)

	return err
}

func DeleteAssetAssign(tx sqlx.Ext, assetAssign models.AssignRequest, assignedID string) error {

	SQL := `UPDATE asset_assignments
            SET returned_date = $1,
                archived_at = NOW()
            WHERE id = $2 AND archived_at IS NULL`

	result, err := tx.Exec(SQL, assetAssign.ReturnedDate, assignedID)

	row, rowErr := result.RowsAffected()
	if rowErr != nil {
		return rowErr
	}
	if row == 0 {
		return errors.ErrAssetAssignedNotFound
	}
	return err
}

func UpdateAssetStatusUnassigned(tx sqlx.Ext, assetID string) error {

	SQL := `UPDATE assets
            SET asset_status = $1
            WHERE id = $2 AND archived_at IS NULL`

	_, err := tx.Exec(SQL, "available", assetID)

	return err
}

func GetAssignedAssets() ([]models.AssetAssignedResponse, error) {

	SQL := `SELECT 
               aa.id AS assigned_id,
               aa.assigned_date AS assigned_date,
               aa.remark AS remark,
               a.id AS asset_id,
               a.asset_name AS asset_name,
               a.asset_type AS asset_type,
               a.asset_brand AS asset_brand,
               a.serial_number AS serial_number,
               e.id AS employee_id,
               e.name AS employee_name,
               e.email AS employee_email,
               e.phone AS employee_phone,
               e.employee_type AS employee_type,
               e.employee_role AS employee_role
            FROM asset_assignments aa
            JOIN assets a ON aa.asset_id = a.id
            JOIN employees e ON aa.employee_id = e.id
            WHERE aa.archived_at IS NULL`

	var assetAssigned []models.AssetAssignedResponse

	err := database.AssetDB.Select(&assetAssigned, SQL)

	return assetAssigned, err
}

func GetEmployeeAssignedAssets(employeeID string) ([]models.EmployeeAssetAssignedResponse, error) {

	SQL := `SELECT 
               a.id AS asset_id,
               a.asset_name AS asset_name,
               a.asset_type AS asset_type,
               a.asset_brand AS asset_brand,
               a.serial_number AS serial_number,
               a.charger_available AS charger_available,
               a.warranty_expiry AS warranty_expiry,
               a.purchase_date AS purchase_date
            FROM asset_assignments aa
            JOIN assets a ON aa.asset_id = a.id
            WHERE aa.employee_id = $1 AND aa.archived_at IS NULL`

	var assetAssigned []models.EmployeeAssetAssignedResponse
	err := database.AssetDB.Select(&assetAssigned, SQL, employeeID)

	return assetAssigned, err
}
