package warning

import (
  "os/exec"
  "fmt"
  "testing"
)
func ExamplePRINT_DEFAULT_ERRORS(){
	out, err := exec.Command("whoami").Output()
	PRINT_DEFAULT_ERRORS(err, "INVALID COMMAND")
	output := string(out[:])
	fmt.Println(output)
  //Output: 
  //[Command successfully executed]
 //gabriel
}

func BenchmarkPRINT_DEFAULT_ERRORS(benchmark *testing.B){
  for i := 0; i < benchmark.N; i++{
  out, err := exec.Command("disfhdisfias").Output();
	PRINT_DEFAULT_ERRORS(err, "INVALID COMMAND");
	output := string(out[:]);
	fmt.Println(output);
  }
}
func BenchmarkPRINT_DEFAULT_ERRORS_BASIC(benchmark *testing.B){
  for i := 0; i < benchmark.N; i++{
  out, err := exec.Command("disfhdisfias").Output();
    if err != nil{
      fmt.Println("INVALID COMMAND");
    }else{
      fmt.Println("Command successfully executed");
    }
	output := string(out[:]);
	fmt.Println(output);
  }
}
