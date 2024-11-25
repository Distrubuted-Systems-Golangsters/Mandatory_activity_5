# Auction

**Distributed Systems, BSc (Autumn 2024) - Mandatory Activity 5 - Auction**

**Date:** Sunday, 18 November 2024

**Group:** GoLangsters

**Contact:** Adot@itu.dk, Csti@itu.dk, Mhiv@itu.dk

## How to run the program:

1. **Clone the Repository:**

## First we start the servers for each server cd server/ 

2. **Start server1 node (In terminal):**

- go run main.go 1 8080

3. **Start server2 node (In a different terminal):**

- go run main.go 2 8081

4. **Start server3 node (In a different terminal):**

- go run main.go 3 8082

## Now we start the clients for each client cd client/ 

4. **Start a client node (In a different terminal):**

- go run main.go

5. **Note it is important that the servers number and port follow these exact names and that all three servers are started, aswell as at least one client, before continuing to the next step**

6. **From the client side type "start [time in secounds]" in the terminal to start an auction. Say you type "start 60" this will start an auction ending in 60 secounds**

7. **After starting an auction you can from every started client type an integer in the terminal say "5" to bid in the auction**

8. **Every client can type "status" in the terminal to query the current status of the auction**
