package main

// server端
import (
	"fmt"
	// "net/url"
	"encoding/json"
	"net/http"
    "encoding/csv"
	// "io"
    "os"
	// "time"
	// "io/ioutil"
	"go_code/project22/bean"
	"strconv"
)

// calculate1:获得国内航班行李数据
func routing1(w http.ResponseWriter, req *http.Request) {
	fmt.Println("routing1 is running...")
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	// http://127.0.0.1:5500/getPrice?flight_type=0&passengers_type=1&class_type=3&has_card=1&number=go&the_fly_count=1&flight_price=&luggage-count=1&x=&y=&z=&weight=
	// http://127.0.0.1:5500/getPrice?flight_type=0&passengers_type=2&class_type=3&has_card=1&number=many&the_fly_count=1&flight_price=1000&luggage-count=1&x=1&y=&z=1&weight=1
	// 获取form表单中的数据
	flightType, _ := req.Form["flight_type"]
	passengersType, _ := req.Form["passengers_type"]
	classType, _ := req.Form["class_type"]
	hasCard, _ := req.Form["has_card"]
	number, _ := req.Form["number"]
	theFlyCount, _ := req.Form["the_fly_count"]
	flightPrice, found7 := req.Form["flight_price"]
	luggageCount, _ := req.Form["luggage_count"]
	X, _ := req.Form["x"]
	Y, _ := req.Form["y"]
	Z, _ := req.Form["z"]
	Weight, _ := req.Form["weight"]
	// http://127.0.0.1:5500/getPrice?flight_type=0&passengers_type=1&class_type=3&has_card=1&number=go&the_fly_count=1&flight_price=&luggage-count=1&x=&y=&z=&weight=
	// 输出显示
	fmt.Println(req)
	fmt.Printf("%T ,%s\n", flightType, flightType)
	fmt.Println(passengersType)
	fmt.Println(classType)
	fmt.Println(hasCard)
	fmt.Println(number)
	fmt.Println(theFlyCount)
	fmt.Println(luggageCount)
	fmt.Println(flightPrice)
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
	fmt.Println(Weight)
	if !(found7) {
		fmt.Fprint(w, "不好意思，机票费用不能为空")
		return
	}

	// 格式转换 :[]string转[]int
	// string转int
	// 10进制，int64
	ft, _ := strconv.ParseInt(flightType[0], 10, 0)
	pt, _ := strconv.ParseInt(passengersType[0], 10, 0)
	ct, _ := strconv.ParseInt(classType[0], 10, 0)
	hc, _ := strconv.ParseInt(hasCard[0], 10, 0)
	tfc, _ := strconv.ParseInt(theFlyCount[0], 10, 0)
	fp, _ := strconv.ParseInt(flightPrice[0], 10, 0)
	lc, _ := strconv.ParseInt(luggageCount[0], 10, 0)
	x := []int64{}
	y := []int64{}
	z := []int64{}
	weight := []int64{}
	for _, v := range X {
		n, _ := strconv.ParseInt(v, 10, 0)
		x = append(x, n)
	}
	for _, v := range Y {
		n, _ := strconv.ParseInt(v, 10, 0)
		y = append(y, n)
	}
	for _, v := range Z {
		n, _ := strconv.ParseInt(v, 10, 0)
		z = append(z, n)
	}
	for _, v := range Weight {
		n, _ := strconv.ParseInt(v, 10, 0)
		weight = append(weight, n)
	}
	// http://127.0.0.1:5500/getPrice?flight_type=0&passengers_type=1&class_type=1&has_card=1&number=go&the_fly_count=1&flight_price=1000&luggage-count=1&x=1&y=1&z=1&weight=1
	fmt.Println("-----------------转换后的数据为-------------------")
	// fmt.Println(ft+"-"+pt+"-"+ct+"-"+hc+"-"+num+"-"+tfc+"-"+fp+"-"+lp+"-"+x+"-"+y+"-"+z+"-"+weight)
	fmt.Printf("地区：%s\n",flightType)
    fmt.Printf("乘客类型：%d\n",pt)
	fmt.Printf("舱位：%d\n",ct)
	fmt.Printf("VIp：%d\n",hc)
	fmt.Printf("来回：%s\n",number[0])
	fmt.Printf("来回数量：%d\n",tfc)
	fmt.Printf("行李数量：%d\n",lc)
	fmt.Printf("机票价格：%d\n",fp)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(weight)
	// 获取返回值bean对象
	result := bean.NewBaseJSONBean()
	// 获取国内航班行李托运计算对象
	ticket := bean.NewTicket(bean.FlightType(ft), bean.PassengersType(pt), bean.ClassType(ct), bean.HasCard(hc), bean.Number(number[0]), tfc, fp, lc, x, y, z, weight)
	fmt.Println(ticket)
	// 处理返回数据

	result.Data = map[string]int{
       "price":int(ticket.Calclate1()),
    }
	result.Code = 200
	// if userName == "zhangsan" && password == "123456" {
	//     result.Code = 100
	//     result.Message = "登录成功"
	// } else {
	//     result.Code = 101
	//     result.Message = "用户名或密码不正确"
	// }
	//向客户端返回JSON数据
	fmt.Println(result.Data )
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

// calculate2 :获得国际航班行李数据
func routing2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("routing2 is running...")
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	fmt.Println(req)
	// getPrice2?passengers_type=1&class_type=0&origin-target-area=0&out-number=go&the_fly_count=1&out_flight_price=1000&flight_type=1&x=0&y=0&z=0&weight=0
	flightType, _ := req.Form["flight_type"]
	passengersType, _ := req.Form["passengers_type"]
	classType, _ := req.Form["class_type"]
	originTargetArea, _ := req.Form["origin_target_area"]
	num, _ := req.Form["out_number"]
	theFlyCount, _ := req.Form["the_fly_count"]
	flightPrice, _ := req.Form["out_flight_price"]
	luggageCount, _ := req.Form["luggage_count"]
	X, _ := req.Form["x"]
	Y, _ := req.Form["y"]
	Z, _ := req.Form["z"]
	Weight, _ := req.Form["weight"]
	// 格式转换 :[]string转[]int
	// string转int
	// 10进制，int64
	ft, _ := strconv.ParseInt(flightType[0], 10, 0)
	pt, _ := strconv.ParseInt(passengersType[0], 10, 0)
	ct, _ := strconv.ParseInt(classType[0], 10, 0)
	ota, _ := strconv.ParseInt(originTargetArea[0], 10, 0)
	tfc, _ := strconv.ParseInt(theFlyCount[0], 10, 0)
	fp, _ := strconv.ParseInt(flightPrice[0], 10, 0)
	lc, _ := strconv.ParseInt(luggageCount[0], 10, 0)
	x := []int64{}
	y := []int64{}
	z := []int64{}
	weight := []int64{}
	for _, v := range X {
		n, _ := strconv.ParseInt(v, 10, 0)
		x = append(x, n)
	}
	for _, v := range Y {
		n, _ := strconv.ParseInt(v, 10, 0)
		y = append(y, n)
	}
	for _, v := range Z {
		n, _ := strconv.ParseInt(v, 10, 0)
		z = append(z, n)
	}
	for _, v := range Weight {
		n, _ := strconv.ParseInt(v, 10, 0)
		weight = append(weight, n)
	}
	// http://127.0.0.1:5500/getPrice?flight_type=0&passengers_type=1&class_type=1&has_card=1&number=go&the_fly_count=1&flight_price=1000&luggage-count=1&x=1&y=1&z=1&weight=1
	fmt.Println("-----------------转换后的数据为-------------------")
	// fmt.Println(ft+"-"+pt+"-"+ct+"-"+hc+"-"+num+"-"+tfc+"-"+fp+"-"+lp+"-"+x+"-"+y+"-"+z+"-"+weight)
	fmt.Printf("%T ,%s\n", flightType, flightType)
	fmt.Println(pt)
	fmt.Println(ct)
	fmt.Println(ota)
	fmt.Println(num[0])
	fmt.Println(tfc)
	fmt.Println(lc)
	fmt.Println(fp)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
    fmt.Println(weight)
    writeCSV()
	// 获取返回值bean对象
	result := bean.NewBaseJSONBean()
	outticket := bean.NewOutTicket(bean.FlightType(ft), bean.PassengersType(pt), bean.ClassType(ct), bean.OriginTargetArea(ota), bean.Number(num[0]), tfc, fp, lc, x, y, z, weight)
	fmt.Println(outticket)
	result.Data = int(outticket.Calclate2())
	result.Code = 200
    result.Message = "购买成功！"
    fmt.Printf("价格为：")
	fmt.Println(result.Data )
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

//
func main() {
  
	// 第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	// 获得国内航班的接口，路由
	http.HandleFunc("/getPrice1", routing1)
	// 获得国际航班的接口，路由
	http.HandleFunc("/getPrice2", routing2)
	// 服务器要监听的主机地址和端口号
	err1 := http.ListenAndServe("localhost:5500", nil)
	if err1 != nil {
		fmt.Println("url的问题")
		return
	}
}


func writeCSV(){
    
    file, err := os.OpenFile("./testdata.csv", os.O_RDWR|os.O_CREATE, 0666)
    if err != nil{
        fmt.Println("读csv有问题！")
        return
    }
    file.Close()
    ss := csv.NewWriter(file)
    ss.Write([]string{"123q21e4141", "234234", "345345", "234234"})
    ss.Flush()
	// rfile, _ := os.Open("testdata.csv")
	// r := csv.NewReader(rfile)
	// strs, _ := r.Read()
	// for _, str := range strs {
	// 	fmt.Print(str, "\t")
	// }

}