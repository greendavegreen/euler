package main

import (
	"log"
	"os"
	"github.com/greendavegreen/euler/p054/hand"
)

func main() {
	file, err := os.Open("p054_poker.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	h := hand.New([]string{"3C", "3D", "3H", "3S", "2H"})
	h.Display()

	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	cards := strings.Fields(line)
	//
	//	h1 := hand.New(cards[0:5])
	//	h1.toSring()
	//	h2 := hand.New(cards[5:])
	//	h2.toSring()
	//	fmt.Println()
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}
}
