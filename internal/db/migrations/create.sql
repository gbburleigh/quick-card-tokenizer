CREATE TABLE IF NOT EXISTS tokens (
    token TEXT PRIMARY KEY,
    pan TEXT,
    expiry TEXT,
    cardholder TEXT,
    created_at DATETIME
)