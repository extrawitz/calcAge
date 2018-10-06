package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//Request input
	fmt.Print("When were you born?: ")
	text, _ := reader.ReadString('\n')
	//Normalize input
	text = strings.Replace(text, "\n", "", -1)
	//Convert input string to int
	inputvalue, _ := strconv.Atoi(text)
	fmt.Println("Your age:", getAge(inputvalue))
}
func getAge(birthyear int) (age int) {
	thisyear := time.Time.Year(time.Now())
	current := (thisyear - birthyear)
	return current
}
