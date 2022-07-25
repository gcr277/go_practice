package monster

import (
	"testing" // 引入go的test框架
)

func TestStore(t *testing.T){
	monster1 := monster{
		Name : "Bokoblin",
		Age : 2,
		Skill : "throw",
	}
	filePath := "./save"
	err := (&monster1).Store(filePath)
	if err != nil{
		t.Fatalf("存档失败:%v\n", err)
	}
}
func TestRestore(t *testing.T){
	var monster2 monster
	filePath := "./save"
	err := (&monster2).Restore(filePath)
	if err != nil{
		t.Fatalf("读档失败:%v\n", err)
	}
	t.Logf("读取数据为:%v\n", monster2)
}