package gost

// A broadcaster is the part dealing with receiving
// task and sending them to all the intersted workers.
type Broadcaster interface {
    Init(config Config) error       // init the broadcaster
    Close()
    Broadcast(Task) error           // do your job, broadcaster.
    // TODO Register(Channel, Topic) 
}
