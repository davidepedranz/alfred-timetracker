package alfred

import (
	"fmt"
	"log"
	"os"
)

func PrintError(message string, err error) {
	log.Printf("%s: %v", message, err)
	fmt.Println(message)
	os.Exit(1)
}
