package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func getTextResultByHour(name string, hour int) string {
	if hour>=0 && hour<5 {
		return fmt.Sprintf("Доброй ночи, %s", name)
	} else if hour>=5 && hour<10 {
		return fmt.Sprintf("Доброе утро, %s",name)
	} else if hour>=10 && hour<17{
		return fmt.Sprintf("Добрый день, %s",name)
	} else if hour>=17 && hour<=23{
		return fmt.Sprintf("Добрый вечер, %s",name)
	} else {
		panic("Unknown hour")
	}
}

func main() {
	hour:=time.Now().Hour()
	fmt.Println("Current hour - ",hour)
	resp, err := http.Get("http://worldtimeapi.org/api/timezone/Europe/London")
	if err == nil {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		fmt.Println(getTextResultByHour("kristina",hour))
		loc, _ := time.LoadLocation(result["timezone"].(string))
		time.LoadLocation(result["timezone"].(string))
		fmt.Println("London hour - ",time.Now().In(loc).Hour())
		return
	} else {
		panic("Unknown error")
	}
	defer resp.Body.Close()
}


