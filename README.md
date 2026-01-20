Toy-RTB: High-Performance Go Bidder with Cloud Bigtable
=======================================================

A simulation of a **Demand-Side Platform (DSP)** "Hot Path" bidder. This project demonstrates sub-100ms real-time bidding (RTB) logic using **Golang** for high-concurrency request handling and **Google Cloud Bigtable** for ultra-low latency user profile lookups.

🚀 Architecture Overview
------------------------

In AdTech, the "Golden Rule" is a **100ms round-trip latency budget**. This project implements the **Critical Path** (Hot Path) where every millisecond counts.

*   **Language:** Golang (utilizing Goroutines for parallel request processing).
    
*   **Database:** Google Cloud Bigtable (NoSQL) for $O(1)$ key-value lookups.
    
*   **Throughput:** Optimized for thousands of Queries Per Second (QPS).
    

### The Data Flow

1.  **Ingestion (Offline):** seed.go populates Bigtable with user segments (e.g., user\_123 is interested in tech).
    
2.  **Request (Real-Time):** An Ad Exchange sends a Bid Request via HTTP.
    
3.  **Lookup:** The Go server performs a single-row lookup in Bigtable using the user\_id.
    
4.  **Decision:** If a high-value segment is found, the server responds with a bid price; otherwise, it returns a "No-Bid" (0.00).
    

🛠️ Project Structure
---------------------

*   main.go: The high-concurrency Bidder server.
    
*   seed.go: Administrative script to populate the Bigtable database with mock user data.
    
*   go.mod: Dependency management for the Google Cloud SDK.
    

⚡ Performance Benchmarks
------------------------

Using hey for load testing, this system achieves the following P99 latencies (tested in a simulated environment):

| **Metric** | **Result** |
| :--- | :--- |
| **P99 Latency** | 8.9ms |
| **P90 Latency** | 4.7ms |
| **Max Throughput** | 3,707 QPS |
| **Average Latency** | ~2.6ms |

💻 Local Development (Emulator)
-------------------------------

To run this project for free without incurring GCP costs, we use the **Bigtable Emulator**.

### 1\. Prerequisites

```bash
gcloud components install bigtable-emulator
```

### 2\. Start the Environment

**Terminal 1: Start Emulator**

```bash
gcloud beta emulators bigtable start --host-port=localhost:8086
```

**Terminal 2: Run Bidder**

```bash
export BIGTABLE_EMULATOR_HOST=localhost:8086  go run main.go
```

**Terminal 3: Seed & Test**

```bash
export BIGTABLE_EMULATOR_HOST=localhost:8086  &
go run seed.go  &
hey -n 1000 -c 10 "http://localhost:8080/bid?user_id=user_123"
```

📈 Future Roadmap
-----------------

*   \[ \] **Cold Path Integration:** Add a Pub/Sub sink to stream bid logs to BigQuery for analytics.
    
*   \[ \] **Bid Shading:** Implement a basic algorithm to adjust bid prices based on historical win rates.
    
*   \[ \] **OpenRTB Compliance:** Upgrade JSON structures to match the OpenRTB 2.5/3.0 specification.
