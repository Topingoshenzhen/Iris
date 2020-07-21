package main  // 这是每个程序必须加上的包

import (     // 导入需要使用的标准库
	"fmt"
)
import分为这几种方式：
//相对路径
import "./lib/bolt"  //当前代码文件夹这个目录里的lib/bolt目录，但是不建议这么import

//绝对路径
import "log"  			//这个为常用的import

//点操作
import (
	. "fmt"  		//这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，
					// 例如：之前是fmt.Print，点操作直接Print()
)

//别名操作
import (
	f "fmt" 		//别名操作调用包函数时前缀变成了重命名的前缀，即：之前为fmt.Print，别名操作直接f.Print()
)             		//这种方法应该适合多次导入同一个包，然后用不同的别名替换

//下划线操作
import (
	_ "github.com/mysql"
)
//下滑线 “_” 操作其实只是引入该包。当导入一个包时，它所有的 init() 函数就会被执行，但有些时候并非真的需要使用这些包，
// 仅仅是希望它的 init() 函数被执行而已。这个时候就可以使用 “_” 操作引用该包了。
// 即使用 “_” 操作引用包是无法通过包名来调用包中的导出函数，而是只是为了简单的调用其 init() 函数。


func main(){    // 理解为主程序入口函数
	/*
	这个多行注释用法，println带自动换行的打印，格式化输出每打印一个变量就自动换行，(多行注释快捷键：ctrl + shift + /)
	print 单行打印，如果出现多个print，后面一个print直接覆盖前面的print内容，只打印一行
	 */
	fmt.Println("go")  // 单行注释的用法
	fmt.Print("gogogogogogogogogogogo")
}
    printf()函数一次性输出多个变量，必须加格式符才可以按照指定格式输出，换行必须要加换行符
// 当前文件夹空白的地方，shift + 右键 选择命令窗口打开 按tab选择你要运行的文件，然后直接运行
// 在当前文件夹地址栏上方 - 输入cmd - 回车 - 按tab选择你要运行的文件
// 以上两种方式是在当前文件夹通过cmd运行的两种方式，在该窗口可以tab选择文件直接运行,or 按照以下编译运行方式
// 编译：go build 文件名.go可生成同名的*.exe文件于当前文件夹
// 运行：.\文件名.exe(运行编译结果文件) or go run 文件名.go(直接运行文件)

// 变量的定义与声明：
// var 变量名 类型 / var 变量名 类型 = 值

// 输出格式："%格式",变量名  用法应该跟python差不多。%表示一个占位符，百分号里面可以加换行符号，间隔符
// 输入: Scanf

func main() {   // go难道与python一样，赋值这些只是引用，地址指针指向了值？
	var age int
	fmt.Println("请输入年龄：")   // 输入界面提示
	fmt.Scanf("%d", &age)     // 输入的值赋给变量，必须加&代表变量地址指向了输入值所在的地址
	fmt.Println("age = ", age)    // 变量的值已经指向了输入值
	fmt.Println(&age)                  //  此处打印了变量的地址
	fmt.Printf("%p", &age)    //  指针类型输出，输出了变量地址
}
//字符串本质是一个结构体 struct
//struct{
//	pointer // 指向一块连续的空间，存储所有的字符串
//	len int // 描述字符串的长度
//}
//
// print系列(与窗口有关)：输出 -- println/printf 输入 -- scanf/scan
//sprint系列(与字符串有关)：
//      s:=fmt.Sprintln("hello abc")
//
//      fprint系列(与文件输入有关)
//      // 创建文件，把数据格式到文件
//      f,_:= os.OpenFile(name:'1.txt',os.O_CREATE,perm:666)
//      defer f.close() // 把内存数据刷新到磁盘文件，关闭文件释放资源
//      fmt.Fprintf(f,"name=%s age=%d","gplang",6)
//
//      f,_ := os.Open(name:"1.txt")
//      var name string
//      var age int
//      fmt.Fsacnf(f,format:"name=%s age=%d",&name1,&age)  // 按照固定格式读取文件
//      fmt.Fsacnln(&name1,&age)



变量命名规范：跟python差不多，并且一样区分大小写.大驼峰，小驼峰命名法，下划线分隔法
基础数据类型：
		整型：int(有符号的整型，包括：正整数，负整数，0) 8位(bit)为一个字节(byte)
		32位系统中，int是四个字节，范围 -2^32 - 2^31 - 1
		64位系统中，int是八个字节，范围 -2^63 - 2^63 - 1

  		uint(无符号的整型，包括：正整数，0)
		32位系统中，int是四个字节，范围 0 - 2^32 - 1
		64位系统中，int是八个字节，范围 0 - 2^64 - 1


		浮点类型：go提供了两种精度的浮点型：float32(精确到小数点后7位)，float64(精确到小数点后15位)
				由于go中涉及到的关于数学运算的包中，都要求使用float64这个类型，建议都使用float64
				variable := 66666.66666666666666666   // 自动提到类型默认使用float64
				fmt.Printf("%T",variable)           // %T输出变量的类型，%.2f 保留两位小数并且同时进行四舍五入


		布尔类型(bool)：布尔类型取值要么是真(true)要么是假(false)，用关键字bool来进行定义
				布尔默认值为：false
				var isResult bool
				isResult = true
				fmt.Println(isResult)
				fmt.Printf("%t",isResult)  //  布尔类型只能使用%t特定的格式符输出类型
				b = bool(1)

		字符类型：用单引号引起来的单个字符为字符类型，ASCII是一套字符集
				var ch byte = 'a'


		字符串类型：双引号引起来的字符是字符串
				var variable string = "i"  // 变量只包含了一个字符，但隐藏着一个字符串的结束标志 '\0'
				fmt.Printf("%s",variable)  // %s打印出了结束标志 '\0'前面的内容
		字符串中的个数用 len(字符串)来获取，强调：go语言中一个汉字占用三个字符，为了与Linux系统保持一致


		强制类型转换：类型转换将一种数据类型强制转换成其他类型，建议低类型转换成高类型，以保证精度。
		              高类型转换成低类型导致丢失精度或者数据溢出
				var num float64 = 3.15
				fmt.Printf("%d",int(num))   // 浮点型转换成整形
				string(int) // 将数字转换成对应的ASCII码里的字符

		fmt包的格式输入输出：(Scanf需要指定格式符号，Scan(&variable)不需要)
				b% 转换成二进制
				o% 转换成八进制
				x% 转换成十六进制(大小写输出结果一样，只是分大小写)
				// 1，定义变量
				var name string
				var age   int
				// 2，给出提示录入
				fmt.Println("请输入姓名：")
				fmt.Scan(&name)
				fmt.Println("请输入年龄：")
				fmt.Scan(&age)
				// 3，完成输入
				// 4，打印输出结果
				fmt.Printf("您好，%s, 您的年龄是%d",name,age)


常量定义使用：常量的定义是通过关键字const完成的，并且建议常量名称大写，定义后常量的值不允许修改,地址不允许打印
		const 常量名(建议为大写) 类型 = 值    // 这个作为固定模板，一般不更改，:=不适用于常量定义


运算符：应该很多用法等同python
		算数运算符：+ - * / %(取余)

		自增自减预算符：++(自增加，不同于C语言有两种，GO只有i++这一种用法,使用完了再增加）
						--(自减法，不同于C语言有两种，GO只有i--这一种用法,使用完了再减法)

		赋值运算符：=，+=，-=，*=，、=，%=，python用法完全一样

		关系运算符：比较的结果是布尔类型(返回TRUE,FALSE)，用法完全等同与python

		逻辑运算符：&&与，||或，！非，运算符结果为布尔类型，逻辑运算符两边一般只能接布尔类型或者关系运算

		单目运算符：运算所需的变量为一个的运算符，即运算中只有一个操作数。双目运算符：运算中只有两个操作数
		++，--，！，&(取地址)                                                  eg.算数运算符+关系运算符

		运算符优先级：单目运算符 > 双目运算符，算数运算符 > 比较运算符 > 逻辑运算符 > 赋值运算符
		逻辑与 > 逻辑或

		eg. 判断闰年：被400整除or被4整除不能被100整除
		fmt.Println("请输入年份：")
		var year int
		fmt.Scan(&year)
		b: = year % 400 ==0 || year % 4 == 0 && year % 100 != 0

分支结构：
		if结构：
		If  表达式1 {  //表达地可以加括号，可以不加括号
			…………….
		} else if 表达式2 {
			…………….
		} else {
			…………….
		}

		switch结构： 可以与if搭配来使用,switch里面可以嵌套if，if里面可以嵌套switch
		switch 变量或表达式 {   // case全部是表达式，这里可以为空白，不写变量或表达式
				case 值(可以写),变量或表达式:         多个
					if 表达式 {
					   嵌套这个情况开发中有多种应用例子,例如 cobol
					} else if 表达式 {
						..................................
					} else {
						..................................
					}
				case 值,变量或表达式: // 每个case执行完毕直接跳出了switch
				if 表达式 {
					..................................
				} else if 表达式 {
					..................................
				} else {
					..................................
				}
				case 值,变量或表达式:
					fallthrough  // fallthrough代表直接跳到下一个case继续执行，不执行本case
				case 值,变量或表达式:
					if 表达式 {
						..................................
					} else if 表达式 {
						..................................
					} else {
						..................................
					}
				case 值,变量或表达式:
					if 表达式 {
						..................................
					} else if 表达式 {
						..................................
					} else {
						..................................
					}
				default:  // case以外的情况都走这条路
					if 表达式 {
						..................................
					}
				}


		for循环：// for  {...}   这种循环一直停不下来
		var i int
		for i=1; i<value; i++ {  // 模板
			..............................................
		}
		break, continue每个语言用法一样

		嵌套循环无非内循环，外循环，明白谁嵌套谁。

		打印九九乘法表：

		func main() {
			for i:= 1;i<=9;i++ {
				for j:=1;j<=i;j++ {
					fmt.Printf("%d*%d=%d,\t",j,i,j*i)
				}
				fmt.Print("\n")
			}
		}


函数：函数基本语法：
					func 函数名() {
						函数体
					}

go里面形参，实参如何用与python一样。go不同的是固定参数在前面，不定参数在后面，位置错了编译器直接报错

func 函数名(形参 参数数据类型) 类型必须要填。如果是多个形参，形参之间用逗号分隔

传参调用：函数名(实参)  实参的个数，类型必须与实参保持一致。如果是固定参数必须传值，如果是不定参数按照需要传值

固定参数指上面这个，不指定参数格式：func 函数名(args...int)  args是参数名，必须紧跟着...  类型必须填
其实不指定参数是个集合，通过集合来传值，下标从0开始,这里引申一下，遍历集合可以用range
for a,v := range args {  //range暂定的标准用法，返回两个返回值，下标 + 值
	fmt.Println(a,v)   // a代表集合里元素的下标，v为该下标对应的值
}
如果定义了一个变量而没有使用，编译有报错 declared not use，暂时不需要使用这个变量可用如下匿名变量,该变量丢弃数据不处理
for _,v:= range args{  // 该匿名变量格式为 _
	fmt.Println(v)
}
函数返回值：返回单个值，多个值  return 将值返回，然后定义一个变量来接收
func Result(a int,b int) int {  // last int 是指定函数返回的数据的类型
	var i int
	i = a + b
	return i
}
func main() {
	var u int
	u = Result(0,0)   //  函数返回值可以被变量接收
}
func Result(a int,b int) (i int) {  // i在这里可直接定义，这样不需要在函数里定义
	//var i int    这里因为上面已经定义了可以注释
	i = a + b
	return        //  已经指明了返回值的变量名，这里格式为return，不需要接变量名
	// return i
}
return多个返回值，接收格式为：a,b,c,r,u = result() 返回值位置，个数必须跟定义一致



标准的函数格式：
func 函数名(参数定义位置) (返回值定义位置) {  // 参数必须在这里定义,返回值可以在这里定义
	函数体  //如果返回值上面有定义，这里可不定，如果上面只有类型，这里必须定义
}

函数作用域：局部变量，全局变量的区别，go里面与python差不多，同名变量，函数里优先使用局部变量,函数体外使用全局变量

延迟调用：defer，格式：defer + 操作(方法都可以)，延迟执行这个操作。如果一个函数哟多个defer语句，它们以
LIFO(后进先出)的顺序执行，最后一个defer语句最先执行。eg. defer func(),defer可直接放在函数体前面，函数体后面必须加()

递归函数(recursion)：
var r int = 1

func main() {
	TestDemo(6)
	fmt.Println(r)
}

func TestDemo(i int) {
	if i == 1 {
		return
	}
	r *= i
	TestDemo(i - 1)
}

复合类型：数组定义：var 数组名 [元素数量] 类型 or 数组名 := [元素数量] 类型
var array[i] int = [i]int{0,......,i-1}   //整个为初始化整个数组
array := [1000000]int{100000,10000000}    //部分赋值，按顺序从下标为0的元素逐步赋值100000,10000000，其他元素都为初始值0
array := [1000000]int{1000:100, 10000:100} //指定下标初始化，只给下标为1000,10000的元素赋值，{}里前面为下标，后面为数值
如果定义了一个数组,暂时不能确定长度，通过初始化来确定该数组长度
array := [...]int{0,0} // 前面...代表长度未定义,{}里面确定了长度，不过谁会这么无聊这么设置，一般不定长度后面不接{}
数组遍历：通过for i:=0; i<len(array); i++ {}，for v,i := range array {} 遍历
数组作为函数参数：func 函数名(数组) {
						函数体
				  }
				  调用：函数名(数组)
				  eg：
				  func main() {
				  		var ber [5]int = [5]int {0,1,2,3,4}
				  		getPrint(ber)
				  		//如果下面修改了元素的值，这里打印出每个元素发现不影响原数组
				  		for i:=0;i<len(ber); i++ {
				  			fmt.Println(ber[i])
						}
				  }
				  func getPrint(array [5]int) {  // 这里参数格式为定义数组格式：数组名 [元素个数]元素类型
				  		for i:=0; i<len(ber); i++ {
				  			fmt.Println(array[i]) //打印出每个元素
							array[i] = value 	 // 函数中修改数组中的值，不影原数组
						}
				  }

二维数组规范：var array[i][u] int // i为总行数, u为总列数，想象成表通过行下标+列下标定位元素
全部初始化：  var array[i][u] int = [i][u]int{{},{}} // i为什么值，int后{}里面就有几个{},每个{}里有u个元素
部分初始化：  var array[i][u] int = [i][u]int{{},{}} // 参考数组部分初始化,int{}里的{}个数不强制等于u，未写出也有初始值
指定元素初始化：var attay[i][i] int = [i][u]int{0:{0:val,u:value}} // 等同于二维数组定义里面参考单维数组指定元素初始化
																   //  	int{}里的{}个数也不强制等于u，未写出也有初始值
如果暂时没有确定i的值，可以用...代替，u不可以代替，不然会报错下标超界
二维数组遍历：for......len()，for......range()
			len(array) 获取的是i的值
			len(array[0]) 获取的是u的值 = 制定下标行里有几个元素
			for i:=0; i<len(array); i++ {
				for o:=0; o<len(array[i]); o++ {
					fmt.Println(array[0][o])
				}
			}
			不管是for len，for range遍历都要两层循环才能打印出每一个单独的元素
			for _, v := range array {
				// 外循环打印一整行数据
				for i, data := range v {
				    // 内循环打印单个值
				}
			}

切片：应该等同于python切片？
		var 切片名[]数据类型 or 切片名:= []数据类型{} // {}里面可以加入值代表初始化

		make函数创建切片: make(切片类型,长度,容量) 长度是已经初始化的空间。容量是已经开辟的空间，
		包括已经初始化的空间与空闲的空间
		i := make([]int,3,6)  // 6代表分配了6个长度空间(6可以省略不写)，3代表切片类型[]int只占用了3个，其余没有占用
		cap()函数获取对象的容量，用法cap(),比len()更全面,cap不能用于map(因为字典没有容量),，切片，数组，channle都可以用
		append函数在尾部追加数据，格式：变量名 = append(变量名, 0,val,value) // 变量可以是列表，切片等,
		// 注意：追加的数据 必须与变量名对应的变量类型里面的元素类型保持一致(=申明的时候指定类型)，否则导致报错
		注意,go里面，变量名[下标] = value这种形式一般都可以改变值
切片初始化：
		a := []int{0,9,8,7,6,5,4,3,2,1} // 初始化
        a = append(a,10000)   //代表在最后添加10000
        a = append(a,value,value,value,value,value,value) // append初始化,如果a未赋值，等于在初始a后面追加数据
        a := make([]int,6,16)  a[下标] = value //这种初始化
        for i=0, i<len(a); i++ {
		}

切片的遍历：
a := []int{0,9,8,7,6,4,3,2,1,0}
for i=0, i<len(a); i++ {
	fmt.Println(a[i])
}
for _,v := range a {  //下标可忽略
	fmt.Println(i,v)
}

切片截取操作：go里面只有定义成切片类型，才可以截取操作？python貌似任何变量直接截取
//第一个数为截取的起始位置
//第二个数为截取的终止位置(不包括该位置)
//第三个数为用来计算容量，容量为切片中最多存储多少个元素,容量未注明，即 容量 = 长度
//容量 = 第三个数减去第一个数
//长度 = 第二个数减去第一个数
u := a[0:v:r]
u[:]  取全部
u[i:] 取从i起始位置到最后全部区间
u[:i] 取0到下标为i之间全部(不包括i)
切片截取后返回新切片，对新切片值进行修改，会影响原切片：
              u := []int{0,9,8,7,6,5,4,3,2,1,0}
              i := u[0:3] // 切片i指向了u对应的区间，并没有开辟新区间给i
              i[下标] = value // 如果对切片i值进行修改，会影响原切片u
切片扩容：append一般扩容方式为：上一次容量*2，如果超过了1024字节 每次扩容上一次的1/4
eg. u := []int{0,9,8,7,6,5,4,3,2,1,0} 用了一次append,cap(u)就直接*2
copy函数： copy(切片1,切片2)   拷贝的长度为两个切片中长度较小的长度值
copy函数将切片2的内容从切片1起始的位置开始，替换切片1这个区间的值，= 从下标0开始按照规范长度替换
eg.
	u := []int{0,9,8,7,6,5,4,3,2,1,0}
	i := []int{{0,0,0,0,0,0,0,0,0,0}}
	copy(u,i)
	fmt.Println(u) //   最后u结果为[0,0,0,0,0,0,0,0,0,0,0]
	fmt.Println(i)
切片作为函数参数：
基本语法：
		func 函数名(切片名) {
			函数体
		}
		函数名(切片)  // 在函数中修改切片的值，影响到原切片

func main() {
	u ;= []int{0,6,9,6,0} // 定义一个切片类型
	Init(u) // 切片机作为参数传递
}

func Init(参数名称 参数类型) {
	for操作打印每个元素
	如果函数里对切片元素更改了，原切片被修改
}

sort：
func main() {
	q := []int{9,8,7,6,5,4,3,2,1,0}
	var temp int
	for i:=0; i<len(q)-1; i++ {
		for u:=0; u<len(q)-1-i; u++ {
			if q[u]>q[u+1] {
				temp = q[u]
				q[u] = q[u+1]
				q[u+1] = temp
			}
		}
	}
	fmt.Println(q)
}




func main() {
	q := []int{9,8,7,6,5,4,3,2,1,0}
	for i:=0; i<len(q)-1; i++ {
		min := q[i]
		minIndex := i
		for u:=i+1; u<len(q)-1; u++{
			if min > q[u] {
				min = q[u]
				minIndex = u
			}
		}
		//
		if minIndex != i {
			q[i],q[minIndex] = q[minIndex],q[i]
		}
		//
		fmt.Println(q)
	}
}

MAP（字典结构）：GO语言里面的字典机构为键值对，并且键值对排列无序,键唯一不重复 = python

var map名称 map[键的类型]值得类型 eg. var m map[int]string

map名称 := map[键的类型]值得类型{}

map名称 := make(map[键的类型]值得类型,value,value)

var m map[int]string = map[int]string {6:"哈哈",2:"哈有"}

通过key取值，map名字[键] eg,a['键'] = value

delete(map名,键) // 通过key删除某个值 eg, delete(u,i)

通过for range 遍历字典来取值

for key,value := range a {

}

通过key取值判断字典是否存在该键值对
变量1,变量2 := map名字[键]  //如果变量1存储的值等于键对应的值，并且变量2的值为true，则代表存在

eg, var u map[int]string = map[int]string{1:"haha",2:'UUUUUU'}
    u := map[int]string{1:"haha",2:'UUUUUU'}
    value, ok := u[1]  // 变量名不唯一, ok为布尔类型，否则这个不成立,报错：non-bool used as if condition
    if ok {
    	//代表该键值对存在
    }else {
    	//该键值对不存在
    }

map作为参数传参：如果函数体里面修改了map，原map会被改变

func main() {
	var m map[int]string = map[int]string{}
	PrintMap(u)
}

func PrintMap(m map[int]string) {
	// 如果函数体里面修改了map，原map会被改变
	a['键'] = value
}

结构体：
定义：

type 结构体名 struct {
	成员名 数据类型
	...............
}

初始化：
顺序初始化 or 制定成员初始化 or 通过‘结构体变量.成员’初始化
type Student struct {
	id int
	name string
	age int
	addr string
}

var u Student = Student{}  // 结构体是一个数据类型，这里是定义变量u为Struct类型并且初始化
u := Student{} //自动推导类型进行定义初始化
var u Student //简单定义一个变量为Student类型，不进行初始化传值

eg:
func main() {
type Student struct {
name string
Id int
age int
addr string
}
//顺序初始化
//var u Student = Student{"User",10000,25,"深圳"}
u := Student{"User",10000,25,"深圳"}  //定义方式跟上面一样
fmt.Println(u)
//部分初始化
var i Student = Student{name:"UTR",age:10000} //指定成员变量初始化
fmt.Println(i)
//指定变量初始化，这种方式灵活多变使用非常多
var stu Student
stu.Id = 100
stu.name = "RRRRRRRRRRRR"
stu.age = 1000000
stu.addr = "深圳南山区"
fmt.Println(stu)
}

结构体数组定义：数组里面的元素都为结构体类型
	var 结构体数组名[下标] 结构体类型
	结构体数组名[下标].成员 = 值  // 修改结构体成员的值
	for......len()进行遍历，for......range进行遍历

type Student struct {
name string
Id int
age int
addr string
}  //这里定义在主函数外面理解为全局变量吧
func main() {
var arr [2]Student = [2]Student{    // 定义一个数组 ，元素为结构体
Student{"Usr",10000,100,"深圳"},
Student{"Urr",100,100,"深圳取"},
}   //初始化该数组
fmt.Println(arr)
fmt.Println(arr[0])
fmt.Println(arr[0].addr)
arr[0].addr = "深圳南山区" // 修改结构体成员的值
fmt.Println(arr[0].addr)
fmt.Println(arr)            // 修改结构体成员的值影响原结构体
}


结构体切片定义：切片里面的元素都为结构体类型
	//var stu := []Student{
	arr := []Student{   //这是定义切片的方法
		Student{"Usr",10000,100,"深圳"},
		Student{"Urr",100,100,"深圳取"},
	}
	其他方法参考数组操作方法，操作方法一样
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[0].addr)
	arr[0].addr = "深圳南山区"
	fmt.Println(arr[0].addr)
	fmt.Println(arr)

	arr = append(arr,Student{"UURRII",10000,10000,"深圳软件园"})
	fmt.Println(arr) // append追加数据
    // append(参数1,参数2),第二个参数如果为切片类型，格式为：append(参数1,参数2...) 带...

