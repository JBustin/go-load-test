# go-load-test

Load Test CLI - Curl or Browser like

<img width="795" alt="Capture d’écran 2021-11-05 à 14 53 02" src="https://user-images.githubusercontent.com/2632709/140521293-228157a3-03ee-406b-a532-6dca71fcb185.png">

## Find the binaries for your OS

In `build/` directory, download the binary you need for your OS.

## Or install with Golang

Required: [Golang](https://golang.org/doc/install).

Then,

```sh
git clone https://github.com/JBustin/go-load-test.git
cd go-load-test
make install
# move executable "gload" to a bin directory
# or copy the path of the binary inside $PATH
```

## Usage

```
./gload -f test.json
```

Make your own json test file.

## Json

- `isBrowser` boolean (default: false)
- `isSerie` boolean (default: false)
- `hits` int (default: 100)
- `waitMs` int (default 1000)
- `concurrency` int (default 50)
- `logLevel` string (default "error", other values "info", "debug")
- `timeoutMs` int (default 20000)
- `scrap` boolean (default false)
- `urls` array (no default value)
- `headers` map (default empty)
