## To Avoid Errors
When working with multiple nodes in the same local environment, you might encounter the following errors:

**1. Signature Mismatch Error:**

```bash
[lncli] rpc error: code = Unknown desc = verification failed: signature mismatch after caveat verification
```
This error occurs due to issues with macaroons. Instead of using a single `admin.macaroon` for all nodes, create individual macaroons for each node: `alice.macaroon`, `bob.macaroon` and `charlie.macaroon`.

These macaroons should be placed in the folder:
```bash
C:\Users\{yourusername}\AppData\Local\lnd\data\chain\bitcoin\simnet
```


**2. RPC Server Not Ready Error:**

```bash
[lncli] rpc error: code = Unknown desc = the RPC server is in the process of starting up, but not yet ready to accept calls
```

To resolve this, check the following:

- Ensure all configurations in `btcd.conf` and `lnd.conf` are correct, including `listen`, `rpcrest`, and `rpclisten` ports. Refer to the configuration part [Lightning Network ](https://github.com/ChaimaaNairi/Lightning-Nostr-Prototype/blob/main/LightningNetwork.md).

- Verify that `btcd` is running. You can start btcd with the following command:
```bash
btcd --txindex --simnet --rpcuser=username --rpcpass=password
```

## .bashrc
To avoid repetition, you can save the necessary information in the `.bashrc` file using aliases for each Lightning Network node:
```bash
alias lncli-alice="lncli --rpcserver=localhost:10001 --macaroonpath=data/chain/bitcoin/simnet/alice.macaroon"
alias lncli-bob="lncli --rpcserver=localhost:10002 --macaroonpath=data/chain/bitcoin/simnet/bob.macaroon"
alias lncli-charlie="lncli --rpcserver=localhost:10003 --macaroonpath=data/chain/bitcoin/simnet/charlie.macaroon"
```

