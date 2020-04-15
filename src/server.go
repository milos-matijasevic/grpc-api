package main

import (
	"context"
	"./proto"
	"net"
	//"time"
	"fmt"
	
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
	Timestamp string `gorm:"type:timestamp without time zone;not null"`
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
  Total int32
}

type ResultAdvertiserEventNum struct {
  Device  string
  Total int32
}

var db *gorm.DB
var err error

func main(){
	db, err = gorm.Open("sqlite3", "./video_metrics.sqlite3?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	fmt.Println(db.HasTable("video_metrics"))
	db.AutoMigrate(&Video_metric{})
	
	//fmt.Println(db.HasTable(&Video_metric{}))
	//var metric Video_metric
	//db.First(&metric, "device = ?", "mobile")
	//fmt.Println(metric)
	
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

func (s *server)GetMoneyForEachPublisher(ctx context.Context, request *proto.Request)(*proto.Response, error){
	var reply proto.Response
	startDate := request.GetStartDate()
	endDate :=  request.GetEndDate()
	rows, err := db.Table("video_metrics").Select("advertiser_id as publisher, sum(publisher_revenue) as total").Where("timestamp BETWEEN ? AND ?", startDate, endDate).Group("advertiser_id").Rows()
	if err != nil {
		return nil, err
	}
	reply.PublisherRevenueMap = make(map[int32]int32)
	var res ResultPublisherRevenue
	
	for rows.Next() {	
		db.ScanRows(rows, &res)
		reply.PublisherRevenueMap[res.Publisher] = res.Total
	}

	return &reply, nil
}

func (s *server)GetNumEventsForAdvertiser(ctx context.Context, request *proto.Request)(*proto.Response, error){
	var reply proto.Response
	startDate := request.GetStartDate()
	endDate :=  request.GetEndDate()
	eventName := request.GetEventName()
	advertiser := request.GetAdvertiser()
	rows, err := db.Table("video_metrics").Select("device as device, count(*) as total").Where("event_name = ? AND advertiser_id = ? AND timestamp BETWEEN ? AND ?", eventName, advertiser, startDate, endDate).Group("device").Rows()
	if err != nil {
		return nil, err
	}
	reply.DeviceEventNumMap = make(map[string]int32)
	var res ResultAdvertiserEventNum
	
	for rows.Next() {	
		db.ScanRows(rows, &res)
		reply.DeviceEventNumMap[res.Device] = res.Total
	}

	return &reply, nil
}




