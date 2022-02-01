### Goでプロジェクト作成

```
cd ./backend/go
go mod init github.com/gushikem01/cryptcurrency_autotrade
```

### Goのバージョンアップ

```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
go version
```

### sqlcとgolang-migrateでPostgreSQLを扱う

https://zenn.dev/tchssk/articles/a701d3ce5f9b6b

### anyenv インストール

```
brew install anyenv
export PATH="$HOME/.anyenv/bin:$PATH"
echo eval "$(anyenv init -)" >> ~/.bash_profile
anyenv install --init
```

* run

```
go run ./server/cmd/api/main.go
```

* go env

go env GOROOT
go env GOPATH
go env GOBIN
go env PATH

* export

```
export GOROOT=$HOME/.anyenv/envs/goenv/versions/1.16.0/
export GOPATH=$HOME/go/1.16.0
export PATH=${GOROOT}/bin:$PATH
export GOBIN=
```

* brew install

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

* golangci-lint install

```
homebrew install golangci-lint
```
