package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
    "math/rand"
    "time"
    "strconv"
)

func init(){
    rand.Seed(time.Now().UTC().UnixNano())
}

func main(){
    //Steps:
    // 1: Pick a number between 1 and 10
    answer := strconv.Itoa(rand.Intn(10) + 1)
    fmt.Printf("Answer: %s\n", answer)

    // 2: Ask the user to choose a number between 1 and 10
    fmt.Println("I have picked a number between 1 and 10. Can you guess it?\n\n")
    fmt.Println("Please Choose a number between 1 and 10")

    input := bufio.NewScanner(os.Stdin)
    userAnswer := ""
    for input.Scan(){
        if input.Text() != "" {
            userAnswer = strings.TrimSpace(input.Text())
            // 3: Check the user number against the answer.

            // if correct, end game. If not, try again.
            if userAnswer == answer {
                fmt.Println("You guessed correctly! Congrats!!!")
                break
            }else{
                fmt.Printf("%s is not correct.\n", userAnswer)
            }
        }
    }
}
