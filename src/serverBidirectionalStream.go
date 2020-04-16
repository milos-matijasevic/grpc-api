package main

import (
	"./protobi"
	"net"
	"time"
	"fmt"
	"io"
	"flag"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


type Video_metric struct {
	Id int `gorm:"primary_key;not null"`
	Event_name string `gorm:"type:character varying;not null"`
	Device string `gorm:"type:character varying;not null"`
	Session_id *string `gorm:"type:character varying"`
	Impression_id *string `gorm:"type:character varying"`
	Timestamp time.Time `gorm:"type:timestamp without time zone;not null"`
	Price *int
	Publisher_revenue *int
	Our_revenue *int
	Currency *int `gorm:"type:video_metrics_currency_enum"`
	Video_id *int
	Site_id *int
	Line_item_id *int
	Advertiser_id *int
}

type server struct {}

type ResultPublisherRevenue struct {
  Publisher  int32
  Revenue int64
}

type ResultAdvertiserEventNum struct {
  Device  string
  EventNum int32
}

var db *gorm.DB
var err error

func main(){
	dbName := flag.String("dbType", "postgres", "a string")
	
	flag.Parse()
	if *dbName == "postgres" {
		db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=grpcchallenge password=root sslmode=disable")
	} else if *dbName == "sqlite3" {
		db, err = gorm.Open("sqlite3", "./video_metrics.sqlite3?parseTime=true")
	} else {
		fmt.Println("dbType should be postgres or sqlite")
		return
	}

	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	fmt.Println(db.HasTable("video_metrics"))
	db.AutoMigrate(&Video_metric{})
	
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	
	srv := grpc.NewServer()
	proto.RegisterFServiceServer(srv, &server{})
	reflection.Register(srv)
	
	if e:= srv.Serve(listener); e!= nil {
		panic(err)
	}
}

func (s *server)GetMoneyForEachPublisher(stream proto.FService_GetMoneyForEachPublisherServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		var reply proto.Response
		layout := "2006-01-02 15:04:05.00"
		startDate, err := time.Parse(layout, request.GetStartDate())
		if err != nil {
			return err
		}
		endDate, err :=  time.Parse(layout, request.GetEndDate())
		if err != nil {
			return err
		}
		rows, err := db.Table("video_metrics").Select("advertiser_id as publisher, sum(publisher_revenue) as revenue").Where("timestamp BETWEEN ? AND ?", startDate, endDate).Group("advertiser_id").Rows()
		if err != nil {
			return err
		}
		var res ResultPublisherRevenue
		
		for rows.Next() {	
			db.ScanRows(rows, &res)
			reply.Publisher = res.Publisher
			reply.Revenue = res.Revenue
			if err := stream.Send(&reply); err != nil {
				return err
			}
		}
	}
}

func (s *server)GetNumEventsForAdvertiser(stream proto.FService_GetNumEventsForAdvertiserServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		var reply proto.Response
		layout := "2006-01-02 15:04:05.00"
		startDate, err := time.Parse(layout, request.GetStartDate())
		if err != nil {
			return err
		}
		endDate, err :=  time.Parse(layout, request.GetEndDate())
		if err != nil {
			return err
		}
		eventName := request.GetEventName()
		advertiser := request.GetAdvertiser()
		rows, err := db.Table("video_metrics").Select("device as device, count(*) as event_num").Where("event_name = ? AND advertiser_id = ? AND timestamp BETWEEN ? AND ?", eventName, advertiser, startDate, endDate).Group("device").Rows()
		if err != nil {
			return err
		}
		var res ResultAdvertiserEventNum
		
		for rows.Next() {	
			db.ScanRows(rows, &res)
			reply.Device = res.Device
			reply.EventNum = res.EventNum
			if err := stream.Send(&reply); err != nil {
				return err
			}
		}
	}
}
