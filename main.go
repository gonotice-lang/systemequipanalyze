package main

import (
	"analyzes/models"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// ResultError - Information errors in goroutines
type ResultError struct {
	res ResErr
	err error
}

// ResErr - Information errors
type ResErr struct {
	ErrorName string
	NumbOccur int
}

// InfoMacOs - Informations Mac OS
type InfoMacOs interface {
	InfoOS() (*OSInfo, []*ResultError) // returns name os and arrays error
	SystemEquip() *SysEq               // returns system equip information
	NetInterfacesInfo() string         // returns net interfaces
	RoutingTableInfo() string          // returns table routing
	ArpTable() string                  // returns arp table routing
	PortInfo(string) []*PortInf        // returns port info information
	FilesSystemMount() string          // returns files system mounted information
	ProcInfo() []*PrcInf               // returns information launched process
}

// SystemEquip - system equipment information
type SystemEquip interface {
	HarwareEquip() (*models.DataHardware, error)
	ProcEquip() (string, error)
	RAMEquip() (*models.DataMem, error)
	StorageEquip() (*models.DataStorage, error)
	DisplayEquip() (*models.DataDisplay, error)
	USBEquip() (*models.DataUSB, error)
	NetworkEquip() (*models.DataNetwork, error)
	AirPortEquip() (*models.DataAirPort, error)
	EthernetEquip() (*models.DataEthernet, error)
	PciEquip() (*models.DataPci, error)
}

// MacOSInfo - structure macos informations
type MacOSInfo struct {
	OSInfo           *OSInfo
	SystemEquip      *SysEq
	NetInrerfaces    string
	RoutingTable     string
	ArpTable         string
	PortInfo         *PortInf
	FilesSystemMount string
	ProcInfo         *PrcInf
}

// OSInfo - Infomation OS
type OSInfo struct {
	OsName   string
	KernVer  string
	NodeName string
	OsArch   string
}

// SysEq - Structure system information
type SysEq struct {
	ProcInfo     string
	HardwareInfo string
	MemoryInfo   string
	StorageInfo  string
	DisplayInfo  string
	NetworkInfo  string
	AirPortInfo  string
	EthernetInfo string
	PCIInfo      string
	UsbInfo      string
}

// PortInf - information port
type PortInf struct {
	Proto       string
	RecvQ       uint32
	SenQ        uint32
	LocalAddr   string
	ForeignAddr string
	State       string
}

// PrcInf - Launched process Infomations
type PrcInf struct {
	ProcPid  uint32
	ProcTty  string
	ProcTime string
	ProcCmd  string
}

func osDetect() ([]byte, error) {
	cmd := exec.Command("uname")

	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	return stdout, nil
	/*res, err := osDetect()
	if err != nil {
		log.Println(err)
	}
	fmt.Print(string(res))*/
}

// InfoOS - returns name os and error
func (oi *OSInfo) InfoOS() (*OSInfo, []*ResultError) {
	osInfoCmd := "uname"
	osInfoCmdKey := []string{"-s", "-v", "-n", "-mp"}
	resOsInfo := make([]string, 0, 4)

	outErr := make([]*ResultError, 0, 4)

	chOsInfo := make(chan string, 1)

	for gn, vk := range osInfoCmdKey {
		go func(gn int, vk string) {
			cmd := exec.Command(osInfoCmd, vk)
			stdout, err := cmd.Output()
			if err != nil {
				outErr = append(outErr, &ResultError{
					res: ResErr{
						ErrorName: "Undefined key value " + vk,
						NumbOccur: gn,
					},
					err: err,
				})
			}

			chOsInfo <- string(stdout)
		}(gn, vk)

		resOsInfo = append(resOsInfo, <-chOsInfo)
	}

	sp := strings.Split(resOsInfo[3], " ")
	if sp[0] == sp[1] {
		resOsInfo[3] = sp[1]
	}

	oi = &OSInfo{
		OsName:   resOsInfo[0],
		KernVer:  resOsInfo[1],
		NodeName: resOsInfo[2],
		OsArch:   resOsInfo[3],
	}

	return oi, outErr
}

// HarwareEquip - General information equipment hardware
func HarwareEquip() (*models.DataHardware, error) {
	var hard *models.DataHardware

	cmd := exec.Command("system_profiler", "-json", "SPHardwareDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &hard)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return hard, nil
}

// ProcEquip - processor information
func ProcEquip() (string, error) {
	cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("\"%s\" %v", cmd, err)
	}

	return string(stdout), nil
}

// RAMEquip - RAM memory information
func RAMEquip() (*models.DataMem, error) {
	var mem *models.DataMem

	cmd := exec.Command("system_profiler", "-json", "SPMemoryDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &mem)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return mem, nil
}

// StorageEquip - storage information
func StorageEquip() (*models.DataStorage, error) {
	var stor *models.DataStorage

	cmd := exec.Command("system_profiler", "-json", "SPStorageDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &stor)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return stor, nil
}

// DisplayEquip - display information
func DisplayEquip() (*models.DataDisplay, error) {
	var disp *models.DataDisplay

	cmd := exec.Command("system_profiler", "-json", "SPDisplaysDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &disp)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return disp, nil
}

// USBEquip - usb equipments information
func USBEquip() (*models.DataUSB, error) {
	var usb *models.DataUSB

	cmd := exec.Command("system_profiler", "-json", "SPUSBDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &usb)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return usb, nil
}

// NetworkEquip - Network information
func NetworkEquip() (*models.DataNetwork, error) {
	var netw *models.DataNetwork

	cmd := exec.Command("system_profiler", "-json", "SPNetworkDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &netw)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return netw, nil
}

// AirPortEquip - AirPort information
func AirPortEquip() (*models.DataAirPort, error) {
	var airp *models.DataAirPort

	cmd := exec.Command("system_profiler", "-json", "SPAirPortDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &airp)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return airp, nil
}

// EthernetEquip - Ethernet information
func EthernetEquip() (*models.DataEthernet, error) {
	var eth *models.DataEthernet

	cmd := exec.Command("system_profiler", "-json", "SPEthernetDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &eth)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return eth, nil
}

// PciEquip - PCI information
func PciEquip() (*models.DataPci, error) {
	var pci *models.DataPci

	cmd := exec.Command("system_profiler", "-json", "SPPCIDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &pci)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return pci, nil
}

func main() {
	var resOsInfo *OSInfo

	resOsInfo, resErr := resOsInfo.InfoOS()
	for _, valr := range resErr {
		log.Printf("Error: %s in function number \"%d\", %v\n", valr.res.ErrorName, valr.res.NumbOccur, valr.err)
	}
	fmt.Print(resOsInfo.KernVer)

	resProcInfo, err := ProcEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resProcInfo)

	resRAMInfo, err := RAMEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resRAMInfo)

	resStorInfo, err := StorageEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resStorInfo)

	resDispInfo, err := DisplayEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resDispInfo)

	resHardInfo, err := HarwareEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resHardInfo)

	resUsbInfo, err := USBEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resUsbInfo)

	resNetworkInfo, err := NetworkEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resNetworkInfo)

	resAirPortInfo, err := AirPortEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resAirPortInfo)

	resPciInfo, err := PciEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resPciInfo)

	resEthernetInfo, err := EthernetEquip()
	if err != nil {
		log.Printf("Failed execution command: %v", err)
	}
	fmt.Print(resEthernetInfo)
}
