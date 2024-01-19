package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var pl = fmt.Println

var hangmanArr = [8]string{
	" +-------+\n" +
		" |       |\n" +
		"         |\n" +
		"         |\n" +
		"         |\n" +
		"         |\n" +
		"       =====\n",

	" +--------+\n" +
		" |        |\n" +
		" 0        |\n" +
		"          |\n" +
		"          |\n" +
		"          |\n" +
		"        =====\n",

	" +--------+\n" +
		" |        |\n" +
		" 0        |\n" +
		" |        |\n" +
		"          |\n" +
		"          |\n" +
		"        =====\n",

	" +--------+\n" +
		" |        |\n" +
		" 0        |\n" +
		"/|        |\n" +
		"          |\n" +
		"          |\n" +
		"        =====\n",

	" +--------+\n" +
		" |        |\n" +
		" 0        |\n" +
		"/|\\       |\n" +
		"          |\n" +
		"          |\n" +
		"        =====\n",

	" +--------+\n" +
		" |        |\n" +
		" 0        |\n" +
		"/|\\       |\n" +
		" |        |\n" +
		"          |\n" +
		"        =====\n",
	" +--------+\n" +
		" |        |\n" +
		" 0        |\n" +
		"/|\\       |\n" +
		" |        |\n" +
		"/         |\n" +
		"        =====\n",
	" +--------+\n" +
		" |        |\n" +
		" X        |\n" +
		"/|\\       |\n" +
		" |        |\n" +
		"/ \\       |\n" +
		"        =====\n",
}

var words = []string{
	"apple", "banana", "cherry", "dog", "elephant",
	"fish", "grape", "happiness", "cream", "jazz",
	"kangaroo", "lemon", "mountain", "notebook", "orange",
	"penguin", "quasar", "rabbit", "sunshine", "tiger",
	"umbrella", "violet", "waterfall", "xylophone", "yellow",
	"zebra", "astronomy", "butterfly", "computer", "diamond",
	"monkey", "fluff",
}

var randomWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentAttempt := 0
	wordToGuess := getRandomWord()
	placeholder := getLettersPlaceholders(wordToGuess)
	gamePlaying := true

	for gamePlaying {
		// Clear terminal window
		clearShell()

		// Show game board
		displayGameBoard(currentAttempt)

		// Show placeholder
		pl("Word:", placeholder)

		// Get a letter from the user
		// fmt.Printf("\n [DEBUG] word: %s", wordToGuess)
		fmt.Print("\nGuess a letter: ")

		userGuess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		// Removing delimiter from user input
		userGuess = strings.TrimSuffix(userGuess, "\n")

		// A. If they guessed letter in word
		// Add to correctLetters
		if strings.Contains(wordToGuess, userGuess) {
			guessedLetters += userGuess

			// Updating placeholder
			letterIndexes := getLetterIndexes(wordToGuess, userGuess)
			placeholder = replacePlaceholder(placeholder, userGuess, letterIndexes)
		} else {
			wrongGuesses = append(wrongGuesses, userGuess)
			currentAttempt++
		}
	}
}

// Display current game board.
func displayGameBoard(attempt int) {
	if attempt >= len(hangmanArr) {
		pl("YOU DIED")
	} else {
		pl(hangmanArr[attempt])
	}
}

// Get a random word from words array.
func getRandomWord() string {
	milliseconds := time.Now().UnixMilli()
	rand.Seed(milliseconds)
	wordIndex := rand.Intn(len(words))

	return words[wordIndex]
}

// Get letters placeholder from chosen random word.
func getLettersPlaceholders(word string) string {
	splitWord := strings.Split(word, "")
	var placeholder string

	for i := 0; i < len(splitWord); i++ {
		placeholder += "_ "
	}

	return placeholder
}

// Clears STDIO.
func clearShell() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Get all letter's indexes in a word
func getLetterIndexes(wordToGuess string, letter string) []int {
	var indexes []int
	letterIndex := 9999

	for letterIndex != -1 {
		letterIndex = strings.IndexAny(wordToGuess, letter)
		wordToGuess = strings.Replace(wordToGuess, letter, "_", 1)

		if letterIndex != -1 {
			indexes = append(indexes, letterIndex)
		}
	}

	return indexes
}

// Replace placeholder with guessed letters
func replacePlaceholder(placeholder string, letter string, indexes []int) string {
	newPlaceholder := strings.Split(placeholder, " ")
	for _, val := range indexes {
		newPlaceholder[val] = letter
	}

	return strings.Join(newPlaceholder, " ")
}
