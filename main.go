package main

import (
	"beauty-shop/auth"
	"beauty-shop/delivery"
	repo "beauty-shop/repository"
	serv "beauty-shop/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*func InitProductApi(base repo.BaseRepository) delivery.BookHandler {
	repository := repo.ProvideBookRepository(base)
	service := serv.ProvideBookUseCase(repository)
	api := delivery.ProviderBookHandler(service)
	return api
}*/

func InitUserApi(base repo.BaseRepository) delivery.UserHandler {
	repository := repo.ProvideUserRepository(base)
	service := serv.ProvideUserUseCase(repository)
	api := delivery.ProviderUserHandler(service)

	return api
}

func InitProductApi(base repo.BaseRepository) delivery.ProductHandler {
	repository := repo.ProvideProductRepository(base)
	service := serv.ProvideProductUseCase(repository)
	api := delivery.ProvideProductHandler(service)

	return api
}

func initDB() repo.BaseRepository {
	er := godotenv.Load("data.env")

	if er != nil {
		log.Fatalf("Error loading .env file")
	}

	dbDriver := os.Getenv("APP_DB_DRIVER")
	username := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	host := os.Getenv("APP_DB_HOST")
	database := os.Getenv("APP_DB_NAME")
	dbPort := os.Getenv("APP_DB_PORT")

	repos, err := repo.Model.Initialize(dbDriver, username, password, dbPort, host, database)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return repos
}

func initRouter(base repo.BaseRepository) *gin.Engine {
	r := gin.Default()
	main := r.Group("beauty-shop")
	//setBookRouter(main, base)
	setProductRouter(main, base)
	setLoginRegister(main, base)

	return r
}

func setProductRouter(r *gin.RouterGroup, repository repo.BaseRepository) {
	productHandler := InitProductApi(repository)

	product := r.Group("product", auth.TokenAuthMiddleware())
	product.GET("/findAll", productHandler.GetAll)
	product.GET("/find/:id", productHandler.GetById)
	product.POST("/create", productHandler.Create)
	product.POST("/update/:id", productHandler.Update)
	product.POST("/delete/:id", productHandler.Delete)
}

func setLoginRegister(r *gin.RouterGroup, repository repo.BaseRepository) {
	userHandler := InitUserApi(repository)

	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)
	r.POST("/uploadPhoto", delivery.UploadPhotoHandler)
}

/*func setBookRouter(r *gin.RouterGroup, base repo.BaseRepository) {
	bookHandler := InitProductApi(base)

	book := r.Group("book", auth.TokenAuthMiddleware())
	book.GET("/findAll", bookHandler.FindAll)
	book.GET("/find/:id", bookHandler.FindByID)
	book.POST("/create", bookHandler.Create)
	book.PUT("/update/:id", bookHandler.Update)
	book.DELETE("/delete/:id", bookHandler.Delete)
}*/

func main() {
	base := initDB()
	r := initRouter(base)

	err := r.Run(":8010")
	if err != nil {
		panic(err)
	}
}
