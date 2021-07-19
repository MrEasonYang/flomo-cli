# flomo-cli
A Golang based command line interface of [flomo](https://flomoapp.com/).

[中文说明](https://easonyang.com/2021/07/17/flomo-cli-cn-readme/)

# Features
- Type and save to flomo using command line.
- Editor mode supports, able to use vim/neovim/emacs to compose the memo.
- Shell pipes supports.

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
### Basic
Just type a memo with the save command into the flomo command:
```shell
flomo save ${Your memo content}
```
And that's it!
### Using shell pipes
Just like usual command line programs, you can use command like `cat` with the shell pipe to save an entire file into flomo:
```shell
cat memo.txt | flomo
```
### Editor mode
Instead of input the memo with the command, you can use editors to compose the memo:
```shell
# Open vim to compose the memo.
flomo vim 

# Open neovim to compose the memo.
flomo nvim 

# Open emacs to compose the memo.
flomo emacs
```
Currently, the flomo-cli only support `vim/neovim/emacs`, the other inputs will trigger an error due to the security consideration.

## Clear the temporary files
How is it possible to use the editors to compose? Well, the answer is a bit tricky. When the editor command is received, the flomo-cli will invoke the specified editor to open a temporary file under `~/.flomo-tmp` and wait until the editor's job is done. Once you quit the editor, the flomo-cli will try to load the content of the temporary file so that the input could be sent to the flomo, the file mentioned before will be deleted then.
However, if the editor is invoked concurrently or the flomo-cli instance is terminated unexpectedly, the temporary file will not be deleted and occupy the disk space. To solve this issue, just run the command below to clear all the useless temporary files:
```shell
flomo clear
```

## Setup alias
In case of unexpected saving actions, flomo-cli force to use save keyword to send memo to flomo currently. You can setup an alias in your .zshrc/.bashrc etc. to simplify the command:
```shell
alias flomo="flomo save"
```

# Contribution
Contributions are welcome, just remember to lint your code.

# License
[MIT](https://github.com/MrEasonYang/flomo-cli/blob/main/LICENSE)