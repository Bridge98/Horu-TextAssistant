// A GTTS made in Golang by Bridge98;
package searchsolutions

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"fmt"
	"log"
	"runtime"
)

func ExecuteWin(search string) string {
	var gts bytes.Buffer //google to search --> with buffer;
	var s []string
	url := "https://www.google.com/search?q="
	
	fmt.Println("Posso cercare in internet il problema se vuoi:")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		search = scanner.Text()
		search += "\n"
	}

	s = append(s, search)

	for i := range s {
		url += s[i]
		if i != len(s)-1 {
			url += "+"
		}
	}

	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", url) //rundll32 --> windows

	err := cmd.Run() //run the cmd command
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gts.String())

	return url
}


func ExecuteLin(search string) string {
	var gts bytes.Buffer //google to search --> with buffer;
	var s []string
	url := "https://www.google.com/search?q="
	
	fmt.Println("Posso cercare in internet il problema se vuoi:")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		search = scanner.Text()
		search += "\n"
	}

	s = append(s, search)

	for i := range s {
		url += s[i]
		if i != len(s)-1 {
			url += "+"
		}
	}

	cmd := exec.Command("xdg-open", url) //xdg-open --> linux

	err := cmd.Run() //run the cmd command
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gts.String())

	return url
}

func SearchSolutions(search string) string {

	var err error

	switch runtime.GOOS {

		case "windows":
			ExecuteWin(search)
		case "linux":
			ExecuteLin(search)
		default:
			err = fmt.Errorf("Piattaforma o sistema non supportati, consultare le istruzioni di Horu.")
	}

	if err != nil {
		log.Fatal(err)
	}

	return search
}
