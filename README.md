# Secure Card Tokenization in Go

Robust and secure card tokenization implemented in Go, protecting sensitive payment data.

## Key Features

*   **Secure Hashing with Salt:** Card data is hashed using SHA256 with a randomly generated salt for strong one-way encryption.
*   **Token Expiration:** Tokens can be configured to expire after a specified duration, enhancing security and limiting exposure.
*   **Database Persistence:** Tokens and associated metadata (masked PAN, expiry, cardholder) are stored in a SQLite database.
*   **Migration Support:** Uses SQL migration files to manage database schema changes.

## Getting Started

1.  **Prerequisites:**
    *   Go 1.18 or later (https://go.dev/doc/install](https://go.dev/doc/install))

2.  **Clone the Repository:**

    ```bash
    git clone [https://github.com/yourusername/tokenization-project.git](https://github.com/yourusername/tokenization-project.git) # Replace with your repo URL
    cd tokenization-project
    ```

3.  **Run the Project:**

    ```bash
    go run main.go
    ```

    This will initialize the database (creating `tokens.db` if it doesn't exist) and run the example tokenization process.

## Usage

The primary functionality is exposed through the `tokenizer` package. Here's a basic example:

```go
  package main
  
  import (
      "fmt"
      "log"
      "time"
  
      "[github.com/yourusername/tokenization-project/db [invalid URL removed]"
      "github.com/yourusername/tokenization-project/tokenizer]([invalid URL removed])"
      "[github.com/yourusername/tokenization-project/util [invalid URL removed]"
  )
  
  func main() {
      database := db.InitDB()
      defer database.Close()
  
      card := util.CardData{PAN: "1234567890123456", Expiry: "12/24", Cardholder: "John Doe"}
      token, err := tokenizer.TokenizeCard(card, database, time.Hour*24) // Token expires in 24 hours
      if err != nil {
          log.Fatal(err)
      }
      fmt.Println("Generated Token:", token)
  
      retrievedCard, err := tokenizer.RetrieveCardData(token, database)
      if err != nil {
          log.Fatal(err)
      }
      fmt.Println("Retrieved Card (Masked):", retrievedCard)
  }
```
