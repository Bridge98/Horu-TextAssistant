// Mapping a conversion from English Alphabet to Morse Alphabet in Golang by Bridge98;
package morse

import (
	"os"
	"bufio"
)

func Morse(question string) string { //alfabeto morse.
	morse := make(map[string]string)
	morse = map[string]string{
		"a":  ".-",
		"b":  "-...",
		"c":  "-.-.",
		"d":  "-..",
		"e":  ".",
		"f":  "..-.",
		"g":  "--.",
		"h":  "....",
		"i":  "..",
		"j":  ".---",
		"k":  "-.-",
		"l":  ".-..",
		"m":  "--",
		"n":  "-.",
		"o":  "---",
		"p":  ".--.",
		"q":  "--.-",
		"r":  ".-.",
		"s":  "...",
		"t":  "-",
		"u":  "..-",
		"v":  "...-",
		"w":  ".--",
		"x":  "-..-",
		"y":  "-.--",
		"z":  "--..",
		"ä":  ".-.-",
		"ö":  "---.",
		"ü":  "..--",
		"Ch": "----",
		"0":  "-----",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		".":  ".-.-.-",
		",":  "--..--",
		"?":  "..--..",
		"!":  "..--.",
		":":  "---...",
		"\"": ".-..-.",
		"'":  ".----.",
		"=":  "-...-",
	}

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		question = scanner.Text()
	}

	var str string

	for _, v := range question {
		str += morse[string(v)]
		str += " "
	}

	return str
}
