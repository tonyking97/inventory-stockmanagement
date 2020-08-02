package service

import (
	"../core"
	"../models"
	"../proto/inventory"
	"context"
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
	log.Println("Incoming Get Category Request")

	var categories []*inventory_pb.Category

	if req.Id != "" {
		category, err := core.GetCategory(req.Id)
		if err != nil {
			return &inventory_pb.GetCategoryResponse{
				Success:  false,
				Category: nil,
			}, nil
		}
		categories = append(categories, &inventory_pb.Category{
			Id:   category.Id,
			Name: category.Name,
		})
	} else {
		categoriesArray, err := core.GetCategories()
		if err != nil {
			return &inventory_pb.GetCategoryResponse{
				Success:  false,
				Category: nil,
			}, nil
		}
		for _, category := range categoriesArray {
			categories = append(categories, &inventory_pb.Category{
				Id:   category.Id,
				Name: category.Name,
			})
		}
	}

	res := &inventory_pb.GetCategoryResponse{
		Success:  true,
		Category: categories,
	}

	return res, nil
}

func (*server) AddCategory(ctx context.Context, req *inventory_pb.AddCategoryRequest) (*inventory_pb.AddCategoryResponse, error) {
	log.Println("Incoming Add Category Request")

	data := &models.Category{
		Name:         req.Name,
		Description:  req.Description,
		Parent:       req.Parent,
		Unit:         req.Unit,
		Manufacturer: req.Manufacturer,
		Brand:        req.Brand,
	}

	err := core.AddCategory(*data)

	if err != nil {
		log.Println("Error -> ", err)
		return &inventory_pb.AddCategoryResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &inventory_pb.AddCategoryResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *server) EditCategory(ctx context.Context, req *inventory_pb.EditCategoryRequest) (*inventory_pb.EditCategoryResponse, error) {
	log.Println("Incoming Edit Category")

	data := &models.EditCategory{
		Name:         req.Name,
		Description:  req.Description,
		Parent:       req.Parent,
		Unit:         req.Unit,
		Manufacturer: req.Manufacturer,
		Brand:        req.Brand,
	}

	err := core.EditCategory(req.Id, *data)

	if err != nil {
		return &inventory_pb.EditCategoryResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &inventory_pb.EditCategoryResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *server) AddItem(ctx context.Context, req *inventory_pb.AddItemRequest) (*inventory_pb.AddItemResponse, error) {
	log.Println("Incoming Add Item Request")

	data := &models.Item{
		CreatedBy:        req.CreatedBy,
		Name:             req.Name,
		Status:           req.Status,
		Size:             req.Size,
		SKU:              req.Sku,
		BarCode:          req.BarCode,
		Category:         req.Category,
		Unit:             req.Unit,
		SalesInformation: req.SalesInformation,
		MRP:              req.Mrp,
		CostPrice:        req.CostPrice,
		SellingPrice:     req.SellingPrice,
		TrackInventory:   req.TrackInventory,
		Stock:            req.Stock,
		ReorderPoint:     req.ReorderPoint,
		PreferredVendor:  req.PreferredVendor,
	}

	err := core.AddItem(*data)

	if err != nil {
		return &inventory_pb.AddItemResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &inventory_pb.AddItemResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *server) EditItem(ctx context.Context, req *inventory_pb.EditItemRequest) (*inventory_pb.EditItemResponse, error) {
	log.Println("Incoming Edit Item")

	data := &models.EditItem{
		UpdatedTime:     time.Now().Unix(),
		Name:            req.Name,
		Status:          req.Status,
		Size:            req.Size,
		SKU:             req.Sku,
		BarCode:         req.BarCode,
		Category:        req.Category,
		Unit:            req.Unit,
		ReorderPoint:    req.ReorderPoint,
		PreferredVendor: req.PreferredVendor,
	}

	err := core.EditItem(req.Id, *data)

	if err != nil {
		log.Println("Error -> ", err)
		return &inventory_pb.EditItemResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &inventory_pb.EditItemResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *server) GetItem(ctx context.Context, req *inventory_pb.GetItemRequest) (*inventory_pb.GetItemResponse, error) {
	log.Println("Incoming Get Item Request")

	var items []*inventory_pb.Item

	if req.Id != "" {
		item, err := core.GetItem(req.Id)
		if err != nil {
			return &inventory_pb.GetItemResponse{
				Success: false,
				Item:    nil,
			}, nil
		}
		items = append(items, &inventory_pb.Item{
			Id:   item.Id,
			Name: item.Name,
		})
	} else {
		itemsArray, err := core.GetItems()
		if err != nil {
			return &inventory_pb.GetItemResponse{
				Success: false,
				Item:    nil,
			}, nil
		}
		for _, item := range itemsArray {
			items = append(items, &inventory_pb.Item{
				Id:   item.Id,
				Name: item.Name,
			})
		}
	}

	res := &inventory_pb.GetItemResponse{
		Success: true,
		Item:    items,
	}

	return res, nil
}

func (s *server) GetQuantityAdjustment(ctx context.Context, req *inventory_pb.GetQuantityAdjustmentRequest) (*inventory_pb.GetQuantityAdjustmentResponse, error) {
	log.Println("Incoming Get Quantity Adjustment Request")

	if req.Id == "" {
		return &inventory_pb.GetQuantityAdjustmentResponse{
			Success:           false,
			Error:             "Item Id not found",
			QuantityAvailable: 0,
			Unit:              "",
		}, nil
	}

	stock, err := core.GetQuantity(req.Id)
	if err != nil {
		return &inventory_pb.GetQuantityAdjustmentResponse{
			Success:           false,
			Error:             "Invalid Item Id",
			QuantityAvailable: 0,
			Unit:              "",
		}, nil
	}

	res := &inventory_pb.GetQuantityAdjustmentResponse{
		Success:           true,
		Error:             "",
		QuantityAvailable: float32(stock.AvailableQuantity),
		Unit:              stock.Unit,
	}

	return res, nil
}

func (s *server) GetValueAdjustment(ctx context.Context, req *inventory_pb.GetValueAdjustmentRequest) (*inventory_pb.GetValueAdjustmentResponse, error) {
	log.Println("Incoming Get Quantity Adjustment Request")

	if req.Id == "" {
		return &inventory_pb.GetValueAdjustmentResponse{
			Success:             false,
			Error:               "Item Id not found",
			CurrentSellingPrice: 0,
			CurrentCostPrice:    0,
			CurrentMrp:          0,
		}, nil
	}

	cost, err := core.GetCost(req.Id)
	if err != nil {
		log.Println(err)
		return &inventory_pb.GetValueAdjustmentResponse{
			Success:             false,
			Error:               "Invalid Item Id",
			CurrentSellingPrice: 0,
			CurrentCostPrice:    0,
			CurrentMrp:          0,
		}, nil
	}

	res := &inventory_pb.GetValueAdjustmentResponse{
		Success:             true,
		Error:               "",
		CurrentSellingPrice: float32(cost.SellingPrice),
		CurrentCostPrice:    float32(cost.CostPrice),
		CurrentMrp:          float32(cost.Mrp),
	}

	return res, nil
}

func (s *server) AdjustQuantity(ctx context.Context, req *inventory_pb.AdjustQuantityRequest) (*inventory_pb.AdjustQuantityResponse, error) {
	log.Println("Incoming Adjust Quantity Request")

	data := &models.Stock{
		Date:             req.Date,
		Quantity:         int64(req.NewQuantity),
		QuantityAdjusted: int64(req.QuantityAdjusted),
		Reason:           req.Reason,
		ReferenceNumber:  req.ReferenceNumber,
	}

	err := core.AdjustQuantity(req.Id, *data)

	if err != nil {
		return &inventory_pb.AdjustQuantityResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &inventory_pb.AdjustQuantityResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *server) AdjustValue(ctx context.Context, req *inventory_pb.AdjustValueRequest) (*inventory_pb.AdjustValueResponse, error) {
	log.Println("Incoming Adjust Value Request")

	data := &models.Cost{
		Date:                 req.Date,
		CostPrice:            int64(req.NewCostPrice),
		AdjustedCostPrice:    int64(req.AdjustedCostPrice),
		SellingPrice:         int64(req.NewSellingPrice),
		AdjustedSellingPrice: int64(req.AdjustedSellingPrice),
		Mrp:                  int64(req.NewMrp),
		AdjustedMrp:          int64(req.AdjustedMrp),
		Reason:               req.Reason,
		ReferenceNumber:      req.ReferenceNumber,
	}

	err := core.AdjustValue(req.Id, *data)

	if err != nil {
		return &inventory_pb.AdjustValueResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &inventory_pb.AdjustValueResponse{
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
