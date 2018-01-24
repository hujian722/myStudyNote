package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	ctx, cancel := context.WithCancel(context.Background())
	cerr := make(chan error)
	go func() {
		cerr <- ping()
	}()
	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()
	select {
	case err := <-cerr:
		fmt.Println("err:", err)
	case <-ctx.Done():
		fmt.Println("done:")
	}
	time.Sleep(time.Second * 6)
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	cerr := make(chan error)
	go func() {
		cerr <- ping()
	}()
	select {
	case err := <-cerr:
		fmt.Println("err:", err)
	case <-ctx.Done():
		fmt.Println("done:")
	}
	time.Sleep(time.Second * 6)
}
func ping() error {
	time.Sleep(time.Second * 4)
	fmt.Println("ping over:")
	return nil
}
func Afunc(ctx context.Context, name string) {
	for {

		select {
		case <-ctx.Done():
			fmt.Println(name + " is done")
			return
		default:
			fmt.Println(name + " is running")
			time.Sleep(time.Second * 1)

		}
	}
}
func Bfunc() {
	ctx, _ := context.WithCancel(context.Background())
	//defer cancel()

	go func() {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("no startup timeout")
		case <-ctx.Done():
			fmt.Println("done")
		}
	}()
}