结构体map定义：map元素value为架构体类型

m := make(map[int]Student)   //定义map，键值对的值为结构体类型
m[1] = Student{"Urr",100,100,"深圳取"} //通过map[键] = Student{value}来初始化
m[2] = Student{"Usr",10000,100,"深圳"}
fmt.Println(m)
delete(m,2)  //删除map中的值
fmt.Println(m)

结构体作为函数参数传参:

	func 函数名(结构体){
       // 如果函数体里面修改了结构体成员的值，不改变原结构体
       //eg. 结构体.成员名 = value，原结构体的值不被改变
	}

	函数名(结构体)









指针：一种特别变量存储的是变量的地址，
定义：var 指针变量名 *类型
eg, var a int = 10
var p *int = &a  // 定义一个指针类型指向了变量a
fmt.Println(*p)  // 打印出*p的值为10，证明可以通过指针间接访问变量
*p = 222         // 通过指针间接修改变量对应的内存空间
fmt.Println("a=",a) // a得值被修改为了a

var p *int
fmt.Println(p)  \\ 打印出nil，p指向了内存地址编号为0的区间

p := new(int) // new(数据类型) 函数 创建开辟一个内存区域,,不需要管理空间的释放
*p = 10000   // 通过指针间接修改变量对应的内存空间
fmt.Println(*p) // 值被修改了

指针作为参数传参，
func swap(a *int, b *int) {  //  参数类型这里为指针类型：*int
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 60
	swap(&a,&b)  //指针作为函数参数，为地址传递
	fmt.Println(a)
	fmt.Println(b)
}

数组指针：定义指针，指向数组。函数里面如果通过数组指针修改数组里的元素直接影响原数组该元素值

arr := [10]int{10,9,8,7,6,5,4,3,2,1}
var p *[10]int = &arr
//如果直接打印*p获取了整个数组的值
//如果是获取数组某一个元素的地址值,(*p){2},*p必须加小括号，因为{}优先级大于*
//go语言里面可以直接p[下标]访问数字指针类型的元素



