package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"systemequipanalyze/models"
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
	InfoOS() (*OSInfo, []*ResultError)  // Name os and arrays error
	SystemEquip() *SysEq                // System equip information
	NetInterfacesInfo() string          // Net interfaces
	RoutingTableInfo() string           // Table routing
	ArpTable() string                   // Arp table routing
	PortInfo(string) []*PortInfo        // Port info information
	FilesSystemMount() string           // Files system mounted information
	ProcessInfo() []*ProcessInfo        // Information launched process
	DiskUsage() []*models.DiskUsageInfo // Disk usage information
	ProcEquip() (string, error)         // Processor name information
}

// SystemEquip - system equipment information
type SystemEquip interface {
	HarwareEquip() (*models.DataHardware, error)
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
	PortInfo         *PortInfo
	FilesSystemMount string
	ProcessInfo      *ProcessInfo
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

// PortInfo - information port
type PortInfo struct {
	Proto       string
	RecvQ       uint32
	SenQ        uint32
	LocalAddr   string
	ForeignAddr string
	State       string
}

// ProcessInfo - Launched process Infomations
type ProcessInfo struct {
	ProcPid  uint32
	ProcTty  string
	ProcTime string
	ProcCmd  string
}

// ARPInfo - ARP table information
type ARPInfo struct {
	NameIP    string
	MacInf    string
	Interface string
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

// PowerEquip - Power information
func PowerEquip() (*models.DataPower, error) {
	var power *models.DataPower

	cmd := exec.Command("system_profiler", "-json", "SPPowerDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &power)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return power, nil
}

// PrinterEquip - Printer information
func PrinterEquip() (*models.DataPrinter, error) {
	var printer *models.DataPrinter

	cmd := exec.Command("system_profiler", "-json", "SPPrintersDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &printer)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return printer, nil
}

func removeIndex(s []string, count, index int) []string {
	return append(s[:index], s[index+count:]...)
}

// DiskUsage -  Disk usage information
func DiskUsage() ([]*models.DiskUsageInfo, error) {
	//var keys string
	var newDiskUs = new(models.DiskUsageInfo)
	var resDiskUsage []*models.DiskUsageInfo

	cmd := exec.Command("df", "-H")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("\"%s\" %v", cmd, err)
	}

	// delete last line break
	if len(stdout) > 0 {
		stdout = stdout[:len(stdout)-1]
	}

	// line splitting by break
	res := strings.Split(string(stdout), "\n")
	// delete first line
	res = removeIndex(res, 1, 0)

	re, _ := regexp.Compile("map auto_home")

	for _, val := range res {
		// replace for deleting space
		val = re.ReplaceAllString(val, "map_auto_home")
		valFields := strings.Fields(val)
		for range valFields {
			newDiskUs = &models.DiskUsageInfo{
				FileSys:     valFields[0],
				Size:        valFields[1],
				Used:        valFields[2],
				Avail:       valFields[3],
				Capacity:    valFields[4],
				IUsed:       valFields[5],
				IFree:       valFields[6],
				ProcentUsed: valFields[7],
				Mounted:     valFields[8],
			}
		}
		resDiskUsage = append(resDiskUsage, newDiskUs)
	}
	return resDiskUsage, nil
}

// NetStatConnInfo - Internet connect information
func NetStatConnInfo(key, arg string) ([]*models.NetStatConn, error) {
	var cmd *exec.Cmd

	var newNetStatConn = new(models.NetStatConn)
	var resNetStatConn []*models.NetStatConn

	cmd = exec.Command("netstat", key, arg)
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("\"%s\" %v", cmd, err)
	}
	// delete last line break
	if len(stdout) > 0 {
		stdout = stdout[:len(stdout)-1]
	}

	res := strings.Split(string(stdout), "\n")
	res = removeIndex(res, 2, 0)

	for _, val := range res {
		// replace for deleting space
		valFields := strings.Fields(val)
		// add element in the end
		if len(valFields) == 5 {
			valFields = append(valFields, "-")
		}
		for range valFields {
			newNetStatConn = &models.NetStatConn{
				Proto:       valFields[0],
				RecvQ:       valFields[1],
				SendQ:       valFields[2],
				LocalAddr:   valFields[3],
				ForeignAddr: valFields[4],
				State:       valFields[5],
			}
		}

		resNetStatConn = append(resNetStatConn, newNetStatConn)
	}

	return resNetStatConn, err
}

// NetStatRouteInfo - Table routing information
func NetStatRouteInfo(key, arg string) (*models.NetStatRoute, error) {
	var cmd *exec.Cmd
	var verIP string

	switch arg {
	case "inet":
		verIP = "IPv4"
	case "inet6":
		verIP = "IPv6"
	}

	var newRouteInfo = new(models.RouteInfo)
	var resRouteInfo []*models.RouteInfo

	cmd = exec.Command("netstat", key, arg)
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("\"%s\" %v", cmd, err)
	}
	// delete last line break
	if len(stdout) > 0 {
		stdout = stdout[:len(stdout)-1]
	}

	res := strings.Split(string(stdout), "\n")
	res = removeIndex(res, 4, 0)

	for _, val := range res {
		// replace for deleting space
		valFields := strings.Fields(val)
		// add element in the end
		if len(valFields) == 4 {
			valFields = append(valFields, "-")
		}
		for range valFields {
			newRouteInfo = &models.RouteInfo{
				Dst:     valFields[0],
				Gateway: valFields[1],
				Flags:   valFields[2],
				Netif:   valFields[3],
				Expire:  valFields[4],
			}
		}

		resRouteInfo = append(resRouteInfo, newRouteInfo)
	}

	resNetStatRoute := &models.NetStatRoute{
		VerIP:     verIP,
		RouteInfo: resRouteInfo,
	}

	return resNetStatRoute, err
}

// ARPTableInfo - information ARP Table
func ARPTableInfo() ([]*ARPInfo, error) {
	var templ = []string{"at", "on"}
	var resArp []*ARPInfo

	cmd := exec.Command("arp", "-a")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("\"%s\" %v", cmd, err)
	}

	if len(stdout) > 0 {
		stdout = stdout[:len(stdout)-1]
	}

	res := strings.Split(string(stdout), "\n")
	for _, vs := range res {
		for _, vres := range templ {
			vs = strings.ReplaceAll(vs, vres, "")
		}
		newVRes := strings.Fields(vs)

		for i := 0; i < len(newVRes); i++ {
			if newVRes[i] == "?" {
				s := newVRes[i] + newVRes[i+1]
				newVRes[i] = s
				newVRes = append(newVRes[:i+1], newVRes[i+2:]...)
			}
		}
		resArp = append(resArp, &ARPInfo{
			NameIP:    newVRes[0],
			MacInf:    newVRes[1],
			Interface: newVRes[2],
		})
	}

	return resArp, err
}

