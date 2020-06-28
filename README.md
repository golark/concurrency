#### simple concurrency with go routines:
 - 2/2 simple go routine to handle connections concurrently
 - /2 limit the number of connections
 
 ### go routine lifetime
 - /2 go routine signalling finish
 - /2 waiting for go routines
 - /2 cancelling go routines
 - /2 returning error from go routines
 
 ### benchmarking
 - /3 benchmarking with go routines
 - /3 execution of a task with concurrency vs single thread of execution comparison
 - /2 stack growth of go routines
 - /5 garbage collection and go routines
 - /5 channel buffering for fast execution 
 - /5 limiting number of go routines executing concurrently
 
 ### architectures ( channels & go routines )
 - /1 pipeline
 - /1 select statement and infinite loop
 
 
 
 /37 total + /13 extra points that can be added  concurrency
 
##### Directory structure and files

    .
    ├── handleconn                  # handle incoming connections
    |   ├── example1                # read/write n bytes at a time, should better handle the read/write buffering
    |   ├── example2                # improves example 1 by using buffer to read/write
    ├── README.md                   # this file
 
 
