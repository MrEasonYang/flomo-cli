# flomo-cli
A Golang based command line interface of [flomo](https://flomoapp.com/).
[中文说明](https://easonyang.com/2021/07/17/flomo-cli-cn-readme/)

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
Below platforms are supported:
- Apple Intel AMD64
- Apple Silicon
- Linux AMD64

## Install from release
If you are using Windows or dislike Homebrew, you can download the appropriate program from [Release](https://github.com/MrEasonYang/flomo-cli/releases) and setup the environment by yourself. 

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

## Setup alias
In case of unexpected saving actions, flomo-cli force to use save keyword to send memo to flomo currently. You can setup an alias in your .zshrc/.bashrc etc. to simplify the command:
```shell
alias flomo="flomo save"
```

# Contribution
Contributions are welcome, just remember to lint your code.

# License
[MIT](https://github.com/MrEasonYang/flomo-cli/blob/main/LICENSE)