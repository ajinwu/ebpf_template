CLANG ?= clang
CFLAGS := -O2 -g -Wall -Werror -I /usr/include/x86_64-linux-gnu -v $(CFLAGS)


build: mod_tidy generate
	go build .


mod_tidy:
	go mod tidy

generate: mod_tidy
generate: export BPF_CLANG := $(CLANG)
generate: export BPF_CFLAGS := $(CFLAGS)
generate:
	go generate ./pkg/ebpf-go...


clean:
	rm pkg/ebpf-go/bpf_bpfeb.*
	rm pkg/ebpf-go/bpf_bpfel.*

	