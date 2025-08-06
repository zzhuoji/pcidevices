

## 注册后 node 节点不显示


## kubelet 重启后, 需要二次注册 


### ListAndWatch 267 
```shell
time="2024-09-10T11:41:37Z" level=info msg="monitoredDevices: map[0000:7d:00.2:/proc/1/root/dev/vfio/29]"
time="2024-09-10T11:41:37Z" level=info msg="health Event: REMOVE        \"/var/lib/kubelet/device-plugins/kubevirt-huawei.com-HNS_GE_10GE_25GE_RDMA_NETWORK_CONTROLLER.sock\""
time="2024-09-10T11:41:37Z" level=info msg="monitoredDevices: map[0000:7d:00.2:/proc/1/root/dev/vfio/29]"
time="2024-09-10T11:41:37Z" level=info msg="health Event: REMOVE        \"/var/lib/kubelet/device-plugins/kubevirt-huawei.com-HNS_GE_10GE_25GE_RDMA_NETWORK_CONTROLLER.sock\""
time="2024-09-10T11:41:37Z" level=info msg="monitoredDevices: map[0000:7d:00.2:/proc/1/root/dev/vfio/29]"
time="2024-09-10T11:41:37Z" level=info msg="health Event: REMOVE        \"/var/lib/kubelet/device-plugins/kubevirt-huawei.com-HNS_GE_10GE_25GE_RDMA_NETWORK_CONTROLLER.sock\""
time="2024-09-10T11:41:37Z" level=info msg="monitoredDevices: map[0000:7d:00.1:/proc/1/root/dev/vfio/28 0000:7d:00.3:/proc/1/root/dev/vfio/30]"
{"component":"","level":"info","msg":"huawei.com/HNS_GE_10GE_25GE_RDMA_NETWORK_CONTROLLER device plugin failed to deregister: rpc error: code = Unavailable desc = transport is closing","pos":"device_manager.go:265","reason":"rpc error: code = Unavailable desc = transport is closing","timestamp":"2024-09-10T11:41:37.721354Z"}
time="2024-09-10T11:41:37Z" level=info msg="huawei.com/HNS_GE_10GE_25GE_RDMA_NETWORK_CONTROLLER: Successfully restarted device plugin server. Terminating."
panic: close of closed channel

goroutine 1515 [running]:
github.com/harvester/pcidevices/pkg/deviceplugins.(*PCIDevicePlugin).ListAndWatch(0x400219e410, 0x1bc9100?, {0x2061ac0, 0x4004cf3230})
        /go/src/github.com/harvester/pcidevices/pkg/deviceplugins/device_manager.go:267 +0x5f0
kubevirt.io/kubevirt/pkg/virt-handler/device-manager/deviceplugin/v1beta1._DevicePlugin_ListAndWatch_Handler({0x1bc9100, 0x400219e410}, {0x205fe60, 0x4001b503c0})
        /go/pkg/mod/kubevirt.io/kubevirt@v1.1.1/pkg/virt-handler/device-manager/deviceplugin/v1beta1/api.pb.go:1157 +0x120
google.golang.org/grpc.(*Server).processStreamingRPC(0x40007625a0, {0x2065680, 0x4004a62d00}, 0x4004c97320, 0x4004b005d0, 0x3314280, 0x0)
        /go/pkg/mod/google.golang.org/grpc@v1.56.3/server.go:1629 +0xf74
google.golang.org/grpc.(*Server).handleStream(0x40007625a0, {0x2065680, 0x4004a62d00}, 0x4004c97320, 0x0)
        /go/pkg/mod/google.golang.org/grpc@v1.56.3/server.go:1716 +0x780
google.golang.org/grpc.(*Server).serveStreams.func1.1()
        /go/pkg/mod/google.golang.org/grpc@v1.56.3/server.go:947 +0xb4
created by google.golang.org/grpc.(*Server).serveStreams.func1 in goroutine 1570
        /go/pkg/mod/google.golang.org/grpc@v1.56.3/server.go:958 +0x14c
```