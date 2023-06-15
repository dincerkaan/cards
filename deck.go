package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suites := range cardSuites {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suites)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] //ilk kısım elede kalan kartlar. Başlangıçtan ldeki kart sayısına kadar, ikinci değer kalan kartları temsil eder.
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileuserName string) error {
	return ioutil.WriteFile(fileuserName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileuserName string) deck {
	bs, err := ioutil.ReadFile(fileuserName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
