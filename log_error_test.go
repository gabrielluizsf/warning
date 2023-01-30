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

func ExamplePANIC_ERROR(){
  filename := "file.adah";
  open, err := os.Open(filename);
    PANIC_ERROR(err);
  defer open.Close();
  /* Output:
   2023/01/30 13:36:33 open file.adah: no such file or directory
panic: open file.adah: no such file or directory


goroutine 1 [running]:
log.Panicln({0xc00008ef48?, 0x9?, 0x0?})
        $HOME/.asdf/installs/golang/1.19.3/go/src/log/log.go:402 +0x65
github.com/theGOURL/warning.PANIC_ERROR(...)
        $HOME/projects/warning/log_error.go:16
main.Testing_the_PANIC_ERROR_function()
        $HOME/projects/warning/internal/main.go:21 +0x72
main.main()
        $HOME/projects/warning/internal/main.go:9 +0x17
exit status 2
*/
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