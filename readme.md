# Celestia Domain Rollup by Upnode

A blockchain built with Rollkit and Ignite CLI to daemonstrate domain system rollup

## Requirements

1. Install go 1.20.3

```bash
# Install go
cd $HOME
ver="1.20.3"
wget "https://golang.org/dl/go$ver.linux-amd64.tar.gz"
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf "go$ver.linux-amd64.tar.gz"
rm "go$ver.linux-amd64.tar.gz"
echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> $HOME/.bash_profile
source $HOME/.bash_profile
```

2. Install Celestia light node and connect to blocksapcerace-0 testnet

```bash
cd $HOME 
git clone https://github.com/celestiaorg/celestia-node.git 
cd celestia-node 
git checkout v0.9.4 
make build

rm -f /usr/local/bin/celestia
sudo mv $HOME/celestia-node/build/celestia /usr/local/bin/

make cel-key

rm -f /usr/local/bin/cel-key
sudo mv $HOME/celestia-node/cel-key /usr/local/bin/
```

3. Start Celestia light node

```bash
celestia light start --core.ip https://rpc-blockspacerace.pops.one --gateway --gateway.addr 127.0.0.1 --gateway.port 26659 --p2p.network blockspacerace
```

4. Run init-local.sh

```bash
bash init-local.sh
```

Accept anything it has prompted and you will have your celestia domain running in your machine

Then open a new terminal

## Commands

### Register a domain name

Note: only .tia TLD is supported

```bash
celestia-domaind tx celestiadomain register-domain upnode.tia --from celestia-domain-key --keyring-backend test --chain-id celestia-domain
```

### Register a subdomain

```bash
celestia-domaind tx celestiadomain register-domain sub.upnode.tia --from celestia-domain-key --keyring-backend test --chain-id celestia-domain
```

### Set primary domain

```bash
celestia-domaind tx celestiadomain set-primary-domain upnode.tia --from celestia-domain-key --keyring-backend test --chain-id celestia-domain
```

## Queries

### Get domain information

```bash
celestia-domaind query celestiadomain domain upnode.tia
```

```
domain: upnode.tia
expiration: "1715801930"
owner: celestia148lrvfkc26xls02yhskut2vy0pxf4r64zkkj2h
```

**Subdomain example**

```bash
celestia-domaind query celestiadomain domain sub.upnode.tia
```

```
domain: sub.upnode.tia
expiration: "1715801960"
owner: celestia148lrvfkc26xls02yhskut2vy0pxf4r64zkkj2h
```

### List owned domains

Note: replace celestia address with your address

```bash
celestia-domaind query celestiadomain domains celestia148lrvfkc26xls02yhskut2vy0pxf4r64zkkj2h
```

```
domains:
- domain: upnode.tia
  expiration: "1715801930"
  owner: celestia148lrvfkc26xls02yhskut2vy0pxf4r64zkkj2h
  parent: tia
- domain: sub.upnode.tia
  expiration: "1715801960"
  owner: celestia148lrvfkc26xls02yhskut2vy0pxf4r64zkkj2h
  parent: upnode.tia
pagination:
  next_key: null
  total: "2"
```

### Get primary domain name

Note: replace celestia address with your address

```bash
celestia-domaind query celestiadomain primary-domain celestia148lrvfkc26xls02yhskut2vy0pxf4r64zkkj2h
```

```
domain:
  domain: upnode.tia
  expiration: "1715801930"
  owner: celestia148lrvfkc26xls02yhskut2vy0pxf4r64zkkj2h
  parent: tia
```

## How to get your address

The easiest way is to register a domain and query that domain. That's the point of domain, to simplify complex address.

Another way is to use `celestia-domaind keys list --keyring-backend=test`

```bash
celestia-domaind keys list --keyring-backend=test
- address: celestia16836wtdzhgehqc7wcn082ffnhm9d9vhdlug2dc
  name: celestia-domain-key
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ApeFVo3kvS5eaOjfg9m/+sLRBPf1abUDIphMzZOuBf8x"}'
  type: local
```
