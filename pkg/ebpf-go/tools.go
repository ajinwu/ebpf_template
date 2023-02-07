package ebpfgo

import (
	"fmt"
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

type CiliumEBPFRuntime struct {
	Objects *bpfObjects
	Links   map[string]link.Link
}

func (c *CiliumEBPFRuntime) RemoveMemoryLimit() error {
	return rlimit.RemoveMemlock()
}

func (c *CiliumEBPFRuntime) LoadBpfObjects(opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}
	// spec.Maps["map_payload_buffer"].Pinning = ebpf.PinByName
	return spec.LoadAndAssign(c.Objects, opts)
}

func New() *CiliumEBPFRuntime {
	return &CiliumEBPFRuntime{
		Objects: &bpfObjects{},
		Links:   make(map[string]link.Link),
	}
}

func (c *CiliumEBPFRuntime) CreateCiliumEBPFRuntime(objname, PinPath string) error {
	err := c.LoadBpfObjects(&ebpf.CollectionOptions{
		Maps: ebpf.MapOptions{
			PinPath: PinPath,
		},
	})
	if err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	return nil
}

func (c *CiliumEBPFRuntime) CreateLink(objname string) error {
	var err error
	// TODO 修改这里

	// c.Links[objname], err = link.AttachCgroup(link.CgroupOptions{
	// 	Path:    objname,
	// 	Attach:  ebpf.AttachCGroupInetEgress,
	// 	Program: c.Objects.CountEgressPackets,
	// })

	// c.Links[objname], err = link.Kprobe(objname, c.Objects.KprobeExecve, nil)

	// c.Links[objname], err = link.AttachTracing(link.TracingOptions{
	// 	Program: c.Objects.TcpClose,
	// })
	// c.Links[objname], err = link.AttachXDP(link.XDPOptions{
	// 	Program:   c.Objects.XdpProgFunc,
	// 	Interface: 2,
	// })
	if err != nil {
		return err
	}
	return nil
}

func (c *CiliumEBPFRuntime) Close() error {
	if c.Objects != nil {
		if err := c.Objects.Close(); err != nil {
			return err
		}
	}
	if c.Links != nil {
		for _, link := range c.Links {
			if err := link.Close(); err != nil {
				return fmt.Errorf("Link close error: %w\n", link.Close())
			}
		}
	}
	return nil
}
