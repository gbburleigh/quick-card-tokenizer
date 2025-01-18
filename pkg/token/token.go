package token

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/gbburleigh/quick-card-tokenizer/internal/util"
)

func Generate(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}
	return hex.EncodeToString(b), nil
}

func Tokenize(cardData util.CardData, db *sql.DB) (string, error) {
	token, err := Generate(16)
	if err != nil {
		return "", err
	}

	stmt, err := db.Prepare("INSERT INTO tokens(token, pan, expiry, cardholder, created_at) VALUES(?,?,?,?,?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token, cardData.PAN, cardData.Expiry, cardData.Cardholder, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}

func Query(token string, db *sql.DB) (util.CardData, error) {
	row := db.QueryRow("SELECT pan, expiry, cardholder FROM tokens WHERE token = ?", token)

	var cardData util.CardData
	err := row.Scan(&cardData.PAN, &cardData.Expiry, &cardData.Cardholder)
	if err == sql.ErrNoRows {
		return util.CardData{}, fmt.Errorf("token not found")
	} else if err != nil {
		return util.CardData{}, err
	}

	maskedCard := util.CardData{
		PAN:        Mask(cardData.PAN),
		Expiry:     cardData.Expiry,
		Cardholder: cardData.Cardholder,
	}
	return maskedCard, nil
}

func Mask(pan string) string {
	if len(pan) <= 4 {
		return pan
	}
	return strings.Repeat("*", len(pan)-4) + pan[len(pan)-4:]
}
