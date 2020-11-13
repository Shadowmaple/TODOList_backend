package matter

type MatterInfo struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Priority int8   `json:"priority"` // 优先级，0/1/2/3 => 无优先级/低优先级/中优先级/高优先级
	Date     string `json:"date"`     // 设定的日期：2020-11-01
	Time     string `json:"time"`     // 设定的时间："12:02"
}
