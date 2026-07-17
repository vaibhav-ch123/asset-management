package models

import "time"

type AssignRequest struct {
	EmployeeID   string `json:"employeeID"`
	AssetID      string `json:"assetID"`
	AssignedDate string `json:"assignedDate"`
	ReturnedDate string `json:"returnedDate"`
	Remark       string `json:"remark"`
}

type AssetAssignedResponse struct {
	ID           string    `db:"assigned_id" json:"assignedId"`
	AssignedDate time.Time `db:"assigned_date" json:"assignedDate"`
	Remark       string    `db:"remark" json:"assignedRemark"`

	AssetID      string `db:"asset_id" json:"assetId"`
	AssetName    string `db:"asset_name" json:"assetName"`
	AssetType    string `db:"asset_type" json:"assetType"`
	AssetBrand   string `db:"asset_brand" json:"assetBrand"`
	SerialNumber string `db:"serial_number" json:"serialNumber"`

	EmployeeID   string `db:"employee_id" json:"employeeId"`
	Name         string `db:"employee_name" json:"employeeName"`
	Email        string `db:"employee_email" json:"employeeEmail"`
	Phone        string `db:"employee_phone" json:"employeePhone"`
	EmployeeType string `db:"employee_type" json:"employeeType"`
	EmployeeRole string `db:"employee_role" json:"employeeRole"`
}

type EmployeeAssetAssignedResponse struct {
	ID               string     `db:"asset_id" json:"assetId"`
	AssetName        string     `db:"asset_name" json:"assetName"`
	AssetType        string     `db:"asset_type" json:"assetType"`
	AssetBrand       string     `db:"asset_brand" json:"assetBrand"`
	SerialNumber     string     `db:"serial_number" json:"serialNumber"`
	ChargerAvailable bool       `db:"charger_available" json:"chargerAvailable"`
	PurchaseDate     *time.Time `db:"purchase_date" json:"purchaseDate"`
	WarrantyExpiry   *time.Time `db:"warranty_expiry" json:"warrantyExpiry"`
}
