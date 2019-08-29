package main

import (
	"github.com/kataras/iris"
	"nimohunter.com/model"
)

func main() {
	app := iris.New()
	// 从 "./views" 目录加载HTML模板
	// 模板解析 html 后缀文件
	// 此方式使用 `html/template` 标准包 (Iris 的模板引擎)
	app.RegisterView(iris.HTML("./views", ".html"))

	// 方法：GET
	// 路径：http://localhost:8080
	app.Get("/welcome", func(ctx iris.Context) {
		// {{.message}} 和 "Hello world!" 字符串变量绑定
		ctx.ViewData("message", "Hello world..")
		// 映射 HTML 模板文件路径 ./views/hello.html
		ctx.View("index.html")
	})

	search_result := createFakeData()
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Hits", search_result.Hits)
		ctx.ViewData("Start", search_result.Start)
		ctx.ViewData("Items", search_result.Items)
		ctx.View("template.html")
	})

	// 绑定端口并启动服务.
	app.Run(iris.Addr(":8082"))
}

func createFakeData() model.SearchResult {
	search_result := model.SearchResult{
		Hits:  5,
		Start: 0,
	}

	test_item := model.Item{
		Name:       "林YY",
		Marriage:   "未婚",
		Age:        "26岁",
		Xingzuo:    "魔羯座(12.22-01.19)",
		Height:     "165cm",
		Weight:     "50kg",
		Income:     "月收入:5-8千",
		Occupation: "职业技术教师",
		Education:  "高中及以下",
	}

	for i := 0; i < 5; i++ {
		search_result.Items = append(search_result.Items, test_item)
	}
	return search_result
}
