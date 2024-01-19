package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentAttempt := 0
	wordToGuess := getRandomWord()
	placeholder := getLettersPlaceholders(wordToGuess)

	for true {
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

		// Add to correctLetters
		if strings.Contains(wordToGuess, userGuess) {
			guessedLetters += userGuess

			// Updating placeholder
			letterIndexes := getLetterIndexes(wordToGuess, userGuess)
			placeholder = replacePlaceholder(placeholder, userGuess, letterIndexes)
		} else {
			currentAttempt++
		}

		// Check if player lost
		if currentAttempt == len(hangmanArr) {
			pl("You lost.", "The word was", wordToGuess)
			break
		}

		// Check if player won
		if !strings.ContainsAny(placeholder, "_") {
			clearShell()
			pl("You won! the word was", wordToGuess)
			break
		}
	}
}

// Display current game board.
func displayGameBoard(attempt int) {
	if attempt < len(hangmanArr) {
		pl(hangmanArr[attempt])
	} else {
		pl(hangmanArr[len(hangmanArr)-1])
	}
}

// Get a random word from words array.
func getRandomWord() string {
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
