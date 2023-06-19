package warning

import (
	"fmt"
	"os/exec"
	"os/user"
	"strings"
	"testing"
)

func TestPRINT_DEFAULT_ERRORS(test *testing.T) {
	SO_USER, _ :=  user.Current()
	SO_USERNAME := SO_USER.Username
	out, err := exec.Command("whoami").Output()
	PRINT_DEFAULT_ERRORS(err, "INVALID COMMAND")
	output := strings.TrimSpace(string(out))

	if !strings.Contains(output, SO_USERNAME) {
		test.Errorf("Nome de usuário incorreto. Esperado: %s, Obtido: %s", SO_USERNAME, output)
	}
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
      fmt.Println("Command successfully executed 🐛");
    }
	output := string(out[:]);
	fmt.Println(output);
  }
}
