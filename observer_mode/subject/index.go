package subject

import (
	"github.com/hzlpypy/design_patterns/observer_mode/observer"
	"github.com/hzlpypy/design_patterns/observer_mode/weather_api"
)

// subject 主题

type SubjectInterface interface {
	// 注册观察者
	RegisterObserver(name string, oi observer.ObserverInterface)
	// 删除观察者
	RemoveObserver(name string)
	// 通知观察者
	NotifyObserver(wd *weather_api.WeatherData, name string)
	// 接收消息传递，if name is null，will notify all observer
	ReceiveMessage(wd *weather_api.WeatherData, name string)
}

// 主题天气 weatherData
type WeatherData struct {
	si        SubjectInterface
	obServers map[string]observer.ObserverInterface
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		obServers: map[string]observer.ObserverInterface{},
	}
}

func (w *WeatherData) SetSi(si SubjectInterface) {
	w.si = si
}

func (w *WeatherData) RegisterObserver(name string, oi observer.ObserverInterface) {
	w.obServers[name] = oi
}

func (w *WeatherData) RemoveObserver(name string) {
	delete(w.obServers, name)
}

func (w *WeatherData) NotifyObserver(wd *weather_api.WeatherData, name string) {
	if name == "" {
		for _, v := range w.obServers {
			v.Display(wd)
		}
	} else {
		for k, v := range w.obServers {
			if k != name {
				continue
			}
			v.Display(wd)
		}
	}
}

func (w *WeatherData) ReceiveMessage(wd *weather_api.WeatherData, name string) {
	w.NotifyObserver(wd, name)
}
