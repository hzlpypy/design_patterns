package main

// 观察者模式
import (
	"fmt"
	"github.com/hzlpypy/design_patterns/observer_mode/observer"
	"github.com/hzlpypy/design_patterns/observer_mode/subject"
	"github.com/hzlpypy/design_patterns/observer_mode/weather_api"
	"time"
)

// 展示方式 observer(布告板)
type CurrentConditionDisplay struct {
	oi observer.ObserverInterface
}

func (c *CurrentConditionDisplay) SetSi(oi observer.ObserverInterface) {
	c.oi = oi
}

func NewCurrentConditionDisplay(w *subject.WeatherData) *CurrentConditionDisplay {
	ccd := &CurrentConditionDisplay{}
	ccd.SetSi(ccd)
	w.RegisterObserver("CurrentConditionDisplay", ccd)
	return ccd
}

func (c *CurrentConditionDisplay) Display(wd *weather_api.WeatherData) {
	fmt.Println(fmt.Sprintf("CurrentConditionDisplay:%v", wd))
}

type ThirdPartyDisplay struct {
	oi observer.ObserverInterface
}

func NewThirdPartyDisplay(w *subject.WeatherData) *ThirdPartyDisplay {
	tpd := &ThirdPartyDisplay{}
	tpd.SetSi(tpd)
	w.RegisterObserver("ThirdPartyDisplay", tpd)
	return tpd
}

func (t *ThirdPartyDisplay) SetSi(oi observer.ObserverInterface) {
	t.oi = oi
}

func (t *ThirdPartyDisplay) Display(wd *weather_api.WeatherData) {
	fmt.Println(fmt.Sprintf("ThirdPartyDisplay:%v", wd))
}

// 自定义的展示方式（布告板）
func display1(w weather_api.WeatherData) {
	fmt.Println(w.Humidity)
}

func display2(w weather_api.WeatherData) {
	fmt.Println(w.Pressure)
}

func main() {
	// custom 各个展示方式的的天气信息
	//var fs []func(weather_api.WeatherData)
	//fs = append(fs, display1, display2)
	//w := custom.NewWDS()
	//for _,f := range fs{
	//	w.RegisterDisplayWayFuncs(f)
	//}
	//w.SetNotifyBehaviorInterface(w)
	//w.PerformShow()
	// 声明主题
	w := subject.NewWeatherData()
	w.SetSi(w)
	// 声明观察者
	NewCurrentConditionDisplay(w)
	NewThirdPartyDisplay(w)
	go w.ReceiveMessage(&weather_api.WeatherData{Humidity: weather_api.GetHumidity()}, "CurrentConditionDisplay")
	time.Sleep(2 * time.Second)
	w.RemoveObserver("CurrentConditionDisplay")
	go w.ReceiveMessage(&weather_api.WeatherData{Humidity: weather_api.GetHumidity()}, "")
	time.Sleep(2 * time.Second)
}
