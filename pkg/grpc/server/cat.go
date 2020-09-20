package server

import (
	"context"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/api/server"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/model"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/service"
	"github.com/google/uuid"
)

type CatServer struct {}

func (s *CatServer) ReadAll(ctx context.Context, req *server.ReadAllRequest) (*server.ReadAllResponse, error) {
	var cats []model.Cat
	err := service.GetCats(ctx, &cats)
	if err != nil {
		return nil, err
	}
	var formattedCats []*server.Cat
	for _, v := range cats {
		formattedCat := server.Cat{
			Id:         v.ID.String(),
			Name:       v.Name,
			Color:      v.Color,
			CategoryId: v.CategoryID,
		}
		formattedCats = append(formattedCats, &formattedCat)
	}
	return &server.ReadAllResponse{
		Cats: formattedCats,
	}, nil
}

func (s *CatServer) Read(ctx context.Context, req *server.ReadRequest) (*server.ReadResponse, error) {
	var cat model.Cat
	err := service.GetCatByID(ctx, &cat, req.Id)
	if err != nil {
		return nil, err
	}
	return &server.ReadResponse{
		Cat: &server.Cat{
			Id:         cat.ID.String(),
			Name:       cat.Name,
			Color:      cat.Color,
			CategoryId: cat.CategoryID,
		},
	}, nil
}

func (s *CatServer) Create(ctx context.Context, req *server.CreateRequest) (*server.CreateResponse, error) {
	newCat := model.Cat{
		Name:       req.Cat.Name,
		Color:      req.Cat.Color,
		CategoryID: req.Cat.CategoryId,
	}
	err := service.CreateCat(ctx, &newCat)
	if err != nil {
		return nil, err
	}
	return &server.CreateResponse{
		Cat: &server.Cat{
			Id:         newCat.ID.String(),
			Name:       newCat.Name,
			Color:      newCat.Color,
			CategoryId: newCat.CategoryID,
		},
	}, nil
}

func (s *CatServer) Update(ctx context.Context, req *server.UpdateRequest) (*server.UpdateResponse, error) {
	cat := model.Cat{
		Pure:       model.Pure{
			ID: uuid.MustParse(req.Cat.Id),
		},
		Name:       req.Cat.Name,
		Color:      req.Cat.Color,
		CategoryID: req.Cat.CategoryId,
	}
	err := service.UpdateCat(ctx, &cat)
	if err != nil {
		return nil, err
	}
	return &server.UpdateResponse{Cat: &server.Cat{
		Id:         cat.ID.String(),
		Name:       cat.Name,
		Color:      cat.Color,
		CategoryId: cat.CategoryID,
	}}, nil
}

func (s *CatServer) Delete(ctx context.Context, req *server.DeleteRequest) (*server.DeleteResponse, error) {
	err := service.DeleteCat(ctx, &model.Cat{}, req.Id)
	if err != nil {
		return nil, err
	}
	return &server.DeleteResponse{Deleted: true}, nil
}
