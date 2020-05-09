package bean


// BaseJSONBean ：定义一个结构体，储存get的数据
type BaseJSONBean struct { 
    Code    int         `json:"code"` 
    Data    interface{} `json:"data"` 
    Message string      `json:"message"` 
} 
//NewBaseJSONBean new一个对象返回
func NewBaseJSONBean() *BaseJSONBean { 
    return &BaseJSONBean{} 
}