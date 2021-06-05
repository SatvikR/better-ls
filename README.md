# Better ls

<div align="center">
  
<img width="1078" alt="better-ls screenshot" src="https://user-images.githubusercontent.com/49799352/120881575-f7495580-c586-11eb-8238-4c1d55ed8106.png">

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
| `-h`                | View help message                     |

## License

Copyright (c) 2021, Satvik Reddy

Licensed under the
[BSD 2-Clause License](https://github.com/SatvikR/better-ls/blob/main/LICENSE)
