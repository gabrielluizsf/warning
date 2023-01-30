package main

import (
  "os"
  "github.com/theGOURL/warning"
)

func main(){
  Testing_the_PANIC_ERROR_function();
}

func Testing_the_LOG_ERRORS_function(){  
  filename := "file.adah";
  open, err := os.Open(filename);
    warning.LOG_ERRORS(err, filename + " was not exists");
  defer open.Close();
}
func Testing_the_PANIC_ERROR_function(){  
  filename := "file.adah";
  open, err := os.Open(filename);
    warning.PANIC_ERROR(err);
  defer open.Close();
}