func subSlices(slice []string) [][]string {
	var ss [][]string
	for _, e := range slice {
		if e[0] != 9 || len(ss) == 0 {
			ss = append(ss, make([]string, 0, 3))
		}
		end := len(ss) - 1
		e = strings.ReplaceAll(e, "\t", "")
		ss[end] = append(ss[end], e)
	}

	return ss
}

// NetworkIntInfo - information networks interfaces
func NetworkIntInfo() error {
	var resNetInfInfo []*models.NetIntInfo

	cmd := exec.Command("ifconfig")
	stdout, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("\"%s\" %v", cmd, err)
	}

	// delete last line break
	if len(stdout) > 0 {
		stdout = stdout[:len(stdout)-1]
	}

	reInf, _ := regexp.Compile(`[A-Za-z_0-9]*`)
	reFlag, _ := regexp.Compile(`flags=[0-9]*<([A-Z0-9]*_*,*)*>`)
	reMtu, _ := regexp.Compile(`mtu\s[0-9]*`)
	reInet, _ := regexp.Compile(`^inet\s`)
	//reInet6, _ := regexp.Compile(`([0-9A-Fa-f%a-z0-9]{1,4}:*){8}|\s:{2}1|fe80::1%[a-z0-9]*`)

	res := strings.Split(string(stdout), "\n")

	for _, sub := range subSlices(res) {
		var (
			resOptions    string
			resNd6Options string
			resMedia      string
			resStatus     string
		)

		resInf := reInf.FindAllString(sub[0], 1)
		resFlag := reFlag.FindAllString(sub[0], 1)
		resMtu := reMtu.FindAllString(sub[0], 1)

		// checking matching options value
		matchOpts, _ := regexp.MatchString(`^options=[0-9]*<([A-Z0-9]*_*,*)*>`, sub[1:2][0])
		if matchOpts {
			resOptions = sub[1:2][0]
		}

		resEther, resConfigMem := []string{}, []string{}
		resInetAddr := &models.Inet{}
		resInet6Addr := []*models.Inet6{}

		for s, v := range sub {
			// search and assignment ether value
			if strings.Contains(v, "ether") {
				fEther := strings.Fields(v)
				ether := append(make([]string, 0, 3), fEther[1])
				resEther = ether
			}
			// checking and distribution configuration fields
			if strings.Contains(sub[s], "Configuration:") {
				for _, vs := range sub[s+1:] {
					if vs[0] == 32 {
						vs = strings.Replace(vs, " ", "", 8)
					}
					resConfigMem = append(resConfigMem, vs)
					if strings.Contains(vs, "ifmaxaddr") {
						break
					}
				}
			}
			// checking and distribution inet fields
			if reInet.MatchString(v) {
				fInet := strings.Fields(v)

				netmask, broadcast := "", ""
				inetAddr := fInet[1]

				for k := range fInet {
					if fInet[k] == "netmask" {
						netmask = fInet[k+1]
					}
					if fInet[k] == "broadcast" {
						broadcast = fInet[k+1]
					}
				}

				resInetAddr = &models.Inet{
					InetAddr:  inetAddr,
					Netmask:   netmask,
					Broadcast: broadcast,
				}
			}

			// checking and distribution inet6 fields
			if strings.Contains(v, "inet6") {
				// resInet6 := reInet6.FindAllString(v, 1)
				fInet6 := strings.Fields(v)

				var resScopeid string
				var resParamPrefix string

				for k, val := range fInet6[4:] {
					if fInet6[4:][k] == "scopeid" {
						resScopeid = fInet6[4:][k+1]
					}
					if val != "scopeid" {
						resParamPrefix = resParamPrefix + " " + val
					} else {
						break
					}
				}

				resInet6Addr = append(resInet6Addr, &models.Inet6{
					Inet6Addr: fInet6[1],
					Prefixlen: fInet6[3] + resParamPrefix,
					ScopeID:   resScopeid,
				})
			}

			// checking and distribution nd6 options fields
			matchNd6Opts, _ := regexp.MatchString(`^nd6\soptions=[0-9]*<([A-Z0-9]*_*,*)*>`, v)
			if matchNd6Opts {
				resNd6Options = strings.ReplaceAll(v, "nd6 ", "")
			}

			// checking and distribution media fields
			if strings.Contains(v, "media: ") {
				resMedia = strings.ReplaceAll(v, "media: ", "")
			}

			// checking and distribution status fields
			if strings.Contains(v, "status: ") {
				resStatus = strings.ReplaceAll(v, "status: ", "")
			}

		}

		resNetInfInfo = append(resNetInfInfo, &models.NetIntInfo{
			NameInterface: resInf[0],
			Flags:         resFlag[0],
			Mtu:           strings.ReplaceAll(resMtu[0], "mtu ", ""),
			Options:       resOptions,
			Ether:         resEther,
			ConfigMember:  resConfigMem,
			Inet:          resInetAddr,
			Inet6:         resInet6Addr,
			Nd6Options:    resNd6Options,
			Media:         resMedia,
			Status:        resStatus,
		})
	}

	for _, res := range resNetInfInfo {
		fmt.Println(res.NameInterface)
		fmt.Println(res.Flags)
		fmt.Println(res.Mtu)
		fmt.Println(res.Options)
		fmt.Println(res.Ether)
		fmt.Println(res.ConfigMember)
		fmt.Println(*res.Inet)
		for _, resInet6 := range res.Inet6 {
			fmt.Println(resInet6.Inet6Addr, resInet6.Prefixlen, resInet6.ScopeID)
		}
		fmt.Println(res.Nd6Options)
		fmt.Println(res.Media)
		fmt.Println(res.Status)
	}

	return nil
}

