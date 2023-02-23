package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Rates struct {
	EUR float64
	USD float64
}

type Data struct {
	amount string
	base   string
	date   string
	rates  map[string]interface{}
}

func get_api(url string) map[string]interface{} {
	//api 호출
	//ch 전송
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var mapdata map[string]interface{}           // JSON 문서의 데이터를 저장할 공간을 맵으로 선언
	err = json.Unmarshal([]byte(data), &mapdata) // doc를 바이트 슬라이스로 변환하여 넣고, data의 포인터를 넣어줌
	if err != nil {
		panic(err)
	}
	return mapdata
}
func get_usd_rate(ratesdata map[string]interface{}) float64 {
	rates := ratesdata["rates"]
	//어설션을 하는 이유 => 코드만으로는 외부 데이터 타입을 확정할 수 없기 때문에
	ratemap, ok := rates.(map[string]interface{}) //map 어설션
	if !ok {
		panic("type error")
	}
	usdRate, ok := ratemap["USD"].(float64) //float 어설션
	if !ok {
		panic("type error")
	}
	return usdRate
}
func job(ch *chan float64, wait *sync.WaitGroup, count *int, n int) {
	defer wait.Done()
	result := get_api("https://api.frankfurter.app/latest?from=KRW")
	usdrate := get_usd_rate(result)
	*ch <- usdrate
	*count += 1
	if *count == n {
		close(*ch)
	}

}

func main() {
	ch := make(chan float64)
	var wg sync.WaitGroup

	n := 10
	rate_sum := 0.0
	count := 0

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go job(&ch, &wg, &count, n)
	}

	for i := range ch {
		rate_sum += i
	}
	wg.Wait()
	fmt.Println(rate_sum / float64(n))
}
