package warning

import (
  "log"
	"os"
	"testing"
)

func ExampleLOG_ERRORS(){
  filename := "file.adah";
  open, err := os.Open(filename);
    LOG_ERRORS(err, filename + " was not exists");
  defer open.Close();
  //Output:
  // file.adah was not exists open file.adah: no such file or directory
}


func BenchmarkLOG_ERRORS(benchmark *testing.B) {
  for i := 0; i < benchmark.N; i++{
    filename := "file.adah";
    open, err := os.Open(filename);
      LOG_ERRORS(err, filename + " was not exists");
    defer open.Close();
  }
}

func BenchmarkLOG_ERRORS_BASIC(benchmark *testing.B) {
  for i := 0; i < benchmark.N; i++{
    filename := "file.adah";
    open, err := os.Open(filename);
    if err != nil{
      log.Println(err, filename + " was not exists");
    }
    defer open.Close();
  }
}

func BenchmarkPANIC_ERROR(benchmark *testing.B) {
  for i := 0; i < benchmark.N; i++{
    filename := "file.adah";
    open, err := os.Open(filename);
      PANIC_ERROR(err);
    defer open.Close();
  }
}

func BenchmarkPANIC_ERROR_BASIC(benchmark *testing.B) {
  for i := 0; i < benchmark.N; i++{
    filename := "file.adah";
    open, err := os.Open(filename);
    if err != nil{
      log.Panicln(err);
    }
    defer open.Close();
  }
}
