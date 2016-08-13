package main
import "github.com/kataras/iris"

func main() {
    iris.Get("/hi_json", func(c *iris.Context) {
        c.JSON(iris.StatusOK, iris.Map{
            "Name":  "Iris",
            "Born":  "13 March 2016",
            "Stars": 4270,
        })
    })
    iris.Listen(":8081")
}
