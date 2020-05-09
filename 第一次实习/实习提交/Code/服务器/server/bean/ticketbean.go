package bean

import "fmt"

// FlightType 定义航班类型
type FlightType int64
const (
	// 不能使用下划线命名法，使用驼峰命名法
	demosticType      FlightType = 0
	internationalType FlightType = 1
)
// PassengersType 定义乘客类型
type PassengersType int64
const (
	// 不能使用下划线命名法，使用驼峰命名法
	adultChild PassengersType = 1
	baby       PassengersType = 2
)

// ClassType 舱位
type ClassType int64

const (
	// 不能使用下划线命名法，使用驼峰命名法
	firstClass  ClassType = 3
	secendClass ClassType = 2
	thirdClass  ClassType = 1
)

// HasCard :是否有会员卡
type HasCard int64

const (
	// 不能使用下划线命名法，使用驼峰命名法
	noCard                 HasCard = 1
	platinumGoldSliverCard HasCard = 2
	goldSliverCard         HasCard = 3
	starGoldCard           HasCard = 4
)

// OriginTargetArea ：定义国际飞行区域的结构体
type OriginTargetArea int64

const (
	origin1 OriginTargetArea = 1
	origin2 OriginTargetArea = 2
	origin3 OriginTargetArea = 3
	origin4 OriginTargetArea = 4
	origin5 OriginTargetArea = 5
)

// Number : 乘坐次数
type Number string

const (
	justGo Number = "go"
	goBack Number = "goAndBack"
	many   Number = "many"
)

// NewTicket : new一个Ticket对象
func NewTicket(ft FlightType, pt PassengersType, ct ClassType, hc HasCard, num Number, tfc int64, fp int64, lp int64, x []int64, y []int64, z []int64, weight []int64) *ticket {
	return &ticket{ft, pt, ct, hc, num, tfc, fp, lp, x, y, z, weight}
}

// NewOutTicket : new一个Ticket对象
func NewOutTicket(ft FlightType, pt PassengersType, ct ClassType, ota OriginTargetArea, num Number, tfc int64, fp int64, lp int64, x []int64, y []int64, z []int64, weight []int64) *outticket {
	return &outticket{ft, pt, ct, ota, num, tfc, fp, lp, x, y, z, weight}
}

// 国内航班行李票价计算结构体
type ticket struct {
	ft     FlightType
	pt     PassengersType
	ct     ClassType
	hc     HasCard
	num    Number
	tfc    int64
	fp     int64
	lc     int64
	x      []int64
	y      []int64
	z      []int64
	weight []int64
}

// 国际航班行李票价计算结构体
type outticket struct {
	ft     FlightType
	pt     PassengersType
	ct     ClassType
	ota    OriginTargetArea
	num    Number
	tfc    int64
	fp     int64
	lc     int64
	x      []int64
	y      []int64
	z      []int64
	weight []int64
}

