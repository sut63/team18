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

type StatusOpinions struct {
	StatusOpinion []StatusOpinion
}

type StatusOpinion struct {
	Opinion string
}

type Statuss struct {
	Status []Status
}

type Status struct {
	StatusDescription string
}

type StatusRooms struct {
	StatusRoom []StatusRoom
}

type StatusRoom struct {
	StatusName string
}

type StatusReserves struct {
	StatusReserve []StatusReserve
}

type StatusReserve struct {
	Status string
}

type StatusCheckins struct {
	StatusCheckin []StatusCheckin
}

type StatusCheckin struct {
	Status string
}

type TypeRooms struct {
	TypeRoom []TypeRoom
}

type TypeRoom struct {
	TypeName string
}

type Promotions struct {
	Promotion []Promotion
}

type Promotion struct {
	PromotionName string
	Discount      float64
}

type Customers struct {
	Customer []Customer
}

type Customer struct {
	name     string
	email    string
	password string
}

type CounterStaffs struct {
	CounterStaff []CounterStaff
}

type CounterStaff struct {
	name     string
	email    string
	password string
}

type Furnituretypes struct {
	FurnitureType []FurnitureType
}

type FurnitureType struct {
	type_name string
}

type Furnitures struct {
	Furniture []Furniture
}

type Furniture struct {
	furniture_name string
}

type DataRooms struct {
	DataRoom []DataRoom
}

