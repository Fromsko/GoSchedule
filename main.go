package main

import (
	"Toch/core"
	"Toch/define"
	"Toch/utils"
	"Toch/web"

	"github.com/go-rod/rod"
)

var (
	ServerPort = ":2000" // 服务端口
	ClassName  = "教育学x班" // 班级名称
)

func main() {
	// 初次运行
	if utils.FristRun() {
		Task()
	} else {
		web.StartServer(ServerPort)
	}

	// 定时任务
	go web.AutoTask("0 0 */12 * * ?", Task)

	// 阻塞主程序
	select {}
}

// 任务
func Task() {
	Browser := core.InitWeb(
		define.LoginPageURL,
		"",
	)
	Img := core.InitImg()
	// 登录
	if loginStatus := Browser.Login(); loginStatus {
		Browser.NextPage()
		Browser.Extract(func(downloadPage *rod.Page, selectName string) {
			// 初始化课表
			cname, doc := core.InitCname(
				core.ReadHTML(
					Browser.Html(downloadPage),
				),
			)
			// 解析数据
			cname.Resolve(doc)
			// 写入文件
			result := cname.WriteFile(
				selectName,
				ClassName,
			)
			// 存储图片
			core.SaveImg(Img.Create(result))
		})
		// 退出
		Browser.Logout()
		web.RestartServer(ServerPort)
	}
}
