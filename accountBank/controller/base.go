package controller

import (
	"github.com/go-pg/pg"
	"GoLearn/accountBank/config"
)

type Controller struct {
	 // DB intance
	 DB *pg.DB

	 // configuration
	 Config config.Config
}

func NewController() *Controller  {
	var c Controller
	return &c
}