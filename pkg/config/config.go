package config

import (
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/json"
)

type PciSpec struct {
	VendorId  string `yaml:"vendorId" json:"vendorId"`
	DeviceId  string `yaml:"deviceId" json:"deviceId"`
	AddressId string `yaml:"addressId" json:"addressId"`
}

// node name - pci filter list
var NodeSupportedPciDevices = make(map[string][]PciSpec)

func ParsePciSpec(v string) []PciSpec {
	var ps []PciSpec
	err := json.Unmarshal([]byte(v), &ps)
	if err != nil {
		logrus.Errorf("parse pci spec str: %s, error: %v", v, err)
	}
	return ps
}