// Calclate1 : 计算票价
func (t *ticket) Calclate1() (ticketPrice int64) {
	ticketPrice = 0
	// 首先判断行李尺寸有没有超
	// 行李个数
	var length = int(t.lc)
	fmt.Println("行李数为", length)
	if length == 1 && t.x[0] <= 100 && t.y[0] <= 60 && t.z[0] <= 40 && t.x[0]+t.y[0]+t.z[0] <= 203 {
		fmt.Println("第一件行李尺寸允许")
	} else if length != 1 && t.x[0] <= 100 && t.y[0] <= 60 && t.z[0] <= 40 && t.x[0]+t.y[0]+t.z[0] <= 203 {
		fmt.Println("第一件行李尺寸允许")
		if t.x[1] <= 100 && t.y[1] <= 60 && t.z[1] <= 40 && t.x[1]+t.y[1]+t.z[1] <= 203 {
			fmt.Println("第2件行李尺寸允许")
		} else {
			if t.hc == platinumGoldSliverCard {
				if t.weight[1] < 30 {
					ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
				}
			} else if t.hc == goldSliverCard {
				if t.weight[1] < 20 {
					ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
				}
			} else if t.hc == starGoldCard {
				if t.weight[1] < 10 {
					ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
				}
			}

		}
	} else if length != 1 && !(t.x[0] <= 100 && t.y[0] <= 60 && t.z[0] <= 40 && t.x[0]+t.y[0]+t.z[0] <= 203) {

		if t.x[1] <= 100 && t.y[1] <= 60 && t.z[1] <= 40 && t.x[1]+t.y[1]+t.z[1] <= 203 {
			if t.ct == firstClass {
				if t.weight[0] <= 40 {
					ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
				}
			} else if t.ct == secendClass {
				if t.weight[0] <= 30 {
					ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
				}
			} else if t.ct == thirdClass {
				if t.weight[0] <= 20 {
					ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
				}
			}
		} else {
			if t.ct == firstClass {
				if t.weight[0] <= 40 {
					if t.weight[1] <= 30 {
						ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
					}
				} else {
					if t.weight[1] <= 30 {
						ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
					}
				}
			} else if t.ct == secendClass {
				if t.weight[0] <= 30 {
					if t.weight[1] <= 20 {
						ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
					}
				} else {
					if t.weight[1] <= 20 {
						ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
					}
				}
			} else if t.ct == thirdClass {
				if t.weight[0] <= 20 {
					if t.weight[1] <= 10 {
						ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
					}
				} else {
					if t.weight[1] <= 10 {
						ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
					}
				}
			}
		}

	} else {
		fmt.Println("第一件尺寸超了")
		if t.ct == firstClass {
			fmt.Println("头等舱")
			if t.weight[0] <= 40 {
				ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
			}
		} else if t.ct == secendClass {
			fmt.Println("商务舱")
			if t.weight[0] <= 30 {
				ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
			}
		} else if t.ct == thirdClass {
			fmt.Println("经济舱")
			if t.weight[0] <= 20 {
				ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
			}
		}
	}
	// 判断重量
	if t.ft == demosticType { // 国内航班
		fmt.Println("国内航班")
		if t.pt == adultChild { //成人儿童票
			fmt.Println("成人票")
			if t.ct == firstClass { //头等舱
				fmt.Println("头等舱")
				switch t.hc {
				case noCard: // 无VIP卡
					if length == 0 {
						ticketPrice = 0
						return
					} else if length == 1 {
						if t.weight[0] <= 40 {
							return
						} else if t.weight[0] > 40 {
							ticketPrice = int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 40 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else {
							ticketPrice += int64(float32(t.fp)*0.015) * (lastLuggage + t.weight[0])

						}

					}
				case platinumGoldSliverCard: //有金金金卡
					fmt.Println("金金金卡")
					if length == 0 {
						ticketPrice = 0
						return
					} else if length == 1 {
						if t.weight[0] <= 40 {
							return
						} else {
							ticketPrice = int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {

						if t.weight[0] <= 40 && t.weight[1] <= 30 {
							return
						} else if t.weight[0] <= 40 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 40 && t.weight[1] <= 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 40 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 40 && t.weight[1] <= 30 {
							ticketPrice = int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 40 && t.weight[1] > 30 {
							ticketPrice = int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 40 && t.weight[1] <= 30 {
							ticketPrice = int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 40 && t.weight[1] > 30 {
							ticketPrice = int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}
				case goldSliverCard: //有金银卡
					fmt.Println("金银卡卡")
					if length == 0 {
						ticketPrice = 0
						return
					} else if length == 1 {
						if t.weight[0] <= 30 {
							return
						} else {
							ticketPrice = int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {

						if t.weight[0] <= 40 && t.weight[1] <= 20 {
							return
						} else if t.weight[0] <= 40 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 40 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 40 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 40 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 40 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 40 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 40 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}
				case starGoldCard: // 有星星金卡
					fmt.Println("星星金卡卡")

					if length == 0 {
						ticketPrice = 0
						return
					} else if length == 1 {
						if t.weight[0] <= 40 {
							return
						} else {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {

						if t.weight[0] <= 40 && t.weight[1] <= 10 {
							return
						} else if t.weight[0] <= 40 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 40 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 40 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 40 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 40 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 40 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 40 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}

				}

			} else if t.ct == secendClass { //商务舱
				fmt.Println("商务藏")
				switch t.hc {
				case noCard: // 无VIP卡
					fmt.Println("没有卡")
					for i := 0; i < length; i++ {
						if length == 0 {
							ticketPrice = 0
							return
						} else if length == 1 {
							if t.weight[0] <= 30 {
								return
							} else {
								ticketPrice += int64(float32(t.fp) * 0.015)
								return
							}
						} else {
							ticketPrice += int64(float32(t.fp) * 0.015)
							return
						}
					}
				case platinumGoldSliverCard: // 金金金卡
					fmt.Println("金金金卡")
					var length = len(t.x)
					if length == 0 {
						ticketPrice = 0
					} else if length == 1 {
						if t.weight[0] <= 30 {
							return
						} else {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {

						if t.weight[0] <= 30 && t.weight[1] <= 30 {
							return
						} else if t.weight[0] <= 30 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 30 && t.weight[1] <= 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 30 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 30 && t.weight[1] <= 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 30 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 30 && t.weight[1] <= 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 30 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}
				case goldSliverCard: //金金卡银卡
					fmt.Println("金金银卡")
					var length = len(t.x)
					if length == 0 {
						ticketPrice = 0
					} else if length == 1 {
						if t.weight[0] <= 30 {
							return
						} else {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {

						if t.weight[0] <= 30 && t.weight[1] <= 20 {
							return
						} else if t.weight[0] <= 30 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 30 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 30 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 30 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 30 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 30 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 30 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}
				case starGoldCard: //星星金卡
					fmt.Println("星星卡")
					var length = len(t.x)
					if length == 0 {
						ticketPrice = 0
					} else if length == 1 {
						if t.weight[0] <= 30 {
							return
						} else {
							ticketPrice = +int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {

						if t.weight[0] <= 30 && t.weight[1] <= 10 {
							return
						} else if t.weight[0] <= 30 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 30 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 30 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 30 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 30 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 30 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 30 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}

				}

			} else if t.ct == thirdClass { //经济舱
				fmt.Println("经济舱")
				switch t.hc {
				case noCard: // 无VIP卡
					fmt.Println("没卡")
					var length = len(t.x)
					if length == 0 {
						ticketPrice = 0
					} else if length == 1 {
						if t.weight[0] <= 20 {
							fmt.Println("重量没超")
							return
						} else {
							fmt.Println("重量超了")
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else {
						fmt.Println("两件行李以上")
						ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
						return
					}

				case platinumGoldSliverCard: // 金金金卡
					fmt.Println("金金金卡")
					var length = len(t.x)
					if length == 0 {
						fmt.Println("没有行李")
						ticketPrice = 0
					} else if length == 1 {
						fmt.Println("一个行李")
						if t.weight[0] <= 20 {
							fmt.Println("重量小于20 ")
							return
						} else {
							fmt.Println("重量大于20")
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {
						fmt.Println("两件行李")
						if t.weight[0] <= 20 && t.weight[1] <= 30 {
							fmt.Println("1重量小于20，2重量小于30")
							return
						} else if t.weight[0] <= 20 && t.weight[1] > 30 {
								fmt.Println("1重量小于20，2大于30")
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 20 && t.weight[1] <= 30 {
							fmt.Println("1重量大于20，2小于30")
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 20 && t.weight[1] > 30 {
							fmt.Println("1重量大于20，2大于30")
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						fmt.Println("三件以上行李")
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 20 && t.weight[1] <= 30 {
							fmt.Println("1重量小于20，2小于30")
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 20 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 20 && t.weight[1] <= 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 20 && t.weight[1] > 30 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}
				case goldSliverCard: //金金卡银卡
					fmt.Println("金金银卡")
					var length = len(t.x)
					if length == 0 {
						fmt.Println("没有行李")
						ticketPrice = 0
					} else if length == 1 {
						fmt.Println("一件行李")
						if t.weight[0] <= 20 {
							fmt.Println("重量没超20")
							return
						} else {
							fmt.Println("重量超20了")
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {
						fmt.Println("两个行李")
						if t.weight[0] <= 20 && t.weight[1] <= 20 {
							return
						} else if t.weight[0] <= 20 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 20 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 20 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						fmt.Println("多个行李")
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 20 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 20 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 20 && t.weight[1] <= 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 20 && t.weight[1] > 20 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}

					}
				case starGoldCard: //星星金卡
					fmt.Println("仅仅星卡")
					var length = len(t.x)
					if length == 0 {
						fmt.Println("没有行李")
						ticketPrice = 0
					} else if length == 1 {
						if t.weight[0] <= 20 {
							return
						} else {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						}
					} else if length == 2 {

						if t.weight[0] <= 20 && t.weight[1] <= 10 {
							return
						} else if t.weight[0] <= 20 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[1]
							return
						} else if t.weight[0] >= 20 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * t.weight[0]
							return
						} else if t.weight[0] > 20 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1])
							return
						}

					} else {
						var lastLuggage int64 = 0
						for i := 2; i < length; i++ {
							lastLuggage += t.weight[i]
						}
						if t.weight[0] <= 20 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * lastLuggage
							return
						} else if t.weight[0] <= 20 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[1] + lastLuggage)
							return
						} else if t.weight[0] >= 20 && t.weight[1] <= 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
							return
						} else if t.weight[0] > 20 && t.weight[1] > 10 {
							ticketPrice += int64(float32(t.fp)*0.015) * (t.weight[0] + t.weight[1] + lastLuggage)
							return
						}
					}
				}
			}

		} else if t.pt == baby {
			fmt.Println("儿童票")
			var lastLuggage int64 = 0
			for i := 1; i < length; i++ {
				lastLuggage += t.weight[i]
			}
			if length == 1 {
				if t.weight[0] <= 10 {
					return
				} else {
					ticketPrice = int64(float32(t.fp)*0.015) * t.weight[0]
					return
				}
			} else {
				if t.weight[0] <= 10 {
					ticketPrice = 0
					return
				} else {
					ticketPrice = int64(float32(t.fp)*0.015) * (t.weight[0] + lastLuggage)
					return
				}
			}

		}
	} else {
		fmt.Println("flightType error!")
	}
	return
}

// Calclate : 计算票价
func (t *outticket) Calclate2() (ticketPrice int64) {

	// ft     FlightType
	// pt     PassengersType
	// ct     ClassType
	// ota    OriginTargetArea
	// num    Number
	// tfc    int64
	// fp     int64
	// lc     int64
	// x      []int64
	// y      []int64
	// z      []int64
	// weight []int64

	ticketPrice = 0
	var length = int(t.lc)
	if t.pt == adultChild {
		fmt.Println("成人/儿童")
		if t.ota == origin1 {
			fmt.Println("区域一")
			if length == 0 {
				fmt.Println("没有行李")
				return
			} else if length == 1 {
				fmt.Println("有一个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 {
							ticketPrice = 0
							return
						} else {
							ticketPrice = 980
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 {
							ticketPrice = 980
							// 150MY
							return
						} else {
							ticketPrice = 1400
							// 220my
							return
						}
					}
				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 0
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 28 {
							ticketPrice = 380
							// 150my
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							ticketPrice = 980
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 980
							// 150MY
							return
						} else {
							ticketPrice = 1400
							// 220my
							return
						}
					}

				} else {
					fmt.Println("B区域的经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 0
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 28 {
							ticketPrice = 380
							// 150my
							return
						} else if t.weight[0] > 28 && t.weight[0] <= 32 {
							ticketPrice = 980
							// 150my
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							ticketPrice = 1400
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 980
							// 150MY
							return
						} else {
							ticketPrice = 1400
							// 220my
							return
						}
					}

				}
			} else if length == 2 {
				fmt.Println("两个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							ticketPrice = 0
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件")
							ticketPrice = 980
							return

						} else {
							fmt.Println("超出两件")
							ticketPrice = 980 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							ticketPrice = 980
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件叶超长度")
							ticketPrice = 1400
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 1400 * 2
							return
						}
					}

				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 380
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 980
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 980 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 980
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 1400
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 1400
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 1400 * 2
							return
						}
					}

				} else {
					fmt.Println("B类经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.x[0]+t.y[0]+t.z[0] <= 158 {
							fmt.Println("长度没超")
							if t.weight[0] <= 23 && t.weight[1] <= 23 {
								return
							} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
								fmt.Println("超出一件")
								ticketPrice = 380
								return
							} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
								fmt.Println("超出一件")
								ticketPrice = 980
								return
							} else {
								fmt.Println("超出两件")
								ticketPrice = 980 * 2
								return
							}
						} else {
							fmt.Println("长度超了")
							if t.weight[0] <= 23 && t.weight[1] <= 23 {
								ticketPrice = 980
								return
							} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
								fmt.Println("超出一件")
								ticketPrice = 1400
								return
							} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
								fmt.Println("超出一件")
								ticketPrice = 1400
								return
							} else {
								fmt.Println("超出两件")
								ticketPrice = 1400 * 2
								return
							}
						}
					}
				}

			} else {
				fmt.Println("超出三件")
				ticketPrice = int64(3000 * length)

			}
		} else if t.ota == origin1 {
			fmt.Println("区域二")
			if length == 0 {
				fmt.Println("没有行李")
				return
			} else if length == 1 {
				fmt.Println("有一个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 {
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							if t.weight[0] <= 64 {
								return
							} else {
								ticketPrice = 1
							}
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 {
							ticketPrice = 690
							// 150MY
							return
						} else {
							ticketPrice = 1100
							// 220my
							return
						}
					}
				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 280
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 32 {
							ticketPrice = 690
							// 150my
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							if t.weight[0] <= 46 {
								return
							} else {
								ticketPrice = 1
								return
							}
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 690
							// 150MY
							return
						} else {
							ticketPrice = 1100
							// 220my
							return
						}
					}

				} else {
					fmt.Println("B区域的经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 280
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 32 {
							ticketPrice = 690
							// 150my
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							if t.weight[0] <= 46 {
								return
							} else {
								ticketPrice = 1
								return
							}
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 690
							// 150MY
							return
						} else {
							ticketPrice = 1100
							// 220my
							return
						}
					}

				}
			} else if length == 2 {
				fmt.Println("两个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件")
							ticketPrice = 690
							return

						} else {
							fmt.Println("超出两件")
							ticketPrice = 1100
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							ticketPrice = 690
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件叶超长度")
							ticketPrice = 690
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 1100
							return
						}
					}

				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 280
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 690
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 690 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 690
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 1100
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 1100
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 2200
							return
						}
					}

				} else {
					fmt.Println("B类经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 280
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 690
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 690 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 690
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 1100
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 1100
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 2200
							return
						}
					}

				}
			} else {
				fmt.Println("三件")
				ticketPrice = int64(1590 * length)
			}
		} else if t.ota == origin3 {
			fmt.Println("区域3")
			if length == 0 {
				return
			} else if length == 1 {
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("没超长")
						if t.weight[0] <= 32 {
							fmt.Println("没超重")
							return
						} else {
							ticketPrice = 520
							return
						}
					} else {
						fmt.Println("超长")
						if t.weight[0] <= 32 {
							fmt.Println("没超重")
							ticketPrice = 520
							return
						} else {
							ticketPrice = 520
							return
						}
					}

				} else if t.ct == firstClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("没超长")
						if t.weight[0] <= 23 {
							fmt.Println("没超重")
							return
						} else if t.weight[0] > 23 {
							ticketPrice = 520
							return
						}
					} else {
						fmt.Println("超长")
						if t.weight[0] <= 23 {
							fmt.Println("没超重")
							ticketPrice = 520
							return
						} else {
							ticketPrice = 520
							return
						}
					}

				} else {
					fmt.Println("A区与经济舱")
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("没超长")
						if t.weight[0] <= 23 {
							fmt.Println("没超重")
							return
						} else if t.weight[0] > 23 {
							ticketPrice = 520
							return
						}
					} else {
						fmt.Println("超长")
						if t.weight[0] <= 23 {
							fmt.Println("没超重")
							ticketPrice = 520
							return
						} else {
							ticketPrice = 520
							return
						}
					}
				}

			} else if length == 2 {
				fmt.Println("两个行李的时候")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 && t.x[1]+t.y[1]+t.z[1] <= 158 {
						fmt.Println("没超长")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							fmt.Println("没超重")
							return
						} else {
							ticketPrice = 520
							return
						}
					} else {
						fmt.Println("超长")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							fmt.Println("没超重")
							ticketPrice = 520
							return
						} else {
							ticketPrice = 520
							return
						}
					}

				} else if t.ct == firstClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 && t.x[1]+t.y[1]+t.z[1] <= 158 {
						fmt.Println("没超长")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							fmt.Println("没超重")
							return
						} else if t.weight[0] > 23 || t.weight[1] > 23 {
							ticketPrice = 520
							return
						}
					} else {
						fmt.Println("超长")
						if t.weight[0] <= 23 && t.weight[1] > 23 {
							fmt.Println("没超重")
							ticketPrice = 520
							return
						} else {
							ticketPrice = 520
							return
						}
					}

				} else {
					fmt.Println("A区与经济舱")
					fmt.Println("经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 && t.x[1]+t.y[1]+t.z[1] <= 158 {
						fmt.Println("没超长")
						if t.weight[0] <= 23 {
							fmt.Println("超了一个行李")
							ticketPrice = 1170
							return
						} else if t.weight[0] > 23 {
							ticketPrice = 520 + 1170
							return
						}
					} else {
						fmt.Println("超长")
						if t.weight[0] <= 23 {
							fmt.Println("没超重")
							ticketPrice = 520 + 1170
							return
						} else {
							ticketPrice = 520 + 1170
							return
						}
					}
				}

			} else {
				if length == 3 {
					ticketPrice = 1170 * 1
				} else if length == 4 {
					ticketPrice = 1170 * 2
				} else {
					ticketPrice = int64(length * 1590)
				}

			}

		} else if t.ota == origin4 {
			fmt.Println("区域4")
			if length == 0 {
				fmt.Println("没有行李")
				ticketPrice = 0
				return
			} else if length == 1 {
				fmt.Println("有一个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 {
							ticketPrice = 0
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							ticketPrice = 1040
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 {
							ticketPrice = 1040
							// 150MY
							return
						} else {
							ticketPrice = 2050
							// 220my
							return
						}
					}
				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 0
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 28 {
							ticketPrice = 1040
							// 150my
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							ticketPrice = 2050
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 1040
							// 150MY
							return
						} else {
							ticketPrice = 2050
							// 220my
							return
						}
					}

				} else {
					fmt.Println("A区域的经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 0
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 28 {
							ticketPrice = 1040
							// 150my
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							ticketPrice = 2050
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 1040
							// 150MY
							return
						} else {
							ticketPrice = 2050
							// 220my
							return
						}
					}

				}
			} else if length == 2 {
				fmt.Println("两个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							ticketPrice = 0
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件")
							ticketPrice = 1040
							return

						} else {
							fmt.Println("超出两件")
							ticketPrice = 1040 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							ticketPrice = 1040
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件叶超长度")
							ticketPrice = 1040 + 2050
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 2050 * 2
							return
						}
					}

				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 0
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 490
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 1040
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 1040 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 1040
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 1040 + 2050
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 1040 + 2050
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 4100
							return
						}
					}

				} else {
					fmt.Println("B类经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 0
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 490
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 1040
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 1040 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 1040
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 1040 + 2050
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 1040 + 2050
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 2050 * 2
							return
						}
					}

				}

			} else {
				fmt.Println("三件")
				if length == 3 {
					ticketPrice = 1380 * 1
				} else if length == 4 {
					ticketPrice = 1380 * 2
				} else {
					ticketPrice = int64(length * 1590)
				}

			}
		} else {
			fmt.Println("区域5")
			if length == 0 {
				fmt.Println("没有行李")
				return
			} else if length == 1 {
				fmt.Println("有一个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 {
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							if t.weight[0] <= 64 {
								return
							} else {
								ticketPrice = 1
							}
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 {
							ticketPrice = 520
							// 150MY
							return
						} else {
							ticketPrice = 830
							// 220my
							return
						}
					}
				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 0
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 32 {
							ticketPrice = 210
							// 150my
							return
						} else {
							// 因为一个大于32的可以拆分成两个小的，只要满足免费两个就行了
							ticketPrice = 520
							return

						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 520
							// 150MY
							return
						} else {
							ticketPrice = 830
							// 220my
							return
						}
					}

				} else {
					fmt.Println("A区域的经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 {
							ticketPrice = 0
							// 60my
							return
						} else if t.weight[0] > 23 && t.weight[0] <= 32 {
							ticketPrice = 210
							// 150my
							return
						} else {
							ticketPrice = 520
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 {
							fmt.Println("重量没超")
							ticketPrice = 520
							// 150MY
							return
						} else {
							ticketPrice = 830
							// 220my
							return
						}
					}

				}
			} else if length == 2 {
				fmt.Println("两个行李")
				if t.ct == firstClass {
					fmt.Println("头等舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件")
							ticketPrice = 210
							return

						} else {
							fmt.Println("超出两件")
							ticketPrice = 420
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 32 && t.weight[1] <= 32 {
							ticketPrice = 520
							return
						} else if (t.weight[0] > 32 && t.weight[1] <= 32) || (t.weight[0] <= 32 && t.weight[1] > 32) {
							fmt.Println("超出一件叶超长度")
							ticketPrice = 520 + 830
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 1660
							return
						}
					}

				} else if t.ct == secendClass {
					fmt.Println("悦享经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 210
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 520
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 520 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 520
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 830
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 830
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 830 * 2
							return
						}
					}

				} else {
					fmt.Println("B类经济舱")
					if t.x[0]+t.y[0]+t.z[0] <= 158 {
						fmt.Println("长度没超")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 0
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 210
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 520
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 520 * 2
							return
						}
					} else {
						fmt.Println("长度超了")
						if t.weight[0] <= 23 && t.weight[1] <= 23 {
							ticketPrice = 520
							return
						} else if (t.weight[0] <= 28 && t.weight[0] > 23 && t.weight[1] <= 23) || (t.weight[0] <= 23 && t.weight[1] > 23 && t.weight[1] <= 28) {
							fmt.Println("超出一件")
							ticketPrice = 520 + 830
							return
						} else if (t.weight[0] <= 32 && t.weight[0] > 28 && t.weight[1] <= 28) || (t.weight[0] <= 28 && t.weight[1] > 28 && t.weight[1] <= 32) {
							fmt.Println("超出一件")
							ticketPrice = 520 + 830
							return
						} else {
							fmt.Println("超出两件")
							ticketPrice = 830 * 2
							return
						}
					}

				}

			} else {
				fmt.Println("三件")
				if length == 3 {
					ticketPrice = 1380 * 1
				} else if length == 4 {
					ticketPrice = 1380 * 2
				} else {
					ticketPrice = int64(length * 1590)
				}
			}
		}
	} else {
		fmt.Println("婴儿票")
		return
	}

	return
}
