
> PCI 设备与 usb 设备都不支持热插拔, 需要硬重启重新调度这一过程
> 调度过去后除非解绑，否则不允许删除 claims, (两边怎么一致呢?怎么知道设备已经挂载出去了,删的话，要把 vm 对应的 yaml 也删除,再关机才行)


### pci 设备 bind 与 unbind 核心逻辑
- 记录原始 driver ghw.PCI()
- attemptToEnablePassthrough() 判断是否为 vfio-pci 驱动
  - 在 /sys/bus/pci/drivers/vfio-pci 下匹配对应设备的 address
    如果不是则先 unbind origin driver, 然后 bind 到 vfio-pci driver (就是写入 /sys/bus/pci/drivers/%s 驱动的 bind 与 unbind 文件)



### usb 设备 
- usbdevice Register 函数里面进行初始化 handler.reconcile()
  - 利用 deviceplugins.WalkUSBDevices 获取所有 map[int][]*USBDevice
    - 遍历/sys/bus/usb/devices目录, 扫描所有USB设备并解析设备信息, 返回按Vendor ID分组的设备映射
      解析设备的uevent文件, 提取设备属性（Vendor、Product、Bus、DeviceNumber等,解析PCI地址映射
    - 使用fsnotify监控/dev/bus/usb/目录变化, 当设备插拔时触发调度逻辑, 处理设备热插拔事件
- 最后回到 Register 阶段, 将上述获取到的 USBDevice 元数据转换成 CR 资源


