package main

import "github.com/jasonsoft/napnap"

func main() {
	nap := napnap.New()
	nap.SetRender("views/*")

	router := napnap.NewRouter()
	// display client page
	router.Get("/chat", func(c *napnap.Context) {
		c.Render(200, "chat.html", nil)
	})
	// websocket
	router.Get("/ws", napnap.WrapHandler(websocketEndpoint()))
	nap.Use(router)

	nap.Run(":8080")

}
