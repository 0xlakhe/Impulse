package main

import (
	"context"
	"fmt"
	"log"

	"github.com/0xlakhe/Impluse/internal/app"
	"github.com/0xlakhe/Impluse/internal/config"
	"github.com/0xlakhe/Impluse/internal/database"
	"github.com/0xlakhe/Impluse/internal/server"
	"github.com/joho/godotenv"
)

func main(){		
	_=godotenv.Load()	
	cfg:=config.Load()
	ctx:=context.Background()
	dbPool,err:=database.NewPostgresPool(ctx,cfg.DatabaseURL)

	if err!=nil{
		log.Fatal(err)
	}
	db:=database.New(dbPool)
	application:=app.New(cfg,db)
	defer application.Shutdown(ctx)
	s:=server.New(application);
	fmt.Printf("Server running of http://localhost: %s",cfg.Port)
	err=s.Start();
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("Server running of http://localhost: %s",cfg.Port)
}