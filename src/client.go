package main

import (
	"context"
	"./proto"
	"fmt"
	"encoding/json"
	"os"
	"bufio"
	"strings"
	
	"google.golang.org/grpc"
)

var client proto.FServiceClient


func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	c, _, _ := reader.ReadRune()
	for c == '\n' {
		c, _, _ = reader.ReadRune()
	}
	line, _ := reader.ReadString('\n')
	line = string(c) + strings.Replace(line, "\r\n", "", -1)
	return line
}
func GetMoneyForEachPublisher(){
	fmt.Println("Enter start date in format yyyy-mm-dd hh-mm-ss.SS")
	fmt.Println("Enter end date in format yyyy-mm-dd hh-mm-ss.SS")
	start := ReadLine()
	end := ReadLine()

	response, err := client.GetMoneyForEachPublisher(context.Background(), &proto.Request{StartDate: start, EndDate: end})
	if err != nil {
		panic(err)
	}
	publisherRevenueMapJson, err := json.Marshal(response.PublisherRevenueMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(publisherRevenueMapJson))
}

func GetNumEventsForAdvertiser(){
	var advertiser int32
	
	fmt.Println("Enter start date in format yyyy-mm-dd hh-mm-ss.SS")
	fmt.Println("Enter end date in format yyyy-mm-dd hh-mm-ss.SS")
	fmt.Println("Enter event name")
	fmt.Println("Advertiser id")

	start := ReadLine()
	end := ReadLine()
	eventName := ReadLine()
	fmt.Scanf("%d", &advertiser)
	var c byte
	fmt.Scanf("%c", &c)

	response, err := client.GetNumEventsForAdvertiser(context.Background(), &proto.Request{StartDate: start, EndDate: end, EventName: eventName, Advertiser: advertiser})
	if err != nil {
		panic(err)
	}
	deviceEventNumMapJson, err := json.Marshal(response.DeviceEventNumMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(deviceEventNumMapJson))
}


func main(){
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client = proto.NewFServiceClient(conn)
	
	
	var option int
	fmt.Println("Menu: ")
	fmt.Println("1. Get money for each publisher")
	fmt.Println("2. Get number events for advertiser")
	fmt.Println("3. Exit")
	fmt.Scanf("%d", &option)
	for option != 3 {
		
		switch ; option {
			case 1:GetMoneyForEachPublisher()
			case 2:GetNumEventsForAdvertiser()
			case 3: return
			default:
			
		}
		fmt.Println("Menu: ")
		fmt.Println("1. Get money for each publisher")
		fmt.Println("2. Get number events for advertiser")
		fmt.Println("3. Exit")
		fmt.Scanf("%d", &option)
	}
	
	//start := "2020-02-09 00:00:17.28"
	//end := "2020-02-09 00:01:35.04"
	
}