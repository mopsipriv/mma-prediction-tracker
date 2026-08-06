# 🥊 MMA Prediction Tracker

A Go-backend service for tracking, comparing, and benchmarking UFC fight predictions between human analysis and AI models (GPT-4, Gemini, Claude).

## 🚀 Key Features

- **Predictor Benchmarking:** Compares accuracy and probability calibration (Brier Score) between human picks and various AI models.
- **Categorized Analytics:** Tracks accuracy broken down by weight classes, fighter stats, and win methods.
- **Clean Architecture:** Idiomatic Go implementation with explicit error handling, direct SQL (`database/sql` + `pgx`), and layered structure.

## 🛠 Tech Stack

- **Language:** Go 1.22+
- **Database:** PostgreSQL 16
- **Containerization:** Docker & Docker Compose
- **Database Driver:** `pgx` (v5)

## 🏎 Quick Start

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/mopsipriv/mma-prediction-tracker.git](https://github.com/mopsipriv/mma-prediction-tracker.git)
   cd mma-prediction-tracker

2. Start PostgreSQL container:

docker compose up -d

3. Apply Database Schema:

cat database/schema.sql | docker exec -i mma_tracker_db psql -U postgres -d mma_tracker

4. Run the backend:

go run main.go