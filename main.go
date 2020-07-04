package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/golang/go/src/path"
	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

const version = "1.0.0"

var (
	resRoot = kingpin.Flag("response", "The root directory that stores the files used for the response.").Required().Short('r').ExistingDir()
	resType = kingpin.Flag("type", "Response object type.: 'json' or 'html'").Short('t').String()
	port    = kingpin.Flag("port", "Port number for starting the server.").Short('p').Default("8080").Int()
)

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	e := echo.New()
	e.GET("/*", app)
	addr := fmt.Sprintf("localhost:%v", *port)
	e.Logger.Fatal(e.Start(addr))
}

func app(c echo.Context) error {
	resFile := ""

	req := c.Request()
	if filepath.Ext(req.URL.Path) == "."+*resType {
		resFile = path.Join(*resRoot, req.URL.Path)
	} else {
		fileList, _ := ioutil.ReadDir(*resRoot + req.URL.Path)
		if len(fileList) == 0 {
			resFile = path.Join(*resRoot, req.URL.Path)
		}

		for _, f := range fileList {
			if filepath.Ext(f.Name()) == "."+*resType {
				resFile = path.Join(*resRoot, req.URL.Path, f.Name())
				break
			}
		}
	}

	// HTTP status 404: File not found error
	content, err := ioutil.ReadFile(resFile)
	if err != nil {
		msg := fmt.Sprintf("file not found: %v", resFile)
		return c.String(http.StatusNotFound, msg)
	}

	// HTTP status 200: generate response object
	switch *resType {
	case "json":
		return c.JSONBlob(http.StatusOK, content)
	case "html":
		return c.HTMLBlob(http.StatusOK, content)
	default:
		return c.String(http.StatusOK, string(content))
	}
}