type DataRoom struct {
	price      float64
	roomnumber string
	roomdetail string

	promotion  int
	statusroom int
	typeroom   int
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
	controllers.NewCheckinController(v1, client)
	controllers.NewCheckoutController(v1, client)
	controllers.NewCustomerController(v1, client)
	controllers.NewDataRoomController(v1, client)
	controllers.NewPromotionController(v1, client)
	controllers.NewReserveRoomController(v1, client)
	controllers.NewStatusController(v1, client)
	controllers.NewStatusReserveController(v1, client)
	controllers.NewStatusRoomController(v1, client)
	controllers.NewTypeRoomController(v1, client)
	controllers.NewCounterStaffController(v1, client)
	controllers.NewFurnitureDetailController(v1, client)
	controllers.NewFurnituretypeController(v1, client)
	controllers.NewFurnitureController(v1, client)
	controllers.NewFixRoomController(v1, client)
	controllers.NewStatusCheckinController(v1, client)
	controllers.NewStatusOpinionController(v1, client)

	// Set StatusRoom Data
	statusrooms := StatusRooms{
		StatusRoom: []StatusRoom{
			StatusRoom{"ว่าง"},
			StatusRoom{"ไม่ว่าง"},
		},
	}
	for _, s := range statusrooms.StatusRoom {
		client.StatusRoom.
			Create().
			SetStatusName(s.StatusName).
			Save(context.Background())
	}

	//set opinion data for checkout
	statusopinion := StatusOpinions{
		StatusOpinion: []StatusOpinion{
			StatusOpinion{"พอใจมาก"},
			StatusOpinion{"รู้สึกเฉยๆ"},
			StatusOpinion{"ไม่พอใจ"},
		},
	}
	for _, sc := range statusopinion.StatusOpinion {
		client.StatusOpinion.
			Create().
			SetOpinion(sc.Opinion).
			Save(context.Background())
	}

	// set status Data for checkout
	statuss := Statuss{
		Status: []Status{
			Status{"ชำระเรียบร้อย"},
			Status{"ยังไม่ชำระ"},
			Status{"ชำระไม่ครบ"},
		},
	}
	for _, sc := range statuss.Status {
		client.Status.
			Create().
			SetDescription(sc.StatusDescription).
			Save(context.Background())
	}

	// Set TypeRoom Data
	typerooms := TypeRooms{
		TypeRoom: []TypeRoom{
			TypeRoom{"Standard"},
			TypeRoom{"Superior"},
			TypeRoom{"Deluxe"},
			TypeRoom{"Suite"},
		},
	}

	for _, t := range typerooms.TypeRoom {
		client.TypeRoom.
			Create().
			SetTypeName(t.TypeName).
			Save(context.Background())
	}

	// Set Promotion Data
	promotions := Promotions{
		Promotion: []Promotion{
			Promotion{"ปีใหม่", 1200.50},
			Promotion{"สงกรานต์", 500},
			Promotion{"ฮาโลวีน", 350},
		},
	}

	for _, p := range promotions.Promotion {
		client.Promotion.
			Create().
			SetPromotionName(p.PromotionName).
			SetDiscount(p.Discount).
			Save(context.Background())
	}

	// Set Customers Data
	customer := Customers{
		Customer: []Customer{
			Customer{"Bos", "bos@gmail.com", "bos123"},
			Customer{"Noi", "noi@gmail.com", "noi666"},
			Customer{"Best", "best@gmail.com", "best33"},
			Customer{"Tongkong", "tongkong@gmail.com", "tong456"},
			Customer{"Ta", "ta@gmail.com", "ta007"},
			Customer{"Film", "film@gmail.com", "film89"},
		},
	}

	for _, c := range customer.Customer {
		client.Customer.
			Create().
			SetName(c.name).
			SetEmail(c.email).
			SetPassword(c.password).
			Save(context.Background())
	}

	// Set StatusReserve Data
	statusreserves := StatusReserves{
		StatusReserve: []StatusReserve{
			StatusReserve{"ยังไม่เข้าพัก"},
			StatusReserve{"เข้าพัก"},
		},
	}
	for _, sr := range statusreserves.StatusReserve {
		client.StatusReserve.
			Create().
			SetStatusName(sr.Status).
			Save(context.Background())
	}

	// Set Statuscheckin Data
	statuscheckins := StatusCheckins{
		StatusCheckin: []StatusCheckin{
			StatusCheckin{"พักอยู่"},
			StatusCheckin{"ออกแล้ว"},
		},
	}
	for _, si := range statuscheckins.StatusCheckin {
		client.StatusCheckIn.
			Create().
			SetStatusName(si.Status).
			Save(context.Background())
	}

	// Set Data Room Data
	datarooms := DataRooms{
		DataRoom: []DataRoom{
			DataRoom{4500, "A001", "ห้องนี้ผีดุ", 1, 1, 1},
			DataRoom{5000, "A002", "ใหม่เอี่ยม", 2, 1, 1},
			DataRoom{5500, "B001", "หลับสบาย", 2, 1, 2},
			DataRoom{6000, "C001", "โคตรหรู", 3, 1, 3},
		},
	}

	for _, dr := range datarooms.DataRoom {
		client.DataRoom.
			Create().
			SetPrice(dr.price).
			SetRoomnumber(dr.roomnumber).
			SetRoomdetail(dr.roomdetail).
			SetPromotionID(dr.promotion).
			SetStatusroomID(dr.statusroom).
			SetTyperoomID(dr.typeroom).
			Save(context.Background())
	}

	// Set Furniture Data
	furnitures := Furnitures{
		Furniture: []Furniture{
			Furniture{"เก้าอี้"},
			Furniture{"โซฟา"},
			Furniture{"เตียง"},
			Furniture{"แอร์"},
			Furniture{"ตู้เย็น"},
		},
	}
	for _, f := range furnitures.Furniture {
		client.Furniture.
			Create().
			SetFurnitureName(f.furniture_name).
			Save(context.Background())
	}

	// Set Furniture type Data
	furnituretypes := Furnituretypes{
		FurnitureType: []FurnitureType{
			FurnitureType{"เครื่องเรือน"},
			FurnitureType{"เครื่องใช้ไฟฟ้า"},
		},
	}
	for _, ft := range furnituretypes.FurnitureType {
		client.FurnitureType.
			Create().
			SetFurnitureType(ft.type_name).
			Save(context.Background())
	}

	// Set CounterStaff Data
	counterstaffs := CounterStaffs{
		CounterStaff: []CounterStaff{
			CounterStaff{"Staff No.1", "staff1@gmail.com", "1234"},
			CounterStaff{"Staff No.2", "staff2@gmail.con", "abcd"},
		},
	}
	for _, ct := range counterstaffs.CounterStaff {
		client.CounterStaff.
			Create().
			SetName(ct.name).
			SetEmail(ct.email).
			SetPassword(ct.password).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
