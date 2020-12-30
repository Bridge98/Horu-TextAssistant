package quests

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"./binary_code"
	"./GTTS/searchsolutions"
	"./istructions"
	"./LANGUAGES/morse"
	"./player"
	"./texts"
)

const standardq = "Che altro posso fare per te?"

func Quests(question string) string { //domande che il nostro amico può capire  -->	 interazione diretta con utente.

	wt := "\nNon ho capito la richiesta, puoi ripetere?\n"
	count := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		question = scanner.Text()
		keyq := strings.ToLower(question)

		if strings.Contains(keyq, "che") && (strings.Contains(keyq, "oggi") || strings.Contains(keyq, "giorno")) {
			fmt.Println("\nOggi è il ", time.Now()) //giorno in tempo reale;
			fmt.Println(standardq,"\n")
		} else if (strings.Contains(keyq, "ore") || strings.Contains(keyq, "ora")) && strings.Contains(keyq, "che") {
			fmt.Println("\nSono le ", time.Now()) //ora in tempo reale;
			fmt.Println(standardq,"\n")
		} else if strings.Contains(keyq, "stai") || strings.Contains(keyq, "bene") || strings.Contains(keyq, "come") {
			fmt.Println("\nTutto bene grazie e tu?\n")
		} else if strings.Contains(keyq, "fai") || strings.Contains(keyq, "facendo") {
			fmt.Println("\nNiente di che, aspetto che tu mi dica di fare qualcosa!\n")
		} else if strings.Contains(keyq, "leggi") || strings.Contains(keyq, "leggimi") {
			fmt.Println("\nNon so ancora leggere necessito di un aggiornamento.")
			fmt.Println("Intanto eccoti il testo di una bellissima canzone: \n")
			texts.Read() //lettore di testi --> .txt;
			fmt.Println("\n",standardq,"\n")
		} else if strings.Contains(keyq, "musica") || strings.Contains(keyq, "canta") || strings.Contains(keyq, "riproduci") || strings.Contains(keyq, "stop") {
			fmt.Println("\nCerto, ecco un brano per te:")
			fmt.Println("(Per terminare digita: ferma, stop, basta)\n")
			player.Song() //mediaplayer --> mp3;
			fmt.Println("\n",standardq,"\n")
		} else if strings.Contains(keyq, "morse") || strings.Contains(keyq, "trascrivi") {
			fmt.Println("\nCerto, inserisci la parola o la frase da convertire: \n")
			fmt.Println(morse.Morse(question)) //conversione di testo in morse;
			fmt.Println("\n",standardq,"\n")
		} else if strings.Contains(keyq, "binario") || strings.Contains(keyq, "binaria") {
			fmt.Println("\nCerto, inserisci il numero da convertire: \n")
			fmt.Println(binary_code.Bin(question)) //converti qualsiasi n in binario;
			fmt.Println(standardq,"\n")
		} else if strings.Contains(keyq, "ciao") || strings.Contains(keyq, "spegniti") || strings.Contains(keyq, "basta così") {
			fmt.Println("\nCercami quando hai bisogno! Ciao!")
			question = ""
			break
		} else {
		
			count++

			if count >= 3 { //se viene posta la domanda sbagliata per tre volte di fila allora ricordo all'utente cosa è in grado di fare Horu;
				wt = ""
				fmt.Println(wt,"\n")
				istructions.Istructions()
				fmt.Println("\n",standardq,"\n")
			} else {
				searchsolutions.SearchSolutions(question)
				fmt.Println(standardq,"\n")
			}
		}

		keyq = question
		question += "\n"
	}

	return question
}