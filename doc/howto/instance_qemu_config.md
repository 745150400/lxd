(instance-qemu-config)=
# How to override QEMU configuration

For VM instances, LXD configures QEMU via a somewhat undocumented configuration
file format passed to QEMU with the `-readconfig` command-line option, with
each instance having a configuration file generated before boot. The generated
configuration file can be found at `/var/log/lxd/[instance-name]/qemu.conf`.

The default configuration works fine for LXD most common use case: Modern UEFI
guests with VirtIO devices. In some situations however, it can be desirable to
override the generated configuration:

- Running an old guest OS that doesn't support UEFI.
- Specify custom virtual devices when VirtIO is not supported by the guest OS .
- Add devices not supported by LXD before the machines boots.
- Remove devices that conflict with the guest OS.

This level of customization can be achieved through the `raw.qemu.conf` configuration
option, which supports a format similar to `qemu.conf` with some additions.
Here's how to override the default `virtio-gpu-pci` GPU driver:

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
    driver = "qxl-vga"
```

The above would replace the corresponding section/key in the generated configuration
file. Since `raw.qemu.conf` is a multi-line configuration option, multiple
sections/keys can be modified.

It is also possible to entirely remove sections/keys by specifying a section
without any keys:

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
```

To remove a key, specify an empty string as the value:

```
raw.qemu.conf: |-
    [device "qemu_gpu"]
    driver = ""
```

The configuration file format used by QEMU allows multiple sections with the
same name. Here's a piece of the configuration generated by LXD:

```
[global]
driver = "ICH9-LPC"
property = "disable_s3"
value = "1"

[global]
driver = "ICH9-LPC"
property = "disable_s4"
value = "1"
```

To specify which section will be overridden, an index can be specified like this:

```
raw.qemu.conf: |-
    [global][1]
    value = "0"
```

Section indexes start at 0 (which is the default value when not specified), so
the `raw.qemu.conf` above example would generate this:

```
[global]
driver = "ICH9-LPC"
property = "disable_s3"
value = "1"

[global]
driver = "ICH9-LPC"
property = "disable_s4"
value = "0"
```

To add new sections, simply specify section names that are not present in the
configuration file.
