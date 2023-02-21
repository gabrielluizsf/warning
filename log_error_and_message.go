package warning

import "log";

func LOG_ERROR_AND_MESSAGE(err error, message string){
  if err != nil{
    log.Fatal(err,message);
  }
}
