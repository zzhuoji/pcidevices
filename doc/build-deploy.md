

### 打包

```shell
# 打包的方式
# 先打二进制包
./scripts/build
# 然后打镜像包
./scripts/package
```

### 部署

> deploy 目录，直接指向 bash entrypoint.sh

```shell
# 找可配置的PCI设备，这里以网卡为例
[root@arm01 ~]# nmcli d s
DEVICE          TYPE         STATE                   CONNECTION 
docker0         bridge       connected (externally)  docker0    
virbr0          bridge       connected (externally)  virbr0     
enp125s0f1      ethernet     disconnected            --         
enp125s0f2      ethernet     disconnected            --         
enp125s0f3      ethernet     disconnected            --         
enp1s0f0np0     ethernet     disconnected            --         
enp1s0f1np1     ethernet     disconnected            --  

[root@arm01 ~]# ethtool -i enp1s0f0np0
driver: mlx5_core
version: 5.0-0
firmware-version: 16.32.1010 (HUA0000000024)
expansion-rom-version: 
bus-info: 0000:01:00.0
supports-statistics: yes
supports-test: yes
supports-eeprom-access: no
supports-register-dump: no
supports-priv-flags: yes

[root@arm01 ~]# lspci -nnks 0000:01:00.0
01:00.0 Ethernet controller [0200]: Mellanox Technologies MT27800 Family [ConnectX-5] [15b3:1017]
        Subsystem: Huawei Technologies Co., Ltd. Device [19e5:d135]
        Kernel driver in use: mlx5_core
        Kernel modules: mlx5_core

# 所以上面网卡 enp1s0f0np0 对应的值为:
vendorId: 15b3
deviceId: 1017
addressId: 0000:01:00.0


kubectl apply -f sample/supported-pci.yaml


# 创建 对应的 pdc 

# 查看
[root@arm01 device-plugins]# lspci -nnk -d 15b3:
01:00.0 Ethernet controller [0200]: Mellanox Technologies MT27800 Family [ConnectX-5] [15b3:1017]
        Subsystem: Huawei Technologies Co., Ltd. Device [19e5:d135]
        Kernel driver in use: vfio-pci
        Kernel modules: mlx5_core
01:00.1 Ethernet controller [0200]: Mellanox Technologies MT27800 Family [ConnectX-5] [15b3:1017]
        Subsystem: Huawei Technologies Co., Ltd. Device [19e5:d135]
        Kernel driver in use: mlx5_core
        Kernel modules: mlx5_core

```