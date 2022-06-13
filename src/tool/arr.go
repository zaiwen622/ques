package tool

// IsInArrayInt 判断是否在数组中
func IsInArrayInt(arr []int32,value int32)bool  {
	//转为结构体
	data :=make(map[any]struct{})
	for _,val := range arr {
		data[val]= struct{}{}
	}

	if _,ok :=data[value];ok{
		return true
	}
	return false
}
