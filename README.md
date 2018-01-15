# Test Shared Memory


```
// doc: https://users.cs.cf.ac.uk/Dave.Marshall/C/node27.html
// from: http://web.archive.org/web/20150524005553/http://advancedlinuxprogramming.com/alp-folder/alp-ch05-ipc.pdf
//
// https://github.com/apple/darwin-xnu/blob/master/bsd/sys/shm.h
// https://github.com/apple/darwin-xnu/blob/master/bsd/sys/ipc.h
#include <stdio.h>
#include <sys/shm.h>
#include <sys/stat.h>

int main () {
    int segment_id;
    char* shared_memory;
    struct shmid_ds shmbuffer;
    int segment_size;
    const int shared_segment_size = 0x6400;

    /* Allocate a shared memory segment. */
    segment_id = shmget (IPC_PRIVATE, shared_segment_size, IPC_CREAT | IPC_EXCL | S_IRUSR | S_IWUSR);

    /* Attach the shared memory segment. */
    shared_memory = (char*) shmat (segment_id, 0, 0);
    printf ("shared memory attached at address %p\n", shared_memory);

    /* Determine the segment’s size. */
    shmctl (segment_id, IPC_STAT, &shmbuffer);
    segment_size = shmbuffer.shm_segsz;
    printf ("segment size: %d\n", segment_size);

    /* Write a string to the shared memory segment. */
    sprintf (shared_memory, "Hello, world.");

    /* Detach the shared memory segment. */
    shmdt (shared_memory);

    /* Reattach the shared memory segment, at a different address. */
    shared_memory = (char*) shmat (segment_id, NULL, 0);
    printf ("shared memory reattached at address %p\n", shared_memory);

    /* Print out the string from shared memory. */
    printf ("%s\n", shared_memory);

    /* Detach the shared memory segment. */
    shmdt (shared_memory);

    /* Deallocate the shared memory segment. */
    shmctl (segment_id, IPC_RMID, 0);

    return 0;
}
```

```
#include<sys/ipc.h>
#include<sys/shm.h>

int shmid;
int shmkey = 12222;//u can choose it as your choice

int main()
{
  //now your main starting
  shmid = shmget(shmkey,1024,IPC_CREAT);
  // 1024 = your preferred size for share memory
  // IPC_CREAT  its a flag to create shared memory

  //now attach a memory to this share memory
  char *shmpointer = shmat(shmid,NULL);

  //do your work with the shared memory
  //read -write will be done with the *shmppointer
  //after your work is done deattach the pointer
  shmdt(&shmpointer, NULL);
}
```