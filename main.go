package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Digite dois Números e a Operação que deseja fazer? EX: 2+2")

	reader := bufio.NewReader(os.Stdin)
	nu, _ := reader.ReadString('\n')

	nu = strings.TrimSpace(nu)
	nu = strings.ReplaceAll(nu, " ", "")

	acheoOperador := strings.IndexAny(nu, "+*/-")
	if acheoOperador == -1 {
		fmt.Println(("Operador Invalido. Tente um desses Operadores: +, -, /, *"))
		return
	}

	n1 := nu[:acheoOperador]
	operador := string(nu[acheoOperador])
	n2 := nu[acheoOperador+1:]

	result := calculos(n1, operador, n2)
	fmt.Printf("O Resultado de %s %s %s é %d\n", n1, operador, n2, result)

	fmt.Println("Pressione Enter para sair...")
	bufio.NewReader(os.Stdin).ReadString('\n')

}

func calculos(n1 string, operador string, n2 string) int {
	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)

	switch operador {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		if num1 == 0 || num2 == 0 {
			fmt.Println("Alerta: Multiplicação por 0")
			return 0
		}
		return num1 * num2

	case "/":
		if num1 == 0 || num2 == 0 {
			fmt.Println("Alerta: Divisão por 0")
			return 0
		}
		return num1 / num2
	default:
		fmt.Println("O Operador Inserido está invalido, tente com um desses operadores: +, -, /, *")
		return 0
	}

}
