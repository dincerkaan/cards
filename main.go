package main

func main() {

	//var card string = "Ace of Spades"
	// card := "Ace of Spades"   //yukardaki kod satırıyla bire bir aynı işlemi yapıyor.
	// card = "Five of Diamonds" // := ifadesi yalnızca bir değişkeni ilk kez yarattığımızda kullanırız. Sonrasında yalnızca = kullanarak değeri değiştiririz.
	//card := newCard()
	// var card string
	// card = newCard() //aynı işlemleri yapan. card değişkenini tanımlama yolları. Fonksiyondan değer alıyoruz.
	// card := newCard()
	// fmt.Print(card)

	// cards := deck{"Ace of Diamond", newCard()} //slice oluşturduk. İlk değeri kendimiz atadık, 2. değeri oluşturudğumuz newCard fonkisoynundan aldık.
	// cards = append(cards, "Six of Spades")     //sliceımızın sonuna yeni bir değer ekledik. Slice arrayden farklı olarak sabit aralıklı değildir.

	// for i, card := range cards { //range fonksiyonuyla cards içinde tuttuğumuz değerlerin hepsine tek tek ulaşıp cardın içine attık sonrasonda ekrana bastırdık.
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 8)

	// fmt.Println("Cards in hand:")
	// hand.print()
	// fmt.Println("Remainings cards")
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("kayitliKartlar")

	// cards := newDeckFromFile("kayitliKartlar")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
