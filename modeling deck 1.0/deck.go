package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//making custome type
type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	stringifiedDeck := d.toString()

	//0666 perm means anyone can read and write this file
	return os.WriteFile(fileName, []byte(stringifiedDeck), 0666)
}

func (d deck) toString() string {
	slicedStringDeck := []string(d)
	stringifiedDeck := strings.Join(slicedStringDeck, ", ")
	return stringifiedDeck
}

func newDeckFromFile(fileName string) deck {
	byteSlicedData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		//exit the program with failure flag
		os.Exit(1)
	}

	stringifiedData := string(byteSlicedData)

	//spliting a string to chunks and convert it to string slice
	stringSlicedData := strings.Split(stringifiedData, ", ")

	//convert string slice to custome type deck
	deckFromFile := deck(stringSlicedData)

	return deckFromFile
}

// inplace method on deck instance
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		randomIndex := r.Intn(len(d) - 1)

		//swapping two card in slice
		d[index], d[randomIndex] = d[randomIndex], d[index]
	}
}
