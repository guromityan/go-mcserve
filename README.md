# go-mcserve
Mock Server



## Usage

### Directory structure

```bash
$ tree response/
response
├── response.json
└── v1
    └── response.json
```



### Start the server

```bash
$ ./mcserve -r response -t json -p 8080
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.16
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on 127.0.0.1:8080
```



### HTTP Request

```bash
$ curl http://localhost:8080/
{
  "test": "test1"
}
```



## Help

```bsah
$ mcserve --help
usage: mcserve --response=RESPONSE [<flags>]

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -r, --response=RESPONSE  The root directory that stores the files used for the response.
  -t, --type=TYPE          Response object type.: 'json' or 'html'
  -p, --port=8080          Port number for starting the server.
      --version            Show application version.

```