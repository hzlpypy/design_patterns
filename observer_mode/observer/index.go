package observer

import "github.com/hzlpypy/design_patterns/observer_mode/weather_api"

// observer 观察者

type ObserverInterface interface {
	// 展示信息
	Display(wd *weather_api.WeatherData)
	//
}
