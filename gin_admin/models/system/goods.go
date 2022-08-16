/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-13 15:46:15
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 16:40:25
 * @FilePath: /gin_admin/models/goods.GO
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name             string `gorm:"type:varchar(20);not null"`
	ParentCategoryID int32  `gorm:"type:int(11);not null;"`
	Level            int32  `gorm:"type:int;not null;default:1"`
	IsTab            bool   `gorm:"type:boolean;not null;default:false"`
}

type Brands struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	gorm.Model
	CategoryID int32 `gorm:"type:int;index:idx_category_brand, unique"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     Brands
}

//声明表名称
// func (GoodsCategoryBrand) TableName() string {
// 	return "goodscategorybrand"
// }

type Banner struct {
	gorm.Model
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `grom:"type:int;default:1;not null"`
}

type Goods struct {
	gorm.Model
	CategoryID int32 `gorm:"type:int;index:idx_category_brand, unique"`
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique"`
	OnSale     bool  `gorm:"default:false;not null"`
	ShipFree   bool  `gorm:"default:false;not null"`
	IsNew      bool  `gorm:"default:false;not null"`
	IsHot      bool  `gorm:"default:false;not null"`

	Name        string  `gorm:"type:varchar(50);not null"`
	GoodsSn     string  `gorm:"type:varchar(50);not null"` //商品编号
	ClickNum    int32   `gorm:"type:int;default:0;not null"`
	SoldNum     int32   `gorm:"type:int;default:0;not null"` //销售量
	FavNum      int32   `gorm:"type:int;default:0;not null"` //欢迎量
	MarketPrice float32 `gorm:"not null"`                    //市场价
	ShopPrice   float32 `gorm:"not null"`
}

type Orders struct {
	gorm.Model
	GoodID   int32   `gorm:"type:int;index:idx_good_id;"`
	Name     string  `gorm:"type:varchar(50);not null"`
	OrderNum string  `gorm:"type:varchar(50);not null"` //订单编号
	ClickNum int32   `gorm:"type:int;default:0;not null"`
	SoldNum  int32   `gorm:"type:int;default:0;not null"`    //销售量
	Price    float32 `gorm:"type:float;default:0;not null;"` //价格
	Status   int32   `gorm:type:tinyint(2);default:0;not null;comment:0.待支付,1.已支付,2退款,3.删除;"`
	Desc     string  `gorm:type:text;defafult:'';not null;comment:订单描述;"`
}
