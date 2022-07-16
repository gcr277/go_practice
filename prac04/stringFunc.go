package main

import(
	"fmt"
	"strconv"
)

func main(){
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}else{
		fmt.Println("err:",err)
	}
}
// len(string)
// strconv.Itoa(int)
// strconv.Atoi(string, error)
// []byte(string)
// string([]byte)
// strconv.FormatInt()
// strings.Contains(str, substring)
// strings.Count(str, substring)
// strings.EqualFold("abc", "ABC") == true //不区分大小写的比较
// ("abc" == "ABC") == false
// strings.Index(str, substring) //返回第一次出现子串的index
// strings.LastIndex(str, substring)
// strings.Replace(str1, str2, str3, n) str // 把str1中的str2替换为str3,
											// n指定替换几个,n=-1表示全替换，
											// 返回新串，str1没有变化
// strings.Split(str1, sep)[]string			// If sep is empty, Split splits after each sequence. 
											// If both s and sep are empty, returns an empty slice.
// strings.ToLower(str)
// strings.ToUpper(str)
// strings.TrimSpace(str)
// strings.Trim(¡¡¡Hello, Gophers!!!", "!¡)	// 去掉开头和结尾的字符集内的字符
// strings.HasPrefix("abc.jpg", "abc")
// strings.HasSuffix("abc.jpg", "jpg")