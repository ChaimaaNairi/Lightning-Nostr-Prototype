# Nodes

In this section, we explore the configuration and transactions of three Lightning Network nodes: **Alice**, **Bob**, and **Charlie**. Operating within the same local environment, these nodes showcase the versatility of Lightning Network transactions.

**Transactions Demonstrated:**
- **Single-hop payments**: Alice to Bob
- **Multiple-hop payments**: Bob to Charlie

<p align="center">
  <img src="images/image.png" alt="Lightning Network Transaction" width="400" height="180"/>
</p>

***Image:** Illustration of single and multiple-hop payments in the Lightning Network.*

After configuring the `lnd.conf` file, we can initiate our [Lightning Network ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/LightningNetwork.md) nodes using the following commands:
```bash
alice$ lnd --rpclisten=localhost:10001 --listen=localhost:10011 --restlisten=localhost:8001
bob$ lnd --rpclisten=localhost:10002 --listen=localhost:10012 --restlisten=localhost:8002
charlie$ lnd --rpclisten=localhost:10003 --listen=localhost:10013 --restlisten=localhost:8003
```

## Alice Node
- Create a wallet for Alice on her Lightning Network node:
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon create
```
- Retrieve information about Alice's wallet and test it:
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon getinfo
```
<img src="images/AliceGetInfo.png" alt="Alice getInfo">

- Set up a Bitcoin address for Alice, which will serve as the address storing Alice’s on-chain balance:
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon newaddress np2wkh
```
<img src="images/AliceBitcoinAddress.png" alt="Alice's Bitcoin address">

- To fund Alice, we need to run this command in a new terminal to ensure that Alice is set as the recipient of all mining rewards:
```bash
alice$ btcd --simnet --txindex --rpcuser=username --rpcpass=password --miningaddr=<ALICE_ADDRESS>
```
<img src="images/AliceRunBtcd.png" alt="">

- Generate <number_of_blocks> blocks to ensure that Alice receives the rewards:
```bash
alice$ btcctl --simnet --rpcuser=username --rpcpass=password generate <number_of_blocks>
```
<img src="images/AliceGenerateBlock.png" alt="">

- Check that segwit is active:
```bash
alice$ btcctl --simnet --rpcuser=username --rpcpass=password getblockchaininfo | grep -A 1 segwit
```
<img src="images/AliceGetChainInfo.png" alt="">

- Check Alice’s wallet balance:
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon walletbalance
```
<img src="images/AliceWalletBalance.png" alt="">

## Bob Node  
- Create a wallet for Bob on her Lightning Network node:
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon create
```
- Retrieve information about Bob's wallet and test it:
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon getinfo
```
<img src="images/BobGetInfo.png" alt="">

- Set up a Bitcoin address for Bob, which will serve as the address storing Bob’s on-chain balance:
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon newaddress np2wkh
```
<img src="images/BobBitcoinAddress.png" alt="">

## Charlie Node
- To create a wallet for Charlie on her Lightning Network node:
```bash
charlie$ lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/charlie.macaroon create
```
- To retrieve information about Charlie's wallet and test it:
```bash
charlie$ lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/charlie.macaroon getinfo
```
<img src="images/CharlieGetInfo.png" alt="">

- Set up a Bitcoin address for Charlie, which will serve as the address storing Charlie's on-chain balance:
```bash
charlie$ lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/charlie.macaroon newaddress np2wkh
```
<img src="images/CharlieBitcoinAddress.png" alt="">

- To fund Charlie, we need to run this command in a new terminal to ensure that Alice is set as the recipient of all mining rewards:
```bash
charlie$ btcd --simnet --txindex --rpcuser=username --rpcpass=password --miningaddr=<CHARLIE_ADDRESS>
```
- Generate <number_of_blocks> blocks to ensure that Charlie receives the rewards:
```bash
charlie$ btcctl --simnet --rpcuser=username --rpcpass=password generate <number_of_blocks>
```
<img src="images/CharlieGenerateBlock.png" alt="">

- Check that segwit is active:
```bash
charlie$ btcctl --simnet --rpcuser=username --rpcpass=password getblockchaininfo | grep -A 1 segwit
```
<img src="images/CharlieGetChainInfo.png" alt="">

- Check Charlie's wallet balance.
```bash
charlie$ lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/charlie.macaroon walletbalance
```
<img src="images/CharlieWalletBalance.png" alt="">
 
## Creating the P2P Network
After funding Alice and Charlie, they now possess some simnet Bitcoin. Now, we need to connect them together.

**Connect Alice to Bob:**
1. **Get Bob's identity pubkey:**
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon getinfo
```
<img src="images/BobGetInfo.png" alt="">

2. **Connect Alice and Bob together:**
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon connect <BOB_PUBKEY>@localhost:10012
```
<img src="images/AliceBobConnection.png" alt="">

3. **Check that Alice has added Bob as a peer:**
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon listpeers
```
<img src="images/AliceListpeers.png" alt="">

4. **Check that Bob has added Alice as a peer:**
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon listpeers
```
<img src="images/BobListpeers.png" alt="">

**Connect Bob to Charlie:**
```bash
charlie$ lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/charlie.macaroon connect <BOB_PUBKEY>@localhost:10012
```
<img src="images/CharlieBobConnection.png" alt="">

- Check that Charlie has added Bob as a peer:
```bash
charlie$ lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon listpeers
```
<img src="images/CharlieListpeers.png" alt="">

- Check that Bob has added Charlie as a peer:
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon listpeers
```
<img src="images/BobListpeersBC.png" alt="">

## Setting up Lightning Network
Before initiating payments, we must establish payment channels from **Alice to Bob** and from **Bob to Charlie**.

- Open the Alice<–>Bob channel:
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon openchannel --node_key=<BOB_PUBKEY> --local_amt=1000000
```
<img src="images/AliceBobOpenChannel.png" alt="">

- We now need to mine six blocks so that the channel is considered valid:
```bash
alice$ btcctl --simnet --rpcuser=udername --rpcpass=password generate 6
```
<img src="images/AliceBobGenerateBlock.png" alt="">

- Check that Alice<–>Bob channel was created:
```bash
    alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon listchannels
```
<img src="images/AliceChannelList.png" alt="">

```bash
    bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon listchannels
```
<img src="images/BobChannelList.png" alt="">

## Single-hop payments: Alice to Bob
- Bob will need to generate an invoice to request payment from Alice:
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon addinvoice --amt=10000
```
<img src="images/BobInvoiceGeneration.png" alt="">

- Send the payment from Alice to Bob:
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon sendpayment --pay_req=<encoded_invoice>
```
<img src="images/AliceToBob_Transaction.png" alt="">

- Check that Alice's channel balance has been decremented accordingly:
```bash
alice$ lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon listchannels
```
<img src="images/AliceListChannels.png" alt="">

- Check that Bob's channel has been credited with the payment amount:
```bash
bob$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon listchannels
```
<img src="images/BobListChannels.png" alt="">

- Bob walletbalance:
```bash
alice$ lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon walletbalance
```
<img src="images/BobWalletBalanceA.png" alt="">

## Multiple-hop payments: Bob to Charlie
Sending multi-hop payments involves routing a payment through multiple channels.

- Set up a channel from Bob<–>Charlie:
```bash
charlie$ lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/charlie.macaroon openchannel --node_key=<BOB_PUBKEY> --local_amt=800000 --push_amt=200000
```


