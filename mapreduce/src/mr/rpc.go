package mr


import "os"
import "strconv"


type ExampleArgs struct {
        X int
}

type ExampleReply {
        Y int
}

type MrArgs struct {

}

type MrReply struct  {
        MapFileName string
        TaskType    string
        Index       int
        NReduce     int
        Files       []string
}
type NotifyReply {

}

type NotifyIntermediateArgs {
        Reduceindex int
        File        string
}

type NotifyMapSuccessArgs {
        File string
}

type NotifyReduceSuccessArgs {
        Reduceindex int
}

func mastersock() string {
  sock := "/var/tmp/456-mr"
  sock += strconv.Itoa(os.Getuid())
  return sock
}
