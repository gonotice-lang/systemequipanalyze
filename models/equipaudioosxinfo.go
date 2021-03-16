package models

// AudioInfo - all audio information struct
type AudioInfo struct {
	AudioInf []*AudioIntfInfo `json:"SPAudioDataType"`
}

// AudioIntfInfo - all audio
type AudioIntfInfo struct {
	ItemsAudio []*AudioIntf `json:"_items"`
	NameAudio  string       `json:"_name"`
}

// AudioIntf - audio
type AudioIntf struct {
	NameIntfAdio     string `json:"_name"`
	Properties       string `json:"_properties,omitempty"`
	AudioOutDevice   string `json:"coreaudio_default_audio_output_device,omitempty"`
	AudioSysteSevice string `json:"coreaudio_default_audio_system_device,omitempty"`
	DeviceManufact   string `json:"coreaudio_device_manufacturer,omitempty"`
	DeviceOutput     string `json:"coreaudio_device_output,omitempty"`
	DeviceSrate      string `json:"coreaudio_device_srate,omitempty"`
	DeviceTrans      string `json:"coreaudio_device_transport,omitempty"`
	OutputSource     string `json:"coreaudio_output_source,omitempty"`
}
