// A File Reader (I/O) for TXT files, made in Golang by Bridge98;
package texts

import (
	"bufio"
	"fmt"
	"os"
)

func Read() {
	g, _ := os.Open("Giorgio_Mastrota.txt")
	defer g.Close()

	scanner := bufio.NewScanner(g)

	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
		str += "\n"
	}
}
