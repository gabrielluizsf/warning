# warning - The Error Library :bug:


## Default Validation

```go
 out, err := exec.Command("disfhdisfias").Output();
    if err != nil{
      fmt.Println("INVALID COMMAND");
    }else{
      fmt.Println("Command successfully executed");
    }
	output := string(out[:]);
	fmt.Println(output);
```

benchmark:

<img src="https://github.com/theGOURL/warning/blob/main/images/BASIC_ERROR.png?raw=true" />

## Using PRINT_DEFAULT_ERRORS function

```go
 out, err := exec.Command("disfhdisfias").Output();
    warning.PRINT_DEFAULT_ERRORS(err,"INVALID COMMAND");
	output := string(out[:]);
	fmt.Println(output);
```
benchmark:
<img src="https://github.com/theGOURL/warning/blob/main/images/PRINT_DEFAULT_ERRORS.png?raw=true" />
