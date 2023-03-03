package warning

import (
  "testing"
  "log"
  "os"
)


func BenchmarkFATAL_ERROR(benchmark *testing.B){
  for i := 0; i < benchmark.N; i++{
    _, err := os.Open("file.txt");
      FATAL_ERROR(err);
  }
}

func BenchmarkStandardErrors(benchmark *testing.B){
  for i := 0; i < benchmark.N; i++{
      _, err := os.Open("file.txt");
        if err != nil{
          log.Fatalln(err);
        }
  }
}
