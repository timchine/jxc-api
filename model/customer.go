package model

type Customer struct {
	Model
	CustomerID        int       `json:"customer_id" gorm:"primaryKey;autoIncrement;column:customer_id;type:int;comment:主键"`
	CustomerShortName string    `json:"customer_short_name" gorm:"column:customer_short_name;type:varchar(50);comment:简称"`
	Intro             string    `json:"intro" gorm:"column:intro;type:varchar(500);comment:简介"`                         //简介
	ShippingAddress   string    `json:"shipping_address" gorm:"column:shipping_address;type:varchar(200);comment:送货地址"` //送货地址
	CustomerType      int       `json:"customer_type" gorm:"column:customer_type;type:int;comment:客户类型"`                // todo 放设置
	Receiver          string    `json:"receiver" gorm:"column:receiver;type:varchar(50);comment:收货人"`
	ReceiverPhone     string    `json:"receiver_phone" gorm:"column:receiver_phone;type:varchar(20);comment:收货人电话"`
	Contact           string    `json:"contact" gorm:"column:contact;type:varchar(50);comment:联系人"`
	ContactPhone      string    `json:"contact_phone" gorm:"column:contact_phone;type:varchar(20);comment:联系人电话"`
	CustomerName      string    `json:"channel_name" gorm:"column:channel_name;type:varchar(100);comment:供应商名称"`
	TaxNumber         string    `json:"tax_number" gorm:"column:tax_number;type:varchar(50);comment:税号"`
	Address           string    `json:"address" gorm:"column:address;type:varchar(300);comment:地址"` //单位地址
	CustomerTel       string    `json:"customer_tel" gorm:"column:customer_tel;type:varchar(20);comment:单位电话"`
	CardNo            string    `json:"card_no" gorm:"column:card_no;type:varchar(50);comment:银行卡号"`
	OpeningBank       string    `json:"opening_bank" gorm:"column:opening_bank;type:varchar(30);comment:开户行"`
	PayMethod         PayMethod `json:"pay_method" gorm:"column:pay_method;type:int;comment:付款方式"`
	Status            int       `json:"status" gorm:"column:status;type:int;comment:状态 1 正常 8 删除；default:1"`
}

type PayMethod int

const (
	// todo 方式 问题
	_                      PayMethod = iota
	PAY_METHOD_FIRST_CARGO           //先货后付
	PAY_METHOD_FIRST_PAY             //先付后获
	PAY_METHOD_DELIVERY              //货到付款
	PAY_METHOD_MONTH                 // 月结
	PAY_METHOD_SEASON                //季结
	PAY_METHOD_YEAR                  // 年结
)
