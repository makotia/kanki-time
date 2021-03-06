package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
	"github.com/makotia/kanki-time/api/config"
	"github.com/makotia/kanki-time/api/util"
	"github.com/valyala/fasttemplate"
)

func main() {
	var addr = flag.String("addr", ":8080", "TCP address to listen to")
	flag.Parse()
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
	}))

	api := e.Group("/api")
	{
		api.GET("/:id", getHandler)
		api.POST("", createHandler)
		api.OPTIONS("*", optionsHandler)
		api.Static("/media", path.Join(config.GetConfig().Server.StaticDir))
		api.GET("/image", getImage)
	}
	e.Logger.Fatal(e.Start(*addr))
}

func getHandler(c echo.Context) (err error) {
	if _, err = os.Stat(path.Join(config.GetConfig().Server.StaticDir, c.QueryParam("id")+".png")); err != nil {
		if os.IsNotExist(err) {
			return c.NoContent(http.StatusNotFound)
		}
	}
	return c.NoContent(http.StatusOK)
}

func getImage(c echo.Context) (err error) {
	var (
		buf bytes.Buffer
		img *image.RGBA
	)
	if img, err = util.GenImage(strings.ReplaceAll(c.QueryParam("Text"), ",", "\n"), c.QueryParam("Type")); err != nil {
		return c.JSON(http.StatusInternalServerError, toMap("", err))
	}
	if err = png.Encode(&buf, img); err != nil {
		return c.JSON(http.StatusInternalServerError, toMap("", err))
	}
	return c.Blob(http.StatusOK, "image/png", buf.Bytes())
}

func optionsHandler(c echo.Context) (err error) {
	return c.NoContent(http.StatusOK)
}

func createHandler(c echo.Context) (err error) {
	type reqJSON struct {
		Text string `json:"Text"`
		Type string `json:"Type"`
	}

	var (
		id      string
		reqBody reqJSON
		buf     bytes.Buffer
		img     *image.RGBA
	)

	io.Copy(&buf, c.Request().Body)

	if err = json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		return c.JSON(http.StatusInternalServerError, toMap(id, err))
	}

	if img, err = util.GenImage(reqBody.Text, reqBody.Type); err != nil {
		return c.JSON(http.StatusInternalServerError, toMap(id, err))
	}
	if id, err = util.SaveImage(img); err != nil {
		return c.JSON(http.StatusInternalServerError, toMap(id, err))
	}
	return c.JSON(http.StatusOK, toMap(id, err))
}

func toMap(id string, err error) map[string]string {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return map[string]string{
		"id":    id,
		"error": errMsg,
	}
}

type loggerConfig struct {
	Format           string `yaml:"format"`
	CustomTimeFormat string `yaml:"custom_time_format"`
	Output           io.Writer

	template *fasttemplate.Template
	colorer  *color.Color
	pool     *sync.Pool
}

func logger() echo.MiddlewareFunc {
	config := loggerConfig{
		Format:           "${time} | ${method} | ${status} | ${path}\n",
		CustomTimeFormat: "15:04:05",
		colorer:          color.New(),
		Output:           os.Stdout,
	}
	config.template = fasttemplate.New(config.Format, "${", "}")
	config.colorer = color.New()
	config.colorer.SetOutput(config.Output)
	config.pool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 256))
		},
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			buf := config.pool.Get().(*bytes.Buffer)
			buf.Reset()
			defer config.pool.Put(buf)

			if _, err = config.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
				switch tag {
				case "time":
					return buf.WriteString(time.Now().Format(config.CustomTimeFormat))
				case "method":
					return buf.WriteString(fmt.Sprintf("%-7s", req.Method))
				case "path":
					p := req.URL.Path
					if p == "" {
						p = "/"
					}
					return buf.WriteString(p)
				case "status":
					n := res.Status
					s := config.colorer.Green(n)
					switch {
					case n >= 500:
						s = config.colorer.Red(n)
					case n >= 400:
						s = config.colorer.Yellow(n)
					case n >= 300:
						s = config.colorer.Cyan(n)
					}
					return buf.WriteString(s)
				case "error":
					if err != nil {
						b, _ := json.Marshal(err.Error())
						b = b[1 : len(b)-1]
						return buf.Write(b)
					}
				case "latency_human":
					return buf.WriteString(stop.Sub(start).String())
				}
				return 0, nil
			}); err != nil {
				return
			}

			if config.Output == nil {
				_, err = c.Logger().Output().Write(buf.Bytes())
				return
			}
			_, err = config.Output.Write(buf.Bytes())
			return
		}
	}
}
