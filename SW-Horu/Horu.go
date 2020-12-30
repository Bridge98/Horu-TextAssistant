package main

import (
	"bufio"
	"fmt"
	"os"

	"./quests"
)

/*

Aggiungere:
- posizione GPS in tempo reale;
- database/func/link a religione di Horu --> Pastafarianesimo!
- migliorare data e ora --> ricerca su data (es. "che giorno era ieri" ecc.) e miglioramento ora per fuso orario tramite ricerca nel web (link) (es. "che ora è in Nepal?" ecc.);
- ricerca web per qualsiasi altra domanda;
- fare altre mappe per parlare in altre lingue --> settare Horu in un'altra lingua come si fa??

Aggiunti:
- linguaggio morse --> linguaggi
- riproduttore di file mp3 --> musica
- istruzioni di Horu --> istruzioni
- conversione binaria --> da n a bin(n)

*/

func main() {
	fmt.Println("\nCiao io sono Horu, il tuo assistente personale.")
	fmt.Println("Come ti chiami?\n")
	var name string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Println("\nCiao", name, "cosa posso fare per te?\n")
	var question string

	fmt.Println(quests.Quests(question))
	fmt.Println()
}
