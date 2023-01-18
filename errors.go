package warning

import "fmt"

// This function prints the errors received by parameter,
//along with the message received by parameter, 
//generating a customized error message for the console.
func PRINT_DEFAULT_ERRORS(ERROR error,MESSAGE string){
	if ERROR != nil{
	fmt.Println("[MESSAGE]:",MESSAGE,"\n[ERROR]:",ERROR);
	}else{
    fmt.Println("[Command successfully executed]");
  }
}
