package util

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// LIFE define if the element is alive
const LIFE uint8 = '*'

// DECEASED define if the element is died
const DECEASED uint8 = '.'

// GetDivMod get the mod between two values
func GetDivMod(x int, y int) int {
	return x - getFloorDiv(x, y)*y
}

func GetFloorMod(a int, b int) int {
	m := a % b
	if (a^b) >= 0 || m == 0 {
		return m
	} else {
		return b + m
	}
}

// PrintGrid is used for print the grid
func PrintGrid(matrix [][]uint8) string {
	var cadena strings.Builder
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cadena.WriteString(string(matrix[i][j]))
		}
		cadena.WriteString("\n")
	}
	cadena.WriteString("\n")
	return cadena.String()
}

// CopyGrid is used for copy the grid
func CopyGrid(array [][]uint8) [][]uint8 {
	duplicate := make([][]uint8, len(array))
	for i := range array {
		duplicate[i] = make([]uint8, len(array[i]))
		copy(duplicate[i], array[i])
	}

	return duplicate
}

/*
// GenerateRandGrid is used for generate randomic grid od string
func GenerateRandGrid(row int, col int) [][]string {
	fmt.Println("hola")
	array := make([][]string, row)
	for i := range array {

			array[i] = newRandomStringArray(col)
	}

	return array
}
*/
// GetValueFromConsole is used for get value string from console
func GetValueFromConsole(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	return text
}

// ClearConsole is used for clear de output from console
func ClearConsole() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// Util functions
func (b *boolgen) boolElement() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func newRandom() *boolgen {
	return &boolgen{src: rand.NewSource(time.Now().UnixNano())}
}

func newRandomBoolArray(col int) []bool {
	row := make([]bool, col)
	random := newRandom()
	for index := 0; index < col; index++ {
		row[index] = random.boolElement()
	}
	return row
}

/*
func newRandomStringArray(col int) []string {
	row := make([]string, col)
	random := newRandom()
	for index := 0; index < col; index++ {
		if random.boolElement() {
			row[index] = LIFE
		} else {
			row[index] = DECEASED
		}

	}
	return row
}
*/
type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

// The Euclidean Algorithm
func getFloorDiv(x int, y int) int {
	r := x / y
	if (x^y) < 0 && (r*y != x) {
		r--
	}
	return r
}

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

// TimeTrack get duration of execution of function
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s started %s", name, start)
	log.Printf("%s took %s", name, elapsed)
}
