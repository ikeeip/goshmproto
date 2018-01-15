package syscall_shm

// +build darwin linux

import (
	"syscall"
	"unsafe"
	"C"
)

const (
	IPC_CREAT	= 001000		/* Create entry if key does not exist */
	IPC_EXCL	= 002000		/* Fail if key exists */
)

//int shmget(key_t key, size_t size, int shmflg);
func Shmget(key int, size int, shmflg int) (shmid int, err error) {
	r0, _, e1 := syscall.Syscall(syscall.SYS_SHMGET, uintptr(key), uintptr(size), uintptr(shmflg))
	shmid = int(r0)
	if e1 != 0 {
		err = syscall.Errno(e1)
	}
	return
}

const (
	SHM_RDONLY	= 010000	/* [XSI] Attach read-only (else read-write) */
	SHM_RND		= 020000	/* [XSI] Round attach address to SHMLBA */
)

//void * shmat(int shmid, const void *shmaddr, int shmflg);
func Shmat(shmid int, shmaddr uintptr, shmflg int) (_shmaddr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(syscall.SYS_SHMAT, uintptr(shmid), shmaddr, uintptr(shmflg))
	_shmaddr = uintptr(r0)
	if e1 != 0 {
		err = syscall.Errno(e1)
	}
	return
}

//int shmdt(const void *shmaddr);
func Shmdt(shmaddr uintptr) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)
	if e1 != 0 {
		err = syscall.Errno(e1)
	}
	return
}

const (
	IPC_RMID	= 0		/* Remove identifier */
	IPC_SET		= 1		/* Set options */
	IPC_STAT	= 2		/* Get options */

)

//TODO: fix me. dirty hack. golang does not have several types in "C" package :(
//https://github.com/apple/darwin-xnu/blob/master/bsd/sys/_types.h
//https://github.com/apple/darwin-xnu/blob/master/bsd/i386/_types.h
//https://github.com/apple/darwin-xnu/blob/master/bsd/sys/_types/_key_t.h
type pid_t int32
type uid_t uint32
type gid_t uint32
type time_t C.long
type mode_t uint16
type key_t int32

type IpcPerm struct {
	uid  uid_t				//	uid_t           uid;   /* Owner's user ID */
	gid  gid_t				//	gid_t           gid;   /* Owner's group ID */
	cuid uid_t				//	uid_t           cuid;  /* Creator's user ID */
	cgid gid_t				//	gid_t           cgid;  /* Creator's group ID */
	mode mode_t				//	mode_t          mode;  /* r/w permission (see chmod(2)) */
	_seq C.ushort			//	unsigned short  _seq;  /* Reserved for internal use */
	_key key_t				//	key_t           _key;  /* Reserved for internal use */
}

type ShmidDs struct {
	shm_perm  	 IpcPerm	//	struct ipc_perm  shm_perm;     /* operation permissions */
	shm_segsz  	 C.int		//	int              shm_segsz;    /* size of segment in bytes */
	shm_lpid  	 pid_t		//	pid_t            shm_lpid;     /* pid of last shm op */
	shm_cpid  	 pid_t		//	pid_t            shm_cpid;     /* pid of creator */
	shm_nattch   C.short	//	short            shm_nattch;   /* # of current attaches */
	shm_atime    time_t		//	time_t           shm_atime;    /* last shmat() time*/
	shm_dtime    time_t		//	time_t           shm_dtime;    /* last shmdt() time */
	shm_ctime    time_t		//	time_t           shm_ctime;    /* last change by shmctl() */
	shm_internal *byte		//	void            *shm_internal; /* sysv stupidity */
}

//int shmctl(int shmid, int cmd, struct shmid_ds *buf);
func Shmctl(shmid int,cmd int, buf *ShmidDs) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_SHMCTL, uintptr(shmid), uintptr(cmd), uintptr(unsafe.Pointer(buf)))
	if e1 != 0 {
		err = syscall.Errno(e1)
	}
	return
}
