package types

type OrderType int

const (
	Normal  OrderType = 0
	Special OrderType = 1
)

func (orderType OrderType) String() string {
	names := [...]string{"L", "S"}

	return names[orderType]
}

func GetOrderTypes() []OrderType {
	return []OrderType{Normal, Special}
}