指针数组：数组元素是指针类型，一个数组中存储的都是指针(也就是地址)
eg.func main() {
	var p[2]*int  // *放在[]后面进行定义
	var i int = 10
	var u int = 60
	p[0] = &i   // 指针指向了变量
	p[1] = &u   // 指针指向了变量
	fmt.Println(p) //打印出全部元素的地址
	fmt.Println(*p[0]) //通过操作指针访问元素的值
	fmt.Println(p[0])  //没有加*获取了元素的地址
	for o,v := range p {  // 指针数组的遍历
		fmt.Println(o)
		fmt.Println(*v)
}

指针与切片：定义指针，指向切片
var u := []int{10,9,8,7,6,5,4,3,2,1,0}
var p *int = &u   // 指针指向了切片类型
fmt.Println((*p)[0]) // 切片类型里面通过下标访问切片元素必须要用(*p)[下标]
(*p)[0] = 600000
fmt.Println((*p)[0])
fmt.Println(*p)




指针与结构体：顶一个指针指向结构体
var Student struct {
	id int
	name string
	age int
	addr string
}

func main() {
	stu := Student{101,"UTR",100,"SZ"}
	var p *Student = & stu
	fmt.Println(p)  // 返回带&符号的结构体值
	fmt.Println(*p) // 返回正常的结构体值
	fmt.Println((*p).name) // 这两种方式都可以直接访问结构体的成员
	fmt.Println(p.age)     // 这两种方式都可以直接访问结构体的成员
	UpdateStruct(p)
	fmt.Println(stu)
}

func UpdateStruct(stu *Student) {
	stu.age = 6000000  // 函数体里面可以直接修改原结构体的
}

多级指针：一个*一级指针，两个*两级指针
	var a int = 10
	var p *int = &a  //一级指针

	var pp **int = &p //这里等同于指针调用指针吧
	**pp = 200
	fmt.Println(a) //最后a的值改变了


浅拷贝：只拷贝了变量值，没有拷贝变量指向存储空间，赋值，参数传参，函数返回值都为浅拷贝
深拷贝：原有变量空间全部都拷贝了一份

切片传参数：go中切片底层是一个结构体
append当容量不足的时候动态分配内存，这块内存是一个新的memory

for range分析：
range的本质是一个函数方法，使用的时候可以加括号使用
更改range得到后的value，不影响元时期切片或者数组



结构体：

type A struct {


	col float
}

type B struct {
	A      //  这里匿名字段代表结构体可以嵌套结构体，
	col int  //如果B，A有同名字段
}

obj := B{} //这里初始化
obj.* = value //这种方式可以修改B+A里面的成员信息
如果有同名变量，需要更改这个栏位的值,obj.col = value,更改了子类的同名变量（因为B有A,等同于A为父类，B为子类）
如果要更改A的同名变量的值，obj.A.col = value

type A struct {

	col float
}

type B struct {
	*A       //  这里指针匿名字段
	col int  //如果B，A有同名字段
}
u := B{} //初始化
u.col = value // 直接更改B or A的成员变量,这些操作与上面一样
如果要打印指针类型结构体，只需要加个*,fmt.Println(*A)

type A struct { //多重继承方法，一层一层嵌套使用

	col float
}

type B struct {
	A
	coli int
}

type C struct {
	B
	colu int
}
r := C{} // 如果有同名变量，eg, r.B.A.colume = value，指定变量更改,其他操作都一样

创建对象的用法：

go通过结构体代表一个类，属性--成员，方法--函数

type StudentInfo struct {
name string
age  int
sex  string
chinese  float64
math  	 float64
english  float64
}

func (s *StudentInfo)SayHello(username string, userAge int, userSex string) {  //前面的(）制定了接收者类型 StudentInfo
//初始化
s.name = username
s.age = userAge
s.sex = userSex
//初始化后取值进行判断
if s.sex != "男生" && s.sex != "女生" { }
//打印出结果
fmt.Println()
}

func (s *StudentInfo)GetScore() {   //前面的(）制定了接收者类型 StudentInfo，代表为StudentInfo这个父类的方法
}
func main() {

	var stu Student   // 定义一个对象?
	stu.SatHello(参数)
	stu.GetScore(参数)  //子类直接调用父类以及子类方法,为方法的继承？

	var su *Stu       //定义struct的指针类型
	su=new(Stu)       // 新建一个对象，不过返回的为指针类型
	su.方法名()       // 这种调用也可以
}





方法的继承：
参考之前的笔记

type Person struct {  //定义一个父类
name string
age   int
sex string
}

func (p *Person)SetValue(userName string, userAge int, userSex string){  //父类的方法了
p.name = userName
p.age = userAge
p.sex = userSex
}

func (p *Person) RepSayHello(h string) {  //父类的方法
fmt.Printf("我叫%s,我的爱好是%s,我的年龄是%d,我的性别为%s",p.name,h,p.age,p.sex)
}

type Rep struct {  //子类 ，继承了父类的全部成员
Person
Hobby string
}

type Pro struct {  //子类 ，继承了父类的全部成员
Person
WorkYear int
}


func (r *Rep) RepSayHello(h string) {  //Rep子类的方法

r.Hobby = h
fmt.Printf("我叫%s,我的爱好是%s,我的年龄是%d,我的性别是%s\n",r.name,r.Hobby,r.age,r.sex)
}


func (p *Pro) ProSayHello() {  //Pro子类的方法
fmt.Println("成功")
}


func main() {
var rep Rep
rep.SetValue("张三",1000,"男生") //子类对象可以调用父类的方法，父类对象只能调用自己的方法无法调用其他方法
rep.RepSayHello("偷拍") //这个子类调用自己的方法，如果父类子类都有这个方法，默认调用子类的，这个叫方法的重写
rep.Person.RepSayHello("偷拍") //这个调用才可以调用父类的同名方法
}
方法值调用：
type Student struct {
	user string
	age int
}
func (p *Student)GetInfo() {
	fmt.Println(*p)
}

func main() {
	stu := Student("张三",10000)
	stu.GetInfo() //如果方法没有定义形参，直接这样调用
	//方法值
	f := stu.GetInfo
	fmt.Printf("%T",f)  //打印出f为 func
	f() //这样也可以获取Student全部的成员
	//方法表达式
	f := (*Student).GetInfo //前面类型必须与方法的接收者类型保持一致
	f(&per) //括号里面必须指定传值的对象，顶一个变量接收，加上取地址符号次啊可以打印值
}


接口：只是一种规范与标准，接口为虚的，把所有的具有共性的方法定义在一起，方法为实的，任何其他类型只要实现了这些方法，
	就是实现了这个接口(interface)
	格式：type inter interface {
		sayHello()//方法名，接口中的方法必须有具体的实现，该方法必须有}
		sayOoooo()//方法有几个，定义必须有几个
	type person struct {  //父类
		name string
		sex string
		age int
	}

	type student struct { //子类
		person
		score int
	}

	type teacher struct { //子类
		person
		subject string
	}

	func (s *student)SayHello(){  //子类方法
		fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁了，我的成绩是%d分\n",s.name,s.sex,s.age,s.score)
		}
	func (t *teacher)SayHello(){  //子类方法
		fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁了，我叫的学科是%s\n",t.name,t.sex,t.age,t.subject)
		}
	//1
	//接口定义
	//接口定义了规则   方法实现了规则
	//接口是虚的    方法是实的
	//格式 type  接口名  interface{方法列表}
	//方法名（参数列表）（返回值列表）
	type Humaner interface {
	//2
	//方法 函数声明 没有具体实现  具体的实现要根据对象不同  实现方式也不同
	//接口中的方法必须有具体的实现
		SayHello()
	}
	func main() {

	var stu student=student{person{"小明","男",18},100}
	var tea teacher=teacher{person{"法师","女",31},"盗墓"}
	//3
	//定义接口类型
	//接口做了同一的处理 先实现接口  在根据接口实现对应的方法
	//在需求改变时 不需要改变接口 只需要修改方法
	var h Humaner

	//fmt.Println(h.SayHello)
	h=&stu    //这种调用必须满足接口里面的全部方法有几个，就定义了几个方法，否则这种就报错：
	h.SayHello()  //does not implement Humaner (missing SayOoooo method)

	fmt.Printf("%p\n",h)
	h=&tea
	h.SayHello()
	h.SayOoooo()  //接口里面定义了几个方法，必须拥有几个方法定义
	fmt.Printf("%p\n",h)

	//stu.SayHello() // 这种调用与之前的借口调用结果一样
	//tea.SayHello() // 这种调用与之前的借口调用结果一样
	}
多态：多种表现形式，多态为同一个接口，使用不同的实例而执行不同的操作
如何实现多态：
	func 函数名(参数 接口类型) {

	}

	type Object struct { //定义了一个父类
		numA int
		numB int
	}

	type Add struct { //子类
		Object
	}

	func (a *Add)GetResult() int {   //每一个运算规则交给自己的类来完成
		return a.numA + a.numB
	}

	type Sub struct { //子类
		Object
	}

	func (r *Sub)GetResult() int {
		return r.numA - r.numB
	}

	type Resultor interface{ //接口类只存储方法，方法必须有定义
		GetResult() int  //这里是否要加上返回类型，取决于同名函数是否有返回值
	}

	// 如果要进一步精简mian里面的结构，定义一个新的类
	type OperatorFactory struct {
	}

	func (o *OperatorFactory)CreateOperator(op string, numA int, numB int) int { //这个类里面负责定义子类对象，等同于
		switch op {                                                             //之前在main里面建立子类对象，如果
		case "+":                                                               //接口对象调用方法,声明定义接口对象后,
			add := Add{Object{numA,numB}}                                        //接口对象 = &子类对象,然后  接口.方法
			return OperatorWho(&add) // 这个是传递子类独享的地址通过多态的方式调用同名方法?
		case "-":
			sub := Sub{Object{numA,numB}}
			return OperatorWho(&sub)
		default:
		return 1012
		}

	}


func OperatorWho(o Resultor) int{  //参数为接口类型，里面的方法可以不同对象调用彼此自己的同名方法,多态
return o.GetResult()
}
	func main() {
		//add := Add{Object{10,60}}
		//i := add.GetResult()
		//fmt.Println(i)
		//var Res Resultor
		//Res = &add
		//i := Res.GetResult()
		//fmt.Println(i)
		//fmt.Println(i)
		//这样每次只选择一个功能，只需要更改 参数不需要更改代码
		var oper OperatorFactory  // 建立新类的对象
		ur := oper.CreateOperator("-", 60, 100) // 通过类的方法传参即可一层层调用达到目标
		fmt.Println(ur)
	}

空接口可以接收任意类型的数据，

type Humaner interface {
	//子集
	sayhi()
}

type Personer interface {
	//超集
	//继承与humaner
	Humaner

	sing(string)
}

type Student struct {
	name string
	sex  string
	age  int
}

func (s *Student) sayhi() {
	fmt.Printf("大家好，我是%s，我是%s生，我的年龄是%d\n",
	s.name, s.sex, s.age)
}

func (s *Student) sing(name string) {
	fmt.Println("我为大家唱首歌", name)
}
func main0601() {

	//接口类型变量定义
	var h Humaner

	var stu Student = Student{"王飞", "男", 35}
	h = &stu
	h.sayhi()

	//接口类型变量定义
	var p Personer
	p = &stu

	//从humaner继承来的
	p.sayhi()
	p.sing("传奇")
}
	func main0602() {

		//接口类型变量定义
		var h Humaner  //子集
		var p Personer //超集
		var stu Student = Student{"王飞", "男", 35}

		h = &stu
		//将一个接口赋值给另外一个接口
		//超集中包含所有子集的方法
		h = p //ok
		h.sayhi()

		//子集不包含超集
		//可以将超集赋值给子集  不能将子集赋值给超集
		//p = h //err

		//p.sayhi()
		//p.sing("红豆")

	}



func main0701() {

	var i interface{}

//接口类型可以接收任意类型数据

//fmt.Printf("%T\n", i)
//i = 10
//fmt.Printf("%T\n", i)
//fmt.Println(i)
//i = 3.14
//fmt.Printf("%T\n", i)
//fmt.Println(i)
//i = "传智播客"
//fmt.Printf("%T\n", i)
//fmt.Println(i)

	i = 10

//接口类型 不能直接进行转换需要使用类型断言
//var a int =20
//var a interface{}
//a = 20

	fmt.Println(i)
}

	func test() {
		fmt.Println("test hello world")
	}
	func main0702() {
		//空接口类型的切片
		var i []interface{}

		//fmt.Printf("%T\n",i)
		i = append(i, 10, 3.14, "传智播客", test)

		for idx := 0; idx < len(i); idx++ {
			fmt.Println(i[idx])
		}
	}

	import "fmt"

	func main0801() {
		var i interface{}
		i = 10.234
		//value,ok:=map[key]
		//值,值的判断:=接口变量.(数据类型)
		value, ok := i.(int)
		if ok {
			fmt.Println("整型数据：", value+10)
		} else {
			fmt.Println("错误")
		}

	}
	func demo() {
		fmt.Println("demo hello world")
	}
	func main() {
		var i []interface{}

		i = append(i, 10, 3.14, "传智播客", demo,123,test)

		for _, v := range i {
		//fmt.Println(idx,v)

		if data, ok := v.(int); ok {
			fmt.Println("整型数据：", data)
		} else if data, ok := v.(float64); ok {
			fmt.Println("浮点型数据：", data)
		} else if data, ok := v.(string); ok {
			fmt.Println("字符串数据：", data)
		} else if data, ok := v.(func()); ok {
			//函数调用
			data()
		}
	}



}
异常的处理：(panic)
例如：map里面的key是否存在用这个先判断异常，如果有异常如何处理。这里要讲如何抛出异常
if value,ok := map名[key]; ok { // ok为布尔类型，否则判断不成立，报错：non-bool used as if condition
	................... // 这种形式判断是否有值存在。如果存在执行这个
} else {
	................... // 如果不存在执行这个
}

error抛出异常需要导入包：errors
然后顶一个err变量接收：err = errors.New("报错信息")
import (
	"fmt"
	"errors"
	)

func Raise(a int, b int) (value int, err error) { // error类型
	if b == 0 {
	err = errors.New("除数不能为0") //变量存储错误信息
	return
	}
	value = a / b
	return
}
func main() {

	a := 10
	b := 0
	if value, err := Raise(a, b); err != nil { //这里的if格式，要么为一个布尔类型变量，要么为表达式（判断err表达式）
	fmt.Println("runtime error",err)
	} else {
	fmt.Println(value)
	}
}

panic：当程序遇见panic是自动终止，好像开发中不用到这个函数，eg：panic("hello world") //异常的格式
recover拦截错误:必须配合函数defer来拦截错误，defer是延迟，多个defer是后进先出原则，最后的defer语句最先运行
eg: defer func()，必须接函数调用，如果放在函数体前面，函数体{}必须紧随其后接()。recover系统自带函数。
func main() {
	a := 10
	b := 20


f := func(a int, b int) {   //接收函数返回值的变量
	fmt.Println(a + b)
}

//如果有传参，defer调用时将函数参数先放在内存中 是一个独立的空间  不会因为改变值允许数据
defer f(a, b) //如果后面值有修改是无效的,不影响f函数结构

a = 123
b = 321
//如果不传递参数,使用外部的变量  如果外部变量修改,会影响函数的值
defer f()

a = 123
b = 321
}


import "fmt"

func demo(i int) {

//错误拦截到出现在错误之前
defer func() {
	//错误拦截 panic异常错误
	err := recover()
	//判断是否出现错误
	if err != nil { //这里可以简化为：if err := recover(); err != nil {}
		fmt.Println(err)
	}

}() // defer接函数调用(function call)，如果加在函数体前面，函数体{}紧跟()

var arr [10]int

var p *int //0x0 nil
//p = new(int)
*p = 123

//数组下标越界
arr[i] = 123

fmt.Println(arr)

}

func main() {
demo(10)

fmt.Println("hello world")
}



建立文件操作： //
import (
"os"
"fmt"
)

	func main() {

	//os.Create(文件名)文件名 可以写绝对路径 也可以写相对路径
	//返回值为 文件指针 错误信息
	//如果文件不存在会创建新文件 如果文件存在会清空源文件内容
	//fp, err := os.Create("D:/a.txt") 这个相对路径
	fp, err := os.Create("u.txt") //这个绝对路径，在本project里面建立文件
	fmt.Println("文件指针：",fp)
	fmt.Println("错误信息 ：",err)

	if err != nil {
	//文件创建失败
	/*
	1、路径不存在
	2、文件权限
	3、程序打开文件上限 65535
	 */
	fmt.Println("文件创建失败")
	return
	}

	//延迟调用关闭文件
	//defer fp.Close()

	//文件创建成功 可以操作文件

	//关闭文件
	//如果打开文件不关闭 会造成内存浪费 程序打开文件上限
	fp.Close()   //文件指针来关闭文件

	}


打开写入文件操作：
import (
"os"
"fmt"
)

func main0601() {

//\反斜杠在程序中表示转义字符 会将后面一个字符进行转义  [\\]表示一个[\]
//在写路径是可以使用/正斜杠来代替\反斜杠
fp, err := os.Create("D:/a.txt")

if err != nil {
fmt.Println("文件创建失败")
return            // 简体为：if fp, err := os.Create("D;/a.txt"); err != nil {}
}
//关闭文件
defer fp.Close()

//将字符串写入到文件中
//\n不会换行  原因是 在window文本文件中换行是以\r\n 在linux中换行是以\n
fp.WriteString("hello world\r\n")
//不会换行
fp.WriteString("性感荷官在线发牌")
}

func main0602() {
fp, err := os.Create("D:/a.txt")
if err != nil {
fmt.Println("文件创建失败")
return
}

defer fp.Close()
//
//slice := []byte{'h', 'e', 'l', 'l', 'o'}
////将字符写入文件中
//fp.Write(slice)  //Write一个一个字符

//slice :=[]byte("性感法师在线讲课")
//
//fp.Write(slice)
//count, _ := fp.Write([]byte("性感法师在线讲课"))
count, err1 := fp.WriteString("性感法师在线讲课")
if err1 != nil {
fmt.Println("写入文件失败")
return
} else {
fmt.Println(count)
}
}

func main() {
//fp, err := os.Create("D:/a.txt")
//os.Open(文件名) 只读方式打开文件
//os.OpenFile(文件名，打开方式，打开权限)    如果文件不存在会报错
//打开方式 os.O_RDONLY（只读方式打开）os.O_WRONLY（只写方式打开）os.O_RDWR（可读可写方式打开）os.O_APPEND（追加方式打开）
//打开权限 0-7 rwx  6 （rw-）读写权限  7 （rwx）读写执行权限
fp, err := os.OpenFile("D:/a.txt", os.O_RDWR, 6)
if err != nil {
fmt.Println("打开文件失败")
return
}

defer fp.Close()

//获取光标流位置
//获取从文件起始到结尾有多少个字符
//count,_:=fp.Seek(0,os.SEEK_END)
//count,_:=fp.Seek(0,io.SeekEnd)
//指定位置写入
//fp.WriteAt([]byte("hello world"),count)
//fp.WriteAt([]byte("itcast"), 0)
//fp.WriteAt([]byte("花儿"),17)

fp.WriteString("hello world\r\n")
fp.WriteString("你瞅啥\r\n")
fp.WriteString("瞅你咋地\r\n")
fp.WriteString("在瞅一个试试\r\n")
fp.WriteString("对不起大哥我错了")

//fp.WriteAt([]byte("瞅你咋地"),14)
}


读取文件：
	func main() {
	//读取文件
	fp1, err1 := os.Open("D:/0824基础班/01_Go语言基础第07天（指针）/03视频/03结构体和函数.avi")
	//写入文件
	fp2, err2 := os.Create("D:/test.avi")

	if err1 != nil || err2 != nil {
	fmt.Println("操作文件失败")
	return
	}
	//关闭文件
	defer fp1.Close()
	defer fp2.Close()

	//拷贝文件
	//通过read块进行文件读取
	//通过write进行写入
	buf := make([]byte, 1024*1024*10) //1KB 1MB  1KB-16MB 缓冲块代表每次读取多少数据
	for {
	//将读取的字符写入到新文件中
	n, err := fp1.Read(buf)
	if err == io.EOF {  //代表文件读完了
	break
	}
	fp2.Write(buf[:n])
	}
	}


//创建缓冲区
r := bufio.NewReader(fp)
dict := make(map[string]string)
for {
//读取两行内容  单词和翻译
word, _ := r.ReadBytes('\n')
trans, err1 := r.ReadBytes('\n')
//将单词和翻译做类型转换 格式化 放在map中

//#abandoner\r\n
//abandoner
//Trans:n. 遗弃者;委付者\r\n
//n. 遗弃者;委付者
dict[string(word[1:len(word)-2])] = string(trans[6 : len(trans)-2])

//到文件结尾  循环结束
if err1 == io.EOF {
break
}
}

//fmt.Println(dict)

//单词查找
word := ""
fmt.Println("请输入您要查找的单词")
fmt.Scan(&word)

//fmt.Println(dict[word])
value, ok := dict[word]
if ok {
fmt.Println(value)
} else {
fmt.Println("未找到该单词的翻译")
}





字符串操作：其他操作可以参考文档地址：https://studygolang.com/pkgdoc
方法来自于strings这个包，
Contains()
str := "itcastheima"
eg, b := strings.Contains(str,"字符串") //判断字符串是否在str里面存在,存在返回布尔true，不存在返回false

Join()
str := []string{"it","itcast","heima"} //str必须是切片指针类型
b := strings.Join(str,"|") // Join第一个参数必须为切片指针类型，第二个参数为用什么联接符号联接各个字符串。返回一个字符串

Index()
str := "hello"
b := strings.Index(str,"字符串") //在str里查找字符或者字符串是否存在以及对应的位置值(下标)，字符返回对应的下标，
                                 //字符串返回第一个字符所在的下标，如果没找到就返回 -1

Repeat()
str := strings.Repeat("go",i)  // 重复一个字符或者字符串i次，
fmt.Println(str)


Replace()
str := "go"
strings.Replace(str,o,O,i) // o替换前，O替换后，i替换的次数(i < 0代表全部替换)

Split()

str := "hello@itcast"
u := strings.Split(str,"分隔符号") // 什么分割符号来截断,最后返回一个切片类型，包括str全部字符只是每一段space分割了

import "strconv"
其他类型转换成字符串类型：
str := strconv.FormatBool(true) // Formatbool这个方法可以将布尔类型转换成字符串
 fmt.Println(str)	// 这里打印的类容为true，不是字符串类型的true，不是布尔类型的true(因为布尔类型打印出来也为true)

str := strconv.Itoa(数字) //整数类型转换成字符串类型


字符串类型转换成其他类型：
b,err := strconv.ParseBool("true")  // 字符串类型转换成布尔类型
if err != nil{
	fmt.Println(err)
}else {
	fmt..Println(b)
}

b, err := strconv.Atoi("字符串")
if err != nil {
	fmt.Println(err)
}else {
	fmt.Println(b)
}





工程管理：
同级目录创建：go路径下main.go *.go文件并列存在 main * 文件，文件里面都为package main，然后main里面直接调用* 里面的函数
	(函数名称首字母大写)，eg, sub()   运行该main在上方编辑里面选择按照directory运行，指定文件夹名称

不同级目录创建：go路径下面存在 main文件 R文件夹，R文件(文件夹名称首字母大写)夹里面存在O文件，O文件里面为package R，
	函数名称首字母大写，main里面调用函数，导入 R文件夹，再通过R.函数名()
	eg. import "Ro"  //文件夹名称
		Ro.Uo()       // 运行该main在上方编辑里面选择按照directory运行，指定文件夹名称

函数：
堆zhan逃逸

Walk：一个写给Golang的Window应用程序库，它主要用于Windows开发,只用于Windows







基础加强版：针对前面的只是进一步深入
指针的进一步强化：
func swap(a, b int)  {   //这种类似于浅拷贝，只是值传递，复制一个值传给参数，而变量本身不被改变
	a, b = b, a			//作用域只限于函数体
	fmt.Println("swap  a:", a, "b:", b)
}

func swap2(x, y *int)  { // 传地址值。变量本身被改变
	*x, *y = *y, *x
}

func main()  {
	a, b := 10, 20

	swap2(&a, &b)			// 传地址值。变量本身被改变
	fmt.Println("swap2: main  a:", a, "b:", b)

}

var a int = 10000
var p *int = &a

a = 20000
fmt.Println("a = ", a)    //a被改变了
fmt.Println("*p = ", *p)  //这个就可以打印出值，a被改变，*p同样被改变了

*p = 666666                // 借助a 变量的地址，操作a对应空间
fmt.Println("a = ", a)    //结果a被改变了


 := new(int)			// 默认类型的 默认值，新建一个指针变量
*p = 1000
fmt.Printf("%d\n", *p)     //打印指针变量的值
fmt.Printf("%T\n", p)      //打印出类型(布尔类型只用t打印)
fmt.Printf("%v\n", p)		// 打印Go语言格式的字符串，这个等于%p，因为打印结果一样的都为地址？

切片的加强：
var u []int = []int{} //初始化
u := []int{} //初始化

rr := [100]int{}
u ;= rr[i:v:q]  //这里用到了切片的截取操作，
len(u)  // len = v -i
cap(u)  //虽然定义了容量为q，切片里面取值cap(u) = q -i，与make不太一样，make(对象，容量)，容量定义=cap取值
u := rr[i:v] //如果q没有给出，cap(u)取值为数组长度-i
make 用来为 slice，map 或 chan 类型分配内存和初始化一个对象(注意：只能用在这三种类型上)
make函数：make(对象,长度,容量)  //容量为多少，cap取值 = 容量定义值，如果容量没有定义，容量 = 长度。并且make创建切片
								// 必须指定长度，其他不指定长度，系统分配一个默认值
out := data[:0]					// 在原切片上截取一个长度为 0 的切片 == make([]string, 0)
copy函数：参考上面

map加强部分：var m1 map[key类型]value类型
var m1 map[int]string			// 声明map ，没有空间，不能直接存储key -- value
m1[100] = "Green"             //这里如果运行了直接抛出一个异常 panic: assignment to entry in nil map
	if m1 == nil {             //如果没有上面的操作，这个if语句才运行，因为没有指定空间初始为空白
	fmt.Println("map is nil ")
	}
cap不能用于字典，	报错：invalid argument m (type map[int]string) for cap
map[key] = value //直接改变了key对应的值为value
判断map的key是否存在着：if v,ok := m[key]; ok{ //初始化表达式 + 条件。
											   // 返回两个值，第一个是value，第二个是bool 代表key是否存在。.
	......  // 如果存在怎么操作
}else {
	...... // 如果不存在怎么操作
}

map[键], 键唯一、无序,不可以为引用类型，在go语言中，值类型和引用类型有以下特点：
a、值类型：基本数据类型，int,float,bool,string,以及数组和struct
特点：变量直接存储值，内存通常在栈中分配，栈在函数调用完会被释放

b、引用类型：指针，slice，map，chan等都是引用类型
特点：变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配，通过GC回收。

值传递：值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
map 本质是指针类型，在函数中修改可以影响到原始map，channle 底层也属于指针类型，channel是goroutine的重要通讯机制，
任意一个goroutine向channel添加或者取数据都影响原始的channel中的数据，interace通过反射修改必须要求interface中
保存的是对象指针，反射部分请同学看直播的反射部分。 数组指针是指针函数修改原始值主要通过指针所以是可以的

引用传递：所谓引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
最早出现c++中，其实在go语言中，只存在值传递，好多人有个误解，说slice，map，chan这些引用类型是引用传递，误区可能就是：
map、slice、channel这类引用类型在传递到函数内部，可以在函数内部对它的值进行修改而引起的误会。

结构体加强：
type U struct{}
var i U = U{...}  // 初始化可以全部初始化-结构体里面全部成员赋值，可以部分初始化，给某些成员赋值
其他初始化按照之前的方法，

// 结构体比较 ，返回布尔值，true or false
var p1 Person = Person{"andy", 'm', 18}
var p2 Person = Person{"andy", 'm', 18}
var p3 Person = Person{"andy", 'm', 181}
fmt.Println("p1 == p2 ?", p1 == p2)
fmt.Println("p1 == p3 ?", p1 == p3)
// 相同类型结构体赋值，直接按照赋值操作，数组也可以比较+赋值
var tmp Person
fmt.Println("tmp", tmp)
tmp = p3
fmt.Println("tmp", tmp)
//结构体指针：这两种用法参考指针的用法
var p1 *Person2 = &Person2{"n1", 'f', 19}
fmt.Println("p1", p1)

var tmp Person2 = Person2{"n1", 'f', 19}
var p2 *Person2
p2 = &tmp
fmt.Println("p2", p2)

//指针作为参数传参
func test2(p *Person2)  {
	fmt.Println("test2:", unsafe.Sizeof(p))
	p.name = "Luffy"
	p.age = 779
	p.sex = 'm'
}

p3 := new(Person2)
p3.name = "n3"
p3.age = 22
p3.sex = 'f'    //初始化
fmt.Printf("p3, type= %T\n", p3)
fmt.Println("p3:", p3)
fmt.Println("main:", unsafe.Sizeof(p3))
test2(p3)   //调用函数并且指针传参
fmt.Println("p3:", p3)



字符串加强：
str := "I love my work and I love my family too"
// 字符串按 指定分割符拆分，
ret := strings.Split(str, " I") // 拆分结果不包括分割符，结果为分隔符号前面的部分+分隔符号后面的部分
for _, s := range ret {
fmt.Println(s)   // 结果为分隔符号前面的部分+分隔符号后面的部分
}

// 字符串按 空格拆分，包括一个或者多个空格
ret := strings.Fields(str)
for _, s := range ret {
fmt.Println(s)
}


// 判断字符串结束标记，第一个参数为检查的字符串，第二个为标记,返回给flg为布尔类型
flg := strings.HasSuffix("test.abc", ".mp3")
fmt.Println(flg)

// 判断字符串起始标记，第一个参数为检查的字符串，第二个为标记,返回给flg为布尔类型
flg = strings.HasPrefix("test.abc", "tes.")
fmt.Println(flg)

文件的加强：stat 获取文件类型
fo, err := os.Create("testFile.xyz")  // 建立文件如果与打开，写入文件存在一个作用域，f, err这个区域f 不可以重名
if err != nil {
	fmt.Println("create err:", err)
	return
}

f, err := os.OpenFile("testFile.xyz", os.O_RDWR, 6)
if err != nil {
fmt.Println("OpenFile err:", err)
return
}
defer fo.Close()  //多个defer可以写在一起
defer f.Close()

for i:=0;i<100;i++ {
	_, err = f.WriteString("######") //如果前面有打开或者建立文件，f, err这里f 可以用匿名变量
	if err != nil {
	fmt.Println("WriteString err:", err)
	return
	} // WriteString每次写入一个字符串，Write每次写入一个字符，这个应该不怎么用
}

f.WriteString("helloworld\r\n") //write helloworld
off, _ := f.Seek(-5, io.SeekEnd) // Seek(): 	修改文件的读写指针位置。
               //参1： 偏移量。 正：向文件尾偏， 负：向文件头偏    光标从指定位置正，或者负移动|参1|个位置

                     //参2： 偏移起始位置：
                     io.SeekStart: 文件起始位置
                     io.SeekCurrent： 文件当前位置
                     io.SeekEnd: 文件结尾位置

//off:	返回值：表示从文件起始位置，到当前文件读写指针位置的偏移量。
fmt.Println("off:", off) // off: 7，好像\r\n也要作为一次指针移动

result, _ = f.WriteAt([]byte("1111"), off)  //  writeAt():  在文件指定偏移位置off，写入 []byte ,  通常搭配 Seek()
fmt.Println("WriteAt result :", result)  // 打印结果为：hellowo1111

// 创建一个带有缓冲区(用户缓冲)的 reader
reader := bufio.NewReader(f)               //导入bufio包
for {
//	buf, err := reader.ReadBytes('\n')    // 读一行数据ReadBytes/ReadString/ReadSlice都传一个参数作为读取结束标志
	buf, _, err := reader.ReadLine()	  // 读一行数据ReadLine这个用法可固定，不需要传参数

	if err != nil && err == io.EOF {         //判断文件结束的标志
		fmt.Println("文件读取完毕")
		return
	} else if err != nil {
		fmt.Println("ReadBytes err:", err)
	}
	fmt.Print(string(buf)+"\n")           // 这里打印每次去读的缓冲区的值,转义字符进行换行
}


copy文件：

目录操作加强：
func main()  {
	// 获取用户输入的目录路径
	fmt.Println("请输入待查询的目录：")
	var path string
	fmt.Scan(&path)          // 获取目录路径：盘符:\文件夹?

	// 打开目录
	f, err := os.OpenFile(path, os.O_WRONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err: ", err)
		return
	}
	defer f.Close()
	// 读取目录项
	info, err:= f.Readdir(-1)	// -1： 读取目录中所有目录项，, Readdir返回全部文件,用切片来存储
	if err != nil {
		fmt.Println("Readdir err: ", err)
		return
	}
	// 变量返回的切片
	for _, fileInfo := range info {
		if fileInfo.IsDir() {			// 是目录
			fmt.Println(fileInfo.Name(), " 是一个目录")
		} else {
			fmt.Println(fileInfo.Name(), " 是一个文件")
		}
	}
}


Open文件的一套流程：open， defer close， read， write
GO里面关于进程，线程，协程，并发.1. 协程比线程更轻量级 2. 协程涉及到资源竞争，需要加锁，加锁就可能出现死锁
Goroutine ：Go程, main函数为一个goroutine,main 函数确实运行在自己的 goroutine 中！更重要的是一旦 main 函数返回，
它将关掉当前正在运行的其他 goroutines。现在有更好的方式阻塞 main 函数直到其他所有 goroutines 都运行完。
通常的做法是创建一个 done channel， main 函数在等待读取它时被阻塞。一旦完成工作，向这个 channel发送数据，程序就会结束了。
func ******() { //匿名函数定义也可以 go func() {} ，直接定义在main里面
	...
}

func main() {
	go ******() // 直接使用go关键字，放置于函数调用前面，产生并且启动一个go程
	/* we are using time sleep so that the main program does not terminate before the execution of goroutine.*/
	time.Sleep(60 * time.Second) //import time包，Sleep代表经过60秒 main主程才结束(这个时间包含了goroutine运行的时间)，
								// 如果第一个参数为负数或者0，代表主程序马上结束。Sleep暂停当前的goroutine
	fmt.Println("这是个go程")
	fmt.Println("main terminate") //如果没有Sleep，只有这句话被打印出来
}
func main() {
	doneChan := make(chan string)

	go func() {
		//
		doneChan <- "Work is all done"
	}()

	<- doneChan //这个地方将通道的值给main,一旦子goroutine运行完成有值给通道，通道传给main 函数，程序就结束了
}

//时间片轮转调度算法
go func() {
	for {
		fmt.Println(" this is goroutine test")
	}
}()

for {
	runtime.Gosched()			// 出让当前go程所占用的cpu 时间片。当再次获得cpu时，从出让位置继续回复执行。
	fmt.Println(" this is main test") //—— 时间片轮转调度算法
}




func test()  {
	defer fmt.Println("ccccccccccccccccc")
	//return
	runtime.Goexit()			// 退出当前go程。结束调用该函数的当前go程。Goexit():之前注册的 defer都生效。
	defer fmt.Println("ddddddddddddddddd")
}

// return：	返回当前函数调用到调用者那里去。 return之前的 defer 注册生效。
func main()  {

	go func() {
		defer fmt.Println("aaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbb")
	}()
}

runtime.GOMAXPROCS():

设置当前 进程使用的最大cpu核数。 返回 上一次调用成功的设置值。 首次调用返回默认值。

Channel：通道类型的定义
一般为 ：make(chan 类型，容量)，定义了容量为带缓冲的通道，，未定义容量为不带缓冲的通道
ch := make(chan string，10000)
// len(ch) ： channel 中剩余未读取数据个数。 cap（ch）： 通道的容量。
fmt.Println("len(ch)=", len(ch), "cap(ch)=", cap(ch))

ch  <- "写端:写入数据到通道类型" // 写端，写入数据到通道，
str <- ch //读端：从通道里面读取数据。   两者必须配合使用缺一不可，并且某一端有操作，另外一端没有在状态，就造成通道堵塞

goroutine为异步执行的，通过sync包 or Channel来实现同步。
无缓冲channel：ch := make(chan string) or ch := make(chan string, 0),通道容量为0， len = 0 。 不能存储数据。
func main()  { // 无缓冲channel具备同步的能力。  读、写同步。（打电话）
ch := make(chan int)

go func() {
	for i:= 0; i<5;i++ {
		fmt.Println("子go程， i=",i)
		ch <- i			// ch <- 0
	}
}()
//time.Sleep(time.Second * 2)
for i:= 0; i<5;i++ {     for与goroutine同步运行的，因为goroutine里面ch每次读取两个i值，然后for里面同步运行两次对应i赋值
	num := <- ch
	fmt.Println("主go程读：", num)
	}
}


缓冲channel：ch := make(chan string,1000000),
通道容量为非0。len(ch) ： channel 中剩余未读取数据个数。 cap（ch）： 通道的容量。
缓冲区可以进行数据存储。存储至 容量上限，阻塞。 具备 异步 能力，不需同时操作channel缓冲区（发短信）

关闭channel，close(ch)  //通道关闭后，写入该通道有错误，如果从其中读取读取为0。
对端可以判断 channel 是否关闭：

if num， ok := <-ch ;  ok == true {

	如果对端已经关闭， ok --> false . num无数据。

	如果对端没有关闭， ok --> true . num保存读到的数据。

可以使用 range 替代 ok：(遍历通道类型)

for num := range ch {		// ch 不能替换为 <-ch
}

默认的channel 是双向的。 var ch chan int		ch = make(chan int)
单向写channel:	var  sendCh chan <-	 int	sendCh = make(chan <- int)	不能读操作
单向读channel:	var  recvCh  <- chan int	recvCh = make(<-chan int)

转换：
	1. 双向channel 可以 隐式转换为 任意一种单向channel

	sendCh  = ch

	2. 单向 channel 不能转换为 双向 channel

	ch = sendCh/recvCh   error！！！

	传参： 传【引用】  //定义函数的时候 经常使用的是单向的channel 传参的时候传递的确实双向的







Timer：创建定时器，指定定时时长，定时到达后。 系统会自动向定时器的成员 C 写 系统当前时间。 （对 chan 的写操作）
type Timer struct {
	C <-chan Time
	r runtimeTimer  //  Timer定时器定时成功后，在定时时间到达之前，会一直阻塞
}
func main()  {
	fmt.Println("当前时间：", time.Now())
	// 创建定时器
	myTimer := time.NewTimer(time.Second * 2) // 创建定时器，指定定时时长
	nowTime := <- myTimer.C  // chan 类型   myTimer.C取值为： 系统会自动向定时器的成员 C 写 系统当前时间
	fmt.Println("现下时间：", nowTime)
}
// 3 种定时方法
func main()  {
// 1 . sleep
	time.Sleep(time.Second)    //暂停当前goroutine

// 2. Timer.C
	myTimer := time.NewTimer(time.Second * 2)		// 创建定时器， 指定定时时长
	nowTime := <- myTimer.C  						// 定时满，系统自动写入系统时间给myTimer.C,（对 chan 的写操作）
	fmt.Println("现下时间：", nowTime)

// 3 time.After  指定定时时长，定时到达后。 系统会自动向定时器的成员 写入 系统当前时间。
	fmt.Println("当前时间：", time.Now())
	nowTime2 := <-time.After(time.Second * 2)
	fmt.Println("现下时间：", nowTime2)  // 返回 可读 chan 。 读取，可获得系统写入时间
}


	//Sleep、NewTimer、After —— time包
	// 定时器的停止和重置
func main()  {

	myTimer := time.NewTimer(time.Second * 10)		// 创建定时器。
	myTimer.Reset(1 * time.Second)		// 重置定时时长为 1
	go func() {
	for {
	<- myTimer.C
	fmt.Println("子go程，定时完毕")
	}
	}()

	//myTimer.Stop()		// 设置定时器停止
	for {
	;       // 这里为主go程？必须有这个循环，不然系统直接什么都不做结束了，子go程无法运行的
	}
}
周期定时：Ticker
type Ticker struct {
	C <-chan Time
	r runtimeTimer
}
func main()  {

	quit := make(chan bool)		// 创建一个判断是否 终止的channel

	fmt.Println("now:    ", time.Now())
	myTicker := time.NewTicker(time.Second)  // 定义一个周期定时器

	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C // 定时时长到达后，系统会自动向 Ticker 的 C 中写入 系统当前时间
			i++  // 并且，每隔一个定时时长后，循环写入 系统当前时间。并且，每隔一个定时时长后，循环写入 系统当前时间。
			fmt.Println("nowTime:", nowTime)
			if i == 3 {
				quit <- true	// 解除 主go程阻塞。
				break // return // runtime.Goexit
			}
		}
	}()
	<-quit		// 子go程 循环获取 <-myTicker.C 期间，一直阻塞
}






非阻塞写，非阻塞听
import (
	"fmt"
	"time"
	"runtime"
	)


func main()  {
	ch := make(chan int)			// 用来进行数据通信的 channel
	quit := make(chan bool)			// 用来判断是否退出的 channel
	//ch2 := make(chan string)
	go func() {					// 写数据
		for i:=0; i<5;i++{
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true			// 通知主go程 退出
		runtime.Goexit()
	}()
	for {	// 主go程 select 监听 chan数据流动,
			select {  //select机制用来处理异步IO问题,golang在语言级别支持select关键字,select后面必须跟通道操作(I/O操作)
				case num := <-ch:		// 不可读，阻塞。可以读，将数据保存至num
					fmt.Println("读到：", num)		// 模拟使用数据

				case <-quit:			// 不可读，阻塞。可以读，将主go程结束。
					break 				// break 跳出 select	不可用
			//runtime.Goexit()	// 终止 主 go 程		不可用
			//return 				// 终止进程
			}
			fmt.Println("============")		// select 自身不带有循环机制，需借助外层 for 来循环监听
	}       // 并且 当select后面的case同时满足时，会随机选择一个可用通道做收发操作
}


select超时处理：
import (
	"fmt"
	"time"
)

func test(ch chan int, quit chan bool)  {

}

func main()  {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {			// 子go 程获取数据
		for {
			select {
				case num := <-ch:
					fmt.Println("num = ", num)
				case <-time.After(3 * time.Second):
					quit <- true
				goto lable
				// return
				// runtime.Goexit()
			}
		}
	lable:
	}()
	for i:=0; i<2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}
	<-quit 			// 主go程，阻塞等待 子go程通知，退出
	fmt.Println("finish!!!")
}






Mutex(互斥锁)： // 使用 “锁” 完成同步 —— 互斥锁
var mutex sync.Mutex		// 创建一个互斥量， 新建的互斥锁状态为 0. 未加锁。 并且锁只有一把。

func printer(str string)  {
	mutex.Lock()			// 访问共享数据之前，加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()			// 共享数据访问结束，解锁
}

func person1()  {				// 先
	printer("hello")
}

func person2()  {				// 后
	printer("world")
}

func main()  {
	go person1()
	go person2()
	for {
		;
	}
}



读写锁：读时共享，写时独占。写锁优先级比读锁高。
import (
	"math/rand"
	"time"
	"fmt"
	"sync"
	)

var rwMutex sync.RWMutex		// 锁只有一把， 2 个属性 r w

func readGo(in <-chan int, idx int)  {
	for {
		rwMutex.RLock()			// 以读模式加锁
		num := <-in
		fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
		rwMutex.RUnlock()		// 以读模式解锁
	}
}

func writeGo(out chan<- int, idx int)  {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex.Lock()			// 以写模式加锁
		out <- num
		fmt.Printf("%dth 写go程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)		// 放大实验现象
		rwMutex.Unlock()
	}
}


func main()  {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())

	// quit := make(chan bool)			// 用于 关闭主go程的channel
	ch := make(chan int)				// 用于 数据传递的 channel

	for i:=0; i<5; i++ {
		go readGo(ch, i+1)
	}
	for i:=0; i<5; i++ {
		go writeGo(ch,i+1)
	}

	for{
		;
	}
}

