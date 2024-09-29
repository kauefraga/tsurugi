# `tsurugi`

[![GitHub latest version](https://img.shields.io/github/v/tag/kauefraga/tsurugi?label=latest+version)](https://github.com/kauefraga/tsurugi/releases)
[![GitHub's license](https://img.shields.io/github/license/kauefraga/tsurugi)](https://github.com/kauefraga/tsurugi/blob/main/LICENSE)
[![GitHub last commit](https://img.shields.io/github/last-commit/kauefraga/tsurugi/main)](https://github.com/kauefraga/tsurugi)

> Cut out distractions with tsurugi! Easily block distracting websites.

Ever wanted to block websites that are making you idle? With tsurugi you can, right from your terminal!

## 👀 What to expect

- Configure easily
  - Set blocks in a [`tsurugi.toml`](https://toml.io/en/) using the CLI
  - Support plain text blocks locally and remotely
- Run `tsurugi` and you're free of distractions
- Interface tailored for better experience and usability
- Prebuilt binaries for Linux and Darwin platforms

> [!IMPORTANT]
> Work in progress...

## 💖 Contributing

Feel free to contribute [opening an issue](https://github.com/kauefraga/tsurugi/issues/new) to report a bug or suggesting a CLI change, an improvement or a feature.

### How to contribute

1. Fork this project
2. Clone your fork on your machine
3. Setup the [dev environment](#how-to-setup-dev-environment)
4. Make your changes and commit them following [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
5. Run `git push` to sync the changes
6. Open a pull request specifying what you did

## 🔎 More about the project

### Why "tsurugi"

> Tsurugi is a Japanese term for a straight, double-edged sword used in antiquity.

I was watching [Tsue to **Tsurugi** no Wistoria](https://myanimelist.net/anime/58059/Tsue_to_Tsurugi_no_Wistoria) (Wistoria: Wand and **Sword**) and thinking about this project to block distractions, then suddenly I thought of "cutting" the distractions...

### Why is it not available for Windows

The key idea of tsurugi is to use the `/etc/hosts` file to block websites.

This file in question exists in most unix-like operating systems such as Linux distributions, MacOS, BSD and others. Windows has this file too, but it's inside the System32 folder, that's a problem because the user would need to execute tsurugi with admin privileges in order to achieve the same result of unix-like systems.

The cross platform solution would be to block websites directly in the network with a proxy server. However, in the first version I want to **keep it as simple as possible**.

### How to setup dev environment

- Have [Go](https://go.dev/) installed (Preferably [1.22.5](go.mod))

Install the dependencies

```sh
go mod download
```

Or just run the project

```sh
go run cmd/main.go
```

### How to build

With [Go](https://go.dev/) installed, building tsurugi should be as easy as running the following command

```sh
go build cmd/main.go -o tsurugi
```

However, running the command below should generate a more lightweight binary

```sh
CGO_ENABLED=0 go build -ldflags='-w -s' cmd/main.go -o tsurugi
```

## 📝 License

This project is licensed under the MIT License - See the [LICENSE](https://github.com/kauefraga/tsurugi/blob/main/LICENSE) for more information.

---

Se você gostou do projeto e ele te ajudou, considere [me apoiar um café](https://www.pixme.bio/kauefraga) ☕💙🇧🇷
