package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// OSCmds - commands os
type OSCmds struct {
	MacOsXCmds MacOsXCmds `yaml:"macoscmds"`
	//LinuxCmds  LinuxCmds  `yaml:"http"`
}

// MacOsXCmds - Mac OS X commands
type MacOsXCmds struct {
	IfConfigCmd   IfConfigCmd   `yaml:"ifconfig"`
	UnameCmd      UnameCmd      `yaml:"uname"`
	DiskUsageCmd  DiskUsageCmd  `yaml:"diskusage"`
	ArpCmd        ArpCmd        `yaml:"arp"`
	SystemInfoCmd SystemInfoCmd `yaml:"systeminfo"`
	SysProfileCmd SysProfileCmd `yaml:"sysprofile"`
	NetStatCmd    NetStatCmd    `yaml:"netstat"`
}

// IfConfigCmd - network interfaces information
type IfConfigCmd struct {
	CmdIfConfig string `yaml:"cmd"`
}

// UnameCmd - Command Uname information OS
type UnameCmd struct {
	CmdUname string `yaml:"cmd"`
	KeyUname string `yaml:"key"`
}

// DiskUsageCmd - Usage disk information
type DiskUsageCmd struct {
	CmdDiskUsage string `yaml:"cmd"`
	KeyDiskUsage string `yaml:"key"`
}

// ArpCmd - Viewing arp table information
type ArpCmd struct {
	CmdArp string `yaml:"cmd"`
	KeyArp string `yaml:"key"`
}

// SystemInfoCmd - Viewing cpu information
type SystemInfoCmd struct {
	CmdSystemInfo string `yaml:"cmd"`
	KeySystemInfo string `yaml:"key"`
	ArgSystemInfo string `yaml:"arg"`
}

// SysProfileCmd - Viewing all system information
type SysProfileCmd struct {
	CmdSysProfile  string `yaml:"cmd"`
	KeySysProfile  string `yaml:"key"`
	ArgsSysProfile string `yaml:"args"`
}

// NetStatCmd - Information network connections services
type NetStatCmd struct {
	CmdNetStat  string `yaml:"cmd"`
	KeysNetStat string `yaml:"keys"`
	ArgsNetStat string `yaml:"args"`
}

/*
type LinuxCmds struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}*/

// NewCmd - new command
func NewCmd(path string) (OSCmds, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return OSCmds{}, err
	}
	conf := OSCmds{}
	if err := yaml.Unmarshal(file, &conf); err != nil {
		return OSCmds{}, err
	}
	return conf, nil
}
