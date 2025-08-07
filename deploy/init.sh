
helm search repo pci -l

helm pull harvester/harvester-pcidevices-controller --version=1.5.1 -d charts/ --untar