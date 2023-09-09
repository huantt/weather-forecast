package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"weather_forecast/handler/collector"
	"weather_forecast/impl/weather_service/weatherapi"
	"weather_forecast/pkg/weatherapi_com"
)

func UpdateWeather(use string) *cobra.Command {
	var weatherApiComKey string
	var days int
	var weatherTemplateFilePath string

	command := &cobra.Command{
		Use: use,
		Run: func(cmd *cobra.Command, args []string) {
			weatherApiService := weatherapi.NewWeatherService(weatherapi_com.NewService(weatherApiComKey))
			handler := collector.NewCollector(weatherApiService)
			err := handler.Collect(context.Background(), "HaNoi", days, weatherTemplateFilePath)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			slog.Info("Updated weather")
		},
	}

	command.Flags().StringVarP(&weatherApiComKey, "weather-api-key", "k", "", "weatherapi.com API key")
	command.Flags().StringVarP(&weatherTemplateFilePath, "template-file", "f", "", "Readme template file path")
	command.Flags().IntVar(&days, "days", 7, "Days of forecast")
	err := command.MarkFlagRequired("weather-api-key")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return command
}
