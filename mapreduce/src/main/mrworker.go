package main

import (
         "https://github.com/jayaraju-bura/distributed_systems_learnings/tree/main/mapreduce/src/mr"
)
import (
        "os"
        "fmt"
        "log"
        "plugin"
)

func main() {
  if len(os.Args) != 2 {
      fmt.Fprintf(os.Stderr, "Usage: mrworker <input.so>")
      os.Exit(1)
  }
  
  mapf, reducef := loadPlugin(os.Args[1])
  mr.Worker(mapf, reducef)
}

func loadPlugin(filename string) (func(string, string) []mr.KeyValue, func(string, []string) string) {
  
  plg, err := plugin.Open(filename)
  if err != nil {
    log.Fatalf("Cannot load plugin %v", filename)
  }
  
  xmapf, err := plg.Lookup("Map")
  if err != nil {
    log.Fatalf("Cannot find Map in %v", filename)
  )
   mapf := xmapf.(func(string, string) []mr.KeyValue)
    
   xreducef, err := plg.Lookup("Reduce")
    if err != nil {
      log.Fatalf("Cannot find Reduce in %v", filename)
    }
    
    reducef := xreducef.(func(string, []string) string)
    return mapf, reducef
    
}
