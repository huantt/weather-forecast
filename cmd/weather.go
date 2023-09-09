package cmd

import (
	"context"
	"github.com/huantt/weather-forecast/handler/collector"
	"github.com/huantt/weather-forecast/impl/weather_service/weatherapi"
	"github.com/huantt/weather-forecast/pkg/weatherapi_com"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

func UpdateWeather(use string) *cobra.Command {
	var weatherApiComKey string
	var city string
	var days int
	var weatherTemplateFilePath string
	var outputFilePath string

	command := &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			weatherApiService := weatherapi.NewWeatherService(weatherapi_com.NewService(weatherApiComKey))
			handler := collector.NewCollector(weatherApiService)
			err := handler.Collect(context.Background(), city, days, weatherTemplateFilePath, outputFilePath)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			slog.Info("Updated weather")
		},
	}

	command.Flags().StringVarP(&weatherApiComKey, "weather-api-key", "k", "", "weatherapi.com API key")
	command.Flags().StringVarP(&weatherTemplateFilePath, "template-file", "f", "", "Readme template file path")
	command.Flags().StringVarP(&outputFilePath, "out-file", "o", "", "Output file path")
	command.Flags().StringVar(&city, "city", "", "City")
	command.Flags().IntVar(&days, "days", 7, "Days of forecast")
	err := command.MarkFlagRequired("weather-api-key")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("template-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("city")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = command.MarkFlagRequired("out-file")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return command
}
