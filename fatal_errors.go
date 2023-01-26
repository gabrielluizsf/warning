package warning

import "log";

//This function generates a log with the error message if there is an error
func FATAL_ERROR(ERROR error){
  if ERROR != nil{
    log.Fatalln(ERROR);
  }
}
