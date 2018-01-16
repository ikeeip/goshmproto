// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types_darwin.go

// +build amd64,darwin

package syscall_shm

const (
	IPC_CREAT	= 0x200
	IPC_EXCL	= 0x400
)

const (
	IPC_RMID	= 0x0
	IPC_SET		= 0x1
	IPC_STAT	= 0x2
)

const (
	SHM_RDONLY	= 0x1000
	SHM_RND		= 0x2000
)

const (
	SizeofIpcPerm	= 0x18
	SizeofShmidDs	= 0x4c
)

type IpcPerm struct {
	Uid	uint32
	Gid	uint32
	Cuid	uint32
	Cgid	uint32
	Mode	uint16
	X_seq	uint16
	X_key	int32
}
type ShmidDs struct {
	Perm		IpcPerm
	Segsz		uint64
	Lpid		int32
	Cpid		int32
	Nattch		uint16
	Pad_cgo_0	[2]byte
	Pad_cgo_1	[8]byte
	Pad_cgo_2	[8]byte
	Pad_cgo_3	[8]byte
	Pad_cgo_4	[8]byte
}
