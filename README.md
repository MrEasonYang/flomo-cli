# flomo-cli
A Golang based command line interface of [flomo](https://flomoapp.com/).

# Installation
## Build from source
```shell
git clone git@github.com:MrEasonYang/flomo-cli.git
cd flomo-cli
go build
```
## Use homebrew
```shell
brew tap MrEasonYang/taps
brew install flomo
```

# Usage
## Configure
Visit [Flomo Settings](https://flomoapp.com/mine?source=incoming_webhook) to obtain the API then put it into flomo-cli:
```shell
flomo set api ${Flomo API}
```
A configuration file named `.flomo-cli.config` will be save in your home directory with the default permission of 0600.
## Save memo
Just type a memo with the save command into the flomo command:
```shell
flomo save ${Your memo content}
```
And that's it!

# License
[MIT](https://github.com/MrEasonYang/flomo-cli/blob/main/LICENSE)