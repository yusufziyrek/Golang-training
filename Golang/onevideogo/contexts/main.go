package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {

	// 3 saniyeden çok süren işlem olursa iptal eder
	/* context, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go getProductDetailsFromExternalServer(10)

	select {
	case productDetail := <-productChannel:
		fmt.Println("Ürün detayları getirildi !!", productDetail)
	case <-context.Done():
		fmt.Println("Timeout occurred, context cancelled !!")
	}

	fmt.Println("End of the main") */

	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "abc123")
	F1(ctx)

}

func F1(ctx context.Context) {
	fmt.Println("F1", ctx.Value("correlation-id"))
	F2(ctx)
}

func F2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	F3(ctx)
}

func F3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))
}

func getProductDetailsFromExternalServer(prodductId int64) {
	time.Sleep(10 * time.Second)

	productChannel <- Product{
		Id:   1,
		Name: "Deep Freeze",
	}

}
