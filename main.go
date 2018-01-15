package main

import (
	"fmt"
	"flag"
	"net/http"
	"unsafe"
	"encoding/binary"
	"github.com/ikeeip/goshmproto/syscall_shm"
)

var (
	listenAddress = flag.String("listen-address", ":8889", "The address to listen on for HTTP requests.")
	metricsPath   = flag.String("telemetry-path", "/metrics", "Path under which to expose metrics.")

	segmentId int
	bufferLen = 1024
	sharedKey = 0xff01
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	//attach to shm segment
	sharedMemory, err := syscall_shm.Shmat(segmentId, 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("shared memory attached at address 0x%016x\n", sharedMemory)

	// Slice memory layout
	// https://golang.org/pkg/reflect/#SliceHeader
	// https://blog.golang.org/go-slices-usage-and-internals
	//var sl = reflect.SliceHeader{sharedMemory, 1024, 1024}
	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{sharedMemory, bufferLen, bufferLen}
	// Use unsafe to turn sl into a []byte.
	b := *(*[]byte)(unsafe.Pointer(&sl))

	dataLen := binary.LittleEndian.Uint64(b[0:8])
	if dataLen > 0 && int(dataLen) <= bufferLen - 8 {
		w.Write(b[8:dataLen])
	}

	//detach from shm segment
	err = syscall_shm.Shmdt(sharedMemory)
	if err != nil {
		panic(err)
	}
	fmt.Printf("shared memory dettached from address 0x%016x\n", sharedMemory)
}

func initShm() {
	fmt.Printf("shared memory key 0x%08x\n", sharedKey)
	//allocate shm segment with 0644 mode
	var err error
	segmentId, err = syscall_shm.Shmget(sharedKey, bufferLen, syscall_shm.IPC_CREAT | 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("shared memory segment id %v\n", segmentId)
}

func shutdownShm() {
	// mark shm segment to be removed
	var err = syscall_shm.Shmctl(segmentId, syscall_shm.IPC_RMID, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("shared memory segment %v marked for remove\n", segmentId)
}

func main() {
	flag.Parse()

	fmt.Printf("Starting test_shm\n")

	initShm()
	defer shutdownShm()

	fmt.Printf("Listening on %v\n", *listenAddress)
	http.HandleFunc(*metricsPath, metricsHandler)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		panic(err)
	}
}
