package tcp

func (e *Config) SetupRoutes() {
	api := e.Framework.Conn.Group("/api")

	character := api.Group("/character")
	character.GET(":id", nil)
	character.PUT(":id", nil)
	character.DELETE(":id", nil)
	character.POST("", CreateCharacter)
}
