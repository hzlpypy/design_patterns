package custom

import "github.com/hzlpypy/design_patterns/observer_mode/weather_api"

// 自己的实现

// 天气展示struct
type weatherDisplayService struct {
	*weather_api.WeatherData
	DisplayWayFuncs []func(weather_api.WeatherData) // 所要展示天气的方式
	NotifyBehaviorInterface
}

// 天气展示行为接口
type NotifyBehaviorInterface interface {
	show(fs []func(weather_api.WeatherData))
}

// 注册订阅
func (w *weatherDisplayService) RegisterDisplayWayFuncs(f func(weather_api.WeatherData)) {
	w.DisplayWayFuncs = append(w.DisplayWayFuncs, f)
}

func (w *weatherDisplayService) PerformShow()  {
	w.NotifyBehaviorInterface.show(w.DisplayWayFuncs)
}

// 发送给消费者（向消费者展示天气信息）
func (w *weatherDisplayService) show(fs []func(ws weather_api.WeatherData))  {
	for _,f := range fs{
		f(*w.WeatherData)
	}
}

func (w *weatherDisplayService) SetNotifyBehaviorInterface(n NotifyBehaviorInterface)  {
	w.NotifyBehaviorInterface = n
}

// New
func NewWDS() *weatherDisplayService  {
	humidity := weather_api.GetHumidity()
	pressure := weather_api.GetPressure()
	temperature := weather_api.GetTemperature()
	return &weatherDisplayService{
		WeatherData: &weather_api.WeatherData{humidity,pressure,temperature},
	}
}