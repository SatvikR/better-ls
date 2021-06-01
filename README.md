# Better ls

<div align="center">
  
![Screenshot from 2021-05-31 18-24-11](https://user-images.githubusercontent.com/49799352/120253891-7106d980-c23d-11eb-905d-341ca19c5409.png)

A better version of ls

</div>

> Inspired by [ls](https://en.wikipedia.org/wiki/Ls) and [exa](https://github.com/ogham/exa)

## Install

```sh
go get github.com/SatvikR/better-ls
```

## Build from source

```sh
# Clone repo
git clone https://github.com/SatvikR/better-ls
cd better-ls

# Compile
go build

# Run (make sure add .exe on windows)
./better-ls
```

## Usage

`better-ls [OPTIONS...] [DIRECTORY]`

| Flag                | Usage                                 |
| ------------------- | ------------------------------------- |
| `-group-dirs-first` | output all the directories at the top |
| `-human-readable`   | list sizes in a human readable format |

## License

Copyright (c) 2021, Satvik Reddy

Licensed under the
[BSD 2-Clause License](https://github.com/SatvikR/better-ls/blob/main/LICENSE)
