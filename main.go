package main

import (
	ebpfgo "ebpf/pkg/ebpf-go"
	"log"
	"os"
	"path"
	"time"
)

const bpfFSPath = "/sys/fs/bpf"

func main() {
	e := ebpfgo.New()
	// TODO 这里修改
	objname := "tp/syscalls/sys_enter_read"
	PinPath := path.Join(bpfFSPath, objname)
	if err := os.MkdirAll(PinPath, os.ModePerm); err != nil {
		log.Fatalf("failed to create bpf fs subpath: %+v", err)
	}
	err := e.CreateCiliumEBPFRuntime(objname, PinPath)
	if err != nil {
		log.Fatal(err)
	}
	// e.CreateLink(objname)
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
	}
}
