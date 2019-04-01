package main // 代码包声明语句。
import "fmt" //系统包用来输出的
import "github.com/devfeel/dotweb"
import "adv_show"

func startServer() error {
	app := dotweb.New()
	app.HttpServer.GET("/index", func(ctx dotweb.Context) error {
		return ctx.WriteString("welcome to my first web!")
	})
	//路由注册部分，后期可以单独的拿出去。
	//@广告展示部分
	app.HttpServer.GET("/show", adv_show.ShowAdv)
	app.HttpServer.GET("/redis", openRedis)

	//begin server
	err := app.StartServer(8000)
	return err
}
func openRedis(ctx dotweb.Context) error {
	app := dotweb.New()
	redis := app.cache.NewRedisCache("localhost:7379")
	if redis != nil {
		ctx.WriteString("连接redis成功")
	}
	return nil
}
func main() {
	// 打印函数调用语句。用于打印输出信息。
	startServer()
	fmt.Println("helloworld")
}
