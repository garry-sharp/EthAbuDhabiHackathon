# Price Aggregastor Tool (ETH Abu Dhabi Hackathon)

A backend tool that interacts with dexes on the XDC network. It pulls a list of tokens from https://raw.githubusercontent.com/XSwapProtocol/xdc-token-list/master/testnet.tokenlist.json.

Gets the best buy/sell prices and pushed them to a proxy contract so when you perform a trade it automatically performs the best trade.

Total time to build is around 13 hours.

A video demo is also available here https://youtu.be/Eb7X0eis6zo

Garry Sharp and Ben Liams

## Not in scope

- The actual trade making
- A frontend to make the trade
- Proper error handling

## What I'd do differently (if I had time)

- Not have a hardcoded private key (this should be shipped to a signatory). The key hardcoded is the first derived address from the Test Test Test ... hd wallet mnemonic. It also has some test funds preloaded.
- Not use a map of maps. I'd use a multi key map or something similar like redis or a DB
- Write tests for golang
- Write hardhat tests for the code
- Better validation using something like cobra for CLI arg parsing.
- Dockerise the application.
- Proper error handling (lots of errors are just ignored into a '\_' variable in the go code.)

## Running the code.

The simplest way would be to do `go run main/main.go --deploy --batch` from the `pkg` folder.
