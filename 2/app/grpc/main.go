package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_logRepo "github.com/bxcodec/go-clean-arch/log/repository/mysql"
	_movieGrpcDelivery "github.com/bxcodec/go-clean-arch/movie/delivery/grpc"
	_movieRepo "github.com/bxcodec/go-clean-arch/movie/repository/api"
	_movieUcase "github.com/bxcodec/go-clean-arch/movie/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	logRepo := _logRepo.NewMysqlLogRepository(dbConn)
	movieRepo := _movieRepo.NewAPIMovieRepository()
	movieUsecase := _movieUcase.NewMovieUsecase(movieRepo, logRepo, timeoutContext)

	list, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("SOMETHING HAPPEN")
	}

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*32), grpc.MaxSendMsgSize(1024*1024*32))
	_movieGrpcDelivery.NewGrpcMovieHandler(server, movieUsecase)
	reflection.Register(server)

	err = server.Serve(list)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
	fmt.Println("GRPC run at :3000")
}
