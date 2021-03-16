package models

// BluetoothDataInfo - bluetooth informations data
type BluetoothDataInfo struct {
	BluetoothInfo []*BluetoothInfo `json:"SPBluetoothDataType"`
}

// BluetoothInfo - bluetooth informations
type BluetoothInfo struct {
	BlueVer          string                       `json:"apple_bluetooth_version"`
	DeviceTitle      []map[string]DeviceTitle     `json:"device_title"`
	IncPortsTitle    []map[string]SetsSerialPorts `json:"incoming_serial_ports_title"`
	LocalDeviceTitle *LocalDeviceTitle            `json:"local_device_title"`
	OutPortsTitle    []map[string]SetsSerialPorts `json:"outgoing_serial_ports_title"`
	ServicesTitle    []map[string]ServicesTitle   `json:"services_title"`
}

type DeviceTitle struct {
	DeviceAddr     string `json:"device_addr"`
	ClassDevice    string `json:"device_classOfDevice"`
	CoreSpec       string `json:"device_core_spec"`
	FmVersion      string `json:"device_fw_version"`
	IsConfigured   string `json:"device_isconfigured"`
	IsConnect      string `json:"device_isconnected"`
	IsPaired       string `json:"device_ispaired"`
	DeviceString   string `json:"device_majorClassOfDevice_string"`
	Manufacturer   string `json:"device_manufacturer"`
	ClassOfDevice  string `json:"device_minorClassOfDevice_string"`
	ProductID      string `json:"device_productID"`
	DeviceServices string `json:"device_services"`
	SupportsEDR    string `json:"device_supportsEDR"`
	SupportsESCO   string `json:"device_supportsESCO"`
	SupportsSSP    string `json:"device_supportsSSP"`
	VendorID       string `json:"device_vendorID"`
}

type SetsSerialPorts struct {
	DeviceAuth    string `json:"device_authentication,omitempty"`
	DeviceChannel string `json:"device_channel,omitempty"`
	GenAddr       string `json:"general_address,omitempty"`
}

type LocalDeviceTitle struct {
	GeneralAddress string `json:"general_address"`
	AutoseekKey    string `json:"general_autoseek_keyboard"`
	AutoSeekPoint  string `json:"general_autoseek_pointing"`
	Chipset        string `json:"general_chipset"`
	Connectable    string `json:"general_connectable"`
	ClassComposite string `json:"general_device_class_composite"`
	ClassMajor     string `json:"general_device_class_major"`
	ClassMinor     string `json:"general_device_class_minor"`
	Discover       string `json:"general_discoverable"`
	FwVersion      string `json:"general_fw_version"`
	HciRevision    string `json:"general_hci_revision"`
	HciVersion     string `json:"general_hci_version"`
	LmpSubversion  string `json:"general_lmp_subversion"`
	LmpVersion     string `json:"general_lmp_version"`
	Mfg            string `json:"general_mfg"`
	GenName        string `json:"general_name"`
	GenPower       string `json:"general_power"`
	GenProductID   string `json:"general_productID"`
	RemoteWake     string `json:"general_remoteWake"`
	ServiceClass   string `json:"general_service_class"`
	Handoff        string `json:"general_supports_handoff"`
	InstHotspot    string `json:"general_supports_instantHotspot"`
	LowEnergy      string `json:"general_supports_lowEnergy"`
	CompleteStr    string `json:"general_type_complete_string"`
	MajorStr       string `json:"general_type_major_string"`
	GenVendorID    string `json:"general_vendorID"`
}

type ServicesTitle struct {
	FTPRootFolder string `json:"service_FTPRootFolder,omitempty"`
	OBEXFolder    string `json:"service_OBEXFolder,omitempty"`
	OBEXHandle    string `json:"service_OBEXHandle_Other,omitempty"`
	OBEXReceive   string `json:"service_OBEXReceive,omitempty"`
	ServiceState  string `json:"service_state"`
}
