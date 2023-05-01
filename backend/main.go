package main

import (
	"joosum-backend/pkg/configs"
	"joosum-backend/pkg/routes"

	"joosum-backend/pkg/utils"

	"github.com/gorilla/mux"

	_ "joosum-backend/docs" // load Swagger docs

	"github.com/create-go-app/net_http-go-template/app/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs for Golang net/http Template.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	//db connnect

		 * MongoDB 연결 예시
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		mongoURI := "mongodb://localhost:27017"

		client, err := mongodb.GetMongoClient(ctx, mongoURI)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		// MongoDB 클라이언트 연결을 종료합니다.
		err = mongodb.DisconnectMongoClient(ctx)
		if err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB")


	// Initialize a new router.
	router := mux.NewRouter()

	// List of app routes:
	routes.PublicRoutes(router)
	routes.PrivateRoutes(router)
	routes.SwaggerRoutes(router)
	//Initialize a new router.
	//router := mux.NewRouter()
	//
	//// List of app routes:
	//routes.PublicRoutes(router)
	//routes.PrivateRoutes(router)
	//routes.SwaggerRoutes(router)
	//
	//// Register middleware.
	//router.Use(mux.CORSMethodMiddleware(router)) // enable CORS
	//
	//// Initialize server.
	//server := configs.ServerConfig(router)
	//
	//// Start API server.
	//utils.StartServerWithGracefulShutdown(server)

	router := gin.Default()

	router.GET("/", controllers.GetMainPage)
	router.GET("/api/v1/users", controllers.GetUsers)
	router.GET("/api/v1/users/:id", controllers.GetUser)
	router.POST("/api/v1/users", controllers.CreateUser)
	router.PUT("/api/v1/users/:id", controllers.UpdateUser)
	router.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	router.Run(":5001") // listen and serve on 0.0.0.0:5001 (for windows "localhost:5001")
}
