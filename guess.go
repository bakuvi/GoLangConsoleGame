package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("guess random number")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("you have", 10-guesses, "guesses")
		fmt.Println("make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("your number is too low")
		} else if guess>target {
			fmt.Println("ur number is too high")
		} else {
			success = true
			fmt.Println("u have guessed!")
			break
		}

	}
	if !success{
		fmt.Println("yoa have spent all your tries, it was", target)
	}
}
