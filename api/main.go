package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/daniel-dsouza/microtest/pb"
	"google.golang.org/grpc"

	"github.com/labstack/echo"
)

func main() {
	conn, err := grpc.Dial("gcd-service:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	gcdClient := pb.NewGCDServiceClient(conn)

	e := echo.New()
	e.GET("/gcd/:a/:b", func(c echo.Context) error {
		a, err := strconv.ParseUint(c.Param("a"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad a")
		}

		b, err := strconv.ParseUint(c.Param("b"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad b")
		}

		req := &pb.GCDRequest{A: a, B: b}

		res, err := gcdClient.Compute(context.Background(), req)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, fmt.Sprint(res.Result))

	})
	e.Logger.Fatal(e.Start(":3000"))

}
