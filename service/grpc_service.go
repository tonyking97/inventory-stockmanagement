package service

import (
	"../db"
	"../models"
	"../proto/inventory"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) Ping(ctx context.Context, req *inventory_pb.PingRequest) (*inventory_pb.PingResponse, error) {
	log.Println("Incoming Ping Request.")
	response := "pong from go server"
	res := &inventory_pb.PingResponse{
		Pong: response,
	}
	return res, nil
}

func (s *server) GetCategory(ctx context.Context, req *inventory_pb.GetCategoryRequest) (*inventory_pb.GetCategoryResponse, error) {
	log.Println("Incoming Get Product Request")

	var categoryNames []*inventory_pb.CategoryName
	categoryCollection := db.CategoryCollectionInit()
	filter := bson.D{{}}
	projection := bson.D{
		{"name", 1},
	}
	findOptions := options.Find().SetProjection(projection)

	cur, err := categoryCollection.Find(db.Ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var category models.Category
		err := cur.Decode(&category)
		if err != nil {
			log.Fatal(err)
		}
		categoryNames = append(categoryNames, &inventory_pb.CategoryName{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	res := &inventory_pb.GetCategoryResponse{
		Success:      true,
		CategoryName: categoryNames,
	}

	return res, nil
}

func (*server) AddCategory(ctx context.Context, req *inventory_pb.AddCategoryRequest) (*inventory_pb.AddCategoryResponse, error) {
	log.Println("Incoming Add Product Request")

	categoryCollection := db.CategoryCollectionInit()

	data := &models.Category{
		Id:           uuid.New().String(),
		Time:         time.Now().Unix(),
		Name:         req.Name,
		Description:  req.Description,
		Parent:       req.Parent,
		Unit:         req.Unit,
		Manufacturer: req.Manufacturer,
		Brand:        req.Brand,
	}

	insertResult, err := categoryCollection.InsertOne(db.Ctx, data)

	if err != nil {
		log.Println("Error -> ", err)
		return &inventory_pb.AddCategoryResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	log.Println("Inserted "+req.Name+" in MongoDB - ", insertResult)

	return &inventory_pb.AddCategoryResponse{
		Success: true,
		Error:   "",
	}, nil
}

func GRPCInit() {
	go startGRPCServer()
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	inventory_pb.RegisterInventoryServiceServer(s, &server{})

	reflection.Register(s)

	log.Println("Starting gRPC Service...")

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to start gRPC Service.", err)
	}
}
