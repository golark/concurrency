Includes golang concurrency design patterns and best practices:
- Golang concurrency architecture such as pipeline, fan in/out, race-to-finish.
- Go routine lifetime control with self cancelling go routines or external signals to complete go routines.
- Returning error from Go routines though context or error channels.
- Benchmarking to demonstrate the effective go routine spin, limiting number of workers to achieve execution speed,
analysing performance bottlenecks.

#### simple concurrency with go routines:
 - 2/2 simple go routine to handle connections concurrently
 - 1/2 limit the number of connections
 
 ### go routine lifetime
 - 1/2 go routine signalling finish
 - /2 waiting for go routines
 - 2/2 cancelling go routines
 - 1/2 returning error from go routines
 - 3/3 timeout from go routine ( extra )
 - 1/3 context timeout (extra)
 
 ### benchmarking
 - /3 benchmarking with go routines
 - /3 execution of a task with concurrency vs single thread of execution comparison
 - /2 stack growth of go routines
 - /5 garbage collection and go routines
 - /5 channel buffering for fast execution 
 - /5 limiting number of go routines executing concurrently
 
 ### architectures ( channels & go routines )
 - 1/1 pipeline
 - 1/1 select statement and infinite loop
 - 2/2 limiting number of workers ( extra ) 
  
 /37 total + /13 extra points that can be added  concurrency
 
 @todo: time.After in case statement, stack growth ( downsides of using time.After in case statement )
 @todo: release semaphores as  a FILO manner
 
##### Directory structure and files

    .
    ├── handleconn                  # handle incoming connections
    |   ├── example1                # read/write n bytes at a time, should better handle the read/write buffering
    |   ├── example2                # improves example 1 by using bufio to read/write, fixed the issue with writing more than necessary bytes, logs when closing connection
    |   ├── example3                # uses io.Copy
    |   ├── example4                # limit number of connections
    ├── architecture                #
    |   ├── simplepipeline          # a barebone pipeline demonstrator
    |   ├── simplepipeline2         # cancel go routines
    |   ├── limitedworkers          # limited workers
    |   ├── workersrace             # multiple workers working on same task, first to finish wins
    ├── lifetime                    # 
    |   ├── closuretimeout          # simple query to dB with timeout packed in closure
    |   ├── internaltimeout         # forselect loop that times out
    |   ├── externaltimeout         # external channel signals timeout
    |   ├── exitgoroutines          # exit with exit channel closure broadcast
    |   ├── returnerror             # 
    |   ├── contexttimeout          # timeout if context expires
    |   ├── contextcancel           # go routine cancelled by another go routine
    ├── README.md                   # this file
