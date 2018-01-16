// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types_linux.go

// +build amd64,linux

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
	SizeofIpcPerm	= 0x1c
	SizeofShmidDs	= 0x58
)

type IpcPerm struct {
	Key		int32
	Uid		uint32
	Gid		uint32
	Cuid		uint32
	Cgid		uint32
	Mode		uint32
	Seq		uint16
	Pad_cgo_0	[2]byte
}
type ShmidDs struct {
	Perm		IpcPerm
	Segsz		int32
	Atime		int64
	Dtime		int64
	Ctime		int64
	Cpid		int32
	Lpid		int32
	Nattch		uint16
	Unused		uint16
	Pad_cgo_0	[4]byte
	Unused2		*byte
	Unused3		*byte
}
