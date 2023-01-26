package warning

import (
  "testing"
  "log"
  "os"
)

func ExampleFATAL_ERROR(){
    _, err := os.Open("file.txt");
      FATAL_ERROR(err);
  //Output:
  //2023/01/26 10:06:50 open file.txt: no such file or directory
//exit status 1
}

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
