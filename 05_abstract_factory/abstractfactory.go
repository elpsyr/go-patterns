package abstractfactory

import "fmt"

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct {
}

func (R *RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

type RDBDetailDAO struct {
}

func (R *RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

type RDBDAOFactory struct {
}

func (R *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (R *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

//-----------以上为RDB存储

type XMLMainDAO struct {
}

func (X *XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

type XMLDetailDAO struct {
}

func (X *XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save\n")
}

type XMLDAOFactory struct {
}

func (X *XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (X *XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}
