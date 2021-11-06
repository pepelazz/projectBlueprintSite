package main

import (
	"github.com/otiai10/copy"
	"github.com/pepelazz/nla_framework"
	t "github.com/pepelazz/nla_framework/types"
	"github.com/pepelazz/projectBlueprintSite/projectTemplate/docs/legalEntity"
	"github.com/pepelazz/projectBlueprintSite/projectTemplate/utils"
	"os"
)

func main() {
	nla_framework.Start(getProject(), nil)
}

func getProject() t.ProjectType {
	p := &t.ProjectType{
		Name: "CompanyName",
	}
	p.Config.Vue.QuasarVersion = 2

	p.Docs = []t.DocType{
		legalEntity.GetDoc(p),
	}
	// названием базы маленькими буквами, без пробелов
	p.Config.Postgres = t.PostrgesConfig{DbName: "nla_db", Port: 5657, Password: "xvzDV4curLidx8IWZJ6czDHQ1qa7wjfL", TimeZone: "Asia/Novosibirsk"}
	p.Config.WebServer = t.WebServerConfig{Port: 3101, Url: "https://example.ru", Path: "/home/deploy/projectName", Ip: "85.210.890.567", Username: "root"}
	// TODO: надо прописать настройки почтового сервера для отправки email
	p.Config.Email = t.EmailConfig{Sender: "info@mail.ru", Password: "password", Host: "smtp.mail.ru", Port: 465, SenderName: "CompanyName"}
	p.Config.Logo = "https://cdn.pixabay.com/photo/2017/05/05/00/15/kokopelli-2285538_960_720.png"
	p.Config.Docker.AfterCopy = []string{
		"COPY ./src/webServer/site /webServer/site",
		"COPY ./src/webServer/templates /webServer/templates",
	}
	// формируем routes для Vue
	p.FillVueBaseRoutes()
	p.Vue.UiAppName = "CompanyName"

	// боковое меню для Vue
	p.Vue.Menu = []t.VueMenu{
		//{DocName: "client_order"},
		{Url: "users", Text: "Пользователи", Icon: "image/users.svg", Roles: []string{utils.RoleAdmin}},
		{DocName: "legal_entity"},
		{Text: "Справочники", Icon: "image/directory.png", IsFolder: true, LinkList: []t.VueMenu{{DocName: "legal_entity"}}},
	}
	p.FillSideMenu()

	p.OverridePathForTemplates = map[string]string{
		"/webServer/main.go": "./tmpl/webServer/main.go",
		"/webClient/quasar.conf.js": "./tmpl/webClient/quasar.conf.js",
	}

	// копируем файлы проекта (которые не шаблоны)
	if _, err := os.Stat("./sourceFiles"); !os.IsNotExist(err) {
		err := copy.Copy("./sourceFiles", "../")
		utils.CheckErr(err, "Copy sourceFiles")
	}

	return *p
}
