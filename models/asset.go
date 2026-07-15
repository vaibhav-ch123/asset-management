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

	Laptop   *UpdateLaptopDetail   `json:"laptopDetail"`
	Monitor  *UpdateMonitorDetail  `json:"monitorDetail"`
	Mouse    *UpdateMouseDetail    `json:"mouseDetail"`
	Keyboard *UpdateKeyboardDetail `json:"keyboardDetail"`
	Phone    *UpdatePhoneDetail    `json:"phoneDetail"`
}

type AssetResponse struct {
	ID               string          `json:"id"`
	AssetName        string          `json:"assetName"`
	AssetType        string          `json:"assetType"`
	AssetBrand       string          `json:"assetBrand"`
	SerialNumber     string          `json:"serialNumber"`
	PurchaseDate     string          `json:"purchaseDate"`
	WarrantyExpiry   string          `json:"warrantyExpiry"`
	AssetStatus      string          `json:"assetStatus"`
	ChargerAvailable bool            `json:"chargerAvailable"`
	Laptop           *LaptopDetail   `json:"laptopDetail,omitempty"`
	Monitor          *MonitorDetail  `json:"monitorDetail,omitempty"`
	Mouse            *MouseDetail    `json:"mouseDetail,omitempty"`
	Phone            *PhoneDetail    `json:"phoneDetail,omitempty"`
	Keyboard         *KeyboardDetail `json:"keyboardDetail,omitempty"`
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
	ID       string `db:"id" json:"id"`
	AssetID  string `db:"asset_id" json:"assetId"`
	Wireless bool   `db:"wireless" json:"wireless"`
}

type KeyboardDetail struct {
	ID       string `db:"id" json:"id"`
	AssetID  string `db:"asset_id" json:"assetId"`
	Wireless bool   `db:"wireless" json:"wireless"`
}

type PhoneDetail struct {
	ID              string `db:"id" json:"id"`
	AssetID         string `db:"asset_id" json:"assetId"`
	RAM             int    `db:"ram_gb" json:"ramGb"`
	Storage         int    `db:"storage_gb" json:"storageGb"`
	OperatingSystem string `db:"operating_system" json:"operatingSystem"`
}

type UpdateLaptopDetail struct {
	RAM              *int    `json:"ramGb"`
	Storage          *int    `json:"storageGb"`
	OperatingSystem  *string `json:"operatingSystem"`
	ScreenResolution *string `json:"screenResolution"`
	Processor        *string `json:"processor"`
}

type UpdateMonitorDetail struct {
	ScreenSize       *string `json:"screenSize"`
	ScreenResolution *string `json:"screenResolution"`
}

type UpdateMouseDetail struct {
	Wireless *bool `json:"wireless"`
}

type UpdateKeyboardDetail struct {
	Wireless *bool `json:"wireless"`
}

type UpdatePhoneDetail struct {
	RAM             *int    `json:"ramGb"`
	Storage         *int    `json:"storageGb"`
	OperatingSystem *string `json:"operatingSystem"`
}

type AssetWithSpecs struct {
	// Asset table
	ID               string     `db:"id"`
	AssetName        string     `db:"asset_name"`
	AssetType        string     `db:"asset_type"`
	AssetBrand       string     `db:"asset_brand"`
	SerialNumber     string     `db:"serial_number"`
	PurchaseDate     *time.Time `db:"purchase_date"`
	WarrantyExpiry   *time.Time `db:"warranty_expiry"`
	AssetStatus      string     `db:"asset_status"`
	ChargerAvailable bool       `db:"charger_available"`

	// Laptop Specs
	LaptopID          *string `db:"laptop_id"`
	RAM               *int    `db:"ram_gb"`
	Storage           *int    `db:"storage_gb"`
	LaptopOS          *string `db:"laptop_os"`
	ScreenResolution  *string `db:"laptop_screen_resolution"`
	Processor         *string `db:"processor"`

	// Monitor Specs
	MonitorID                 *string `db:"monitor_id"`
	MonitorScreenSize         *string `db:"screen_size"`
	MonitorScreenResolution   *string `db:"monitor_screen_resolution"`

	// Mouse Specs
	MouseID         *string `db:"mouse_id"`
	MouseWireless   *bool   `db:"mouse_wireless"`

	// Keyboard Specs
	KeyboardID           *string `db:"keyboard_id"`
	KeyboardWireless     *bool   `db:"keyboard_wireless"`

	// Phone Specs
	PhoneID          *string `db:"phone_id"`
	PhoneRAM         *int    `db:"phone_ram_gb"`
	PhoneStorage     *int    `db:"phone_storage_gb"`
	PhoneOS          *string `db:"phone_os"`
}