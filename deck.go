package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

//create type which is a slice of string
type deck []string

func newDeck() deck { // function name is newDeck which returns a type of slice deck
	cards := deck{} //initialize empty slice of string

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) { //a function with multiple return values
	return d[:handSize], d[handSize:]

}

//create a receiver function name of print
// think of receiver function as methods attached to the type
func (d deck) print() {
	for i, card := range d { //what is this d?
		fmt.Println(i, card)
	}
}

func (d deck) toString() string { // converting slice of string to string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log the error & quit the program
		log.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
