package wordle

import "testing"

func TestNewWordleState(t *testing.T) {
	word := "HELLO"
	ws := newWordleState(word)
	wordleAsString := string(ws.word[:])

	t.Log("Created wordleState:")
	t.Logf("    word: %s", wordleAsString)
	t.Logf("    guesses: %v", ws.guesses)
	t.Logf("    currGuess: %v", ws.currGuess)

	if wordleAsString != word {
		t.Errorf("Expected word to be %s, got %s", wordleAsString, word)
	}
}

func statusToString(status letterStatus) string {
	switch status {
	case none:
		return "none"
	case correct:
		return "correct"
	case present:
		return "present"
	case absent:
		return "absent"
	default:
		return "unknown"
	}
}


// // string converts a guess into a string
// func (g *guess) string() string {
//     str := ""
//     for _, l := range g {
//         if 'A' <= l.char && l.char <= 'Z' {
//             str += string(l.char)
//         }
//     }
//     return str
// }

func TestNewGuess(t *testing.T) {
	wordToGuess := "YIELD"
	guess := newGuess(wordToGuess)

	t.Logf("New guess: %s", guess.string())

	// Check that the letter and status are correct
	for i, l := range guess {
		t.Logf("    letter %d: %c, %s", i, l.char, statusToString(l.status))

		if l.char != wordToGuess[i] || l.status != none {
			t.Errorf(
				"letter[%d] = %c, %s; want %c, none",
				i,
				l.char,
				statusToString(l.status),
				wordToGuess[i],
			)
		}
	}
}


func TestUpdateLettersWithWord(t *testing.T) {
	guessWord := "YIELD"
	guess := newGuess(guessWord)

	var word [wordSize]byte
	copy(word[:], "HELLO")
	guess.updateLettersWithWord(word)

	statuses := []letterStatus{
		absent,  // "Y" is not in "HELLO"
		absent,  // "I" is not in "HELLO"
		present, // "E" is in "HELLO" but not in the correct position
		correct, // "L" is in "HELLO" and in the correct position
		absent,  // "D" is not in "HELLO"
	}

	// Check that statuses are correct //letterSatus
	for i := range guess {
		if guess[i].status == statuses[i]{
			t.Logf("Success. Expected %d, Received %d", guess[i].status, statuses[i])
		} else {
			t.Errorf("Expected %d, but got %d", guess[i].status, statuses[i])
		}
	}
}


func TestIsWordGuessed(t *testing.T) {
	ws := newWordleState("HELLO")
	g := newGuess("HELLO")

	g.updateLettersWithWord(ws.word)
	ws.appendGuess(g)

	if !ws.isWordGuessed() {
		t.Errorf("isWordGuessed() should return true")
	}
}

func TestShouldEndGame(t *testing.T) {
	ws := newWordleState("HELLO")
	g := newGuess("HELLO")

	g.updateLettersWithWord(ws.word)
	ws.appendGuess(g)

	if ws.shouldEndGame() {
		t.Log("game over")
	}

	if !ws.shouldEndGame() { 
		t.Error("isWordGuessed() should return true")
	}
}