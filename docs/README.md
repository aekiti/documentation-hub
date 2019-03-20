# INTRODUCTION
## Welcome to æternity's Documentation Hub! (work in progress)

This Documentation Hub will be your best technical guide on æternity as a whole based on the model of the “Mastering Bitcoin” book, mainly focused on dApp (aepp) developer and miner. It should help you to clear the haze around æternity’s technology and whole ecosystem with a clear structure and well written documentation. æternity’s Documentation Hub should become your guide through the seemingly complex ecosystem of æternity, providing the knowledge you need for participating and mastering æternity. Whether you’re building the next killer aepp, investing in a startup, or simply curious about the technology itself, it will provide all essential details to get you started and to accompany you through your personal journey till æternity.

If you don’t know about æternity yet, please feel free to visit [aeternity.com](www.aeternity.com)

If you’d like to see a specific improvement you can either submit a pull request or file an issue. 
Please follow our [contribution guidelines](https://github.com/aeternity/aeternity/blob/master/CONTRIBUTING.md).

# [æternity Node](https://github.com/aeternity/aeternity/blob/master/README.md)
- [**Threat Model**](https://github.com/aeternity/aetmodel/blob/master/ThreatModel.md)
- [**Infrastructure Management**](https://github.com/aeternity/infrastructure/blob/master/README.md)
- [**Contribution**](https://github.com/aeternity/aeternity/blob/master/CONTRIBUTING.md)
- [**Issue Template**](https://github.com/aeternity/aeternity/blob/master/ISSUE_TEMPLATE.md)

# [æternity Node Usage API's](https://github.com/aeternity/protocol/blob/master/node/api/README.md#overview)
- [**Encoding Scheme**](https://github.com/aeternity/protocol/blob/master/node/api/api_encoding.md)
- [**Account Management**](https://github.com/aeternity/protocol/blob/master/node/api/account_api_usage.md) 
- [**Spending Tokens**](https://github.com/aeternity/protocol/blob/master/node/api/spend_api_usage.md) 
- [**Oracle**](https://github.com/aeternity/protocol/blob/master/node/api/oracle_api_usage.md) 
- [**AENS Naming System**](https://github.com/aeternity/protocol/blob/master/node/api/naming_system_api_usage.md) 
- [**Smart Contract**](https://github.com/aeternity/protocol/blob/master/node/api/contract_api_usage.md) 
- [**State Channels**](https://github.com/aeternity/protocol/blob/master/node/api/channels_api_usage.md) 
- [**Mining**](https://github.com/aeternity/protocol/blob/master/node/api/mining_api_usage.md)

# DEVELOPER TOOLS
## æternity SDK's
- [**Go**](https://github.com/aeternity/aepp-sdk-go/README.md)
    - [CLI](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli.md)
        -  Query
            -  [State of Chain](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_chain.md)
            -  [Blocks of Chain](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_chain_play.md)
            -  [Top Block of Chain](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_chain_top.md)
            -  [Status and Version](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_chain_version.md)
        -  [Config](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_config.md)
            -  [Config Profile](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_config_profile.md)
            -  [Config Profile Activate](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_config_profile_activate.md)
            -  [Config Profile Create](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_config_profile_create.md)
            -  [Config Profile List](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_config_profile_list.md)
        -  [Inspect](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_inspect.md)
        -  [Name](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_name.md)
            -  [Name Claim](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_name_claim.md)
        -  [Wallet](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_wallet.md)
            -  [Wallet Address](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_wallet_address.md)
            -  [Wallet Balance](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_wallet_balance.md)
            -  [Wallet Create](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_wallet_create.md)
            -  [Wallet Save](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_wallet_save.md)
            -  [Wallet Spend](https://github.com/aeternity/aepp-sdk-go/blob/develop/doc/cli/aecli_wallet_spend.md)
--------
- [**JavaScript**](https://github.com/aeternity/aepp-sdk-js/blob/develop/README.md)
    - [Change Log](https://github.com/aeternity/aepp-sdk-js/blob/develop/CHANGELOG.md)
    - [API](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api.md)
        - [Account](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/account.md)
            - [Memory](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/account/memory.md)
            - [Selector](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/account/selector.md)
        - [Accounts](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/accounts.md)
        - [Ae](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/ae.md)
            - [Aens](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/ae/aens.md)
            - [Aepp](https://github.com/aeternity/aepp-sdk-js/commit/275f516b163e27d40180361274775fb93dacd20f)
            - [Contract](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/ae/contract.md)
            - [Oracle](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/ae/oracle.md)
            - [Universal](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/ae/universal.md)
            - [Wallet](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/ae/wallet.md)
        - [Chain](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/chain.md)
            - [Node](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/chain/node.md)
        - [Contract](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/contract.md)
            - [Node](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/contract.md)
        - [Node](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/node.md)
        - [Oracle](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/oracle.md)
            - [Node](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/oracle/node.md)
        - RPC
            - [Client](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/rpc/client.md)
        - [Transactions](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/tx.md)
            - [Builder](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/tx/builder.md)
                - [Helpers](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/tx/builder/helpers.md)
                - [Schema](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/tx/builder/schema.md)
            - [Transaction Module](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/tx/tx.md)
            - [Validator](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/tx/validator.md)
        - Utils 
            - [Bytes](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/utils/bytes.md)
            - [Crypto](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/utils/crypto.md)
            - [Keystore](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/utils/keystore.md)
            - [Swagger](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/api/utils/swagger.md)
    - Bin 
        - [Sophia Contract Compiler](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/bin/aecontract.md)
        - [æternity Crypto Helper Script](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/bin/aecrypto.md)
        - [AE Token Spending Script](https://github.com/aeternity/aepp-sdk-js/blob/develop/docs/bin/aewallet.md)
-----
- [**Python**](https://github.com/aeternity/aepp-sdk-python)
    - [Install](https://github.com/aeternity/aepp-sdk-python/blob/develop/INSTALL.md)
    - [CLI](https://github.com/aeternity/aepp-sdk-python/blob/develop/command-line-tool.md) 
-----
## CLI tools
- [**JavaScript**](https://github.com/aeternity/aepp-cli-js)
- [**Stratum**](https://github.com/aeternity/aestratum_client)
    - [Library](https://github.com/aeternity/aestratum_lib)
- [**Websocket**](https://github.com/aeternity/websocket_client)
-----
## Middleware
- [**æternity Middleware**](https://github.com/aeternity/aepp-middleware/blob/develop/README.md)
-----
## Frameworks
- [**ForgAE**](https://github.com/aeternity/aepp-forgae-js/blob/develop/README.md)
    - Tutorials
        - [Deploying a Smart Contract on æternity](https://github.com/aeternity/tutorials/blob/master/smart-contract-deployment-in-forgae.md)
        - [Deploying a Smart Contract with init parameters](https://github.com/aeternity/tutorials/blob/master/deploy-with-init-params.md)
- [**æpp Components**](https://github.com/aeternity/aepp-components/blob/develop/README.md)
    - [address](http://aeternity.com/aepp-components/#aeaddress)
    - [backdrop](http://aeternity.com/aepp-components/#aebackdrop)
    - [badge](http://aeternity.com/aepp-components/#aebadge)
    - [button](http://aeternity.com/aepp-components/#aebutton)
    - [button-group](http://aeternity.com/aepp-components/#aebuttongroup)
    - [card](http://aeternity.com/aepp-components/#aecard)
    - [check](http://aeternity.com/aepp-components/#aecheck)
    - [dropdown](http://aeternity.com/aepp-components/#aedropdown)
    - [flip](http://aeternity.com/aepp-components/#aeflip)
    - [icon](http://aeternity.com/aepp-components/#aeicon)
    - [identicon](http://aeternity.com/aepp-components/#aeidenticon)
    - [input](http://aeternity.com/aepp-components/#aeinput)
    - [input-plain](http://aeternity.com/aepp-components/#aeinputplain)
    - [list](http://aeternity.com/aepp-components/#aelist)
    - [list-item](http://aeternity.com/aepp-components/#aelistitem)
    - [panel](http://aeternity.com/aepp-components/#aepanel)
    - [phraser](http://aeternity.com/aepp-components/#aephraser)
    - [qrcode](http://aeternity.com/aepp-components/#aeqrcode)
    - [slider](http://aeternity.com/aepp-components/#aeslider)
    - [text](http://aeternity.com/aepp-components/#aetext)
    - [toolbar](http://aeternity.com/aepp-components/#aetoolbar)
    - [view](http://aeternity.com/aepp-components/#aeview)
    - [address-input](http://aeternity.com/aepp-components/#aeaddressinput)
    - [amount](http://aeternity.com/aepp-components/#aeamount)
    - [amount-input](http://aeternity.com/aepp-components/#aeamountinput)
    - [app-icon](http://aeternity.com/aepp-components/#aeappicon)
    - [banner](http://aeternity.com/aepp-components/#aebanner)
    - [divider](http://aeternity.com/aepp-components/#aedivider)
    - [filter-item](http://aeternity.com/aepp-components/#aefilteritem)
    - [filter-list](http://aeternity.com/aepp-components/#aefilterlist)
    - [filter-separator](http://aeternity.com/aepp-components/#aefilterseparator)
    - [header](http://aeternity.com/aepp-components/#aeheader)
    - [identity](http://aeternity.com/aepp-components/#aeidentity)
    - [identity-background](http://aeternity.com/aepp-components/#aeidentitybackground)
    - [identity-light](http://aeternity.com/aepp-components/#aeidentitylight)
    - [label](http://aeternity.com/aepp-components/#aelabel)
    - [link](http://aeternity.com/aepp-components/#aelink)
    - [loader](http://aeternity.com/aepp-components/#aeloader)
    - [main](http://aeternity.com/aepp-components/#aemain)
    - [modal](http://aeternity.com/aepp-components/#aemodal)
    - [modal-light](http://aeternity.com/aepp-components/#aemodallight)
    - [overlay](http://aeternity.com/aepp-components/#aeoverlay)
    - [switch](http://aeternity.com/aepp-components/#aeswitch)
    - [textarea](http://aeternity.com/aepp-components/#aetextarea)

## Online Testnet Tools
- [**Smart Contract Editor**](https://testnet.contracts.aepps.com/)
- [**Blockchain Explorer**](https://testnet.explorer.aepps.com/#/)
- **Faucet**
    - [Token Faucet](https://testnet.faucet.aepps.com/)
    - [aepp Faucet](https://github.com/aeternity/aepp-faucet/blob/master/README.md)

## [Sophia Support](https://github.com/mradkov/vscode-sophia/blob/master/README.md)

# æternity FEATURES
- [**Oracles**](https://github.com/aeternity/protocol/blob/master/oracles/oracles.md)
    - [Life Cycle](https://github.com/aeternity/protocol/blob/master/oracles/oracle_life_cycle.md)
    - [State Tree](https://github.com/aeternity/protocol/blob/master/oracles/oracle_state_tree.md)
    - [Transactions](https://github.com/aeternity/protocol/blob/master/oracles/oracle_state_tree.md)
- [**AENS Naming System**](https://github.com/aeternity/protocol/blob/master/AENS.md)
- [**Smart Contracts**](https://github.com/aeternity/protocol/blob/master/contracts/contracts.md)
    - [Aeternity VM](https://github.com/aeternity/protocol/blob/master/contracts/contract_vms.md) 
    - [Sophia Compiler](https://github.com/aeternity/aesophia)
    - [State Tree](https://github.com/aeternity/protocol/blob/master/contracts/contract_state_tree.md)
    - [Transactions](https://github.com/aeternity/protocol/blob/master/contracts/contract_transactions.md)
- [**State Channels**](https://github.com/aeternity/protocol/blob/master/channels/README.md)
- **Consensus**
    - [Cuckoo Cycle (PoW)](https://github.com/aeternity/cuckoo/blob/master/README.md)
        - [GPU Solver](https://github.com/aeternity/cuckoo/blob/master/GPU.md)
    - Bitcoin-NG
- [**Serializazion Formats**](https://github.com/aeternity/protocol/blob/master/serializations.md)

# [æternity Expansions](https://github.com/aeternity/AEXs/blob/master/README.md)
- [**Issue Template**](https://github.com/aeternity/AEXs/blob/master/ISSUE_TEMPLATE.md)
- [**Pull Request Template**](https://github.com/aeternity/AEXs/blob/master/PULL_REQUEST_TEMPLATE.md)
- **AEXS**
    - [AEX1](https://github.com/aeternity/AEXs/blob/master/AEXS/aex-1.md)
    - [AEX2](https://github.com/aeternity/AEXs/blob/master/AEXS/aex-2.md)

# SOPHIA LIBRARY
- **Libraries**
    - [Sophia DateTime](https://github.com/aeternity/aepp-sophia-examples/blob/master/libraries/DateTime/README.md)
    - [Sophia List](https://github.com/aeternity/aepp-sophia-examples/blob/master/libraries/List/README.md)
    - [Sophia Ownable](https://github.com/aeternity/aepp-sophia-examples/blob/master/libraries/Ownable/README.md)
    - [Sophia Fungible Tokens](https://github.com/aeternity/aepp-sophia-examples/blob/master/libraries/FuungibleTokens/README.md)
    - [Sophia NonFungible Tokens](https://github.com/aeternity/aepp-sophia-examples/blob/master/libraries/NonFungibleTokens/README.md)
    - [Sophia Muiltisignature Wallet](https://github.com/aeternity/aepp-sophia-examples/blob/master/libraries/MuiltisignatureWallet/README.md)
    - [Sophia Base Converter](https://github.com/aeternity/aepp-sophia-examples/blob/master/libraries/SophiaBaseConverter/README.md)

# SOPHIA BY EXAMPLE
- **Examples**
    - [Sophia SmartShop](https://github.com/aeternity/aepp-sophia-examples/blob/master/examples/SmartShop/README.md)
    - [Sophia Crypto Hamster](https://github.com/aeternity/aepp-sophia-examples/blob/master/examples/CryptoHamster/README.md)
    - [Sophia SmartRealEstate](https://github.com/aeternity/aepp-sophia-examples/blob/master/examples/SmartRealEstate/README.md)
    - [Sophia Data Provider](https://github.com/aeternity/aepp-sophia-examples/blob/master/examples/DataProvider/README.md)
    - [Sophia Spend-to-Many](https://github.com/aeternity/aepp-sophia-examples/blob/master/examples/SpendToMany/README.md)
    - [Sophia TicTacToe](https://github.com/aeternity/aepp-sophia-examples/blob/master/examples/TicTacToe/README.md)
 
# TUTORIALS
## Development
- **Running Node**
    - [Ubuntu 18.04](https://github.com/aeternity/tutorials/blob/master/run-node-on-ubuntu-1804.md)
    - [macOS Mojave](https://github.com/aeternity/tutorials/blob/master/run-node-on-macos-mojave.md)