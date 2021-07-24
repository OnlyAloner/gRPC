package main

import (
	"context"
	"log"
	"net"

	cml "grpc/contactlist"
	model "grpc/model"
	pb "grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedContactRequestServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := in.GetUser()

	cm, err := cml.NewContactManager()
	if err != nil {
		return nil, err
	}
	var temp model.Contact
	temp.ID = int(user.Id)
	temp.CreateAt = user.Createdat
	temp.Email = user.Email
	temp.Gender = user.Gender
	temp.Name = user.Name
	temp.Phone = user.Phone

	err = cm.AddContact(&temp)

	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: in.User}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := in.User
	cm, err := cml.NewContactManager()
	if err != nil {
		return nil, err
	}
	var temp model.Contact
	temp.ID = int(user.Id)
	temp.CreateAt = user.Createdat
	temp.Email = user.Email
	temp.Gender = user.Gender
	temp.Name = user.Name
	temp.Phone = user.Phone
	err = cm.UpdateContact(temp.ID, &temp)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{User: in.User}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user_id := in.Id
	cm, err := cml.NewContactManager()

	if err != nil {
		return nil, err
	}

	user, err := cm.GetContact(int(user_id))

	if err != nil {
		return nil, err
	}

	var temp *pb.User = new(pb.User)

	temp.Id = uint32(user.ID)
	temp.Createdat = user.CreateAt
	temp.Email = user.Email
	temp.Gender = user.Gender
	temp.Name = user.Name
	temp.Phone = user.Phone

	return &pb.GetUserResponse{User: temp}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	user_id := in.Id
	cm, err := cml.NewContactManager()
	if err != nil {
		return &pb.DeleteUserResponse{Success: false}, err
	}
	err = cm.DeleteContact(int(user_id))
	if err != nil {
		return &pb.DeleteUserResponse{Success: false}, err
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}

func (s *server) GetAll(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	var temp []*pb.User = make([]*pb.User, 0)
	cm, err := cml.NewContactManager()
	if err != nil {
		return nil, err
	}
	contacts, err := cm.GetAllContacts()
	if err != nil {
		return nil, err
	}
	for index := range contacts {

		user := &pb.User{
			Id:        uint32(contacts[index].ID),
			Name:      contacts[index].Name,
			Phone:     contacts[index].Phone,
			Gender:    contacts[index].Gender,
			Email:     contacts[index].Email,
			Createdat: contacts[index].CreateAt,
		}
		temp = append(temp, user)
	}
	return &pb.GetAllResponse{
		User: temp}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterContactRequestServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
