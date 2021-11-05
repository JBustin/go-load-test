# go-load-test

Load Test CLI - Curl or Browser like

## Install

Install [Golang](https://golang.org/doc/install).

Then,

```sh
git clone https://github.com/JBustin/go-load-test.git
cd go-load-test
make install
# move executable "go-load-test" to a bin directory
# or copy the path of the binary inside $PATH
```

## Usage

```
./go-load-test -f test.json
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
