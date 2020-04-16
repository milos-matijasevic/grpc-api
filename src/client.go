package main

import (
	"context"
	"./proto"
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	
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
	line = strings.Replace(string(c) + line, "\r", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	return line
}
func GetMoneyForEachPublisher(){
	fmt.Println("Enter start date in format yyyy-mm-dd hh-mm-ss.SS")
	fmt.Println("Enter end date in format yyyy-mm-dd hh-mm-ss.SS")
	start := ReadLine()
	end := ReadLine()
	request := &proto.Request{StartDate: start, EndDate: end}
	stream, err := client.GetMoneyForEachPublisher(context.Background(), request)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println()
	fmt.Println("Publisher : Revenue")
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(response.Publisher, " : " , response.Revenue)
	}
	fmt.Println()
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
	// To skip newline char
	fmt.Scanf("%c", &c)

	request := &proto.Request{StartDate: start, EndDate: end, EventName: eventName, Advertiser: advertiser}
	stream, err := client.GetNumEventsForAdvertiser(context.Background(), request)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println()
	fmt.Println("Device : Event count")
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(response.Device, " : " , response.EventNum)
	}
	fmt.Println()
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