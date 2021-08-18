package weather_api

type WeatherData struct {
	Humidity float32
	Pressure float32
	Temperature float32
}

// GetTemperature：获取温度
func GetTemperature() float32 {
	return 22.3
}

// GetTemperature：获取湿度
func GetHumidity() float32 {
	return 5.6
}

// GetPressure：获取压强
func GetPressure() float32 {
	return 3.6
}