func main() {
	/*
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

							resPowerInfo, err := PowerEquip()
							if err != nil {
								log.Printf("Failed execution command: %v", err)
							}
							fmt.Print(resPowerInfo)

							resPrinterInfo, err := PrinterEquip()
							if err != nil {
								log.Printf("Failed execution command: %v", err)
							}
							fmt.Print(resPrinterInfo)

							resDiskUsage, err := DiskUsage()
							if err != nil {
								log.Printf("Failed execution command: %v", err)
							}
							for _, res := range resDiskUsage {
								fmt.Println(*res)
							}

							resNetStatConn, err := NetStatConnInfo("-nap", "TCP")
							if err != nil {
								log.Printf("Failed execution command: %v", err)
							}
							for _, res := range resNetStatConn {
								fmt.Println(*res)
							}

							resNetStatRoute, err := NetStatRouteInfo("-rnf", "inet")
							if err != nil {
								log.Printf("Failed execution command: %v", err)
							}

							fmt.Println("IP version: ", resNetStatRoute.VerIP)
							for _, res := range resNetStatRoute.RouteInfo {
								fmt.Println(*res)
							}


		resArpTable, err := ARPTableInfo()
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}
		for _, res := range resArpTable {
			fmt.Println(*res)
		}
	*/
	NetworkIntInfo()
}
