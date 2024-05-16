# Lightning Network - Layer 2  

The Lightning Network Prototype is a solution designed to enhance the scalability of Bitcoin transactions.

#### Why Lightning Network?
Bitcoin's blockchain has limitations in terms of transaction throughput and speed. As more users join the network, these limitations become more apparent, resulting in higher fees and slower confirmation times. The Lightning Network addresses these issues by enabling off-chain transactions, allowing for instant, low-cost micropayments. It aims to enhance scalability, reduce congestion on the blockchain, and improve the overall efficiency of the Bitcoin network.

#### How Lightning Network Works?
The Lightning Network works by creating a network of payment channels on top of the Bitcoin blockchain. These payment channels allow users to transact directly with each other without involving the blockchain for every transaction.

1. **Opening Channels**: Users open payment channels by funding a multisignature wallet on the blockchain. This wallet acts as a shared account between the parties involved.
2. **Off-chain Transactions**: Once the channel is open, users can conduct off-chain transactions by updating the balance in the channel. These transactions are private and instantaneous, occurring outside the blockchain.
3. **Routing Payments**: Payments can be routed through multiple channels in the network, allowing users to send funds to parties they do not have a direct channel with. This routing is facilitated by the network's peer-to-peer architecture.
4. **Closing Channels**: When users are finished transacting, they can close the payment channel by broadcasting the final state to the blockchain. This settles the balances and updates the respective parties' on-chain holdings.
5. **Scalability**: By keeping most transactions off-chain, the Lightning Network improves the scalability of the Bitcoin network, enabling it to handle a higher volume of transactions with lower fees and faster confirmation times.


## Lightning Network Prototype with Simnet Testing

