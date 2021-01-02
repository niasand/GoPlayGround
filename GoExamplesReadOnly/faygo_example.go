package GoExamplesReadOnly
import (
	"time"
    "github.com/henrylee2cn/faygo"
)

type Index struct {
    Id    int    `param:""`
    Title    string    `param:""`
    Paragraph    []string    `param:""`
    Cookie    string    `param:""`   
}

func (i *Index) Serve(ctx *faygo.Context) error {
    if ctx.CookieParam("faygoID") == ""{
        ctx.SetCookie("faygoID", time.Now().String())
    }
    return ctx.JSON(200, i)
    
}

func mainfalse()  {
    app := faygo.New("myapp", "0.1")
    app.GET("/index/:id", new(Index))
    faygo.Run()
}