//go:build linux
// +build linux

package pcideviceclaim

import (
	"github.com/sirupsen/logrus"
	"github.com/u-root/u-root/pkg/kmodule"
)

func loadVfioDrivers() {
	for _, driver := range []string{"vfio-pci", "vfio_iommu_type1"} {
		logrus.Infof("Loading driver %s", driver)
		if err := kmodule.Probe(driver, ""); err != nil {
			logrus.Error(err)
		}
	}
}
