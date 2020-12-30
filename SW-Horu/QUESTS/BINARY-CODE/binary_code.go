// Convert a decimal number to binary number in Golang by Bridge98;
package binary_code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Bin(n string) string { //effettuare calcoli in binario.
	var num int
	fmt.Scan(&num)

	n = strconv.Itoa(num)
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		n = scanner.Text()
	}

	fmt.Println("\n", num, "convertito in binario è uguale a: ", fmt.Sprintf("%b", num))

	return n
}
