# Lightning Network - L2 Prototype 

## Introduction

The Lightning Network Prototype repository aims to provide a hands-on guide to understanding and experimenting with the Lightning Network, a layer 2 scaling solution for Bitcoin and other blockchain-based cryptocurrencies.

## Getting Started

### Setting Up a Simnet Environment

1. **Install Bitcoin and Lightning Software:**
   - Install and set up a Bitcoin node **btcd**.
   - Install Lightning Network software **LND** (Lightning Network Daemon).

##### **LND**
**lnd** is the main component that we will interact with. lnd stands for Lightning Network Daemon, and handles channel opening/closing, routing and sending payments, and managing all the Lightning Network state that is separate from the underlying Bitcoin network itself.

##### **BTCD**
**btcd** represents the gateway that lnd nodes will use to interact with the Bitcoin / Litecoin network. lnd needs btcd for creating on-chain addresses or transactions, watching the blockchain for updates, and opening/closing channels. In our current schema, all three of the nodes are connected to the same btcd instance. In a more realistic scenario, each of the lnd nodes will be connected to their own instances of btcd or equivalent.





