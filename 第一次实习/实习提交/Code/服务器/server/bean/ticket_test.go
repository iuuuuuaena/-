package bean
 
import(
	"testing"
	"reflect"
	"fmt"
)

func TestPrice1(t * testing.T){
	ft:= 0
	pt := 1
	ct := 1
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 1
	x := []int64{0}
	y := []int64{0}
	z := []int64{0}
	weight := []int64{0}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 0

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}
func TestPrice2(t * testing.T){
	ft:= 0
	pt := 1
	ct := 3
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 1
	x := []int64{100}
	y := []int64{60}
	z := []int64{40}
	weight := []int64{40}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 0

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}
func TestPrice3(t * testing.T){
	ft:= 0
	pt := 2
	ct := 1
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 1
	x := []int64{100}
	y := []int64{60}
	z := []int64{40}
	weight := []int64{40}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 600

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}
func TestPrice4(t * testing.T){
	ft:= 0
	pt := 1
	ct := 3
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 1
	x := []int64{110}
	y := []int64{60}
	z := []int64{40}
	weight := []int64{40}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 600

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}
func TestPrice5(t * testing.T){
	ft:= 0
	pt := 1
	ct := 3
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 1
	x := []int64{100}
	y := []int64{60}
	z := []int64{40}
	weight := []int64{41}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 600

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}

func TestPrice6(t * testing.T){
	ft:= 0
	pt := 1
	ct := 1
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 1
	x := []int64{120}
	y := []int64{60}
	z := []int64{40}
	weight := []int64{40}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 600

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}


func TestPrice7(t * testing.T){
	ft:= 0
	pt := 1
	ct := 3
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 2
	x := []int64{100,100}
	y := []int64{60,60}
	z := []int64{40,60}
	weight := []int64{40,30}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 0

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}


func TestPrice8(t * testing.T){
	ft:= 0
	pt := 1
	ct := 3
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 2
	x := []int64{120,100}
	y := []int64{60,60}
	z := []int64{40,60}
	weight := []int64{40,30}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 1050

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}

func TestPrice9(t * testing.T){
	ft:= 0
	pt := 1
	ct := 3
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 2
	x := []int64{100,100}
	y := []int64{60,60}
	z := []int64{40,60}
	weight := []int64{40,30}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 0

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}

func TestPrice10(t * testing.T){
	ft:= 0
	pt := 1
	ct := 3
	hc := 1
	var num  Number = justGo
	var tfc int64= 1
	var fp int64 = 1000
	var lc int64 = 2
	x := []int64{100,100}
	y := []int64{60,60}
	z := []int64{40,60}
	weight := []int64{40,30}
	ticket := NewTicket(FlightType(ft), PassengersType(pt),ClassType(ct), HasCard(hc), num, tfc, fp, lc, x, y, z, weight)

	fmt.Printf("ft:%d\tpt:%d\tct:%d\thc:%d\tnum:%s\ttcf:%d\tfp:%d\tlc:%d\tx:%v\ty:%v\tz:%v\tweight:%v\n",ft,pt,ct,hc,num,tfc,fp,lc,x,y,z,weight)
	
	get := ticket.Calclate1()
	var want int64 = 0

	// 用反射比较测试结果
	if !reflect.DeepEqual(get,want){
		// 如果不等
		t.Errorf("wang :%v  bu got :%v",want,get)
	}

}