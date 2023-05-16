package wordle

import (
	"errors"
	//"fmt"
	"github.com/j-zarr/wordle/words"
)

const (
	maxGuesses = 6
	wordSize   = 5
)

type wordleState struct {
	// word is the word that the user is trying to guess
	word [wordSize]byte
	// guesses holds the guesses that the user has made
	guesses [maxGuesses]guess
	// currGuess is the index of the available slot in guesses
	currGuess int
}

// guess is an attempt to guess the word
type guess [wordSize]letter

type letter struct {
	// char is the letter that this struct represents
	char byte
	// status is the state of the letter (absent, present, correct)
	status letterStatus
}

// letterStatus can be none, correct, present, or absent
type letterStatus int

const (
	// none = no status, not guessed yet
	none letterStatus = iota
	// absent = not in the word
	absent
	// present = in the word, but not in the correct position
	present
	// correct = in the correct position
	correct
)

// newWordleState builds a new wordleState from a string.
// Pass in the word you want the user to guess.
func newWordleState(word string) wordleState {
	w := wordleState{}
	wordChars := []byte{}

	for i := range word {
		wordChars = append(wordChars, word[i])
	}

	w.word = [5]byte(wordChars)

	return w
}

func newLetter(char byte) letter {
	return letter{char: char, status: none}
}

func newGuess(str string) guess {

	var g guess

	for i, c := range str {
		g[i] = newLetter(byte(c)) //letter = {0,0}  guess = [{0,0}, {0,0}, {0,0}, {0,0}]
	}
	return g
}

// updateLettersWithWord updates the status of the letters in the guess based on a word
func (g *guess) updateLettersWithWord(word [wordSize]byte) {
	// g array of letter , letter {char byte, status letterStatus}
	for i := range g {
		letter := &g[i]
		if letter.char == word[i] {
			letter.status = correct
		} else {
			for _, b := range word {
				if letter.char == b {
					letter.status = present
				}
			}

		}
		if letter.status == 0 {
			letter.status = absent
		}
	}

}

func (g *guess) string() string {
	str := ""
	for _, l := range g {
		if 'A' <= l.char && l.char <= 'Z' {
			str += string(l.char)
		}
	}
	return str
}

// appendGuess adds a guess to the wordleState. It returns an error
// if the guess is invalid.
func (w *wordleState) appendGuess(g guess) error {

	if w.currGuess >= maxGuesses {
		return errors.New("maximum number of guesses reached")

	} else if len(g.string()) != wordSize {
		return errors.New("invalid guess - guess must of length 5")
		
	} else {
		// for _, l := range g {
			// if l.char < 'A' || l.char > 'Z' {
			if !words.IsWord(g.string()) {
				return errors.New("invalid guess - invalid word")
			}
			// }
		//}
	}


	w.guesses[w.currGuess] = g
	w.currGuess ++
	return nil
}


// isWordGuessed returns true when the latest guess is the correct word
func (w *wordleState) isWordGuessed() bool {

	var b [5]byte

	for i, l := range w.guesses[w.currGuess] {
		b[i] = l.char
	}

	return  b == w.word
}

func (w *wordleState) shouldEndGame() bool {

	return w.isWordGuessed() || w.currGuess >= 6
}

