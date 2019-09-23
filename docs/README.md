### Table of Contents
1. [About](#About)
2. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying pfcd via pfcctl](#PfcctlConfig)
    4. [Mining](#Mining)
3. [Help](#Help)
    1. [Network Configuration](#NetworkConfig)
    2. [Wallet](#Wallet)
4. [Contact](#Contact)
    1. [Community](#ContactCommunity)
5. [Developer Resources](#DeveloperResources)
    1. [Code Contribution Guidelines](#ContributionGuidelines)
    2. [JSON-RPC Reference](#JSONRPCReference)
    3. [Go Modules](#GoModules)
    4. [Module Hierarchy](#ModuleHierarchy)

<a name="About" />

### 1. About

pfcd is a full node Picfight implementation written in [Go](http://golang.org),
and is licensed under the [copyfree](http://www.copyfree.org) ISC License.

This software is currently under active development.  It is extremely stable and
has been in production use since February 2016.

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block.  It
ensures all individual transactions admitted to the pool follow the rules
required into the block chain and also includes the vast majority of the more
strict checks which filter transactions based on miner requirements ("standard"
transactions).

<a name="GettingStarted" />

### 2. Getting Started

<a name="Installation" />

**2.1 Installation**<br />

The first step is to install pfcd.  The installation instructions can be found
[here](https://github.com/picfight/pfcd/tree/master/README.md#Installation).

<a name="Configuration" />

**2.2 Configuration**<br />

pfcd has a number of [configuration](http://godoc.org/github.com/picfight/pfcd)
options, which can be viewed by running: `$ pfcd --help`.

<a name="PfcctlConfig" />

**2.3 Controlling and Querying pfcd via pfcctl**<br />

pfcctl is a command line utility that can be used to both control and query pfcd
via [RPC](http://www.wikipedia.org/wiki/Remote_procedure_call).  pfcd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password or both an RPC limited username and password:

* pfcd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
* pfcctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
OR
```
[Application Options]
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
For a list of available options, run: `$ pfcctl --help`

<a name="Mining" />

**2.4 Mining**<br />
pfcd supports the [getwork](https://github.com/picfight/pfcd/tree/master/docs/json_rpc_api.md#getwork)
RPC.  The limited user cannot access this RPC.<br />

**1. Add the payment addresses with the `miningaddr` option.**<br />

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=DsExampleAddress1
miningaddr=DsExampleAddress2
```

**2. Add pfcd's RPC TLS certificate to system Certificate Authority list.**<br />

`cgminer` uses [curl](http://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `pfcd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**<br />

1. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/user/.pfcd/rpc.cert /usr/share/ca-certificates/pfcd.crt`<br />
2. Add pfcd.crt to /etc/ca-certificates.conf: `# echo pfcd.crt >> /etc/ca-certificates.conf`<br />
3. Update the CA certificate list: `# update-ca-certificates`<br />

**3. Set your mining software url to use https.**<br />

`$ cgminer -o https://127.0.0.1:9109 -u rpcuser -p rpcpassword`

<a name="Help" />

### 3. Help

<a name="NetworkConfig" />

**3.1 Network Configuration**<br />
* [What Ports Are Used by Default?](https://github.com/picfight/pfcd/tree/master/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/picfight/pfcd/tree/master/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/picfight/pfcd/tree/master/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring pfcd with Tor](https://github.com/picfight/pfcd/tree/master/docs/configuring_tor.md)

<a name="Wallet" />

**3.2 Wallet**<br />

pfcd was intentionally developed without an integrated wallet for security
reasons.  Please see [pfcwallet](https://github.com/picfight/pfcwallet) for more
information.

<a name="Contact" />

### 4. Contact

<a name="ContactCommunity" />

**4.1 Community**<br />

If you have any further questions you can find us at:

https://picfight.org/community

<a name="DeveloperResources" />

### 5. Developer Resources

<a name="ContributionGuidelines" />

**5.1 Code Contribution Guidelines**

* [Code Contribution Guidelines](https://github.com/picfight/pfcd/tree/master/docs/code_contribution_guidelines.md)

<a name="JSONRPCReference" />

**5.2 JSON-RPC Reference**

* [JSON-RPC Reference](https://github.com/picfight/pfcd/tree/master/docs/json_rpc_api.md)
    * [RPC Examples](https://github.com/picfight/pfcd/tree/master/docs/json_rpc_api.md#ExampleCode)

<a name="GoModules" />

**5.3 Go Modules**

The following versioned modules are provided by pfcd repository:

* [rpcclient](https://github.com/picfight/pfcd/tree/master/rpcclient) - Implements
  a robust and easy to use Websocket-enabled Picfight JSON-RPC client
* [pfcjson](https://github.com/picfight/pfcd/tree/master/pfcjson) - Provides an
  extensive API for the underlying JSON-RPC command and return values
* [wire](https://github.com/picfight/pfcd/tree/master/wire) - Implements the
  Picfight wire protocol
* [peer](https://github.com/picfight/pfcd/tree/master/peer) - Provides a common
  base for creating and managing Picfight network peers
* [blockchain](https://github.com/picfight/pfcd/tree/master/blockchain) -
  Implements Picfight block handling and chain selection rules
  * [stake](https://github.com/picfight/pfcd/tree/master/blockchain/stake) -
    Provides an API for working with stake transactions and other portions
    related to the Proof-of-Stake (PoS) system
* [txscript](https://github.com/picfight/pfcd/tree/master/txscript) -
  Implements the Picfight transaction scripting language
* [pfcec](https://github.com/picfight/pfcd/tree/master/pfcec) - Provides constants
  for the supported cryptographic signatures supported by Picfight scripts
  * [secp256k1](https://github.com/picfight/pfcd/tree/master/pfcec/secp256k1) -
    Implements the secp256k1 elliptic curve
  * [edwards](https://github.com/picfight/pfcd/tree/master/pfcec/edwards) -
    Implements the edwards25519 twisted Edwards curve
* [database](https://github.com/picfight/pfcd/tree/master/database) -
  Provides a database interface for the Picfight block chain
* [mempool](https://github.com/picfight/pfcd/tree/master/mempool) - Provides a
  policy-enforced pool of unmined Picfight transactions
* [pfcutil](https://github.com/picfight/pfcd/tree/master/pfcutil) - Provides
  Picfight-specific convenience functions and types
* [chaincfg](https://github.com/picfight/pfcd/tree/master/chaincfg) - Defines
  chain configuration parameters for the standard Picfight networks and allows
  callers to define their own custom Picfight networks for testing puproses
  * [chainhash](https://github.com/picfight/pfcd/tree/master/chaincfg/chainhash) -
    Provides a generic hash type and associated functions that allows the
    specific hash algorithm to be abstracted
* [certgen](https://github.com/picfight/pfcd/tree/master/certgen) - Provides a
  function for creating a new TLS certificate key pair, typically used for
  encrypting RPC and websocket communications
* [addrmgr](https://github.com/picfight/pfcd/tree/master/addrmgr) - Provides a
  concurrency safe Picfight network address manager
* [connmgr](https://github.com/picfight/pfcd/tree/master/connmgr) - Implements a
  generic Picfight network connection manager
* [hdkeychain](https://github.com/picfight/pfcd/tree/master/hdkeychain) - Provides
  an API for working with  Picfight hierarchical deterministic extended keys
* [gcs](https://github.com/picfight/pfcd/tree/master/gcs) - Provides an API for
  building and using Golomb-coded set filters useful for light clients such as
  SPV wallets
* [fees](https://github.com/picfight/pfcd/tree/master/fees) - Provides methods for
  tracking and estimating fee rates for new transactions to be mined into the
  network

<a name="ModuleHierarchy" />

**5.4 Module Hierarchy**

The following diagram shows an overview of the hierarchy for the modules
provided by the pfcd repository.

![Module Hierarchy](./assets/module_hierarchy.svg)