channel应用：
import (
"math/rand"
"time"
"fmt"
)

//var value06 int		// 定义全局变量，模拟共享数据
func readGo06(in <-chan int, idx int)  {
	for {
		num := <-in				// 从 channel 中读取数据
		fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
		time.Sleep(time.Second)
	}
}
func writeGo06(out chan<- int, idx int)  {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		out <- num					// 写入channel
		fmt.Printf("%dth 写go程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)		// 放大实验现象
	}
}

func main()  {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	for i:=0; i<5; i++ {			// 5 个 读 go 程
		go readGo06(ch, i+1)
	}
	for i:=0; i<5; i++ {			//
		go writeGo06(ch, i+1)
	}
	for{
		;
	}
}




条件变量：本身不是锁！！！ 但经常与锁结合使用！！

使用流程：

	1.  创建 条件变量： var cond    sync.Cond

	2.  指定条件变量用的 锁：  cond.L = new(sync.Mutex)

	3.  cond.L.Lock()	给公共区加锁（互斥量）

	4.  判断是否到达 阻塞条件（缓冲区满/空）	—— for 循环判断

	for  len(ch) == cap(ch) {   cond.Wait() —— 1) 造成当前阻塞 2) 可以完成解锁 3) 然后再加锁

	5.  访问公共区 —— 读、写数据、打印

	6.  解锁条件变量用的 锁  cond.L.Unlock()

	7.  唤醒阻塞在条件变量上的 对端。 signal()  Broadcast()


mac地址    (ARP协议)Ip ---------> mac
IP地址      ---->确定主机
port端口    ---->确定程序
	1,不能使用系统占用的默认端口(eg.80，访问Web需要。 8080，浏览器占用的端口)，我们一般用5000+端口
	2,65535为端口上限

OSI(理论上的标准，由高层到低层)

应用层：是最靠近用户的OSI层。这一层为用户的应用程序（例如电子邮件、文件传输和终端仿真）提供网络服务。

表示层：可确保一个系统的应用层所发送的信息可以被另一个系统的应用层读取。例如，PC程序与另一台计算机进行通信，
        其中一台计算机使用扩展二一十进制交换码(EBCDIC)，而另一台则使用美国信息交换标准码（ASCII）来表示相同的字符。
        如有必要，表示层会通过使用一种通格式来实现多种数据格式之间的转换。

会话层：通过传输层(端口号：传输端口与接收端口)建立数据传输的通路。主要在你的系统之间发起会话或者接受会话请求
       （设备之间需要互相认识可以是IP也可以是MAC或者是主机名）。

传输层：定义了一些传输数据的协议和端口号（WWW端口80等），常见协议有TCP/UDP协议。 主要是将从下层接收的数据进行分段和传输
        ，到达目的地址后再进行重组。常常把这一层数据叫做段。

网络层：提供路由支持，建立连接和路径选择，这一层常见协议有IP协议、ICMP协议、IGMP协议，通过路由一级一级定位目标主机位置。

数据链路层：定义了如何让格式化数据以帧为单位进行传输，以及如何让控制对物理介质的访问。并且通常还提供错误检测和纠正，
            以确保数据的可靠传输。

物理层：主要定义物理设备标准，例如各种设备，线路的接口。这一层的数据叫比特流。


TCP/IP(实际上的标准，由高层到低层)
应用层：应用程序收到“传输层”的数据，接下来就要进行解读或者格式化压缩数据公下层传输。
		该层的作用就是规定应用程序的数据格式。

传输层：端口表示接收的数据包到底供哪个程序（进程）使用。它其实是每一个使用网卡的程序的编号。
		每个数据包都发到主机的特定端口，所以不同的程序就能取到自己所需要的数据。

网络层：网络层的作用是引进一套新的地址，包含的主要信息是源IP和目的IP。IP帮助我们确定计算机所在的子网络，MAC 地址则将数据
		包送到该子网络中的目标网卡。从逻辑上是先处理网络地址，然后再处理 MAC 地址。

链路层：数据包必须是从一块网卡，传送到另一块网卡。通过网卡能够使不同的计算机之间连接，从而完成数据通信等功能。
		网卡的地址——MAC 地址，就是数据包的物理发送地址和物理接收地址。

三次握手：

1.客户端发送一个带SYN标志的TCP报文到服务器表示连接请求,这是第一个报文段。

2.服务器端接收报文后回应客户端，是三次握手中的第二个报文段，同时带ACK标志和SYN标志。ACK表示对刚才客户端SYN的回应；
同时又发送SYN标志给客户端，询问客户端是否准备好进行数据通讯。

3.客户必须再次回应服务器端一个ACK报文，这是第三个报文段。在这个过程中，客户端和服务器分别给对方发了连接请求，
也应答了对方的连接请求，其中服务器的请求和应答在一个段中发出。

标志 TCP 三次握手建立完成。 —— server：Accept() 返回 。 	—— client：Dial() 返回。




四次挥手：

1. 客户端主动发出关闭连接请求的报文， FIN标志位表示关闭连接请求

2. 服务端发出报文响应关闭连接请求，带有应答ACK标志，标识半关闭完成

3. 服务端也发出关闭连接请求的报文给客户端 ，同样带有FIN标志位

4. 客户端发出报文响应客户端关闭连接请求，带有应答ACK




TCP编程：

listen(network string, laddr *TCPAddr) laddr为本机需要监听的 ip 和port，如果为nil或者未指定ip，则监听广播
import (
	"net"
	"fmt"
	)
// TCP-CS-server
func main()  {
	// 指定服务器 通信协议、IP地址、port。 创建一个用于监听的 socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待客户端建立连接...")
	// 阻塞监听客户端连接请求, 成功建立连接，返回用于通信的socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功建立连接！！！")
	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	conn.Write(buf[:n])		// 读多少写多少。原封不动
	// 处理数据—— 打印
	fmt.Println("服务器读到数据：", string(buf[:n]))
}

// TCP-CS-client
import (
	"net"
	"fmt"
	)

func main()  {
	// 指定 服务器 IP + port 创建 通信套接字。
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 主动写数据给服务器
	conn.Write([]byte("Are you Ready?"))

	buf := make([]byte, 4096)
	// 接收服务器回发的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务器回发：", string(buf[:n]))
}







TCP-CS 并发服务器
import (
"net"
"fmt"
"strings"
)

func HandlerConnect(conn net.Conn) {

	defer conn.Close()

	// 获取连接的客户端 Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接！")

	// 循环读取客户端发送数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		//fmt.Println(buf[:n])
		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务器接收的客户端退出请求，服务器关闭")
			return
		}
		if n == 0 {
			fmt.Println("服务器检测到客户端已关闭，断开连接！！！")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器读到数据：", string(buf[:n]))	// 使用数据

		// 小写转大写，回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main()  {
	// 创建监听套接字
	//listener, err := net.Listen("tcp", "127.0.0.1:8001")
	listener, err := net.Listen("tcp", "192.168.15.78:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 监听客户端连接请求
	for {
		fmt.Println("服务器等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		// 具体完成服务器和客户端的数据通信
		go HandlerConnect(conn)
	}
}
TCP-CS 并发客户端
import (
	"net"
	"fmt"
	"os"
	)

func main()  {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "192.168.15.78:8001")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	// 获取用户键盘输入( stdin )，将输入数据发送给服务器
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			//写给服务器, 读多少，写多少！
			conn.Write(str[:n])
		}
	}()

	// 回显服务器回发的大写数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("检查到服务器关闭，客户端也关闭")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("客户端读到服务器回发：", string(buf[:n]))
	}
}


UDP编程：
import (    //服务器
	"net"
	"fmt"
	"time"
	)

func main()  {
	// 组织一个 udp 地址结构, 指定服务器的IP+port
	// 创建用户通信的 socket
	// 读取客户端发送的数据
	// 返回3个值，分别是 读取到的字节数， 客户端的地址， error
	// 模拟处理数据
	// 提取系统当前时间
	// 回写数据给客户端
}
func main()  {
	// 组织一个 udp 地址结构, 指定服务器的IP+port
	srvAddr, err := net.ResolveUDPAddr("udp","127.0.0.1:8007") //这里的端口被绑定后(func完成后)，其他程序无法再占用，
	if err != nil {												  //只能指定其他端口，eg.func运行后再次运行，端口要改变
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("udp 服务器地只结构，创建完成")

	// 创建用户通信的 socket
	udpConn, err := net.ListenUDP("udp", srvAddr) //TCP编程差不多，套接字编程
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("udp 服务器通信socket创建OVER")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)  //建立缓冲

	// 返回3个值，分别是 读取到的字节数， 客户端的地址， error
	n, cltAddr, err := udpConn.ReadFromUDP(buf)  //read每次读取大小长度为0 - len(切片)数据，返回长度值，所以每次打印读取
	if err != nil {								 //数据时，用buf[:n]取值
		fmt.Println("ReadFormUDP err:", err)
		return
	}
	// 模拟处理数据
	fmt.Printf("服务器读到了 %v 的数据：%s\n", cltAddr, string(buf[:n]))
	// 提取系统当前时间
	daytime := time.Now().String()

	// 回写数据给客户端
	_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
	if err != nil {
		fmt.Println("WriteToUDP err:", err)
		return
	}
}

import (   //客户端
	"net"
	"fmt"
	)

func main()  {
	// 指定 服务器 IP + port 创建 通信套接字。

	for i:=0; i <1000000; i++{
	// 主动写数据给服务器
	// 接收服务器回发的数据
	}

}

import (
	"net"
	"fmt"
	)

func main()  {
// 指定 服务器 IP + port 创建 通信套接字。
	conn, err := net.Dial("udp", "127.0.0.1:8007") //这个端口与服务器的端口一致
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	for i:=0; i <1000000; i++{
		// 主动写数据给服务器
		conn.Write([]byte("Are you Ready?"))

		buf := make([]byte, 4096)
		// 接收服务器回发的数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器回发：", string(buf[:n]))
	}

}

UDP并发服务器：
import (
	"net"
	"fmt"
	"time"
		)

func main() {

	// 组织一个 udp 地址结构, 指定服务器的IP+port
	// 创建用户通信的 socket
	// 读取客户端发送的数据

	for {
		// 返回3个值，分别是 读取到的字节数， 客户端的地址， error
		// --- 主go程读取客户端发送数据
		// 模拟处理数据
		fmt.Printf("服务器读到 %v 的数据：%s\n", cltAddr, string(buf[:n]))
		go func() {					// 每有一个客户端连接上来，启动一个go程 写数据。
			// 提取系统当前时间
			// 回写数据给客户端
		}()
	}
}

import (
	"net"
	"fmt"
	"time"
	)

func main() {

	// 组织一个 udp 地址结构, 指定服务器的IP+port
	srvAddr, err := net.ResolveUDPAddr("udp","127.0.0.1:8007")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("udp 服务器地址结构，创建完程!!!")

	// 创建用户通信的 socket
	udpConn, err := net.ListenUDP("udp",srvAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("服务器通信socket创建完成")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)

	for {
		// 返回3个值，分别是 读取到的字节数， 客户端的地址， error
		n, cltAddr, err := udpConn.ReadFromUDP(buf)		// --- 主go程读取客户端发送数据
		if err != nil {
			fmt.Println("ReadFromUDP err", err)
			return
		}
		// 模拟处理数据
		fmt.Printf("服务器读取到 %v 的数据：%s\n", cltAddr, string(buf[:n]))

		go func() {					// 每有一个客户端连接上来，启动一个go程 写数据。
			// 提取系统当前时间
			daytime := time.Now().String() + "\n"

			// 回写数据给客户端
			_, err := udpConn.WriteToUDP([]byte(daytime), cltAddr)
			if err != nil {
				fmt.Println("WriteToUDP err:", err)
				return
			}
		}()
	}
}








文件属性：这个应该与FTP有关联，这里案例为客户端上载文件服务端
import (
	"os"
	"fmt"
	)

func main()  {
	list := os.Args			// 获取命令行参数，这个作用意义？

	if len(list) != 2 {
		fmt.Println("格式为：go run xxx.go 文件名")
		return
	}
	// 提取文件名
	path := list[1]
	// 获取文件属性
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())
}

文件传输发送端：
import (      // 这里的案例为客户端上载文件到服务端
	"os"
	"fmt"
	"net"
	"io")

func sendFile(conn net.Conn, filePath string)  {
	// 只读打开文件

	// 从本文件中，读数据，写给网络接收端。 读多少，写多少。原封不动。

		// 写到网络socket中

}

func main()  {
	list := os.Args		// 获取命令行参数
	//
	提取 文件的绝对路径
	//提取文件名

	// 主动发起连接请求

	// 发送文件名给 接收端

	// 读取服务器回发的 OK

		// 写文件内容给服务器——借助conn
	}
}



import (
	"os"
	"fmt"
	"net"
	"io")

func sendFile(conn net.Conn, filePath string)  {
	// 只读打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer f.Close()

	// 从本文件中，读数据，写给网络接收端。 读多少，写多少。原封不动。
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送文件完成。")
			} else {
				fmt.Println("os.Open err:", err)
			}
			return
		}
		// 写到网络socket中
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}

func main()  {
	list := os.Args		// 获取命令行参数

	if len(list) != 2 {
		fmt.Println("格式为：go run xxx.go 文件绝对路径")
		return
	}
	// 提取 文件的绝对路径
	filePath := list[1]

	//提取文件名
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	fileName := fileInfo.Name()

	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 发送文件名给 接收端
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	// 读取服务器回发的 OK
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	if "ok" == string(buf[:n]) {
		// 写文件内容给服务器——借助conn
		sendFile(conn, filePath)
	}
}




文件传输接收端：
import (
	"net"
	"fmt"
	"os")

func recvFile(conn net.Conn, fileName string)  {
	// 按照文件名创建新文件

	// 从 网络中读数据，写入本地文件

	// 写入本地文件，读多少，写多少。


	func main()  {
	// 创建用于监听的socket

	// 阻塞监听

	// 获取文件名，保存

	// 回写 ok 给发送端

	// 获取文件内容
	recvFile(conn, fileName)
}
import (
	"net"
	"fmt"
	"os"
	)

func recvFile(conn net.Conn, fileName string)  {
	// 按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	// 从 网络中读数据，写入本地文件
	buf := make([]byte, 4096)
	for {
		n,_ := conn.Read(buf)
		if n == 0 {
			fmt.Println("接收文件完成。")
			return
		}
		// 写入本地文件，读多少，写多少。
		f.Write(buf[:n])
	}
}
func main()  {

	// 创建用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println(" net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(" listener.Accept() err:", err)
		return
	}
	defer conn.Close()

	// 获取文件名，保存
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(" conn.Read err:", err)
		return
	}
	fileName := string(buf[:n])

	// 回写 ok 给发送端
	conn.Write([]byte("ok"))

	// 获取文件内容
	recvFile(conn, fileName)
}



HTTP编程：

web工作方式：

1. 客户端 ——> 访问 www.baidu.com  ——> DNS 服务器。 返回 该域名对应的 IP地址。

2. 客户端 ——> IP + port ——> 访问 网页数据。（TCP 连接。 HTTP协议。）
域名系统（英文：Domain Name System，缩写：DNS）是互联网的一项服务。它作为将域名和IP地址相互映射的一个分布式数据库，能够
使人更方便地访问互联网。DNS使用协议：TCP（当请求大于512字节时）和UDP，端口53[1]。当前，对于每一级域名长度的限制是63个字
符，域名总长度则不能超过253个字符。

IP + port(网页一般指定一个默认页面或者端口供客户端访问，网站的文件夹里面一般有个默认HTML页面)
上网过程大概如下：浏览器本身作为一个客户端，输入URL时，浏览器首先去请求DNS服务器，通过DNS获取相应的域名对应的IP，然后通
过IP地址找到IP对应的服务器后，要求建立TCP连接，等浏览器发送完HTTP Respuest包(请求包)后，服务器接收到请求包才开始处理
请求包，服务器调用自身服务，返回HTTP Response包(响应包)；客户端收到来自服务器的响应开始渲染Response包里面的主体(body)，
等收到全部的内容随后断开与该服务器的TCP连接。(请求包，相应包都为报文)

http和URL：

http 超文本传输协议。规定了 浏览器访问 Web服务器 进行数据通信的规则。  http（明文） -- TLS、SSL -- https（加密）

URL：统一资源定位。 在网络环境中唯一定位一个资源数据。  浏览器地址栏内容。

获取http请求服务器：这个为获取浏览器发过来的请求包，最后打印读取浏览器的请求包
import (
	"net"
	"fmt"
	"os"
	)

func errFunc(err error, info string)  {
	if err != nil {
		fmt.Println(info, err)
		//return					// 返回当前函数调用
		//runtime.Goexit()			// 结束当前go程
		os.Exit(1)					// 将当前进程结束。
	}
}

