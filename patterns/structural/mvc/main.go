// MVC

package main

import "fmt"

// 一个产品结构体
type Product struct {
	name     string
	price    float64
	quantity float64
}

func (product *Product) String() string {
	return fmt.Sprintf("%s<price=%.2f, quantity=%.2f>", product.name, product.price, product.quantity)
}

// Model 接口
type Model interface {
	get(string) interface{}
	showItems() []string
}

// 产品模型结构
type ProductModel struct {
	products map[string]*Product
}

// 创建一个示例 ProductModel 对象
func NewProductModel() *ProductModel {
	return &ProductModel{
		products: map[string]*Product{
			"apple":  &Product{"apple", 1.5, 10.0},
			"eggs":   &Product{"egg", 0.2, 100.0},
			"cheese": &Product{"cheese", 2.0, 10.0},
		},
	}
}

// 从 Model 中取出一个产品
func (model *ProductModel) get(target string) interface{} {
	if product, ok := model.products[target]; ok {
		return product
	} else {
		fmt.Printf("No such a product called %s\n", target)
		return nil
	}
}

// 获取所有产品名称
func (model *ProductModel) showItems() []string {
	result := []string{}

	for key, _ := range model.products {
		result = append(result, key)
	}

	return result
}

// View 接口
type View interface {
	showItemList([]string)
	showItemInformation(interface{})
}

// 一个 View 的具体实现(将数据直接输出到控制台)
type ConsoleView struct{}

func (view *ConsoleView) showItemList(names []string) {
	fmt.Println(names)
}

func (view *ConsoleView) showItemInformation(item interface{}) {
	product := item.(*Product)
	fmt.Printf("%s\n", product)
}

// Controller 结构
type Controller struct {
	model Model
	view  View
}

func NewController(model Model, view View) *Controller {
	return &Controller{
		model: model,
		view:  view,
	}
}

func (controller *Controller) showItemList() {
	controller.view.showItemList(controller.model.showItems())
}

func (controller *Controller) showItemInformation(name string) {
	controller.view.showItemInformation(controller.model.get(name))
}

func main() {
	model := NewProductModel()
	view := &ConsoleView{}
	controller := NewController(model, view)

	controller.showItemList()
	controller.showItemInformation("apple")
	controller.showItemInformation("eggs")
	controller.showItemInformation("cheese")

}
