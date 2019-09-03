package main

import (
	"github.com/kataras/iris"
	"nimohunter.com/service"
	"strconv"
	"strings"
)

func main() {
	app := iris.New()
	// 从 "./views" 目录加载HTML模板
	// 模板解析 html 后缀文件
	// 此方式使用 `html/template` 标准包 (Iris 的模板引擎)
	app.RegisterView(iris.HTML("./views", ".html"))
	elasticService := service.NewElasticService()

	// 方法：GET
	// 路径：http://localhost:8082
	app.Get("/", func(ctx iris.Context) {
		// 映射 HTML 模板文件路径 ./views/hello.html
		ctx.View("index.html")
	})

	app.Post("/search", func(context iris.Context) {
		beauty := strings.TrimSpace(context.PostValue("beauty"))
		age_str := strings.TrimSpace(context.PostValue("age"))
		gender := context.PostValue("gender")
		age, e := strconv.Atoi(age_str)
		if e != nil {
			age = 100
		}
		// Search Service
		search_result := elasticService.AdvanceSearch(age, gender, beauty)
		context.ViewData("Hits", search_result.Hits)
		context.ViewData("Start", search_result.Start)
		context.ViewData("Items", search_result.Items)
		context.View("template.html")
	})

	// 绑定端口并启动服务.
	app.Run(iris.Addr(":8082"))
}
