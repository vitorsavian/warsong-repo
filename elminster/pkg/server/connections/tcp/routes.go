package tcp

func (e *Config) SetupRoutes() {
	api := e.Framework.Conn.Group("/api")

	character := api.Group("/character")
	character.GET("/:id", GetCharacter)
	character.PUT("/:id", UpdateCharacter)
	character.DELETE("/:id", DeleteCharacter)
	character.POST("", CreateCharacter)
}