func main()  {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "net.Listen err:")
	defer listener.Close()


	conn, err := listener.Accept()
	errFunc(err, "Accpet err:")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "conn.Read")

	fmt.Printf("|%s|\n", string(buf[:n]))
}

|GET / HTTP/1.1  // 请求行
Host: 127.0.0.1:8007
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng, "*/*;q=0.8,application/signed-exchange;v=b3
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9
//这里为空行，上面Host - Accept-Language为请求头
//这里紧接着为请求包体

http请求包：

请求行：请求方法（空格）请求文件URL（空格）协议版本（\r\n）

		GET、POST // get是发送请求HTTP协议通过url参数传递进行接收,而post是实体数据,可以通过表单提交大量信息.

请求头：语法格式 ： key ：value

空行：\r\n	 —— 代表 http请求头结束。

请求包体：请求方法对应的数据内容。 GET方法没有内容！！


回调函数就是一个通过函数指针调用的函数。如果你把函数的指针（地址）作为参数传递给另一个函数，当这个指针被用来调用其所
指向的函数时，我们就说这是回调函数。回调函数不是由该函数的实现方直接调用，而是在特定的事件或条件发生时由另外的一方调用的
，用于对该事件或条件进行响应。
机制：定义一个回调函数
      提供函数实现的一方在初始化的时候，将回调函数的函数指针注册给调用者
      当特定的事件或条件发生时，调用者使用函数指针调用回调函数对事件进行处理

http应答包大概步骤：
1. 使用 net/http包 创建 web 服务器

	1） 注册回调函数。

		http.HandleFunc("/itcast", handler)

		参1：用户访问文件位置

		参2：回调函数名 —— 函数必须是 (w http.ResponseWriter, r *http.Request) 作为参数。

	2）绑定服务器监听地址。

		http.ListenAndServe("127.0.0.1:8000", nil)
2. 回调函数：

	本质：函数指针。通过地址，在某一特定位置，调用函数。

	在程序中，定义一个函数，但不显示调用。当某一条件满足时，该函数由操作系统自动调用。
import (
	"net"
	"net/http"
	)

func handler(w http.ResponseWriter, r *http.Request) {
	// w：写回给客户端（浏览器）的数据
	// r: 从 客户端 浏览器 读到的数据
}

func main() {

	//注册回调函数。该回调函数在服务器被访问时，自动被调用。第一个参数是路由匹配的字符串，代表实现了这个路由？
	http.HandleFunc("/itcast",handler) // 因为函数名全球唯一，函数名为地址。第二个handler类型的参数应该为地址传递

	//绑定服务器地址
	http.ListenAndServe("127.0.0.1:8000",nil) //这个函数绑定服务器地址,第二个参数一般为nil，这样
												//系统调用默认的回调函数，这样不用我们指定了
}

HTTP应答包格式：

	状态行：协议版本号（空格）状态码（空格）状态码描述（\r\n）

	响应头：语法格式 ： key ：value

	空行：\r\n

	响应包体： 请求内容存在： 返回请求页面内容

	请求内容不存在： 返回错误页面描述。



Go HTTP服务器：
import (
	"net/http"
	"fmt"
	)

func myHandle(w http.ResponseWriter, r *http.Request) {
	// w : 写给客户端的数据内容
	w.Write([]byte("this is a Web server"))

	// r: 从客户端读到的内容
	fmt.Println("Header:", r.Header)
	fmt.Println("URL:", r.URL)
	fmt.Println("Method:", r.Method)
	fmt.Println("Host:", r.Host)
	fmt.Println("RemoteAddr:", r.RemoteAddr)
	fmt.Println("Body:", r.Body)
}

