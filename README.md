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
  - Create block lists and choose which you want to use
- Run `tsurugi` and you're free of distractions
- Interface tailored for better experience and usability
- Prebuilt binaries for Linux and Darwin platforms

> [!IMPORTANT]
> Work in progress...

## ⚒️ Usage

### Installation

Go to the [latest release](https://github.com/kauefraga/tsurugi/releases) page and download the prebuilt binary for your platform.

### Blocking websites

Create a plain text file with your blocks

```txt
bsky.app
youtube.com
www.youtube.com
instagram.com
twitch.tv
www.twitch.tv
```

And run

```sh
sudo ./tsurugi blocklist.txt
```

If you want to unblock those websites, just run

```sh
sudo ./tsurugi stop
```

Alternatively, you can [create a gist](https://gist.github.com/) and pass url (raw)

```sh
sudo ./tsurugi https://gist.githubusercontent.com/kauefraga/433dc03487e71f8df477e5584d2d2c23/raw/c7c9310ed790e481dbd9b5cd79b96a7ac3f711f9/blocklist.txt
```

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

### Why is it not working

The key idea of tsurugi is to use the `/etc/hosts` file to block websites.

- `/etc/hosts` is being bypassed by your browser or your browser extensions
- Your system is using IPv6
- You are using a VPN
- Browser cache or DNS cache is bypassing `/etc/hosts` new entries

The best solution would be to block websites directly in the network with a proxy server. However, in the first version I want to **keep it as simple as possible**.

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
