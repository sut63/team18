package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team18/app/controllers"
	_ "github.com/team18/app/docs"
	"github.com/team18/app/ent"
)

type RoomTypes struct {
	RoomType []RoomType
}

type RoomType struct {
	name string
}

type RoomStatus struct {
	RoomStatu []RoomStatu
}

type RoomStatu struct {
	name string
}

type Promotions struct {
	Promotion []Promotion
}

type Promotion struct {
	name     string
	discount float64
}

type Customers struct {
	Customer []Customer
}

type Customer struct {
	name  string
	email string
}

// @title SUT SE Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewReserveRoomController(v1, client)
	controllers.NewPromotionController(v1, client)
	controllers.NewCustomerController(v1, client)
	controllers.NewDataRoomController(v1, client)

	// Set RoomType Data
	types := RoomTypes{
		RoomType: []RoomType{
			RoomType{"Standard"},
			RoomType{"Superior"},
			RoomType{"Suite"},
			RoomType{"Duplex"},
			RoomType{"Deluxe"},
		},
	}

	for _, t := range types.RoomType {
		client.TypeRoom.
			Create().
			SetTypeName(t.name).
			Save(context.Background())
	}

	// Set RoomStatus Data
	status := RoomStatus{
		RoomStatu: []RoomStatu{
			RoomStatu{"Aviable"},
			RoomStatu{"Unaviable"},
		},
	}

	for _, s := range status.RoomStatu {
		client.StatusRoom.
			Create().
			SetStatusName(s.name).
			Save(context.Background())
	}

	// Set Promotion Data
	promo := Promotions{
		Promotion: []Promotion{
			Promotion{"New Year Sale!!", 2000.00},
			Promotion{"Summer Sale", 1000.00},
		},
	}

	for _, p := range promo.Promotion {
		client.Promotion.
			Create().
			SetPromotionName(p.name).
			SetDiscount(p.discount).
			Save(context.Background())
	}

	// Set Customers Data
	customer := Customers{
		Customer: []Customer{
			Customer{"sawadee", "example@gmail.com"},
			Customer{"hello", "hellow@gmail.com"},
		},
	}

	for _, c := range customer.Customer {
		client.Customer.
			Create().
			SetName(c.name).
			SetEmail(c.email).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