func main()  {
	// 注册回调函数， 该函数在客户端访问服务器时，会自动被调用
	//http.HandleFunc("/itcast", myHandle)
	http.HandleFunc("/", myHandle)

	// 绑定服务器地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
这个程序运行后在浏览器输入127.0.0.1:8000，页面返回为"this is a Web server",goland打印出如下结果：
Header: map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,***
URL: /
Method: GET
Host: 127.0.0.1:8000
RemoteAddr: 127.0.0.1:55093
Body: {}
Header: map[Accept:[image/webp,image/apng,image/*,*/*;q=0.8] Accept-Encoding:[gzip, deflate, br] Accept-Language:[zh-CN,
URL: /favicon.ico
Method: GET
Host: 127.0.0.1:8000
RemoteAddr: 127.0.0.1:55093
Body: {}




HTTP WEB服务器Flow：

1.  注册回调函数：http.HandleFunc("/", myHandler)

	func myHandler(w http.ResponseWriter, r *http.Request)

		w：给客户端回发数据

		r：从客户端读到的数据

	type ResponseWriter interface {
		Header() Header
		Write([]byte) (int, error)
		WriteHeader(int)
	}

	type Request struct {
		Method string		// 浏览器请求方法 GET、POST…
		URL *url.URL		// 浏览器请求的访问路径
		……
		Header Header		// 请求头部
		Body io.ReadCloser	// 请求包体
		RemoteAddr string	// 浏览器地址
		……
		ctx context.Context
	}

	2. 绑定服务器地址结构：http.ListenAndServe("127.0.0.1:8000", nil)

		参2：通常传 nil 。 表示  让服务端 调用默认的 http.DefaultServeMux 进行处理

import (
	"net/http"
	"fmt"
	"os"
	)

func OpenSendFile(fNmae string, w http.ResponseWriter)  {
	pathFileName := "C:/itcast/test" + fNmae
	f, err := os.Open(pathFileName)
	if err != nil {
		fmt.Println("Open err:", err)
		w.Write([]byte(" No such file or directory !"))
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf)			// 从本地将文件内容读取。
		if n == 0 {
			return
		}
		w.Write(buf[:n])			// 写到 客户端（浏览器）上
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客户端请求：", r.URL)
	OpenSendFile(r.URL.String(), w)
}

func main()  {
	// 注册回调函数
	http.HandleFunc("/", myHandler)
	// 绑定监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}


HTTP WEB客户端Flow：

	1. 获取web服务器数据：

		func Get(url string) (resp *Response, err error)

		返回：http应答包，保存成 struct

		type Response struct {
			Status     string // e.g. "200 OK"
			StatusCode int    // e.g.  200
			Proto      string // e.g. "HTTP/1.0"
			……
			Header Header
			Body io.ReadCloser
			……
		}
	2. defer resp.Body.Close()

	3. for 循环提取 Body 数据：

		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("--Read finish!")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:", err)
			return
		}


http 客户端：
import (
	"net/http"
	"fmt"
	"io"
	)

func main()  {
	// 使用Get方法获取服务器响应包数据
	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000/hello")
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()

	// 获取服务器端读到的数据
	fmt.Println("Status = ", resp.Status)           // 状态
	fmt.Println("StatusCode = ", resp.StatusCode)   // 状态码
	fmt.Println("Header = ", resp.Header)           // 响应头部
	fmt.Println("Body = ", resp.Body)               // 响应包体


	buf := make([]byte, 4096)         // 定义切片缓冲区，存读到的内容
	var result string
	// 获取服务器发送的数据包内容
	for {
		n, err := resp.Body.Read(buf)  // 读body中的内容。
		if n == 0 {
			fmt.Println("--Read finish!")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:", err)
			return
		}

		result += string(buf[:n])     // 累加读到的数据内容
	}
	// 打印从body中读到的所有内容
	fmt.Println("result = ", result)
}


MYSQL:
连接数据库方法：mysql -h127.0.0.1 -P3306 -uroot -proot
其中：-h主机地址(如果mysql客户端在本地，-h可以省略)，-P端口号(如果默认为3306，-P可以省略)，-u用户名 -p密码
简化为：mysql -uroot -proot
关闭数据库：exit;  quit;  \p;   数据库命令最好以分号; 结束。因为MySql中命令必须以;结束，windows不一定要加;
编码方式：
gbk：		简体中文
gb2312：	简体中文
utf8：	    通用字符编码
命令窗口：为什么以mysql命令开始了，因为安装的路径里面有个bin，存放数据库的命令文件
常用的栏位描述(应该都要大写吧)：
null | not null   	空|非空
default	        默认值
auto_increment    	自动增长
primary key       	主键
comment           	备注
engine           	引擎   innodb(这个比较常用)  myisam  memory  引擎是决定数据存储的方式
一个比较复杂的表定义：
create table 数据库名.表名，用于给指定的数据库创建表，这个等同于建立shema，然后操作对象为schema.表名, 建议用这个吧

mysql> create table data.table(  #给data数据库中创建stu表
    -> id int auto_increment primary key comment '主键', 自动增长字段不用插入,插入null即可
    -> name varchar(20) not null comment '姓名',
    -> phone varchar(20) comment '电话号码',
    -> `add` varchar(100) default '地址不详' comment '地址'
    -> )engine=innodb;
这个直接enter应该报错的，前面输入命令：set names gbk; 成功后再建表操作
Query OK, 0 rows affected (0.00 sec)

drop table (if exists)表1，表2
truncate 表名称   --> 直接删除并且重构表
delete

INSERT语句：插入一笔记录到表里面，全部栏位赋值
INSERT INTO table_name
VALUES (value1,value2,value,...)

INSERT INTO table_name (colume1,colume2,colume,...) //这里字段名

// 顺序可以与表结构不一致,可以插入部分字段，不过非空必须给值，
VALUES (value1,value2,value,...)  // 值的顺序必须和插入字段的顺序一致。通过default关键字插入默认值
插入多笔记录：
insert into person (id, personCode, personName , telNumber)
values
(1,'5112403' , '张三1' , '1378902134'),
(2,'5112404' , '张三2' , '1378902135'),
(3,'5112405' , '张三3' , '1378902137');

UPDATE语句:
UPDATE 表名称 set 字段 = 值 [where条件]  //不带where将改变全部数据

DELETE from 表名称 [where条件]  //不带where将删除全部数据,只是删除全部记录,truncate删除并且重建表，drop删除了不重建

 SELECT语句：这个是重点，优良的SQL语句体现在这里，IN/NOT IN替换为EXISTS/not exists提高不少效率

 select * from tablename where colume = 'value'; 这个为一般格式，*哟几个栏位一般返回结果就有几个栏位

全表遍历会非常占资源，素以引入索引的概念,
INDEX优势:
  1，提高数据的检索效率
  2，降低排序成本
缺点：
  1，索引是数据，需要占用磁盘空间
  2，索引提高了查询速度，但是会降低更新速度
  3，最优索引需要花时间研究建立

四种方法添加数据表的索引：
  ALTER TABLE tbl_name ADD PRIMARY KEY(colume)该语句添加一个主键，意味着索引必须是唯一的，并且不能为NULL。
  ALTER TABLE tbl_name ADD UNIQUE index_name (colume)：这条语句创建索引的值必须是唯一的(除了NULL外,NULL可能出现多次 )
  ALTER TABLE tbl_name ADD INDEX index_name(colume)：添加普通索引，索引值可以出现多次
  ALTER TABLE tbl_name ADD FULLTEXT index_name(colume)：该语句指定了索引为FULLTEXT，用于全文索引。
一般创建索引的语句为：
  CREATE [UNIQUE] INDEX 索引名 ON 表名(字段名(长度 )) ：如果CHER，VARCHAR类型，长度可以小于字段实际长度，
  如果BLOB，TEXT类型，必须指定长度。
  Unique约束  唯一约束
  唯一约束保证在一个字段或者一组字段里的数据与表中其它行的数据相比是唯一的。一个表可以有多个UNIQUE,只能有一个PRIMARY KEY
  PRIMARY KEY 拥有自动定义的 UNIQUE 约束，主键列不允许空值，而唯一性索引列允许空值
  主键约束比唯一索引约束严格，当没有设定主键时，非空唯一索引自动称为主键。对于主键和唯一索引的一些区别主要如下：
    1.主键不允许空值，唯一索引允许空值
    2.主键只允许一个，唯一索引允许多个
    3.主键产生唯一的聚集索引，唯一索引产生唯一的非聚集索引
  注：聚集索引确定表中数据的物理顺序，所以是主键是唯一的（聚集就是整理数据的意思）
适合建立索引：
1，主键自动创建唯一索引
2，频繁作为查询条件的
3，外键关系建立索引
4，高并发下倾向创建组合索引
5，查询中排序的字段
6，查询中统计或分组的字段

不适合建立索引：
1，表记录太少
2，经常增删改的表
3，数据重复且分布平均的字段
4，where条件里用不到的字段

CREATE TABLE test (
id INT NOT NULL,
last_name CHAR(30) NOT NULL,
first_name CHAR(30) NOT NULL,
PRIMARY KEY (`id`),
INDEX name(last_name, first_name)
);
这里的索引为组合索引(最左前缀)，查询的时候可以单独在最左查询，也可以组合索引查询，不能单独使用右边查询这样会导致全表遍历
SELECT * FROM test WHERE last_name = 'Wi';
SELECT * FROM test WHERE last_name = 'Wi' AND first_name = 'iuotr';


如果索引了多列，要遵守最左前缀法则。即查询条件列从索引的最左前列开始匹配并且不跳过索引中的列，单列索引不存储null值，复
合索引不存储全为null的值。索引不能存储Null，所以对这列采用is null条件时，因为索引上根本没Null值，不能利用到索引，只能
全表扫描。
Like “%XX“，由于前面是模糊的，所以不能利用索引的顺序，必须一个个去找，看是否满足条件
这些情况导致索引失效：
	1.条件中有or（两个不同字段）
	2.对于多列索引，不是使用的第一部分
	3.如果列类型是字符串，未将在条件中使用的数据用引号引用起来
	4.在索引列上做任何操作（计算、函数、(自动or手动)类型转换）

优化SQL(Optimize Sql)：
应尽量避免在 where 子句中对字段进行表达式或者函数操作，这将导致引擎放弃使用索引而进行全表扫描。

尽量避免在 where 子句中使用 or 来连接条件，否则将导致引擎放弃使用索引而进行全表扫描

尽量避免在 where 子句中使用!=或<>操作符，否则将导致引擎放弃使用索引而进行全表扫描

应尽量避免在 where 子句中对字段进行 null 值判断，否则将导致引擎放弃使用索引而进行全表扫描

in 和 not in 也要慎用，否则会导致全表扫描，对于连续的数值，能用 between 就不要用 in 了

应尽量使用数字型字段，若只含数值信息的字段尽量不要设计为字符型，这会降低查询和连接的性能，并会增加存储开销

在新建临时表时，如果一次性插入数据量很大，那么可以使用 select into 代替 create table，避免造成大量 log ，以提高速度

尽量避免使用游标，因为游标的效率较差，如果游标操作的数据超过1万行，那么就应该考虑改写

尽可能使用 varchar/nvarchar 代替 char/nchar ，因为首先变长字段存储空间小，可以节省存储空间，其次对于查询来说，
在一个相对较小的字段内搜索效率显然要高些

如果使用到了临时表，在存储过程的最后务必将所有的临时表删除，先 truncate table ，然后 drop table ，这样可以避免系统表的较
长时间锁定

Mysql分库(Mysql Sub-treasury)：
定义：一个库里表太多了，导致海量数据，系统性能下降，把原本存储于一个库的表拆分存储到多个库，通常是将表按照功能模块、关系
	密切程度划分出来，部署到不同库上。

原因：写入或者说大数据、频繁的写入操作对master性能影响就比较大，这个时候，单库并不能解决大规模并发写入的问题

优点：减少增量数据写入时的锁对查询的影响。由于单表数量下降，常见的查询操作由于减少了需要扫描的记录，
	使得单表单次查询所需的检索行数变少，减少了磁盘IO，时延变短。

垂直拆分：通常是按照业务功能的使用频次，把主要的、热门的字段放在一起做为主要表；然后把不常用的，
	按照各自的业务属性进行聚集，拆分到不同的次要表中；主要表和次要表的关系一般都是一对一的。

水平拆分(数据分片) ：Mysql单表的容量不超过500W，否则建议水平拆分。

GO语言操作数据库:(将MySql-Front打开，指定要用的数据库名称以表)
package main

import (
"database/sql"
_"github.com/go-sql-driver/mysql"
"fmt"
)

func main(){
//                                   用户名:密码@[连接方式](主机名:端口号)/数据库名
db,_ := sql.Open("mysql","root:root@(127.0.0.1:3306)/stu")
defer db.Close()  // 关闭数据库
// open()在执行时不会真正的与数据库进行连接，只是设置连接数据库需要的参数。
err := db.Ping() //Ping方法才连接数据库
if err != nil {
fmt.Println("数据库链接失败")
return
}else {
fmt.Println("数据库联接成功")
}

//操作一：执行数据操作语句
//sql := "INSERT INTO student VALUES (4,'ttt')"
//result,_ := db.Exec(sql)  // 执行sql语句
//n,_ := result.RowsAffected();  //获取受影响的记录数
//fmt.Printf("%T",result)
//fmt.Println("受影响的记录数是：",n)

//操作二：执行预处理
//stu := [2][2] string{{"8","str"},{"9","rose"}}
//stmt,_ := db.Prepare("INSERT INTO student VALUES (?,?)")
//for _,s := range stu {
//	stmt.Exec(s[0],s[1])
//}

//操作三：单行查询
//var id,name string
//rows := db.QueryRow("SELECT * from student where id = 5")
//rows.Scan(&id,&name)      // 将rows中的数据存到id, name中
//fmt.Println(id,"--", name)

//操作四：多行查询
//var id,name string
//rows,_ := db.Query("SELECT * from student")  //获取所有数据
//for rows.Next() {
//	rows.Scan(&id, &name)
//	fmt.Println(id, "--", name)
//}

}


HTML是 HyperText Mark-up Language 的首字母简写，超文本标记语言，超文本指的是超链接，标记指的是标签，是一种用来制作网页的
语言，这种语言由一个个的标签组成，用这种语言制作的文件保存的是一个文本文件，文件的扩展名为html或者htm。
下面是HTML的特征代码：  <div>这是成对出现的标签</div>。
大致的结构如下： <!-- comment --> //这个为注释的格式
<!COCTYPE html>   // 这为文档声明,申明此文档是一个html5的文档
<html lang="en"> // “<html>”标签和最后一行“</html>”定义html文档的整体，lang指定网页的语言，
	<head>       //  html文档的设置标签，文档的设置及资源的引用都写在这个标签中
		<meta charset="UTF-8">  // <!-- 设置网页编码 -->
		<title>网页标题</title>  // <!-- 设置网页的标题 -->
	</head>
	<body>  //  html文档主体标签，网页上显示的内容，都写在这个标签中
		网页显示内容
	</body>
</html>  //  <!-- html文档结束标签 -->
“<head>”标签和“<body>”标签是它的第一层子元素，“<head>”标签里面负责对网页进行一些设置以及定义标题，设置
包括定义网页的编码格式，外链css样式文件和javascript文件等，设置的内容不会显示在网页上，标题的内容会显示在标题栏，
“<body>”内编写网页上显示的内容。  一个html文件就是一个网页，html文件用编辑器打开显示的是文本，可以用文本的方式编辑它，
如果用浏览器打开，浏览器会按照标签描述内容将文件渲染成网页。

<head>
	<meta charset="UTF-8">
	<!-- 设置网页在移动端观看时，网页不缩放 -->
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<!-- 设置网页在IE上观看时，以IE的最高版本渲染网页 -->
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Document</title>
	<!-- 第三种引入样式的方式：外链式 推荐使用 -->
	<link rel="stylesheet" href="css/main.css">

	<!-- 第二种引入样式的方式：内嵌式 部分推荐使用，一般首页使用 -->
	<style>
		div{width:200px;height:200px;background:red}
	</style>
	<style>
	.news_con{  //下面的子类里属性，如果有同名的就用子类自己的，没有定义的就默认用这个父类的
		width:285px;
		height:310px;
		/* 设置顶部边框 */
		border-top:1px solid #c8dcf2;
		/* 设置底部边框 */
		border-bottom:1px solid #c8dcf2;
	}
	.news_con div{
		/* 宽度width不用设置，默认等于父级的宽度 */
		height:40px;
	}
	/* 层级选择器可以越级选择 */
	.news_con h3{
		/* 设置浮动让h3和a标签可以排列在一行 */
		float:left;
		/* 清除h3标题默认的外边距 */
		margin:0px;
		/* 设置文字颜色 */
		color:#172c45;
		/* 设置文字大小 */
		font-size:16px;
		/* 设置文字字体 */
		font-family: 'Microsoft YaHei';
		/* 设置文字不加粗 */
		font-weight:normal;
		/* 通过设置行高来将单行文字垂直居中 */
		line-height:40px;
	}

	.news_con a{
		float:right;
		/* 设置文字颜色 */
		color:#172c45;
		/* 设置文字大小 */
		font-size:12px;
		/* 设置文字字体 */
		font-family: 'Microsoft YaHei';
		/* 通过设置行高来将单行文字垂直居中 */
		line-height:40px;
		/* 去掉文字默认的下划线 */
		text-decoration:none;
	}
	.news_con p{
		/* 设置文字颜色 */
		color:#737373;
		/* 设置文字大小 */
		font-size:12px;
		/* 设置文字字体 */
		font-family: 'Microsoft YaHei';
		line-height:20px;
		/* 设置首行缩进 */
		text-indent:24px;
		/* 清除p标签默认的外边距 */
		margin:0px;
		margin-top:10px;
	}
	</style>
</head>
&nbsp; 指定一个空格，&ensp; 指定一个TAB，长度为两个空格，&emsp;  指定一个TAB，长度为四个空格
&lt; 实际为< (小于号)，&gt; 实际为> (大于号), <p> </p>之间为一行内容？
<a href="某一个常用标签网页.html">去往某个指定的网页</a> 跳到一个指定的网页。"去往一个指定的网页"显示在网页指定位置
<a href="#">默认不确定的地址</a>  <!-- 设置缺省链接地址 -->
<br> 标签是空标签，只是简单地开始新的一行，而不是分割段落。而当浏览器遇到 <p> 标签时，通常会在相邻的段落之间插入一些垂直
的间距。在 HTML 中，<br> 标签没有结束标签。在 XHTML 中，<br> 标签必须被正确地关闭，比如这样：<br />。
<!-- 创建图片标签 -->
<img src="images/logo.png" alt="传智logo" width="100">  <! 这些图片以及链接必须放在本文件夹image等文件夹下面>
margin:20px;  同时设置四个边的margin值(边框外面的白边，例如下边距为下一行与下边的距离)
padding:100px; 同时设置四个边的padding值(内边框与文字之间的填充，例如换行后该行最后一个字与边框之间)
border:100px solid black; 内边框设置，颜色也可以设置
这个三个参数以上都为同时设置(同时设置顺序为顺时针，上-右-底-左)，分别设置方法为：参数-top/right/bottom/left; 像素值。
<li>标签用于表示文档中列表的项目，例如：一个有序列表（<ol>）和一个无序列表(<ul>)：
有序列表：
<ol>
	<li>Coffe<>
	<li>Coffe<>
	<li>Coffe<>
</ol>
无序列表：
<ul>
	<li>Coffe<>
	<li>Coffe<>
	<li>Coffe<>
</ul>
有序列表打印出三行Coffee，并且每个Coffe前面带有数字seq，无序列打印出三行Coffee，每个Coffe前面只带有符号。
Text 对象 代表 HTML 表单中的文本输入域。在 HTML 表单中 <input type="text"> 每出现一次，Text 对象就会被创建。
该元素可创建一个单行的文本输入字段。当用户编辑显示的文本并随后把输入焦点转移到其他元素的时候，会触发 onchange 事件句柄。
您可以使用 HTML <textarea> 标记来创建多行文本输入。参阅 Textarea 对象。
对于掩码文本输入，把 <input type="text"> 中的 type 设置为 "password"。参阅 Input Password。
您可以通过表单的 elements[] 数组来访问文本输入域，或者通过使用 document.getElementById()。
<div> 可定义文档中的分区或节（division/section）。可以把文档分割为独立的、不同的部分。它可以用作严格的组织工具，
并且不使用任何格式与其关联。如果用 id 或 class 来标记 <div>，那么该标签的作用会变得更加有效。
用法
<div> 是一个块级元素。这意味着它的内容自动地开始一个新行。实际上，换行是 <div> 固有的唯一格式表现。可以通过 <div> 的
class 或 id 应用额外的样式。不必为每一个 <div> 都加上类或 id，虽然这样做也有一定的好处。可以对同一个 <div> 元素应用
class 或 id 属性，但是更常见的情况是只应用其中一种。这两者的主要差异是，class 用于元素组（类似的元素，或者可以理解为
某一类元素），而 id 用于标识单独的唯一的元素。
<form> 标签用于为用户输入创建 HTML 表单。
表单能够包含 input 元素，比如文本字段、复选框、单选框、提交按钮等等。表单还可以包含 menus、textarea、fieldset、legend 和
label 元素。表单用于向服务器传输数据。form 元素是块级元素，其前后会产生折行
float属性：在前端中，很多人会使用float属性去进行图文混排。其实float属性就两个属性，一个是左一个是右，left是表示该元素向
左浮动，right表示该元素是向右浮动。在一般情况下，元素不会自己浮动的。eg. float:left;
getElementById() 方法可返回对拥有指定 ID 的第一个对象的引用。
HTML DOM 定义了多种查找元素的方法，除了 getElementById() 之外，还有 getElementsByName() 和 getElementsByTagName()。
不过，如果您需要查找文档中的一个特定的元素，最有效的方法是 getElementById()。在操作文档的一个特定的元素时，最好给该元素
一个 id 属性，为它指定一个（在文档中）唯一的名称，然后就可以用该 ID 查找想要的元素。
<body>
	<!-- 创建标题标签 -->
	<h1>一级标题</h1>
	<h2>二级标题</h2>
	<h3>三级标题</h3>
	<h4>四级标题</h4>
	<h5>五级标题</h5>
	<h6>六级标题</h6>
	<!-- 创建div标签 -->
	<div>这是第一个div</div>
	<div>这是第二个div</div>
	<div>
		<h3>这是一个h3标题</h3>
		<p>本人叫张山，毕业于某大学&nbsp;&nbsp;&nbsp;&nbsp;计算机科学与技术专业，今年23岁,
		本人性格开朗、稳重、待人真诚、热情。有较强的组织能力和</p>
	</div>
</body>

<head>
<style>
	.box{
		width:200px;
		height:200px;
		background:gold;
		position: fixed;
		left:0px;
		top:100px;
	}
</style>
<script>
	window.onload = function(){
		var oBox = document.getElementById('box');
		var iLeft = 0;

		var timer = setInterval(fnMove,30);

		function fnMove(){
			iLeft += 3;  // 每次位移多少个偏移量
			if(iLeft>600){   // 这个控制总位移的大小
				clearInterval(timer);  // 执行这个代表物体移动停止在当前位置
			}
			oBox.style.left = iLeft + 'px';   // 从左到右移动
		}
	}
</script>
</head>
<body>
	<div class="box" id="box"></div>
</body>
物体在网页内左右不停移动
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="ie=edge">
<title>Document</title>
<style>
	.box{
		width:200px;
		height:200px;
		background:gold;
		position: fixed;
		left:0px;
		top:100px;
	}
</style>
<script>
	window.onload = function(){
		var oBox = document.getElementById('box');
		var iLeft = 0;
		var iSpeed = 5;

		var timer = setInterval(fnMove,30);

		function fnMove(){
			iLeft += iSpeed;

			if(iLeft>600){
				iSpeed = -5;
			}

			if(iLeft<0){
				iSpeed = 5;
			}
			oBox.style.left = iLeft + 'px';
		}
	}
</script>
</head>
<body>
	<div class="box" id="box"></div>
</body>

// 事件委托的写法
$ul.delegate('li','click',function(){
	$(this).css({'background':'red'});
});

<script>
// 直接量的方式定义对象
var person = {
	name:'Tom',
	age:18,
	showName:function(){
		alert('我的名字是：' + this.name )
	},
	showAge:function(){
		alert('我的年龄是：' + this.age)
	}
}
// 使用对象
// 调用对象的属性：
alert(person.name);
// 调用对象的方法：
person.showName();
</script>

<script>
	// 定义一个json数据
	var sJson = '{\
		"name":"tom",\
		"age":18,\
		"hobby":["study","shopping","singing"],\
		"school":{"name":"beida","location":"beijing"}\
	}';

	// 将json数据转化成JavaScript对象，获取对象的属性
	var oJson = JSON.parse(sJson);
	alert( oJson.name );
	alert( oJson.hobby[0] );
	alert( oJson.school.location );
</script>


<script>
// 用var 定义变量，变量有预解析的特性
//alert(iNum01);
	var iNum01 = 12;
	alert(iNum01);

	// let 定义的变量没有预解析，下面这一句出错
	// alert(iNum02);
	let iNum02 = 24;
	alert(iNum02);

	iNum02 = 25;
	alert(iNum02);
	// const定义的变量没有预解析，下面这一句出错
	//alert(iNum03);
	const iNum03 = 36;
	alert(iNum03);
	// const 定义的是常量，它的值不能修改，下面这一句出错
	// iNum03 = 37;
	alert(iNum03)
</script>

结构赋值-字符串模板:
<script>
	let aList = [1,2,3];
	let [a,b,c] = aList; //let 定义的变量没有预解析
	let arr = [1,2,3];
	// 下面这种方式不能复制数组，为啥arr2打印出来了？
	let arr2 = arr;
	// 用for循环复制数组
	let arr3 = [];
	for(var i=0;i<arr.length;i++){
		arr3.push(arr[i]); // push等同于追加?
	};
	// 通过扩展运算符复制数组
	let arr4 = [...arr,5]; //等同于原数组arr + 5
	//alert(a);
	//alert(b);

	let person = {name:"tom",age:18};
	let {name,age} = person;  //这里定义了一个啥？里面有两个变量，不是原来字典的键值对
	alert(name);
	alert(age);

	function fnAlertPerson({name,age}){
		alert(
		`我的名字是：${name},
		我的年龄是：${age}`
		);
	}
	fnAlertPerson(person);
</script>
let person = {
	name:"tom",
	age:18,
	showname:function(){
		setTimeout(()=>{ // 改成箭头函数的形式，当前的this就是指的person对象，如果是一般匿名函数，this指的是window对象
			alert(this.name);
		},2000)
	}
}
person.showname();
function fnAdd(a,b){
	var rs = a + b;
	alert(rs);
}
// 通过匿名函数赋值的形式定义函数
let fnAdd = function(a,b){
	 var rs = a + b;
	 alert(rs);
}
// 匿名函数可以改成箭头函数的形式
let fnAdd =(a,b)=>{
	 var rs = a + b;
	 alert(rs);
}
fnAdd(10,5);
let fnMyalert = function(a){
	alert(a);
}
let fnMyalert2 = function(){
	alert('hi!');
}
//箭头函数的形式
let fnMyalert =a=>{
	alert(a);
}
let fnMyalert2 =()=>{
	alert('hi!');
}
let fnReturn = (a,b)=>a + b;
let fnReturn2 = (a,b)=>({"name":a,"age":b});
let iRs = fnReturn(20,30);
let iRs2 = fnReturn2('tom',18);
alert(iRs);
// 对象可以简写成下面的形式
let person = {
	name,
	age,
	showname(){
		alert(this.name);
	},
	showage(){
		alert(this.age);
	}
}

person.showname();
person.showage();

// 定义一个类
class Person {
	constructor(name,age){
		this.name = name;
		this.age = age;
	}
	showname(){
		alert('我的名称是：'+ this.name)
	}
	showage(){
		alert('我的年龄是：'+ this.age)
	}
}
// 通过类来实例化一个对象
let Andy = new Person('刘德华',55);
// 调用对象上面的方法：
Andy.showname();
Andy.showage();

// 定义类继承Person
class Student extends Person{
	constructor(name,age,school){
		super(name,age);  //  super 关键字用于访问父对象上的函数。
		this.school = school;
	}
	showschool(){
		alert('我的学校是：' + this.school);
	}
}
// 通过Student类来实例化对象
var Xiaoming = new Student('小明',15,'北京一中');

Xiaoming.showname();
Xiaoming.showage();
Xiaoming.showschool();
// concat会返回一个新数组
var aList2 = aList.concat('d','f');
var aList3 = [1,2];

var aList6 = aList.concat(aList3);
React里面<script type="text/babel"> 意义为标签的 type 属性为 text/babel,作用为要里面的代码不解析 javascript
(而默认 type="text/javascript")然后用框架来重新解析里面的代码，react 框架解析用第三方Babel编译器，
browser.js是babel编译器的浏览器版本。因为 React 独有的 JSX 语法，跟 JavaScript 不兼容。凡是使用 JSX 的地方，都要加上
type="text/babel"






beego：流程 => MVC + ORM + 额外的内容(加密解密 + 隐藏域传值 + URL传值 + 删除JS)  用bee运行项目，项目自带热更新（是现在后
台程序常用的一种技术，即在服务器运行期间，可以不停服替换静态资源。替换go文件时会自动重新编译。）


MVC:(M:model, V:view, C:controller)

Model：用于操作数据，例如如果要访问数据库
View：1.view本质就是个html。所以能在浏览器显示. 2.能够接收控制器传递过来的数据
Controller：1.能够给视图传递数据. 2.能够指定视图，联动M & V



Beego运行流程分析
--浏览器发出请求：
  路由拿到请求，并给相应的请求指定相应的控制器
  找到指定的控制器之后，控制器看是否需要查询数据库
  如果需要查询数据库就找model取数据
  如果不需要数据库，直接找view要视图
  控制器拿到视图页面之后，把页面返回给浏览器
--根据文字流程分析代码流程：
  从项目的入口main.go开始
  找到router.go文件的Init函数
  找到路由指定的控制器文件default.go的Get方法
  然后找到指定视图的语法，整个项目就串起来啦。
其中，beego通过内部语法给不同的http请求指定了不同的方法，因为我们是从浏览器地址栏发送的请求，属于get请求，
所以调用的是Get方法。
用bee运行项目，项目自带热更新（是现在后台程序常用的一种技术，即在服务器运行期间，可以不停服替换静态资源。替换go文件时会
自动重新编译。）
  路由的作用：根据不同的请求指定不同的控制器
路由函数： beego.Router("/path",&controller.Main.Controller{})
函数参数：
先分析一下Url地址由哪几部分组成？
http://192.168.110.71:8080/index
http://地址:端口/资源路径
beego.Router("/path",&controller.Main.Controller{})
第一个参数：资源路径，也就是 / (or + 后面的内容)
第二个参数：需要指定的控制器指针
了解上面的内容之后我们来看几个简单的例子：
beego.Router("/", &controllers.MainController{})
beego.Router("/index", &controllers.IndexController{})
beego.Router("/login", &controllers.LoginController{})
beego.Router("/simple",&SimpleController{},"get:GetFunc; post:PostFunc")

  高级路由设置：一般在开发过程中，我们基本不使用beego提供的默认请求访问方法，都是自定义相应的方法。那我们来看一下如何来
自定义请求方法。自定义请求方法需要用到Router的第三个参数。这个参数是用来给不同的请求指定不同的方法。具体有如下几种情况。
一个请求访问一个方法(也是最常用的)，请求和方法之间用  : 隔开，不同的请求用  ; 隔开:
beego.Router("/simple",&SimpleController{},"get:GetFunc;post:PostFunc")
可以多个请求，访问一个方法 ，请求之间用 , 隔开，请求与方法之间用 : 隔开：
beego.Router("/api",&RestController{},"get,post:ApiFunc")
所有的请求访问同一个方法，用 * 号代表所有的请求，和方法之间用 : 隔开：
beego.Router("/api/list",&RestController{},"*:ListFood")
如果同时存在 * 和对应的 HTTP请求，那么优先执行 HTTP请求所对应的方法，例如同时注册了如下所示的路由：
beego.Router("/simple",&SimpleController{},"*:AllFunc;post:PostFunc")
那么当遇到Post请求的时候，执行PostFunc而不是AllFunc。如果用了自定义方法之后，默认请求将不能访问。

	Go操作MySQL数据库：
导包：import "github.com/go-sql-driver/mysql"
连接数据库，用sql.Open()方法,open()方法的第一个参数是驱动名称,第二个参数是用户名:密码@tcp(ip:port)/数据库名称?编码方式,
返回值是连接对象和错误信息，例如：
conn,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
defer conn.Close()//随手关闭数据库是个好习惯
执行数据库操作，这一步分为两种情况，一种是增删改，一种是查询，因为增删改不返回数据，只返回执行结果，查询要返回数据，
所以这两块的操作函数不一样。
查询操作
用的函数是Query(),参数是SQL语句，返回值是查询结果集和错误信息，然后循环结果集取出其中的数据。代码如下：

增删改操作
执行增删改操作语句的是Exec()，参数是SQL语句，返回值是结果集和错误信息，通过对结果集的判断，得到执行结果的信息。以插入数据
为例代码如下：

创建表
创建表的方法也是Exec(),参数是SQL语句，返回值是结果集和错误信息：









ORM框架：ORM框架是Object-RelationShip-Mapping的缩写，中文叫关系对象映射。Beego中内嵌了ORM框架，用来操作数据库。
通过结构体对象生成，操作结构一样的数据库表
	首先要导包：import "github.com/astaxie/beego/orm"
	然后要定义一个结构体：
type User struct{   //这里字段名首字母都要大写，不然私有属性无法直接访问，映射数据库失败
	Id int
	Name string
	Passwd string
}
	然后像数据库中注册表，这一步又分为三步：
	连接数据库
	用RegisterDataBase()函数，第一个参数为数据库别名，也可以理解为数据库的key值，项目中必须有且只能有一个别名为 default
的连接，第二个参数是数据库驱动，这里我们用的是MySQL数据库，所以以MySQL驱动为例，第三个参数是连接字符串，和传统操作数据库
连接字符串一样，格式为：用户名:密码@tcp(ip:port)/数据库名称?编码方式，代码如下：
orm.RegisterDataBase("default","mysql","root:12345@tcp(127.0.0.1:3306)/class1?charset=utf8")
注意：ORM只能操作表，不能操作数据库，所以我们连接的数据库要提前在MySQL终端创建好。
	注册数据库表
用orm.RegisterModel()函数，参数是结构体对象，如果有多个表，可以用  , 隔开，多new几个对象：
orm.RegisterModel(new(User))

	生成表
用orm.RunSyncdb()函数，这个函数有三个参数，第一个参数是数据库的别名和连接数据库的第一个参数相对应。第二个参数是是否强制
更新，一般我们写的都是false，如果写true的话，每次项目编译一次数据库就会被清空一次，fasle的话会在数据库发生重大改变
（比如添加字段）的时候更新数据库。第三个参数是用来说，生成表过程是否可见，如果我们写成课件，那么生成表的时候执行的
SQL语句就会在终端看到。反之看不见。代码如下:
orm.RunSyncdb("default",fasle,true)




import "github.com/astaxie/beego/orm"

o ;= orm.NewOrm()
o := orm.NewOrm()

var user User
var user User

user.Name = "itcast"
user.Passwd = "heima"

id, err := o.Insert(&user)
if err == nil {
	fmt.Println(id)
}

id, err := o.Insert(&user)
if err == nil {
	fmt.Println(id)
}


o := orm.NewOrm()

var user User

user.Name = "itcast"

err, err := o.Read(&user, "Name")
if err != nil {
	beego.Info("查询数据错误"， err)
	return
}




加密三要素

- 明文/密文
- 秘钥
	- 定长的字符串
	- 需要根据加密算法确定其长度
- 算法
	- 加密算法
	- 解密算法
	- 加密算法和解密算法有可能是互逆的, 也有可能相同

常用的两种加密方式

- 对称加密
	- 秘钥: 加密解密使用的是同一个秘钥, 秘钥有一个
	- 特点
		- 双方向保证机密性
		- 加密效率高, 适合加密大数据, 大文件
		- 加密强度不高, 相对于非对称加密
- 非对称加密
	- 秘钥: 加密解密使用的不同的秘钥, 秘钥有两个, 需要使用秘钥生成算法, 得到密钥对
		- 公钥 - 可以公开的秘钥:
		- 公钥加密数据, 解密需要使用私钥
		- 私钥 - 需要妥善保管的秘钥, 知道的人越少越好:
		- 私钥加密, 公钥解密
	- 特点:
		- 数据的机密性只能单方向保证
		- 加密效率低, 适合加密少量数据
		- 加密强度高, 相对于对称加密

加密的步骤：
// 1. 建一个底层使用des的密码接口
// 2. 明文填充
// 3. 创建一个使用cbc分组接口
// 4. 加密
1. des

2. 3des

3. aes

//# 加密流程:
//1. 创建一个底层使用des/3des/aes的密码接口
   //	"crypto/des"
   //	func NewCipher(key []byte) (cipher.Block, error) # -- des
   //	func NewTripleDESCipher(key []byte) (cipher.Block, error)	# -- 3des
   //	"crypto/aes"
   //	func NewCipher(key []byte) (cipher.Block, error) # == aes
//2. 如果使用的是cbc/ecb分组模式需要对明文分组进行填充
//3. 创建一个密码分组模式的接口对象
   //	- cbc
   //	func NewCBCEncrypter(b Block, iv []byte) BlockMode # 加密
   //	- cfb
   //	func NewCFBEncrypter(block Block, iv []byte) Stream # 加密
   //	- ofb
   //	- ctr
//4. 加密, 得到密文




RSA加密：
生成私钥操作流程
	1. 使用rsa中的GenerateKey方法生成私钥
	2. 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	3. 将私钥字符串设置到pem格式块中
	4. 通过pem将设置好的数据进行编码, 并写入磁盘文件中


生成公钥操作流程
	1. 从得到的私钥对象中将公钥信息取出
	2. 通过x509标准将得到 的rsa公钥序列化为字符串
	3. 将公钥字符串设置到pem格式块中
	4. 通过pem将设置好的数据进行编码, 并写入磁盘文件


公钥加密
	1. 将公钥文件中的公钥读出, 得到使用pem编码的字符串
	2. 将得到的字符串解码
	3. 使用x509将编码之后的公钥解析出来
	4. 使用得到的公钥通过rsa进行数据加密


私钥解密
	1. 将私钥文件中的私钥读出, 得到使用pem编码的字符串
	2. 将得到的字符串解码
	3. 使用x509将编码之后的私钥解析出来
	4. 使用得到的私钥通过rsa进行数据解密




单向散列函数（one-waynction）：
有一个输人和一个输出，其中输人称为消息（message），输出称为散列值（hashvalue）。
单向散列函数可以根据消息的内容计算出散列值，而散列值就可以被用来检查消息的完整性。
这里的消息不一定是人类能够读懂的文字，也可以是图像文件或者声音文件。单向散列函数不需要知道消息实际代表的含义。
无论任何消息，单向散列函数都会将它作为单纯的比特序列来处理，即根据比特序列计算出散列值。散列值的长度和消息的长度无关。
无论消息是1比特，还是100MB，甚至是IOOGB，单向散列函数都会计算出固定长度的散列值。以SHA-I单向散列函数为例，
它所计算出的散列值的长度永远是160比特（20字节）。

两个不同的输有不同的散列值。这个也可以理解为哈希。






消息认证码：
消息认证码（message authentication code）是一种确认完整性并进行认证的技术，取三个单词的首字母，简称为MAC。
将“发送者和接收者之间的共享密钥”和“消息，进行混合后计算出的散列值。使用消息认证码可以检测并防止通信过程中的错误、
篡改以及伪装。消息认证码的输入包括任意长度的消息和一个发送者与接收者之间共享的密钥，它可以输出固定长度的数
据，这个数据称为MAC值。根据任意长度的消息输出固定长度的数据，这一点和单向散列函数很类似。但是单向散列函数中计算散列
值时不需要密钥，而消息认证码中则需要使用发送者与接收者之间共享的密钥。消息认证码有很多种实现方法，大家可以暂且这样理解：
消息认证码是一种与密钥相关联的单向散列函数。
消息验证码的使用过程：这个只适用两个人之间验证完整性，通过第三方验证就不一定正确
	1. 发送者Alice与接收者Bob事先共享密钥。
	2. 发送者Alice根据汇款请求消息计算MAC值（使用共享密钥）。
	3. 发送者Alice将汇款请求消息和MAC值两者发送给接收者Bob。
	4. 接收者Bob根据接收到的汇款请求消息计算MAC值（使用共享密钥）。
	5. 接收者Bob将自己计算的MAC值与从Alice处收到的MAC值进行对比。
	6. 如果两个MAC值一致，则接收者Bob就可以断定汇款请求的确来自Alice（认证成功）；如果不一致，则可以断定消息不是来自Alice（认证失败）。








数字签名：
在进行数字签名时也会使用单向散列函数。数字签名是现实社会中的签名（sign）和盖章这样的行为在数字世界中的实现。
使用数字签名可以识别篡改和伪装，还可以防止否认。数字签名的处理过程非常耗时，因此一般不会对整个消息内容直接施加数字签名，\
而是先通过单向散列函数计算出消息的散列值，然后再对这个散列值施加数字签名。
（1）Alice用单向散列函数计算消息的散列值。
（2）Alice用自己的私钥对散列值进行加密。（用私钥加密散列值所得到的密文就是Alice对这条散列值的签名，
由于只有Alice才持有自己的私钥因此, 除了Alice以外，其他人是无法生成相同的签名（密文）的）
（3）Alice将消息和签名发送给Bob。
（4）Bob用Alice的公钥对收到的签名进行解密。（如果收到的签名确实是用Alice的私钥进行加密而得到的密文（签名），那么用
Alice的公钥应该能够正确解密，解密的结果应该等于消息的散列值。如果收到的签名不是用Alice的私钥进行加密而得到的密文，
那么就无法用Alice的公钥正确解密（解密后得到的数据看起来是随机的））
（5）Bob将签名解密后得到的散列值与Alice直接发送的消息的散列值进行对比。（如果两者一致，则签名验证成功；
如果两者不一致，则签名验证失败。）





单向验证中，如果你是客户端，你需要拿到服务器的证书，并放到你的信任库中；如果是服务端，你要生成私钥和证书，
并将这两个放到你的密钥库中，并且将证书发给所有客户端。

1. 服务器要准备的
	生成密钥对(非对称加密)
	将公钥发送给ca, 由ca自己的私钥(非对称加密)对服务端的公钥添加数字签名并且签发证书
	将ca签发的证书和自己的 非对称加密的私钥部署到当前的web服务器
2. 通信流程
	客户端连接服务器, 通过一个域名。域名和IP地址的关系：域名要绑定IP地址，一个域名只能绑定一个IP地址，IP地址需要被域名绑定
	一个IP地址可以被多个域名绑定，客户端访问的域名会别解析成IP地址, 通过IP地址访问web服务器。
	1、客户端向服务端发送SSL协议版本号、加密算法种类、随机数等信息。
	2、服务端给客户端返回SSL协议版本号、加密算法种类、随机数等信息，同时也返回服务器端的证书，即公钥证书
	3、客户端使用服务端返回的信息验证服务器的合法性，包括：
		证书是否过期
		发型服务器证书的CA是否可靠
		返回的公钥是否能正确解开返回证书中的数字签名
		服务器证书上的域名是否和服务器的实际域名相匹配
		验证通过后，将继续进行通信，否则，终止通信
	4、客户端向服务端发送自己所能支持的对称加密方案，供服务器端进行选择
	5、服务器端在客户端提供的加密方案中选择加密程度最高的加密方式。
	6、服务器将选择好的加密方案通过明文方式返回给客户端
	7、客户端接收到服务端返回的加密方式后，使用该加密方式生成产生随机码，用作通信过程中对称加密的密钥，
		使用服务端返回的公钥进行加密，将加密后的随机码发送至服务器
	8、服务器收到客户端返回的加密信息后，使用自己的私钥进行解密，获取对称加密密钥。 在接下来的会话中，
		服务器和客户端将会使用该密码进行对称加密，保证通信过程中信息的安全。

双向认证和单向认证原理基本差不多，只是除了客户端需要认证服务端以外，增加了服务端对客户端的认证，
具体过程如下：
1、客户端向服务端发送SSL协议版本号、加密算法种类、随机数等信息。
2、服务端给客户端返回SSL协议版本号、加密算法种类、随机数等信息，同时也返回服务器端的证书，即公钥证书
3、客户端使用服务端返回的信息验证服务器的合法性，包括：
	证书是否过期
	发型服务器证书的CA是否可靠
	返回的公钥是否能正确解开返回证书中的数字签名
	服务器证书上的域名是否和服务器的实际域名相匹配
	验证通过后，将继续进行通信，否则，终止通信
4、服务端要求客户端发送客户端的证书，客户端会将自己的证书发送至服务端
5、验证客户端的证书，通过验证后，会获得客户端的公钥
6、客户端向服务端发送自己所能支持的对称加密方案，供服务器端进行选择
7、服务器端在客户端提供的加密方案中选择加密程度最高的加密方式
8、将加密方案通过使用之前获取到的公钥进行加密，返回给客户端
9、客户端收到服务端返回的加密方案密文后，使用自己的私钥进行解密，获取具体加密方式，而后，产生该加密方式的随机码，
	用作加密过程中的密钥，使用之前从服务端证书中获取到的公钥进行加密后，发送给服务端
10、服务端收到客户端发送的消息后，使用自己的私钥进行解密，获取对称加密的密钥，在接下来的会话中，
	服务器和客户端将会使用该密码进行对称加密，保证通信过程中信息的安全。

SSL/TLS是世界上应用最广泛的密码通信方法。使用SSL/TLS可以对通信对象进行认证，还可以确保通信内容的机密性。
SSL/TLS中综合运用了之前所学习的对称密码、消息认证码、公钥密码、数字签名、伪随机数生成器等密码技术。严格来说，
SSL（Secure Socket Layer)与TLS（Transport Layer Security）是不同的，TLS相当于是SSL的后续版本。
SSL是安全套接层(secure sockets layer)，TLS是SSL的继任者，叫传输层安全(transport layer security)。说白点，
就是在明文的上层和TCP层之间加上一层加密，这样就保证上层信息传输的安全。如HTTP协议是明文传输，加上SSL层之后，
就有了雅称HTTPS。它存在的唯一目的就是保证上层通讯安全的一套机制。
在大致了解了SSL/TLS之后，我们来整理一下SSL/TLS到底负责哪些工作。我们想要实现的是，通过本地的浏览器访问网络上的web服务器，
并进行安全的通信。用上边的例子来说就是，Alice希望通过web浏览器向Bob书店发送信用卡号。在这里，我们有几个必须要解决的问题。
1. Alice的信用卡号和地址在发送到Bob书店的过程中不能被窃听。
2. Alice的信用卡号和地址在发送到Bob书店的过程中不能被篡改。
3. 确认通信对方的Web服务器是真正的Bob书店。
在这里，（1）是机密性的问题；（2）是完整性的问题；而（3）则是认证的问题。
要确保机密性，可以使用对称加密。由于对称加密算法的密钥不能被攻击者预测，因此我们使用伪随机数生成器来生成密钥。
若要将对称加密的密钥发送给通信对象，可以使用非对称加密算法完成密钥交换。要识别篡改，对数据进行认证，可以使用消息认证码。
消息认证码是使用单向散列函数来实现的。要对通信对象进行认证，可以使用对公钥加上数字签名所生成的证书。
只要用一个“框架”（framework）将这些工具组合起来就可以了。SSL/TIS协议其实就扮演了这样一种框架的角色。






//区块的结构体：区块链理解为一个个区块通过前置哈希联接，并且存储全部区块的bolt数据库
type Block struct {
	//1.版本号
	Version uint64  //无符号，整形，长度为64位 % 8 = 8 字节
	//2. 前区块哈希,实际是区块头哈希，不过交易通过改变梅克尔根数间接改变哈希值
	PrevHash []byte
	//3. Merkel根（梅克尔根，这就是一个哈希值，我们先不管，我们后面v4再介绍）
	MerkelRoot []byte
	//4. 时间戳
	TimeStamp uint64
	//5. 难度值
	Difficulty uint64
	//6. 随机数，也就是挖矿要找的数据
	Nonce uint64

	//a. 当前区块哈希,正常比特币区块中没有当前区块的哈希，我们为了是方便做了简化！
	Hash []byte //这个实际是存在不了
	//b. 数据
	Data []byte //包括coinbase交易（挖矿），普通转账交易
}

GO语言序列有多种方式，请参考书签, 这里只介绍了两个
序列化概念理解：就是将变量从内存中变成可存储或传输的过程称之为序列化，也被称之为serialization，marshalling，flattening等等
，都是一个意思。序列化之后，就可以把序列化后的内容写入磁盘，或者通过网络传输到别的机器上。
反过来，把变量内容从序列化的对象重新读到内存里称之为反序列化。 网络传输类似数据时，记得要考虑粘包问题。

binary序列化
err := binary.Write(&buffer, binary.BigEndian, num), 这个函数的目的是将任意的数据转换成byte字节流，这个过程叫做序列化
write byte form of num into &buffer

err := binary.Read(buf, binary.LittleEndian, &num),  可以通过binary.Read方式进行反序列化，从字节流转回原始结构
read from buf into &num

gob序列化:encoder + decoder,gob是谷歌自己的序列化工具，只能在go语言中使用。常用的场景就是RPC传输数据。
1，结构体中的属性必须大写开头。不然无法序列化
2，序列化的struct与反序列化的struct结构可以不一样。只会匹配属性相同的数据。
import (
	"encoding/gob"
	"bytes"
	"log"
	"fmt"
)
type User struct{	// 结构体中的属性必须大写开头。不然无法序列化
	Id int32
	Name string
	Address string  // 结构体中的属性必须大写开头。不然无法序列化
}

type Student struct{	 // 结构体中的属性必须大写开头。不然无法序列化
	Id int32
	Name string
}
func main(){
	u:=User{1,"nxin","beijing"}  //实例化一个对象
	buf:=new(bytes.Buffer)			//分配内存的作用，一般也用var buf bytes.Buffer
	fmt.Println(buf)				//这时候打印buf为空
	enc:=gob.NewEncoder(buf)		//创建基于buf内存的编码器，实例化一个编码器对象
	err:=enc.Encode(u)				//编码器对结构体编码，序列化的数据写入buf，原来的对象u不改变
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(buf.Bytes())

	us:=User{}
	bu:=bytes.NewBuffer(buf.Bytes())//使用buf里面的数据创建初始化buffer,此时buf里面已经有了序列化的数据
	des:=gob.NewDecoder(bu)			//创建bu的解码器
	des.Decode(&us)					//对buf内容解码，并将解码后的数据写入us的内存中，不改变原来buf的数据
	fmt.Println(us)

	st:=Student{}
	bus:=bytes.NewBuffer(buf.Bytes())//使用buf里面的数据创建初始化buffer,此时buf里面已经有了序列化的数据
	dess:=gob.NewDecoder(bus)		//创建bu的解码器
	dess.Decode(&st)				//对buf内容解码，并将解码后的数据写入us的内存中
	fmt.Println(st)					//序列化的struct与反序列化的struct结构可以不一样。只会匹配属性相同的数据。
									//user, student结构不一致，不影响反序列化结果
}


//Gob序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	//- 使用gob进行序列化（编码）得到字节流
	//1. 定义一个编码器
	//2. 使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
	log.Panic("编码出错!")
	}

	//fmt.Printf("编码后的小明：%v\n", buffer.Bytes())
	return buffer.Bytes()
}

