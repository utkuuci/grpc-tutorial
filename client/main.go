package main

import (
	"context"
	"fmt"
	"io"
	"weather/api"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)
	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("cities: ")
	for _, city := range resp.Items {
		fmt.Printf("\t%s: %s\n", city.GetCityCode(), city.CityName)
	}
	stream_api, err := client.QueryWeather(ctx, &api.WeatherRequest{CityCode: "tr_ank"})
	if err != nil {
		panic(err)
	}
	fmt.Println("weather in Anakara")
	for {
		msg, err := stream_api.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\t tempature: %.2f\n", msg.GetTemperatuer())
	}
	fmt.Println("server stopped sending")
}
