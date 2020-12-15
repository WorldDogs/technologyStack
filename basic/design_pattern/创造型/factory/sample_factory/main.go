package main
import "fmt"
// @Descripttion:
// @version:
// @Author: zsl
// @Date: 2020-12-15 19:27:41
// 工厂模式：简单工厂
// 创建来自同一父类或者接口的对象
// 优点：封装了对象的创建过程，相当于增加了一个代理，无需和具体的工厂类打交道

// Animal 动物
type animal interface {
	name() string // 获取名字
}

// dog
type dog struct {
}

func (d *dog) name() string {
	return "dog"
}

// cat
type cat struct{

}

func (c *cat)name()string{
	return "cat"
}
// 简单工厂
func newAnimal(tp string)animal{
	switch tp {
	case "cat":
		return &cat{}
	case "dog":
		return &dog{}
	default:
		panic("no support type")
	}
}

func main() {
	ani:=newAnimal("cat")
	fmt.Println(ani.name())
}