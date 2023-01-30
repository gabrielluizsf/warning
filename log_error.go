package warning

import (
  "log"
  "fmt"
)
//This function prints the error logs in the console if there is an errors
func LOG_ERRORS(ERROR error, MESSAGE string){
  if ERROR != nil{
    fmt.Println(MESSAGE,ERROR);
  }
}
//This function prints the error on the console, exits the program if there is an error
func PANIC_ERROR(ERROR error){
  if ERROR != nil{
    log.Panicln(ERROR);
  }
}
