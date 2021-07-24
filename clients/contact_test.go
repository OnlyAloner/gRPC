package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

var (
	conn   *grpc.ClientConn
	ctx    context.Context
	cancel context.CancelFunc
	c      pb.ContactRequestClient
	id     int
)

func TestMain(t *testing.M) {
	if conn != nil {
		os.Exit(1)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Couldnt connect: %v", err)
	}

	c = pb.NewContactRequestClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	os.Exit(t.Run())
}

func TestCreate(t *testing.T) {
	user := &pb.User{
		Name:   "Otabek",
		Phone:  "998999048088",
		Gender: "Male",
		Email:  "mad@gmail.com",
	}
	if err := Create(c, user, ctx); err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
	id = int(user.Id)
}

func TestUpdate(t *testing.T) {
	if err := Update(c, &pb.User{
		Id:     15,
		Name:   "Stan Lee",
		Phone:  "998999048088",
		Gender: "Male",
		Email:  "stan@gmail.com",
	}, ctx); err != nil {
		t.Errorf("Failed to Update user: %v", err)
	}
}

func TestGet(t *testing.T) {
	if err := Get(c, 14, ctx); err != nil {
		t.Errorf("Failed to Get user: %v", err)
	}
}

func TestDelete(t *testing.T) {
	if err := Delete(c, 15, ctx); err != nil {
		t.Errorf("failed to delete contact: %v", err)
	}
}

func TestGetAll(t *testing.T) {
	if err := GetAll(c, ctx); err != nil {
		t.Errorf("failed to get all contacts: %v", err)
	}
}
