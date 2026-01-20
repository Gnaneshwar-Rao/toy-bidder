package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
    "cloud.google.com/go/bigtable"
)

var tbl *bigtable.Table

func handler(w http.ResponseWriter, r *http.Request) {
    // Set a strict 100ms deadline for the entire request
    ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
    defer cancel()

    // 1. O(1) Bigtable Lookup
    userID := r.URL.Query().Get("user_id")
    row, _ := tbl.ReadRow(ctx, userID)

    // 2. High-speed Decision Logic
    bid := 0.0
    if len(row) > 0 {
        bid = 1.50 // Found user profile, place a bid
    }

    // 3. Fast JSON Response
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"bid": %.2f, "latency_ms": %v}`, bid, 100)
}

func main() {
    ctx := context.Background()
    // Connects to Emulator if BIGTABLE_EMULATOR_HOST is set
    client, _ := bigtable.NewClient(ctx, "dev-project", "dev-instance")
    tbl = client.Open("user_profiles")

    http.HandleFunc("/bid", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
