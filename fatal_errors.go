package warning

import "log";

func FATAL_ERROR(ERROR error){
  if ERROR != nil{
    log.Fatalln(ERROR);
  }
}
