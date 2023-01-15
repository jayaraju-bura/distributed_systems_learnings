package main

import (
        "https://github.com/jayaraju-bura/distributed_systems_learnings/tree/main/mapreduce/src/mr"
)
import (
        "unicode"
        "strings"
        "stringconv"
)

func Map(filename string, contents string) []mr.KeyValue {
  funtn := func(r rune) bool { !unicode.IsLetter(r)}
  
  kva := []mr.KeyValue{}
  words := strings.FieldFunc(contents, funtn)
  for _, wrd := range words {
    kv := mr.KeyValue{wrd, "1"}
    kva = append(kva, kv)
  }
  
  return kva
}

func Reduce(key string, values []string) string {
  
  return stringconv.Itoa(len(values))
}
