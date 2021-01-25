package models

// DataAirPort - AirPort data items
type DataAirPort struct {
	AirPortData []AirPortData `json:"SPAirPortDataType"`
}

// AirPortData - AirPort item
type AirPortData struct {
	AirPortIntf []AirPortIntf `json:"spairport_airport_interfaces"`   //  Interfaces
	SoftInfo    SoftInfo      `json:"spairport_software_information"` // Software Versions
}

// AirPortIntf - Information AirPorts interfaces
type AirPortIntf struct {
	IntfName     string         `json:"_name"`                                           // Interfaces Name
	AirDropChan  uint8          `json:"spairport_airdrop_channel"`                       // AirDrop Channel
	LocalWireNet []LocalWireNet `json:"spairport_airport_other_local_wireless_networks"` // Other Local Wi-Fi Networks
	CapsAirDrop  string         `json:"spairport_caps_airdrop"`                          // AirDrop
	CapsAutoUn   string         `json:"spairport_caps_autounlock"`                       // Auto Unlock
	CapsWow      string         `json:"spairport_caps_wow"`                              // Wake On Wireless
	CurNetInfo   CurNetInfo     `json:"spairport_current_network_information"`           // Current Network Information
	StatusInfo   string         `json:"spairport_status_information"`                    // Status
	SuppConnCh   []uint16       `json:"spairport_supported_channels"`                    // Supported Channels
	SuppPhyMod   string         `json:"spairport_supported_phymodes"`                    // Supported PHY Modes
	CardType     string         `json:"spairport_wireless_card_type"`                    // Card Type
	ContryCode   string         `json:"spairport_wireless_country_code"`                 // Country Code
	FirmWVer     string         `json:"spairport_wireless_firmware_version"`             // Firmware Version
	WireLocale   string         `json:"pairport_wireless_locale"`                        // Locale
	WireMacAddr  string         `json:"spairport_wireless_mac_address"`                  // MAC Address
}

// LocalWireNet - Information local wire network
type LocalWireNet struct {
	NameWireNet string `json:"_name"`                           // Name Wireless
	NetBssid    string `json:"spairport_network_bssid"`         // BSSID
	NetChan     uint8  `json:"spairport_network_channel"`       // Channel
	ContryCode  string `json:"spairport_wireless_country_code"` // Country Code
	NetPhymode  string `json:"spairport_network_phymode"`       // PHY Mode
	NetType     string `json:"spairport_network_type"`          // Network Type
	SecMode     string `json:"spairport_security_mode"`         // Security
	SignalNoise string `json:"spairport_signal_noise"`          // Signal / Noise
}

// CurNetInfo - Current information wire network
type CurNetInfo struct {
	NameWireNet string `json:"_name"`                          // Name Wireless
	NetBssid    string `json:"spairport_network_bssid"`        // BSSID
	NetChan     uint8  `json:"spairport_network_channel"`      // Channel
	ContryCode  string `json:"spairport_network_country_code"` // Country Code
	NetMcs      uint8  `json:"spairport_network_mcs"`          // MCS Index
	PhyMode     string `json:"pairport_network_phymode"`       // PHY Mode
	Rate        uint16 `json:"spairport_network_rate"`         // Transmit Rate
	NetType     string `json:"spairport_network_type"`         // Network Type
	SecMode     string `json:"spairport_security_mode"`        // Security
	SignalNoise string `json:"spairport_signal_noise"`         // Signal / Noise
}

// SoftInfo - Soft information
type SoftInfo struct {
	CoreVer   string `json:"spairport_corewlan_version"`    // CoreWLAN
	CoreKitWl string `json:"spairport_corewlankit_version"` // CoreWLANKit
	DiagVer   string `json:"spairport_diagnostics_version"` // Diagnostics
	ExtraVer  string `json:"spairport_extra_version"`       // Menu Extra
	FamilyVer string `json:"spairport_family_version"`      // IO80211 Family
	ProfVer   string `json:"spairport_profiler_version"`    // System Information
	UtilVer   string `json:"spairport_utility_version"`     // AirPort Utility
}
