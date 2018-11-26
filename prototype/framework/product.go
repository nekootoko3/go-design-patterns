package framework

type Producter interface {
	Use(str string)
	CreateClone() Producter
}
