package main

import (
    "context"
    "log"
	"fmt"
	"encoding/json"

    "learn-grpc-rpc/common/config"
    "learn-grpc-rpc/common/model"
	"github.com/golang/protobuf/ptypes/empty"

    "google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
 
	return model.NewGaragesClient(conn)
 }

 func serviceUser() model.UsersClient {
    port := config.SERVICE_USER_PORT
    conn, err := grpc.Dial(port, grpc.WithInsecure())
    if err != nil {
        log.Fatal("could not connect to", port, err)
    }

    return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		  Id:       "n001",
		  Name:     "Noval Agung",
		  Password: "kw8d hl12/3m,a",
		  Gender:   model.UserGender(model.UserGender_value["MALE"]),
	  }
  
	//   garage1 := model.Garage{
	// 	  Id:   "q001",
	// 	  Name: "Quel'thalas",
	// 	  Coordinate: &model.GarageCoordinate{
	// 		  Latitude:  45.123123123,
	// 		  Longitude: 54.1231313123,
	// 	  },
	//   }
  
	  user := serviceUser()

	  fmt.Println("\n", "===========> user test")
	  
	  // register user1
	  user.Register(context.Background(), &user1)
	  
	//   // register user2
	//   user.Register(context.Background(), &user2)
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))


  }