### Setting Up a Simnet Environment
1. **Install Bitcoin and Lightning Software:**
   - Install Lightning Network software **LND** - [LND GitHub Repository](https://github.com/lightningnetwork/lnd)
      -  **lnd** is the main component that we will interact with. lnd stands for Lightning Network Daemon, and handles channel opening/closing, routing and sending payments, and managing all the Lightning Network state that is separate from the underlying Bitcoin network itself.
   - Install and set up a Bitcoin node **btcd** - [btcd GitHub Repository](https://github.com/btcsuite/btcd)
      - **btcd** represents the gateway that lnd nodes will use to interact with the Bitcoin/Litecoin network (BTC/LTC). lnd needs btcd for creating on-chain addresses or transactions, watching the blockchain for updates, and opening/closing channels. In our current schema, all three of the nodes are connected to the same btcd instance. In a more realistic scenario, each of the lnd nodes will be connected to their own instances of btcd or equivalent.
2. **Initialize the Simnet:**
   - Simnet (Simulation Network) is a local test network that simulates the Bitcoin network environment.
   - Initialize the Simnet by running the Bitcoin node with the `--simnet` flag.
   ```bash
   btcd --simnet --txindex --rpcuser=username --rpcpass=password
   ```

### Initialize Lightning Nodes
Use Lightning Network software **LND** to initialize Lightning nodes for Alice, Bob, and Charlie respectively on your local machine. Ensure that each node is started with unique port configurations to avoid conflicts.

To set up Lightning Network nodes for Alice, Bob, and Charlie, follow these steps:

1.**Starting lnd (Alice’s node):**
   - Run the following command for Alice's node:
   ```bash
   alice$ lnd --rpclisten=localhost:10001 --listen=localhost:10011 --restlisten=localhost:8001 --datadir=data --logdir=log --debuglevel=info --bitcoin.simnet --bitcoin.active --bitcoin.node=btcd --btcd.rpcuser=username --btcd.rpcpass=password 
   ```
2.**Starting lnd (Bob’s node):**
   - Run the following command for Bob's node:
```bash
bob$ lnd --rpclisten=localhost:10002 --listen=localhost:10012 --restlisten=localhost:8002 --datadir=data --logdir=log --debuglevel=info --bitcoin.simnet --bitcoin.active --bitcoin.node=btcd --btcd.rpcuser=username --btcd.rpcpass=password 
```
3.**Starting lnd (Charlie’s node):**
- Run the following command for Charlie's node:
```bash
charlie$ lnd --rpclisten=localhost:10003 --listen=localhost:10013 --restlisten=localhost:8003 --datadir=data --logdir=log --debuglevel=info --bitcoin.simnet --bitcoin.active --bitcoin.node=btcd --btcd.rpcuser=username --btcd.rpcpass=password 
```

These commands initialize Lightning nodes for Alice, Bob, and Charlie respectively on your local machine. Ensure that each node is started with unique port configurations to avoid conflicts.

### Configuration of btcctl.conf, btcd.conf, and lnd.conf
Managing the configurations of btcctl, btcd, and lnd is essential for setting up and maintaining Lightning Network nodes.

 - **btcctl.conf** configuration - Path: `C:\Users\${name}\AppData\Local\btcctl\btcctl.conf`
```bash
rpcuser=username
rpcpass=password
```

 - **btcd.conf** configuration - Path: `C:\Users\${name}\AppData\Local\btcd\btcd.conf`
```bash
[Network settings]
simnet=1
; Alice
rpclisten=127.0.0.1:8332
; Bob
rpclisten=127.0.0.1:8333
; Charlie
rpclisten=127.0.0.1:8334

;Alice
listen=localhost:10011
;Bob
listen=localhost:10012
;Charlie
listen=localhost:10013

 rpcuser=username
 rpcpass=password
```

 - **lnd.conf** configuration - Path: `C:\Users\${name}\AppData\Local\lnd\lnd.conf`
```bash
[Application Options]
datadir=data
logdir=logs
debuglevel=info

; Alice 
   rpclisten=localhost:10001
; Bob 
   rpclisten=localhost:10002
; Charlie
   rpclisten=localhost:10003

; Alice 
   restlisten=localhost:8001
; Bob 
   restlisten=localhost:8002
; Charlie
   restlisten=localhost:8003

adminmacaroonpath=~/.lnd/data/chain/bitcoin/simnet/admin.macaroon
readonlymacaroonpath=~/.lnd/data/chain/bitcoin/simnet/readonly.macaroon
invoicemacaroonpath=~/.lnd/data/chain/bitcoin/simnet/invoice.macaroon

[Bitcoin]
bitcoin.simnet=true
bitcoin.active=true
bitcoin.node=btcd

[btcd]
btcd.rpcuser=username
btcd.rpcpass=password
```

Now, when we start nodes, we only have to type
```bash
alice$ lnd --rpclisten=localhost:10001 --listen=localhost:10011 --restlisten=localhost:8001
bob$ lnd --rpclisten=localhost:10002 --listen=localhost:10012 --restlisten=localhost:8002
charlie$ lnd --rpclisten=localhost:10003 --listen=localhost:10013 --restlisten=localhost:8003
```

Here, you can find more detailed information, including all steps and transactions related to transactions between nodes, both single and multiple hop payments, in the [Nodes](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/Nodes.md) section.

### Setting Up Lightning Network
1. **Create Payment Channels:**
   - Open payment channels with other Lightning nodes to establish off-chain payment channels.
   - These channels allow for instant and low-cost transactions without touching the main blockchain.
   
   ```bash
   lncli openchannel --node_key=<peer_pubkey> --local_amt=<local_amount>
   ```
2. **Generate Blocks:**
   - Use the Bitcoin node's RPC interface to generate blocks on the Simnet.
   - This creates an environment where we can test Lightning Network transactions.
   ```bash
   btcctl --simnet --rpcuser=username --rpcpass=password generate <number_of_blocks>
   ```
3. **Perform Lightning Transactions:**
   - Once payment channels are established, we can send and receive payments instantly through Lightning transactions.
   - Lightning transactions occur off-chain, providing high scalability and low fees compared to on-chain transactions.
   ```bash
   lncli sendpayment --pay_req=<encoded_invoice>
   ```

### Test Lightning Transactions
   - Test Lightning transactions between our Lightning node and other connected nodes on the Simnet.
   - Monitor transaction speed, cost, and overall performance of the Lightning Network.
   ```bash
   lncli getinfo
   ```

### Closing channels
   - Close payment channels to settle balances and reclaim funds tied up in the Lightning Network.
   - Ensure that all pending transactions are confirmed before initiating the channel closure process. 
   ```bash  
      lncli closechannel --funding_txid=<funding_txid> --output_index=<output_index>
   ```

