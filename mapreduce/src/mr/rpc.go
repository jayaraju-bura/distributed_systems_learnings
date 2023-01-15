package mr


import "os"
import "strconv"


type ExampleArgs struct {
        X int
}

type ExampleReply struct{
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
type NotifyReply struct{

}

type NotifyIntermediateArgs struct{
        Reduceindex int
        File        string
}

type NotifyMapSuccessArgs struct{
        File string
}

type NotifyReduceSuccessArgs struct{
        Reduceindex int
}

func mastersock() string {
  sock := "/var/tmp/456-mr"
  sock += strconv.Itoa(os.Getuid())
  return sock
}
