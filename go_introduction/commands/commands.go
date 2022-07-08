package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("->")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	//\r\n only on Windows, In unix SO used to only \n
	str = strings.Replace(str, "\r\n", "", -1)
	return str, nil
}

func ShowInConsole(exps []float32) {
	var builder strings.Builder

	for _, exp := range exps {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f", exp))

	}
}
