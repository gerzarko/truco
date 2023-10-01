package paquete1


type Card struct {
    breed uint8
    number uint8

}


type Deck struct {
   cards []Card

   
}

//this method adds a card to a valid deck

func (oneCard Deck)addCard(cardToAdd Card, DeckToAdd Deck){

    DeckToAdd.cards = append(DeckToAdd.cards, cardToAdd)

}

//creates a valid deck

func (newDeck Deck) createNewDeck() Deck{

    newDeck = Deck{
    cards: make([]Card,0),
    } 
    return newDeck
}



    

type player struct {
    name string
    age uint8
    handDeck []Deck
    
    }



