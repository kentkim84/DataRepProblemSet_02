package main

import (
	"net/http"
	"log"
	"html/template"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

type PageVariables struct {
	HeadTitle string
	BodyTitle string
	Response string
	Message string
}

// generate a random number
func genRandInt(minNum int, maxNum int) int {
    return rand.Intn(maxNum) + minNum
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// set local variables
	headTitle := "Guessing Game"
	bodyTitle := "Guessing Game"

	// initialise local variables to struct
	MyPageVariables := PageVariables {
		HeadTitle : headTitle,
		BodyTitle : bodyTitle,
	}

	// parse the template
	t, err := template.ParseFiles("index.html")
	if err != nil { // if there is an error
  		log.Print("template parsing error: ", err) // log error message
	}
	t.Execute(w, MyPageVariables)
	// fmt.Fprintf(w, "Execute indexHandler")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// set local variables
	headTitle := "Guessing Game"
	message := "Guess a number between 1 and 20"

	// set random number
	maxNum := 20
	rand.Seed(time.Now().UnixNano())
	randNum := genRandInt(1, maxNum)

	// set up cookie
	cookie := http.Cookie{Name: "target", Value: strconv.Itoa(randNum)}
	_, err := r.Cookie("target")
	if err == nil { // if there is no error
		http.SetCookie(w, &cookie)
	}

	// initialise local variables to struct
	MyPageVariables := PageVariables {
		HeadTitle : headTitle,
		Message : message,
	}

	// parse the template
	t, err := template.ParseFiles("guess.tmpl")
	if err != nil { // if there is an error
  		log.Print("template parsing error: ", err) // log error message
	}
	t.Execute(w, MyPageVariables)
	// fmt.Fprintf(w, "Execute guessHandler")

	// read the value from the input box in the form
	guessNum := r.FormValue("guess")
	iGuessNum, err := strconv.Atoi(guessNum)
	

	// initial try number count
	tryNum := 1

	// store game random number
	var gameRandNum int

	if tryNum == 1 {
		gameRandNum = randNum
	}
	fmt.Fprintf(w, "User guess: %d<br>", iGuessNum)
	fmt.Fprintf(w, "Random number: %d<br>", gameRandNum)

	/* // create slices(=dynamic array)
	numContainer := make([]int, 0)

	// store the guessing number in the number container
	numContainer = append(numContainer, iGuessNum) */

	if (iGuessNum != gameRandNum) {
		if (iGuessNum < gameRandNum) {
			fmt.Fprintf(w, "Your number is lower than the random number<br>")
		} else if (iGuessNum > gameRandNum) {
			fmt.Fprintf(w, "Your number is higher than the random number<br>")
		}
		
		fmt.Fprintf(w, "Number of tries: %d<br>", tryNum)
		fmt.Fprintf(w, "Enter a guessing number<br>")

		/* find := true
		for find {
			for i := 0; i < len(numContainer); i++  {
				if (iGuessNum == numContainer[i]) {
					fmt.Fprintf(w, "Enter a different guessing number")
					//fmt.Scanln(&guessNum)
				} else {
					find = false
				}
				fmt.Fprintf(w, "A number in the container %d", numContainer[i]) // debuging purpose
			}
		}
			
		numContainer = append(numContainer, iGuessNum) */
	}

	if (iGuessNum == gameRandNum) {
		fmt.Fprintf(w, "You have tried: %d<br>", tryNum)
		fmt.Fprintf(w, "Numbers matched<br>")
	}
	tryNum += 1
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/guess/", guessHandler)
    http.ListenAndServe(":8080", nil)
}