//Gob反序列化
func Deserialize(data []byte) Block {

	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	//2. 使用解码器进行解码
	err := decoder.Decode(&block)
	if err != nil {
	log.Panic("解码出错!", err)
	}
	return block
}



//pow
PoW，全称Proof of Work，即工作量证明，又称挖矿。
1、生成挖矿交易，并与其它所有准备打包进区块的交易组成交易列表，生成Merkle根哈希值。
2、将Merkle根哈希值，与区块头其它字段组成区块头，80字节长度的区块头作为Pow算法的输入字符串。
3、不断变更区块头中的随机数Nonce，对变更后的区块头做双重SHA256哈希运算，使得结果哈希值满足给定数量前多少位为0的过程，
4、其中前多少位为0的具体个数，取决于挖矿难度，前导0的个数越多，挖矿难度越大。
(这种说法也成立：与当前难度的目标值做比对，如果小于目标难度，即Pow完成。)
import (
	"crypto/sha256"
	"fmt"
)

data := "helloworld"

for i:= 0; i< 1000000; i++ {
	hash := sha256.Sum256([]byte(data + string(i)))  //  每次的hash值都不一样
	fmt.Printf("hash : %x\n", hash[:])
}




BOLT:
import (
	"fmt"
	"go一期/bolt"  //导入bolt所在的文件夹,里面有大量bolt结构??
	"log"
)

// Blot仅在Go中使用,所以使用前请先安装，bolt数据库类似于redis，是一个以键值（key-value）对形式存储数据的非关系型数据库。
func main() {
	//1. 打开数据库
	//Open的第一个参数为路径,如果数据库不存在则会创建名为my.db的数据库，第二个参数为文件操作权限，第三个参数是可选参数，内部可以配置只读和超时时间等。
	//特别需要注意的地方就是因为boltdb是文件操作类型的数据库，所以只能单点写入和读取，如果多个同时操作的话后者会被挂起直到前者关闭操作为止。
	//Bolt在打开数据库文件的时候会获取一个文件锁，所以多个进程不能同时打开一个数据库文件。打开一个已经Open的Bolt数据库会
	// 导致当前进程挂起直到其他进程关闭该Bolt数据库。
	db, err := bolt.Open("test.db", 0600, nil)

	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败！")
	}
	//Bolt数据库同时只支持一个read-write transaction或者多个read-only transactions。
	//每个独立的transaction以及在这个transaction中创建的所有对象（buckerts，keys等）都不是thread safe的。
	// 如果要在多个routine中处理数据，那么必须在每个routine中单独使用一个transaction或者显式的使用lock以
	// 确保在每个时刻只有一个routine访问这个transaction.
	//read-only的transaction和read-write的transaction不应该相互之间有依赖，
	// 一般来说在同一个goroutine中不要同时打开这两种transaction，
	// 因为read-write transaction需要周期性的re-map数据文件，但是由于read-only transaction打开导致
	// read-write transaction的re-map操作无法进行造成死锁。


	//boltdb的读写事务操作我们可以使用DB.Update()来完成。在Update理用各种方法增删改查
	db.Update(func(tx *bolt.Tx) error) {
		//2. 找到抽屉bucket(如果没有，就创建）
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
		//没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}
			//这里用 bucket, err := tx.CreateBucketIfNotExists([]byte("bucket"))应该效果更good
			//创建了一个bucket，该函数会先判断bucket是否存在，如果不存在会自动创建比上面方法灵活
		}
		//3. 写数据,前面为key，后面为value，在区块链项目中，Key为当前区块的hash值，value为通过gob序列化得到的具体结果值。
		bucket.Put([]byte("11111"), []byte("hello"))

		bucket.Put([]byte("22222"), []byte("world"))

		err = bucket.Put([]byte("user"),[]byte("useriou")) //这个方法只有返回值err?

		//删除bucket操作,argument为bucket的名称，不为对象名称
		err := tx.DeleteBucket([]byte("bucket"))
		//bucket的嵌套
		//func (*Bucket) CreateBucket(key []byte) (*Bucket, error)
		return nil
	})

	//4. 读数据
	db.View(func(tx *bolt.Tx) error) {

		//1. 找到抽屉，没有的直接报错退出
		bucket := tx.Bucket([]byte("b1"))

		if bucket == nil {
			log.Panic("bucket b1 不应该为空，请检查!!!!")
		}

		//2. 直接读取数据
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))
		fmt.Printf("v1 : %s\n", v1)
		fmt.Printf("v2 : %s\n", v2)

		return nil
	})

	//遍历bucket，boltdb以桶中的字节排序顺序存储键。这使得在这些键上的顺序迭代非常快。
	db.View(func(tx *bolt.Tx) error) {
		b := tx.Bucket([]byte("MyBucket"))  //  假设该bucket存在并且有键值存在
		//要遍历键，我们将使用游标Cursor()：
		c := b.Cursor()

		for i, v := c.First(); i != nil; i, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", i, v)
		}

		//  游标Cursor()允许您移动到键列表中的特定点，并一次一个地通过操作键前进或后退。
		//  光标上有以下函数：
		//  First() 移动到第一个健.
		//  Last() 移动到最后一个健.
		//  Seek() 移动到特定的一个健.
		//  Next() 移动到下一个健.
		//  Prev() 移动到上一个健.

		//如果你知道所在桶中一定拥有键，你也可以使用ForEach()来迭代
		b.ForEach(func(i, v []byte) error {
			fmt.Printf("key=%s, value=%s", i, v)

			return nil
		})
	}
}

范围查询
Bolt也可以在两个值之间遍历查询，如两个时间直接、数字之间等。

db.View(func(tx *bolt.Tx) error {
	// Assume our events bucket exists and has RFC3339 encoded time keys.
	c := tx.Bucket([]byte("Events")).Cursor()
	// Our time range spans the 90's decade.
	min := []byte("1990-01-01T00:00:00Z")
	max := []byte("2018-01-01T00:00:00Z")

	// Iterate over the 90's.
	for i, v := c.Seek(min); i != nil && bytes.Compare(i, max) <= 0; i, v = c.Next() {
		fmt.Printf("%s: %s", i, v)
	}
	return nil
})

利用nextsequence()功能，你可以让boltdb生成序列作为你键值对的唯一标识
func (s *Store) CreateUser(u *User) error {
		return s.db.Update(func(tx *bolt.Tx) error {
									// 创建users桶
									b := tx.Bucket([]byte("users"))
									// 生成自增序列
									id, _ = b.NextSequence()
									u.ID = int(id)
									// Marshal user data into bytes.
									buf, err := json.Marshal(u)

									if err != nil {
										return err
									}
									// Persist bytes to users bucket.
									return b.Put(itob(u.ID), buf)
				})
}


// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
		return b
}

type User struct {
	ID int
	...
}

批量读写事务
每一次新的事物都需要等待上一次事物的结束，这种开销我们可以通过DB.Batch()批处理来完成。在批处理过程中如果某个事务失败了,
批处理会多次调用这个函数函数返回成功则成功。如果中途失败了，则整个事务会回滚。
err := db.Batch(func(tx *bolt.Tx) error {
	...
	return nil
	})

Bolt，是一个以键值（key-value）对形式存储数据的非关系型数据库，实现数据持久化。区块链是由创世区块引导的，在区块链项目中
，Key为当前区块的hash值，value为该区块通过gob序列化得到的具体结果值。

过程大概如下：
1.创建DB文件并打开

2.创建存储区块信息的桶

3.把创世区块的信息进行序列化，并存储DB中

4.把创世区块添加到链中，普通区块直接添加在创世区块后面


比特币的地址是标志符,用于发送比特币给其他人进行比特币转账。
比特币地址是一串由字母或数字组成的26位到34位字符串，看起来有些像乱码，例如1N1jtTKD4fNvNcLtGP9RVH3vFdC2cpVoyg，
通过区块链可以查到每个比特币地址的所有转账记录。比特币地址就理解为个人的比特币账户，相当于你的银行卡卡号，
私钥理解为银行卡密码。任何人都可以通过你的比特币地址给你转账比特币。每个人的比特币地址都是独一无二的。下载比特币客户端或
者比特币钱包，也能注册自己的比特币地址。

B转账给A
A有解锁脚本
解锁脚本：是由A签名（私钥的签名）与A的公钥组成，是检验input是否可以使用由某个地址锁定的utxo。外部提供锁定信息，我去检查一下能否解开锁定脚本；
		  保证必须拥有私钥的用户才能对某一笔交易进行解锁
锁定脚本：是由A的公钥与转账金额组成,是用于指定比特币新的主人，在创建output的时候创建，对于这个output来说，他应该是一直在等待一个签名的到来，来检
		  查这个签名能否解开自己锁定的比特

解锁脚本		 	 锁定脚本
<sig>  <PubKey>  +  DUP HASH160 <PubKeyhash>  EQUALVAERIFY  CHECKSIG
比特币将解锁脚本与锁定脚本拼接后这笔数据存入内存栈里面，从左边到右边依次存入。
eg:有一个EXECUTION POINTER依次从这笔数据左边至右边执行，过程按照箭头
<A签名> <A公钥>  +  DUP HASH160 <A公钥哈希结果>  EQUALVAERIFY  CHECKSIG
A签名压入栈底部  -->  A公钥压入栈中  -->  DUP操作赋值一份<A公钥>至栈顶  -->  对栈顶公钥进行hash160操作并且得到哈希结果
-->  将锁定脚本的<A公钥哈希结果>压入栈中  -->  对比栈顶hash160操作并且得到哈希结果与锁定脚本的<A公钥哈希结果>，如果一致,
两个公钥哈希结果都出栈，继续触发CHECKSIG  -->  检测<A签名> <A公钥>是否匹配了
如果返回结果为TRUE，代表解锁成功，用户可以使用该笔输出作为新交易的输入进行交易。如果返回为FALSE或者执行过程中出现问题，则花费这笔输出的请求被拒绝。





