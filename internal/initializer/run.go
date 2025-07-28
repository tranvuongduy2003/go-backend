package initializer

func Run() {
	// Run all initializers here
	// This function will call all the initialization functions in the correct order,
	// ensuring that the application is set up properly before it starts handling requests.
	LoadConfig()
	InitLogger()
	InitMySQL()
	InitRedis()
	r := InitRouter()

	r.Run(":8080") // Start the server on port 8080
}