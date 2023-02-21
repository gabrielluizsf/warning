package warning

import "log";

// This function shows errors in the form of logs along with the message passed by parameter
func LOG_ERROR_AND_MESSAGE(err error, message string){
  if err != nil{
    log.Fatal(err,message);
  }
}
