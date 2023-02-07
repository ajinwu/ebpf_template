package ebpfgo

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf ../ebpf-c/ebpf.c -- -I../ebpf-c/headers