var f = func(int) {}。
因此f可以看成是一个函数类型的变量。这样，可以动态的改变f的功能。匿名函数可以动态的创建，
与之成对比的常规函数必须在包中编译前就定义完毕。匿名函数可以随时改变功能，定义。


闭包是匿名函数与匿名函数所引用环境的组合。匿名函数有动态创建的特性，该特性使得匿名函数不用通过参数传递的方式，
就可以直接引用外部的变量。这就类似于常规函数直接使用全局变量一样，
个人理解为：匿名函数和它引用的变量以及环境，类似常规函数引用全局变量处于一个包的环境。
在使用闭包访问外部变量的时候，考虑这个变量如果发生改变，这样的改变对闭包的运行有什么影响。
i :=0 f := func() int { i+=1  return i},就是一个闭包，类比于常规函数+全局变量+包。f不仅仅是存储了一个函数的返回值，它同时存储了一个闭包的状态。

闭包被返回赋予一个同类型的变量时，同时赋值的是整个闭包的状态，该状态会一直存在外部被赋值的变量f中，直到f被销毁，整个闭包也被销毁。



以太坊：
在以太坊中运用到的有协议或算法：
	ghost协议：一种快速达成共识的协议。以太坊中快速合并分支链策略，用招安的方式（给分支链奖励）的方式，
			   让分支链快速合并到主链中来。Ghost协议中分支链上的区块，一旦合并到主链中来，挖出此区块的矿工，可以获得0到7/8的出块奖励。
	Casper协议：pos算法和pos算法的混合体。加入校验者的角色，交纳保证金后才能成为校验者。
	pos算法
	pow算法
	Hash算法

智能合约：适用于拟定好规则并且不经常变化的业务。公开透明，供所有人审查，写入数据需要支付手续费，查看数据不需要支付手续费。
		  并且一旦运行，无法停止。

以太坊平台一个计算机网络中的软件，能够部署、运行智能合约。以太坊平台用户可以参与，并且发布智能合约。转账或收款
		  并且查看账户以太坊余额。
		  metamask用户的每一个账户，都有唯一性的地址。每个用户可以有多个账户。可以调整转账手续费（油费）。参与智能合约。


比特币以太坊区别：
比特币												 以太坊
选用pow（工作量证明）作为挖矿算法                    以太坊选用pow（工作量证明）和pos（权益证明）混合算法，作为挖矿算法。
比特币做到了货币的去中心化							 以太坊加入了智能合约做到了合同的去中心化，实现了现实业务的去中心化，
													 并且做到了货币的去中心化。
10分钟出一个块										 以太坊相对比特币挖出区块时间更短，所以出现短暂分叉的可能性相对较高。
比特币因为系统扩容导致分叉  						 以太坊因为黑客攻击导致分叉。实现了现实业务的去中心化







web3







node常用语法：
（1）定义变量和常量

方式一：var 变量名 = 值

方式二：let 变量名 = 值

第一种方式定义的变量可以被重新定义，第二种方式定义的变量不能够被重新定义。所以基于安全性考虑，推荐使用let关键字定义变量。

定义常量：const 常量名 = 常量值
注意：1）定义常量时候需要指定具体的值；2）常量一旦定义之后，它的值不能够被改变；

（2）解构
对数组解构：
let arr = [1,2,3,4,5,99] //  定义一个数组

let [a,b,c,d] = arr  //  把数组前面四个元素赋值给a,b,c,d四个变量

console.log(a,b,c,d)  //  打印出四个元素的值 = print

解构对象:
const person = {
	name : 'lily',
	age : 18
}  //  这里将对象定义为常量？？
let {name, age} = person   //这里目标变量的顺序，变量名与结构体一样，所以直接赋值成功
console.log(name, age)

注意：如果解构对象时候目标变量的顺序或变量名与结构体不一样，那么就要明确指定对应的名字。
let {name:newName, age} = person  //这里目标变量newName与结构体不一样，要明确指定对应的名字
console.log(newName, age)  //打印目标变量

当对象作为函数参数的时候，也可以解构。
const person = {
	name : 'lily',
	age : 18
}

function print({name, age}) {
	console.log(`姓名：${name}, 年龄：${age}`) // 注意：这里不能使用单引号
}

print(person)


（3）函数定义
	方式一：传统方式创建函数。
function add(a, b) {
	return a + b
}
	方式二：使用箭头函数。
let add = (a,b) => {
	return a+b
}
	如果函数体只有返回值，那么也可以写成这样：
let add = (a,b) => a+b




（4）类
ECMAScript6支持class关键字定义类，其语法与java语言类似。
class Person {

	cosntructor(name, age) {
			this.name = name
			this.age = age
	}

	say() {
		console.log('may name is ${this.name}, i am ${this.age} years old.')
	}
}
细节问题：
1）在class中可以使用construct定义构造函数。当创建对象时候会自动调用构造函数，不需要自己手动调用。
2）上面构造函数中使用this对Person对象的name和age属性进行了初始化。this代表当前的Person对象。
3）在class中也可以定义方法，但是class中的方法不需要使用function关键字。
与Java语言类似，可以使用extends关键字继承类。构造函数主要用于初始化，析构函数主要用于销毁

class Student extends Person {

	constructor(name, age) {
		super(name, age)  //  super不仅仅是一个关键字，还可以作为函数和对象。这里的super相当于Person类里面的constructor构建函数
		//此时的super里面的this指向子类的实例对象Student，如果有print Student，该Student对象被打印出。
		//函数：在子类继承父类中，super作为函数调用，只能写在子类的构造函数（constructor）里面，super代表的是父类的构造函数，
		//对象： 这里重点理解下对象，概念相对抽象。super作为对象使用时，分为在普通方法中使用和在静态方法中使用，
		//在普通方法找中，super指向父类的原型，即Person.prototype，可以访问到原型上的方法和属性
		//在静态方法之中，这时super将指向父类，而不是父类的原型对象。
	}
}
//上面程序Student类继承了Person类，那么Student类就是Person类的子类。因此Student就拥有了Person类中的所有成员属性和成员方法。
//值得注意的是，在子类构造函数中需要显示地调用父类的构造函数。super(...)的作用就是调用父类的构造函数。
//子类可以定义父类同名的方法。如果通过子类对象调用同名方法，那么默认会调用子类的方法，而不是父类方法。
果要创建class对象，那么可以使用new关键字。
let stu = new Student("mi", 19);
stu.say()

（5）异步编程
node.js的异步编程是通过回调函数来实现的。下面以读取文件为例：
//导入fs模块，该模块是由node.js提供，系统本来就有
let fs = require('fs')  //  注意，由导入必定有导出，每一个require对应一个export。如果require自己定义的库，一定要有export

//异步读取文件，记住，readFile: 异步读取。readFileSync:同步读取
fs.readFile('1.txt', 'utf-8', function(err, data) {
	if (err){
		console.log("读取文件失败！")
		return
	}
	console.log('${data}')
})
//从面代码通过异步方式读取1.txt文件的内容。其中，在readFile函数的第三个参数就是一个回调函数。该回调函数是在文件读取完毕后
// ，系统自动调用。node.js中大量使用了回调函数来实现异步通信。回调函数一般作为函数的最后一个参数传递进来。


（6）模块导入导出
06_模块导出.js
function printA() {
	console.log("printA...")
}

function printB() {
	console.log("printB...")
}

//执行导出
module.exports = {
	printA,
	printB,
}

06_模块导入.js
let ex = require("./06_模块导出")
ex.printA()
ex.printB()
or
let {printA, printB} = require("./06_模块导出")
printA
printB
require可以对导入系统内置模块、也可以导入第三方或自己实现的模块。使用require导入模块，一定要有对应的导出。
如果是自己实现的模块，导入时候要明确指定路径。路径要以./开头，不需要后缀名。

（7）path模块
path模块提供了一些操作文件路径的方法，简化了文件路径的操作。
let path = require("path")
// 返回一个路径的最后一部分
console.log("basename : ", path.basename('D:\\code\\go_workspace\\src\\Project014\\1.txt'))
// 返回一个路径的目录部分
console.log("dirname : ", path.dirname('D:\\code\\go_workspace\\src\\Project014\\1.txt'))
// 返回一个文件的扩展名
console.log("extname : ", path.extname('D:\\code\\go_workspace\\src\\Project014\\1.txt'))
// 拼接给定的路径片段
console.log("join : ", path.join('D:\\code', '\\go_workspace', '\\src', '\\Project014', '\\1.txt'))
// 把一个不规范的路径修改为规范路径
console.log("normalize : ", path.normalize('D:\\code\\\\\\go_workspace\\\src\\\\Project014/1.txt'))
// 基于当前执行目录，返回一个绝对路径
console.log("resolve : ", path.resolve('1.txt'))




（8）fs模块
fs模块提供了一些操作文件的方法。常用方法如下：

fs.stat/fs.statSync：访问文件的元数据，比如文件大小，文件的修改时间。

fs.readFile/fs.readFileSync：异步/同步读取文件。

fs.writeFile/fs.writeFileSync：异步/同步写入文件。

fs.readdir/fs.readdirSync：读取文件夹内容。

fs.unlink/fs.unlinkSync：删除文件。

fs.rmdir/fs.rmdirSync：只能删除空文件夹。如果要删除非空文件夹，可以借助第三方模块fs-extra 来实现。

fs.watchFile：监视文件的变化。
let fs = require("fs")

// 读文件
fs.readFile("1.txt", "utf-8", (err, data) => {
	if (err) {
		console.log("读文件失败！")
		return
	}
	console.log("读到的内容：", data)
})

// 写文件
fs.writeFile("2.txt", "hello world", (err) => {
	if (err) {
		console.log("写文件失败！")
		return
	}
	console.log("写文件成功！")
})

// 读取文件状态
let stat = fs.statSync("1.txt")
console.log("文件大小：", stat.size)
console.log("是否是文件：", stat.isFile())
console.log("是否是目录：", stat.isDirectory())

// 删除文件
fs.unlinkSync("2.txt")

（9）定义Promise

Promise的作用是对异步回调代码包装一下，把原来的一个回调函数拆成2个回调函数，从而提高程序的可读性。
let fs = require("fs")

fs.readFile("1.txt", (err, data) => {
	if (err) {
		console.log("读文件失败！")
		return
	}
	console.log("读取到的内容：", res.toString())
})
如果上述代码使用Promise来实现，代码如下所示：
let fs = require('fs')

//使用promise对readFile函数进行包装
let readFilePromise = new Promise(fucntion (resolve, reject){
	fs.readFile("1.txt", (err, data) => {
		if (err) {
			reject(err)
		} else {
			resolve(data)
		}
	})
})
//调用Promise对象的then以及catch方法
readFilePromise.then(res => {
	console.log("读取的内容：", res.toString())  //  toString方法将读取的数据转成字符串
}).catch(err => {
	console.log(err)
})
调用then方法需要传入回调函数作为参数，回调函数的参数就是resolve(data)中的data参数。
	调用catch方法也要传入回调函数作为参数，回调函数的参数就是reject(err)中的err参数。

（10）async和await
如果需要对多个异步函数封装成Promise，那么需要使用async和await来实现异步调用。
let fs = require("fs")

fs.readFile("1.txt", (err, data) => {
	if (err) {
		console.log("读文件失败！")
		return
	}
	fs.writeFile("2.txt", data, (err) => {
		if (err) {
			console.log("写文件失败！")
			return
		}
		fs.stat("2.txt", (err, stat) => {
			if (err) {
				console.log("读取文件状态失败！")
				return
			}
			console.log(stat)
		})
	})
})
上面代码中有多个异步函数，如果使用Promise来实现，首先要为每一个异步函数都使用一个Promise进行包装。
let fs = require('fs')

// 1.创建一个Promise，对fs.readFile进行封装
let readFilePromise = () => {
	return new Promise(function (resolve, reject) {
		fs.readFile("1.txt", (err, data) => {
			if (err) {
				reject(err)
			} else {
				resolve(data)
			}
		})
	})
}

// 2.创建一个Promise，对fs.writeFile进行封装
let writeFilePromise = (data) => {
	return new Promise(function (resolve, reject) {
		fs.writeFile("2.txt", data, (err) => {
			if (err) {
				reject(err)
			}
		})
	})
}

// 3.创建一个Promise，对fs.stat进行封装
let readStatPromise = () => {
	return new Promise(function (resolve, reject) {
		fs.stat("2.txt", (err, stat) => {
			if (err) {
				reject(err)
			} else {
				resolve(stat)
			}
		})
	})
}
然后再定义一个函数，该函数负责调用上面的readFilePromise、writeFilePromise和readStatPromise函数。
let checkStat = async() => {
	try {
		let data = await readFilePromise()
		await readFilePromise()
		console.log("读写成功！")
		let stat = await readStatPromise()
		console.log(stat)
	} catch (e) {
		console.log(e)
	}
}

checkStat()
在async函数中的每一个函数调用，都需要使用await修饰。async函数必须要等到方法体中所有await修饰的Promise函数执行完后，
async函数才会得到一个resolve状态的Promise对象。如果执行async函数过程中一旦其中一个异步函数发生异常，那么
整个async函数终止执行，并抛出异常对象给catch处理。如果在catch中返回reject()，那么async函数就会得到一个reject状态的Promise对象。



编译合约导出interface以及bytecode
//导入solc编译器
let solc = require('solc')

let fs = require('fs')
//读取合约
let sourceCode = fs.readFileSync('./contracts/SimpleStorage.sol', 'utf-8')  //  读取本工程下面的某一个合约

let output = solc.compile(sourceCode, 1)  //  setting 1 as second parameter activate the optimizer


console.log("output: ", output) //打印出编译结果，里面有bytecode, interface(= ABI) and so on
console.log('abi :', output['contracts'][':SimpleStorage']['interface'])  //  ABI
module.exports = ouput['constracts'][':SimpleStorage']








Docker是采用了(c/s)架构模式的应用程序:
Client dockerCLI :客户端docker命令行
Docker客户端(Docker Client)是用户与Docker进行交互的最主要方式。当在终端输入docker命令时，对应的就会在服务端产生对应的作用，并把结果返回给客户端。Docker Client除了连接本地服务端，通过更改或指定DOCKER_HOST连接远程服务端。
Docker Daemon其实就是Docker 的服务端。它负责监听Docker API请求(如Docker Client)并管理Docker对象(Docker Objects)，如镜像、容器、网络、数据卷等。

REST API : 一套介于客户端与服务端的之间进行通信并指示其执行的接口
Server docker daemon:服务端dacker守护进程等待客户端发送命令来执行

IMAGE-镜像
一个Docker的可执行文件，其中包括运行应用程序所需的所有代码内容、依赖库、环境变量和配置文件等，理解为ISO镜像，通过镜像可以创建一个或多个容器。


CONTAINER-容器
容器是一种轻量级、可移植、并将应用程序进行的打包的技术，使应用程序可以在几乎任何地方以相同的方式运行
•Docker将镜像文件运行起来后，产生的对象就是容器。容器相当于是镜像运行起来的一个实例。
•容器具备一定的生命周期。
•另外，可以借助docker ps命令查看运行的容器，如同在linux上利用ps命令查看运行着的进程那样。
我们就可以理解容器就是被封装起来的进程操作,只不过现在的进程可以简单也可以复杂,复杂的话可以运行1个操作系统.简单的话可以运行1个回显字符串


Repository or Registry:
俗称Docker仓库，专门用于存储镜像的云服务环境.Docker Hub就是一个公有的存放镜像的地方，类似Github存储代码文件。同样的也可以类似Github那样搭建私有的仓库


DATA VOLUMES-数据卷
容器与宿主机之间、容器与容器之间共享存储方式，类似虚拟机与主机之间的共享文件目录。就是将宿主机的某个目录，映射到容器中，作为数据存储的目录，
生产环境使用Docker的过程中，往往需要对数据进行持久化保存，或者需要更多容器之间进行数据共享，
通过操作数据卷（Data Volumes）和数据卷容器（Data Volume Containers）。

数据卷（Data Volumes）：容器内数据直接映射到本地主机环境
数据卷特性
1、数据卷可以在容器之间共享和重用，本地与容器间传递数据更高效；
2、对数据卷的修改会立马有效，容器内部与本地目录均可；
3、对数据卷的更新，不会影响镜像，对数据与应用进行了解耦操作；
4、卷会一直存在，直到没有容器使用。


数据卷容器：
 在多个容器之间共享一些持续更新的数据，最简单的方式是使用数据卷容器。数据卷容器也是一个容器，但是它的目的是专门用来提供数据卷供其他容器挂载。数据卷容器（Data Volume Containers）：使用特定容器维护数据卷。简单点：数据卷容器就是为其他容器提供数据交互存储的容器
如果使用数据卷容器，在多个容器间共享数据，并永久保存这些数据，需要有一个规范的流程才能做得到：
1、创建数据卷容器 2、其他容器挂载数据卷容器
注意： 数据卷容器自身并不需要启动，但是启动的时候依然可以进行数据卷容器的工作。数据卷容器实践包括两部分：创建数据卷容器和使用数据卷容器

数据备份方案： 
1 创建一个挂载数据卷容器的容器
2 挂载宿主机本地目录作为备份数据卷
3 将数据卷容器的内容备份到宿主机本地目录挂载的数据卷中
4 完成备份操作后销毁刚创建的容器


数据恢复方案：
1、创建一个新的数据卷容器（或删除原数据卷容器的内容）
2、创建一个新容器，挂载数据卷容器，同时挂载本地的备份目录作为数据卷
3、将要恢复的数据解压到容器中
4、完成还原操作后销毁刚创建的容器




NETWORK-网络
外部或者容器间如何互相访问的网络方式，如host模式、bridge模式。


默认情况下，容器和宿主机之间网络是隔离的，我们可以通过端口映射的方式，将容器中的端口，映射到宿主机的某个端口上。这样我们就可以通过宿主机的ip+port的方式来访问容器里的内容



通过拍摄快照可以保留虚拟机的状态，以便返回相同的状态。




Dockerfile类似于我们学习过的脚本，将我们在上面学到的docker镜像，使用自动化的方式实现出来。
Dockerfile的作用 
1、找一个镜像： ubuntu
2、创建一个容器： docker run ubuntu
3、进入容器： docker exec -it 容器 命令
4、操作： 各种应用配置....
5、构造新镜像： docker commit

Dockerfile 使用准则 
1、大： 首字母必须大写D
2、空： 尽量将Dockerfile放在空目录中。
3、单： 每个容器尽量只有一个功能。
4、少： 执行的命令越少越好。

Dockerfile 分为四部分:
基础镜像信息 从哪来？
维护者信息 我是谁？
镜像操作指令 怎么干？
容器启动时执行指令 嗨！！！
Dockerfile文件内容：
首行注释信息
指令(大写) 参数


这个理解为创建镜像的一个简单脚本：

# 构建一个基于ubuntu的docker定制镜像
# 基础镜像
FROM ubuntu

# 镜像作者
MAINTAINER panda 1195188926@qq.com

# 执行命令
RUN mkdir hello
RUN mkdir world
RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update
RUN apt-get install nginx -y

# 对外端口
EXPOSE 80


Dockerfile构建过程： 从基础镜像1运行一个容器A
遇到一条Dockerfile指令，都对容器A做一次修改操作
执行完毕一条命令，提交生成一个新镜像2
再基于新的镜像2运行一个容器B
遇到一条Dockerfile指令，都对容器B做一次修改操作
执行完毕一条命令，提交生成一个新镜像3
…
**最终产生我们想要的镜像，这个过程通过docker history 镜像ID 命令可以查看到整个过程。
构建过程中，创建了很多镜像，这些中间镜像，我们可以直接使用来启动容器，通过查看容器效果，从侧面能看到
我们每次构建的效果。提供了镜像调试的能力









任务编排介绍:
我们在工作中为了完成业务目标，首先把业务拆分成多个子任务，然后对这些子任务进行顺序组合，当子任务按照方案执行完毕后，就完成了业务目标。任务编排，就是对多个子任务执行顺序进行确定的过程。
常见的任务编排工具：
单机版： docker compose
集群版：
Docker swarm Docker
Mesos Apache
Kubernetes（k8s） Google


compose是定义和运行多容器Docker应用程序的工具。通过编写，您可以使用YAML文件来配置应用程序的服务。然后，使用单个命令创建并启动配置中的所有服务。要了解更多有关组合的所有特性，请参见特性列表。
docker compose的特点； 
本质：docker 工具
对象：应用服务
配置：YAML 格式配置文件
命令：简单
执行：定义和运行容器


version: '2' 				# compose 版本号
services: 				# 服务标识符
  web1: 				# 子服务命名
  image: nginx 				# 服务依赖镜像属性
  ports: 				# 服务端口属性
    - "9999:80" 			# 宿主机端口:容器端口
    container_name: nginx-web1 		# 容器命名









微服务：
服务（service） 一定要区别于系统，服务一个或者一组相对较小且独立的功能单元，是用户可以感知最小功能集。那么广义上来讲，微服务是一种分布式系统解决方案，推动细粒度服务的使用，这些服务协同工作。 有人把微服务理解为一种细粒度SOA（service-oriented Architecture，面向服务架构），一种轻量级的组件化的小型SOA。



微服务课程的几个重要的组件：
跨语言，跨平台 通讯格式 protobuf，例如各个语言，平台通过各种 协议格式.proto来进行通讯
通讯协议 gRPC
调度管理服务发现 consul
微服务的框架 micro
部署 docker



REST是啥？与RESTful一样么？
REST是一组架构约束条件及原则。满足这些约束条件及原则的应用程序胡总恶化设计是RRSTful




