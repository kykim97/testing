package test

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"test/external"
)

type Order struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Testttttttttttttttttttttttt string `json:"testttttttttttttttttttttttt"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty string `json:"qty"`

}

func (self *Order) onPostPersist() (err error){
	orderPlaced := NewOrderPlaced()
	model.Copy(orderPlaced, self)

	Publish(orderPlaced)

	return nil
}
func (self *Order) onPrePersist() (err error){ return nil }
func (self *Order) onPreUpdate() (err error){ return nil }
func (self *Order) onPostUpdate() (err error){ return nil }
func (self *Order) onPreRemove() (err error){ return nil }
func (self *Order) onPostRemove() (err error){ return nil }


