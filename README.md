🔗 Go URL Shortener (SHA-256)

A simple and efficient URL shortener built with Go that uses SHA-256 hashing to generate unique short URLs. No database required—just fast, deterministic URL mapping using cryptographic hashing.
🚀 Features

    ✅ Built in Go

    🔐 Uses SHA-256 hashing for generating short links

    ⚡ Fast and stateless (ideal for small-scale or demo purposes)

    🌐 Simple REST API using net/http or chi

    📦 Easy to build and run

🛠️ How It Works

    Accepts a long URL via POST request.

    Hashes it using SHA-256.

    Takes the first 8 characters of the hex-encoded hash.

    Returns the shortened URL.

    Redirects incoming requests to the original long URL (stored in memory or optionally persisted).

🧪 Example
Shorten a URL

POST /shorten
Content-Type: application/json

  "url": "https://example.com/very/long/path"

Redirect

GET /9f86d08a(or any short url you get from the output)

➡️ Redirects to: https://example.com/very/long/path


🧰 Tech Stack

    Language: Go

    Framework: chi & CORS

    Hashing: crypto/sha256

🧑‍💻 Running the Project

# Clone the repo
git clone https://github.com/guruorgoru/Url-Shortener
cd url-shortener

# Run the server
make run

Server runs on: http://localhost:8414
🧠 Notes

    SHA-256 is deterministic, so the same input always results in the same short URL.

    Since only the first few characters of the hash are used, there is a small chance of collisions.

    Ideal for personal or educational projects. Production-ready systems should implement collision checks and persistence (e.g., Redis or a DB).

📜 License

MIT License. Feel free to use, modify, and distribute.
✨ Contributions

Pull requests and suggestions are welcome! Fork the repo and make your magic 🪄
