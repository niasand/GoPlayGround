package main

import (
	"mime/multipart"
	"strings"

	"github.com/kataras/iris"
)

const maxSize = 5 << 20

func main() {
	app := iris.Default()
	// http  http://127.0.0.1:8099/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	// http http://127.0.0.1:8099/user/yang
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})
	// http -f POST http://127.0.0.1:8099/user/This/world
	app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.Writef(message)
	})

	// http  http://127.0.0.1:8099/Welcome\?firstname\=ss\&lastname\=122
	app.Get("/Welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		lastname := ctx.URLParam("lastname")
		ctx.Writef("Hello %s %s", firstname, lastname)
	})
	// http   -f POST http://127.0.0.1:8099/form_post  message="haha i am yang" nick="nick name is jerry"
	app.Post("form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick := ctx.FormValueDefault("nick", "anonymous")

		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	// http -f POST http://127.0.0.1:8099/post\?id\=121 message=2 name=3
	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		app.Logger().Infof("id: %s;page:%s;name:%s;message:%s", id, page, name, message)
	})

	app.Post("/upload", iris.LimitRequestBodySize(maxSize), func(ctx iris.Context) {
		ctx.UploadFormFiles("/data", beforesave)
	})

	app.Run(iris.Addr(":8099"))
}

func beforesave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)
	file.Filename = ip + "-" + file.Filename
}
