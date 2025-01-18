package main_test

import (
	"testing"

	"github.com/gbburleigh/quick-card-tokenizer/internal/db"
	"github.com/gbburleigh/quick-card-tokenizer/internal/util"
	"github.com/gbburleigh/quick-card-tokenizer/pkg/token"
)

func TestTokenizeAndRetrieve(t *testing.T) {
	database := db.Create()
	defer database.Close()

	card := util.CardData{PAN: "1234567890123456", Expiry: "12/24", Cardholder: "Test User"}
	newToken, err := token.Tokenize(card, database)

	if err != nil {
		t.Error("Something went wrong while tokenizing")
	}

	if newToken == "" {
		t.Error("Tokenization failed, token is empty")
	}

	retrievedCard, err := token.Query(newToken, database)

	if retrievedCard.PAN != "************3456" {
		t.Errorf("Retrieved PAN is incorrect: got %s, want %s", retrievedCard.PAN, "************3456")
	}
	if retrievedCard.Expiry != card.Expiry {
		t.Errorf("Retrieved Expiry is incorrect: got %s, want %s", retrievedCard.Expiry, card.Expiry)
	}
	if retrievedCard.Cardholder != card.Cardholder {
		t.Errorf("Retrieved Cardholder is incorrect: got %s, want %s", retrievedCard.Cardholder, card.Cardholder)
	}
}
