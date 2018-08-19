package main

import (
	"github.com/jasonsoft/chatroom/pkg/chat"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/log/handlers/console"
	"github.com/jasonsoft/napnap"
)

func main() {
	log.SetAppID("chat") // unique id for the app

	clog := console.New()
	log.RegisterHandler(clog, log.AllLevels...)

	chat.Initialize()

	nap := napnap.New()
	nap.SetRender("./templates")

	router := napnap.NewRouter()
	nap.Use(napnap.NewHealth())

	// display client page
	router.Get("/chat", func(c *napnap.Context) {
		c.Render(200, "chat.html", nil)
	})
	nap.Use(router)

	nap.Use(chat.NewChatRouter())
	nap.Use(napnap.NewNotfoundMiddleware())

	httpEngine := napnap.NewHttpEngine(":8080") //run on port 8080
	nap.Run(httpEngine)

}
