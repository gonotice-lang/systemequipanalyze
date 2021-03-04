package models

// DataDisplay - Display information items
type DataDisplay struct {
	DisplayData []DisplayData `json:"SPDisplaysDataType"`
}

// DisplayData - information display item
type DisplayData struct {
	DisplayName  string          `json:"_name"`
	Vram         string          `json:"_spdisplays_vram"`
	DeviceID     string          `json:"spdisplays_device-id"`   // Device ID
	MetalFamily  string          `json:"spdisplays_metalfamily"` // Metal Family
	DisplayNdrvs []*DisplayNdrvs `json:"spdisplays_ndrvs"`       // Displays
	RevisionID   string          `json:"spdisplays_revision-id"` // Revision ID
	Vendor       string          `json:"spdisplays_vendor"`      // Vendor
	VramShared   string          `json:"spdisplays_vram_shared"` // VRAM (Dynamic, Max)
	Bus          string          `json:"sppci_bus"`              // Bus
	DeviceType   string          `json:"sppci_device_type"`      // Type
	Model        string          `json:"sppci_model"`            // Chipset Model
}

// DisplayNdrvs - Display data and parameters informations
type DisplayNdrvs struct {
	DisplayEdid  string      `json:"_IODisplayEDID"`
	TypeDisName  string      `json:"_name"`
	ProductID    string      `json:"_spdisplays_display-product-id"`
	SerialNumb   string      `json:"_spdisplays_display-serial-number2"`
	VendorID     string      `json:"_spdisplays_display-vendor-id"`
	Week         string      `json:"_spdisplays_display-week"`
	Year         string      `json:"_spdisplays_display-year"`
	DisplayID    string      `json:"_spdisplays_displayID"`
	Path         string      `json:"_spdisplays_displayPath"`
	PortDevice   *PortDevice `json:"_spdisplays_displayport_device"` // Information Display Port
	DisplayRegID string      `json:"_spdisplays_displayRegID"`
	Edid         string      `json:"_spdisplays_edid"`
	Pixels       string      `json:"_spdisplays_pixels"`
	Resolution   string      `json:"_spdisplays_resolution"`
	BrightAmb    string      `json:"spdisplays_ambient_brightness"`
	ConnectType  string      `json:"spdisplays_connection_type"`
	Depth        string      `json:"spdisplays_depth"`
	DisplayType  string      `json:"spdisplays_display_type"`
	Main         string      `json:"spdisplays_main"`
	Mirror       string      `json:"spdisplays_mirror"`
	Online       string      `json:"spdisplays_online"`
	PixelRes     string      `json:"spdisplays_pixelresolution"`
	Resol        string      `json:"spdisplays_resolution"`
}

// PortDevice - Device port information
type PortDevice struct {
	PortDeviceName string `json:"_name"`
	CurBand        string `json:"spdisplays_displayport_current_bandwidth"`
	CurLanes       string `json:"spdisplays_displayport_current_lanes"`
	CurSpread      string `json:"spdisplays_displayport_current_spread"`
	DPCDVer        string `json:"spdisplays_displayport_DPCD_version"`
	ErrLaneZero    string `json:"spdisplays_displayport_errors_lane0"`
	ErrLaneOne     string `json:"spdisplays_displayport_errors_lane1"`
	ErrLaneTwo     string `json:"spdisplays_displayport_errors_lane2"`
	ErrLaneThr     string `json:"spdisplays_displayport_errors_lane3"`
	HDCPCap        string `json:"spdisplays_displayport_hdcp_capability"`
	MaxBand        string `json:"spdisplays_displayport_max_bandwidth"`
	MaxLanes       string `json:"spdisplays_displayport_max_lanes"`
	MaxSpread      string `json:"spdisplays_displayport_max_spread"`
	ASCIIName      string `json:"spdisplays_displayport_sink_ascii_name"`
	ChipVer        string `json:"spdisplays_displayport_sink_chip_version"`
	SinkCount      string `json:"spdisplays_displayport_sink_count"`
	SinkSwVer      string `json:"spdisplays_displayport_sink_sw_version"`
	SinkVen        string `json:"spdisplays_displayport_sink_vendor"`
	VErrLaneZero   string `json:"spdisplays_displayport_valid_error_lane0"`
	VErrLaneOne    string `json:"spdisplays_displayport_valid_error_lane1"`
	VErrLaneTwo    string `json:"spdisplays_displayport_valid_error_lane2"`
	VErrLaneThr    string `json:"spdisplays_displayport_valid_error_lane3"`
}
