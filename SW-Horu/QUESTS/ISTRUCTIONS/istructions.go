// open the TXT file to show how Horu works;
package istructions

import (
	"bufio"
	"fmt"
	"os"
)

func Istructions() { //istruzioni di base --> in costante aggiornamento (file txt).
	g, _ := os.Open("Guide_Horu.txt")
	defer g.Close()

	scanner := bufio.NewScanner(g)

	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
		str += "\n"
	}
}
