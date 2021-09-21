// package main

// import "fmt"
// import "time"
// //ping-pong Game using concurrency 
// func pinger(pinger <-chan int,ponger chan<-int){
// 	for{
// 		//checking anything receive from pinger channel 
// 		<-pinger
// 		fmt.Println("Ping Recieved")
// 		time.Sleep(time.Second)
// 		ponger <-1
// 	}

// }

// func ponger(pinger chan <- int,ponger <-chan int){
// 	for{
// 			//checking anything receive from ponger channel 
// 		<-ponger
// 		fmt.Println("pong Recieved")
// 		time.Sleep(time.Second)
// 		pinger <-1
// 	}

// }

// func main(){
// ping:=make(chan int)
// pong:=make(chan int)
// go pinger(ping,pong)
// go ponger(ping,pong)
// ping<-1
// select {}
// }