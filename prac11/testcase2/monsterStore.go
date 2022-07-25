package monster

import (
	"encoding/json"
	"os"
)

type monster struct{
	Name string
	Age int
	Skill string
} 

func (this *monster)Store(writeFilePath string)error{
	// 序列化
	data, err := json.Marshal(this)
	if err != nil{
		return err
	}
	// 保存到文件
	err = os.WriteFile(writeFilePath, data, 0664)
	return err
}

func (this *monster)Restore(ReadFilePath string)error{
	// 从文件读取
	data, err := os.ReadFile(ReadFilePath)
	if err != nil{
		return err
	}
	// 反序列化
	err = json.Unmarshal(data, this)
	return err
}