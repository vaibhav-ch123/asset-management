package models

import "time"

type CreateAssetRequest struct {
	AssetName        string         `json:"assetName"`
	AssetType        string         `json:"assetType"`
	AssetBrand       string         `json:"assetBrand"`
	SerialNumber     string         `json:"serialNumber"`
	PurchaseDate     string         `json:"purchaseDate"`
	WarrantyExpiry   string         `json:"warrantyExpiry"`
	AssetStatus      string         `json:"assetStatus"`
	ChargerAvailable bool           `json:"chargerAvailable"`
	Laptop           LaptopDetail   `json:"laptopDetail"`
	Monitor          MonitorDetail  `json:"monitorDetail"`
	Mouse            MouseDetail    `json:"mouseDetail"`
	Phone            PhoneDetail    `json:"phoneDetail"`
	Keyboard         KeyboardDetail `json:"keyboardDetail"`
}

type UpdateAssetRequest struct {
	AssetName        *string `json:"assetName"`
	AssetType        *string `json:"assetType"`
	AssetBrand       *string `json:"assetBrand"`
	SerialNumber     *string `json:"serialNumber"`
	PurchaseDate     *string `json:"purchaseDate"`
	WarrantyExpiry   *string `json:"warrantyExpiry"`
	AssetStatus      *string `json:"assetStatus"`
	ChargerAvailable *bool   `json:"chargerAvailable"`
}

type AssetResponse struct {
	ID               string `json:"id"`
	AssetName        string `json:"assetName"`
	AssetType        string `json:"assetType"`
	AssetBrand       string `json:"assetBrand"`
	SerialNumber     string `json:"serialNumber"`
	PurchaseDate     string `json:"purchaseDate"`
	WarrantyExpiry   string `json:"warrantyExpiry"`
	AssetStatus      string `json:"assetStatus"`
	ChargerAvailable bool   `json:"chargerAvailable"`
}

type Asset struct {
	ID               string    `db:"id"`
	AssetName        string    `db:"asset_name"`
	AssetType        string    `db:"asset_type"`
	AssetBrand       string    `db:"asset_brand"`
	SerialNumber     string    `db:"serial_number"`
	PurchaseDate     time.Time `db:"purchase_date"`
	WarrantyExpiry   time.Time `db:"warranty_expiry"`
	AssetStatus      string    `db:"asset_status"`
	ChargerAvailable bool      `db:"charger_available"`

	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	ArchivedAt *time.Time `db:"archived_at"`
}

type LaptopDetail struct {
	ID               string `db:"id" json:"id"`
	AssetID          string `db:"asset_id" json:"assetId"`
	RAM              int    `db:"ram_gb" json:"ramGb"`
	Storage          int    `db:"storage_gb" json:"storageGb"`
	OperatingSystem  string `db:"operating_system" json:"operatingSystem"`
	ScreenResolution string `db:"screen_resolution" json:"screenResolution"`
	Processor        string `db:"processor" json:"processor"`
}

type MonitorDetail struct {
	ID               string `db:"id" json:"id"`
	AssetID          string `db:"asset_id" json:"assetId"`
	ScreenSize       string `db:"screen_size" json:"screenSize"`
	ScreenResolution string `db:"screen_resolution" json:"screenResolution"`
}

type MouseDetail struct {
	ID        string `db:"id" json:"id"`
	AssetID   string `db:"asset_id" json:"assetId"`
	Wireless  bool   `db:"wireless" json:"wireless"`
}

type KeyboardDetail struct {
	ID        string `db:"id" json:"id"`
	AssetID   string `db:"asset_id" json:"assetId"`
	Wireless  bool   `db:"wireless" json:"wireless"`
}

type PhoneDetail struct {
	ID              string `db:"id" json:"id"`
	AssetID         string `db:"asset_id" json:"assetId"`
	RAM             int    `db:"ram_gb" json:"ramGb"`
	Storage         int    `db:"storage_gb" json:"storageGb"`
	OperatingSystem string `db:"operating_system" json:"operatingSystem"`
}
