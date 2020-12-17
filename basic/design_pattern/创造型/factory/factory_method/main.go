package main

import "fmt"

// @Descripttion:
// @version:
// @Author: zsl
// @Date: 2020-12-15 19:27:41
// 工厂模式：工厂方法
// 创建来自同一父类或者接口的对象
// 优点：封装了对象的创建过程，相当于增加了一个代理，无需和具体的工厂类打交道
// 特点：区别于简单工厂，工厂方法对于每个目标类都有一个对应的工厂类，适用于复杂对象的构造，且比简单工厂更符合开闭原则
// 缺点：比简单工厂代码多
// 选择：当创建对象不是很复杂时，尽量选择简单工厂
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
type cat struct {
}

func (c *cat) name() string {
	return "cat"
}

// 动物工厂 创建接口
type iAnimalCreate interface {
	CreateAnimal() animal
}

// 猫工厂类
type catFactory struct {
}

// 猫工厂方法
func (c *catFactory) createAnimal() animal {
	return new(cat)
}

// 狗工厂类
type dogFactory struct {
}

// 狗工厂方法
func (d *dogFactory) createAnimal() animal {
	return new(dog)
}

func main() {
	dogFacry := &dogFactory{}
	dog := dogFacry.createAnimal()
	fmt.Println(dog.name())
}
