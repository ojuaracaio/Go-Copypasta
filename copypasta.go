package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//########## limpar terminal ######################
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	CallClear()
	fmt.Print("A quem gostaria de dedicar a Copypasta?\n")
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	input = input[:len(input)-2]
	CallClear()
	fmt.Printf("Qual Copypasta deseja usar?\n\n1- Para o cego é a luz\n2- Hoje eu torço\n")
	var opcao int
	fmt.Scanf("%d", &opcao)
	CallClear()

	switch opcao {
	case 1:
		paraOCego(input)
	case 2:
		hojeEuTorco(input)
	default:
		fmt.Printf("Opção inválida.")
	}
}

func paraOCego(input string) {
	fmt.Printf("Quem é %s?\n\nPara o cego, é a luz\nPara o faminto, é o pão \nPara o morto, é a vida\nPara o enfermo, é a cura\nPara o prisioneiro, é a liberdade\nPara o solitário, é o companheiro\nPara o viajante, é o caminho\nPara mim, %s é tudo\n", input, input)
}

func hojeEuTorco(input string) {
	input = strings.ToUpper(input)
	fmt.Printf("HOJE EU TORÇO PRA %s\n\nTORÇO PRA %s SE FUDER\n\nVAI TOMAR NO CU %s\n\n................./¯/)\n.............../¯ ./\n............./. . /\n......../´¯/' . '/´¯`•¸,\n..../' /. ./ . ./ . ./¯\\\n..( . ( . ( . ( ¯ ./' . ')\n...\\ . . . . . . . . . . ./\n", input, input, input)
}
