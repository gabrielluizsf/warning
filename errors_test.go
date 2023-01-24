package warning

import (
  "os/exec"
  "fmt"
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
