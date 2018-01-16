// +build ignore

package syscall_shm

/*
#define _LARGEFILE_SOURCE
#define _LARGEFILE64_SOURCE
#define _FILE_OFFSET_BITS 64
#define _GNU_SOURCE
#include <linux/ipc.h>
#include <linux/shm.h>
*/
import "C"

const (
    IPC_CREAT	= C.IPC_CREAT		/* Create entry if key does not exist */
    IPC_EXCL	= C.IPC_EXCL		/* Fail if key exists */
)

const (
    IPC_RMID	= C.IPC_RMID		/* Remove identifier */
    IPC_SET		= C.IPC_SET			/* Set options */
    IPC_STAT	= C.IPC_STAT		/* Get options */

)

const (
    SHM_RDONLY	= C.SHM_RDONLY		/* [XSI] Attach read-only (else read-write) */
    SHM_RND		= C.SHM_RND			/* [XSI] Round attach address to SHMLBA */
)

const (
    SizeofIpcPerm = C.sizeof_struct_ipc_perm
    SizeofShmidDs = C.sizeof_struct_shmid_ds
)

type IpcPerm C.struct_ipc_perm
type ShmidDs C.struct_shmid_ds
