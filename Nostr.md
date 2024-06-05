# Nostr Protocol - Decentralized Social Networking

The Nostr Protocol is a decentralized social networking solution designed to provide users with a censorship-resistant and privacy-focused platform for communication.

#### Why Nostr Protocol?
Centralized social networking platforms often face issues related to censorship, data privacy, and control over user data. The Nostr Protocol addresses these concerns by decentralizing the communication infrastructure, enabling users to interact directly with each other without relying on centralized servers.

#### How Nostr Protocol Works?
The Nostr Protocol operates by establishing a network of relays that facilitate the distribution of messages between users. Here's an overview of its functioning:

1. **Relay Network**: The Nostr Protocol relies on a network of relays distributed across the internet. These relays serve as intermediaries for message propagation, ensuring that messages reach their intended recipients.

2. **Message Distribution**: Users post messages, also known as "notes," to the network of relays. These relays propagate the messages to other relays and clients connected to the network.

3. **Decentralized Identity**: Users can create and manage their identities without the need for a central authority. This allows for greater control over personal data and enhances user privacy.

4. **Censorship Resistance**: By distributing messages across a decentralized network, the Nostr Protocol makes it difficult for any single entity to censor communication. This ensures that users can express themselves freely without fear of censorship.

5. **Privacy**: Users' interactions on the Nostr network are private and secure, as messages are distributed in a peer-to-peer fashion without passing through centralized servers.

## Nostr Prototype  
https://github.com/nbd-wtf/go-nostr



-------------



------------------

- **sk:** stands for the generated private key.
- **pk:** stands for the corresponding public key.
- **nsec:** is the encoded private key.
- **npub:** is the encoded public key.


-------


## Alice client:

![alt text](<Capture d’écran -2.png>)

![alt text](<Capture d’écran -3.png>)

- **Code:** [generate_keys.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/alice/go-nostr/cmd/generate_keys.go)

-------


## Bob client:

![alt text](<Capture d’écran -4.png>)


![alt text](<Capture d’écran -5.png>)

- **Code:** [generate_keys.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/bob/go-nostr/cmd/generate_keys.go)

-------


## Charlie client:

![alt text](<Capture d’écran .png>)


![alt text](<Capture d’écran -1.png>)

- **Code:** [generate_keys.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/charlie/go-nostr/cmd/generate_keys.go)


## single relay

![alt text](<Capture d’écran -6.png>)


- **Code:** [relay_server.go ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/go-nostr/cmd/relay_server.go)
