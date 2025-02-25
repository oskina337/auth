package main

import (
	"auth/internal/db"
	"auth/internal/repository/auth"

	// "auth/internal/repository/auth/model"

	"context"
	"fmt"
	"log"
)

// const grpcPort = 50051

// func main() {
// 	// fmt.Println(color.GreenString("Hello, world! проверка"))
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	reflection.Register(s)
// 	pb.RegisterUserAPIV1Server(s, &service.Server{})

// 	log.Printf("server listening at %v", lis.Addr())

// 	if err = s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

func main() {
	// Инициализация клиента базы данных
	client, err := db.NewClient()
	if err != nil {
		log.Fatalf("Error while initializing DB client: %v", err)
	}
	defer client.Close()

	// Инициализация репозитория
	repo := auth.NewRepository(client)

	// // Добавление нового пользователя
	// usre := model.User{
	// 	Name:     "John Doe",
	// 	Email:    "1244h5n4@example.com",
	// 	Password: "hashedpassword123",
	// 	Role:     model.Role_ADMIN,
	// }

	// userID, err := repo.CreateUser(context.Background(), usre)
	// if err != nil {
	// 	log.Fatalf("Error creating user: %v", err)
	// }

	// // Вывод ID нового пользователя
	// fmt.Printf("User created with ID: %d\n", userID)

	//проверка Get
	users, err := repo.GetUser(context.Background(), 11)
	if err != nil {
		log.Fatalf("Error get user: %v", err)
	}

	fmt.Printf("User get (id %d): %v\n", 11, *users)

	//проверка Create
	// userCreate, err := repo.UpdateUser(context.Background(), &model.User{
	// 	ID:       11,
	// 	Name:     "John Doe",
	// 	Email:    "228n4@example.com",
	// 	Password: "hashedpassword123",
	// 	Role:     model.Role_USER})
	// if err != nil {
	// 	log.Fatalf("Error get user: %v", err)
	// }

	// fmt.Printf("User get (id %d): %v\n", 11, *userCreate)

}
