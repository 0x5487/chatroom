package main

import "github.com/jasonsoft/napnap"

func main() {
	nap := napnap.New()
	nap.SetRender("./templates")

	router := napnap.NewRouter()
	// display client page
	router.Get("/chat", func(c *napnap.Context) {
		c.Render(200, "chat.html", nil)
	})
	// websocket
	router.Get("/ws", napnap.WrapHandler(websocketEndpoint()))
	nap.Use(router)

	httpEngine := napnap.NewHttpEngine(":8080") //run on port 8080
	nap.Run(httpEngine)

}
