# leetcode
My LeetCode journey. Fetching by neovim.

## Installation
**neovim**
```bash
# make sure your neovim version is greate equal than 0.8
# assume you use macOS
brew install neovim
# fetech my configuration
cd ~/.config && git clone git@github.com:yuchanns/nvim.git
# firstime open vim and do initialization
nvim
```
**python stuffs**
```bash
# assume you use macOS
brew install python3
# locate the real python3 path otherwise you will fail below
brew info python3
# in my case python3 installed into /opt/homebrew/opt/python@3.10/libexec/bin
/opt/homebrew/opt/python@3.10/libexec/bin/python -m pip install pynvim \
    keyring browser_cookie3 --user
```
## Startup
First, you need to login your LeetCode account using Chrome.

Then choose a folder (e.g. This repository where it place), open with `nvim`.

Execute `:LeetCodeList`, pick up a problem and start to sovle it!

## Quick Start

- `:LeetCodeList`: browse the problems.
- `:LeetCodeTest`: run the code with the default test case.
- `:LeetCodeSubmit`: submit the code.
- `:LeetCodeSignIn`: manually sign in.

