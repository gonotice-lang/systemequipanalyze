package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
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
func InfoOS() (*models.OSInfo, []*ResultError) {
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

	oi := &models.OSInfo{
		OsName:   resOsInfo[0],
		KernVer:  resOsInfo[1],
		NodeName: resOsInfo[2],
		OsArch:   resOsInfo[3],
	}

	return oi, outErr
}

// HarwareEquip - General information equipment hardware
func HarwareEquip() (*models.DataHardware, error) {
	hard := &models.DataHardware{}

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
	mem := &models.DataMem{}

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
	stor := &models.DataStorage{}

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
	disp := &models.DataDisplay{}

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
	usb := &models.DataUSB{}

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
	netw := &models.DataNetwork{}

	cmd := exec.Command("system_profiler", "-json", "SPNetworkDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &netw)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return netw, err
}

// AirPortEquip - AirPort information
func AirPortEquip() (*models.DataAirPort, error) {
	airp := &models.DataAirPort{}

	cmd := exec.Command("system_profiler", "-json", "SPAirPortDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &airp)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return airp, err
}

// EthernetEquip - Ethernet information
func EthernetEquip() (*models.DataEthernet, error) {
	eth := &models.DataEthernet{}

	cmd := exec.Command("system_profiler", "-json", "SPEthernetDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &eth)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return eth, err
}

// PciEquip - PCI information
func PciEquip() (*models.DataPci, error) {
	pci := &models.DataPci{}

	cmd := exec.Command("system_profiler", "-json", "SPPCIDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &pci)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return pci, err
}

// PowerEquip - Power information
func PowerEquip() (*models.DataPower, error) {
	power := &models.DataPower{}

	cmd := exec.Command("system_profiler", "-json", "SPPowerDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &power)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return power, err
}

// PrinterEquip - Printer information
func PrinterEquip() (*models.DataPrinter, error) {
	printer := &models.DataPrinter{}

	cmd := exec.Command("system_profiler", "-json", "SPPrintersDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &printer)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return printer, err
}

// AppsEquip - information install applications
func AppsEquip() (*models.AppsInfo, error) {
	app := &models.AppsInfo{}

	cmd := exec.Command("system_profiler", "-json", "SPApplicationsDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &app)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return app, nil
}

// AudioEquip - audio equipment
func AudioEquip() (*models.AudioInfo, error) {
	audio := &models.AudioInfo{}

	cmd := exec.Command("system_profiler", "-json", "SPAudioDataType")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed execution command: %v", err)
	}

	json.Unmarshal(stdout, &audio)
	if err != nil {
		return nil, fmt.Errorf("Failed unmarshaling: %v", err)
	}

	return audio, nil
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

// NSConnInfo - Internet connect information
func NSConnInfo(key, arg string) ([]*models.NSConn, error) {
	var cmd *exec.Cmd

	newNetStatConn := new(models.NSConn)
	var resNetStatConn []*models.NSConn

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
			valFields = append(valFields, "")
		}
		for range valFields {
			newNetStatConn = &models.NSConn{
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

// NSRouteInfo - Table routing information
func NSRouteInfo(key, arg string) (*models.NSRoute, error) {
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
			valFields = append(valFields, "")
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

	resNetStatRoute := &models.NSRoute{
		VerIP:     verIP,
		RouteInfo: resRouteInfo,
	}

	return resNetStatRoute, err
}

// ARPTableInfo - information ARP Table
func ARPTableInfo() ([]*models.ARPInfo, error) {
	var templ = []string{"at", "on"}
	var resArp []*models.ARPInfo

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
		resArp = append(resArp, &models.ARPInfo{
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
func NetworkIntInfo() ([]*models.NetIntInfo, error) {
	var resNetInfInfo []*models.NetIntInfo

	cmd := exec.Command("ifconfig")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("\"%s\" %v", cmd, err)
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
				resNd6Options = v
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

	return resNetInfInfo, err
}

// ProcessInfo - information launch process
func ProcessInfo() ([]*models.ProcessInfo, error) {
	var resInfoProcess []*models.ProcessInfo

	cmd := exec.Command("ps", "-e")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("\"%s\" %v", cmd, err)
	}
	// delete last line break
	if len(stdout) > 0 {
		stdout = stdout[:len(stdout)-1]
	}
	resProc := strings.Split(string(stdout), "\n")
	resProc = removeIndex(resProc, 1, 0)

	for _, v := range resProc {
		resProc := strings.Fields(v)
		var (
			resPID      uint64
			resTty      string
			resProcTime string
			resProcCmd  string
		)
		for range resProc {
			resPID, err = strconv.ParseUint(resProc[0], 10, 32)
			if err != nil {
				panic(err)
			}
			lastField := ""
			for _, f := range resProc[3:] {
				lastField = lastField + " " + f
			}
			resTty = resProc[1]
			resProcTime = resProc[2]
			resProcCmd = lastField

		}
		resInfoProcess = append(resInfoProcess, &models.ProcessInfo{
			ProcPid:  resPID,
			ProcTty:  resTty,
			ProcTime: resProcTime,
			ProcCmd:  resProcCmd,
		})
	}
	return resInfoProcess, err
}

func main() {
	/*
		resOsInfo, resErr := InfoOS()
		for _, valr := range resErr {
			log.Printf("Error: %s in function number \"%d\", %v\n", valr.res.ErrorName, valr.res.NumbOccur, valr.err)
		}
		fmt.Print(resOsInfo)
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

		resDisplayInfo, err := DisplayEquip()
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}
		for _, resdispl := range resDisplayInfo.DisplayData {
			fmt.Println(resdispl.DisplayName)
			fmt.Println(resdispl.Vram)
			fmt.Println(resdispl.DeviceID)
			fmt.Println(resdispl.MetalFamily)
			for _, displndrvs := range resdispl.DisplayNdrvs {
				fmt.Println(displndrvs.DisplayEdid)
				fmt.Println(displndrvs.TypeDisName)
				fmt.Println(displndrvs.ProductID)
				fmt.Println(displndrvs.SerialNumb)
				fmt.Println(displndrvs.VendorID)
				fmt.Println(displndrvs.Week)
				fmt.Println(displndrvs.Year)
				fmt.Println(displndrvs.DisplayID)
				fmt.Println(displndrvs.Path)
				fmt.Println(displndrvs.PortDevice)
				fmt.Println(displndrvs.DisplayRegID)
				fmt.Println(displndrvs.Edid)
				fmt.Println(displndrvs.Pixels)
				fmt.Println(displndrvs.Resolution)
				fmt.Println(displndrvs.BrightAmb)
				fmt.Println(displndrvs.ConnectType)
				fmt.Println(displndrvs.Depth)
				fmt.Println(displndrvs.DisplayType)
				fmt.Println(displndrvs.Main)
				fmt.Println(displndrvs.Mirror)
				fmt.Println(displndrvs.Online)
				fmt.Println(displndrvs.PixelRes)
				fmt.Println(displndrvs.Resol)
			}
			fmt.Println(resdispl.RevisionID)
			fmt.Println(resdispl.Vendor)
			fmt.Println(resdispl.VramShared)
			fmt.Println(resdispl.Bus)
			fmt.Println(resdispl.DeviceType)
			fmt.Println(resdispl.Model)
		}
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
		for _, resAir := range resAirPortInfo.AirPortData {
			for _, resPort := range resAir.AirPortIntf {
				fmt.Println(resPort.IntfName)
				fmt.Println(resPort.AirDropChan)
				for _, wirenet := range resPort.LocalWireNet {
					fmt.Println(*wirenet)
				}
				fmt.Println(resPort.CapsAirDrop)
				fmt.Println(resPort.CapsAutoUn)
				fmt.Println(resPort.CapsWow)
				fmt.Println(*resPort.CurNetInfo)
				fmt.Println(resPort.StatusInfo)
				fmt.Println(resPort.SuppConnCh)
				fmt.Println(resPort.SuppPhyMod)
				fmt.Println(resPort.CardType)
				fmt.Println(resPort.ContryCode)
				fmt.Println(resPort.FirmWVer)
				fmt.Println(resPort.WireLocale)
				fmt.Println(resPort.WireMacAddr)
			}
			fmt.Println(*resAir.SoftInfo)
		}

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

		resNSConn, err := NSConnInfo("-nap", "UDP")
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}
		for _, res := range resNSConn {
			fmt.Println(*res)
		}

		resNSRoute, err := NSRouteInfo("-rnf", "inet")
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}
		fmt.Println("IP version: ", resNSRoute.VerIP)
		for _, res := range resNSRoute.RouteInfo {
			fmt.Println(*res)
		}

		resArpTable, err := ARPTableInfo()
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}
		for _, res := range resArpTable {
			fmt.Println(*res)
		}

		resNetInf, err := NetworkIntInfo()
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}
		for _, res := range resNetInf {
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

		resProcessInfo, err := ProcessInfo()
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}
		for _, resProc := range resProcessInfo {
			fmt.Println(*resProc)
		}

		resAppsInfo, err := AppsEquip()
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}

		for _, resApp := range resAppsInfo.Apps {
			fmt.Println(*resApp)
		}

		resAudioInfo, err := AudioEquip()
		if err != nil {
			log.Printf("Failed execution command: %v", err)
		}

		for _, resAudio := range resAudioInfo.AudioInf {
			for _, resAudioIntf := range resAudio.ItemsAudio {
				fmt.Println(*resAudioIntf)
			}
			fmt.Println(resAudio.NameAudio)
		}
	*/
}
