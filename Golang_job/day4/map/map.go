package main

func main() {
	ages := make(map[string]int)
	ages["a"]=1
	ages["b"]=2

	//创建并且初始化
	ages1 := map[string]int{
		"a":1,
		"b":2,
	}
	ages1["c"] = 3
}
