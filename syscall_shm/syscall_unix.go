// +build darwin linux

package syscall_shm

import (
	"syscall"
	"unsafe"
	"C"
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

//int shmctl(int shmid, int cmd, struct shmid_ds *buf);
func Shmctl(shmid int,cmd int, buf *ShmidDs) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_SHMCTL, uintptr(shmid), uintptr(cmd), uintptr(unsafe.Pointer(buf)))
	if e1 != 0 {
		err = syscall.Errno(e1)
	}
	return
}
