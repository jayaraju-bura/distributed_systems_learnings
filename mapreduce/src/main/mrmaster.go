package main


import (
        "https://github.com/jayaraju-bura/distributed_systems_learnings/tree/main/mapreduce/src/mr"
)

import (
        "time"
        "os"
        "fmt"
)

func main() {
  
  if len(os.Args) < 2 {
           fmt.Fprintf(os.Stderr, "Usage : go run mrmaster.go <input_file>")
           os.Exit(1)
  }
  
  master_inst := mr.MakeMaster(os.Args[1:], 10)
  for master_inst.Done() == false {
                    time.Sleep(time.Second)
  }
  
  time.Sleep(time.Second)
}
