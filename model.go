package main

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Products []struct {
	Product struct {
		ID          int    `yaml:"id"`
		Description string `yaml:"description"`
		Color       string `yaml:"color"`
		Category    string `yaml:"category"`
	} `yaml:"product"`
}

type ProductContainer struct {
	mu                 sync.Mutex
	LangMapForProducts map[string]Products
}

func (pc *ProductContainer) FillProductDetailsMap() {
	go pc.FillMapTRProductsDetail()
	go pc.FillMapENProductsDetail()
}

func (pc *ProductContainer) FillMapENProductsDetail() {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	FileEN, err := os.ReadFile("products-en.yaml")
	if err != nil {
		log.Error(err)
		return
	}
	var ENProduct Products
	yaml.Unmarshal(FileEN, &ENProduct)

	pc.LangMapForProducts["en"] = ENProduct
}

func (pc *ProductContainer) FillMapTRProductsDetail() {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	FileTR, err := os.ReadFile("products-tr.yaml")
	if err != nil {
		log.Error(err)
		return
	}

	var TRProducts Products
	yaml.Unmarshal(FileTR, &TRProducts)

	pc.LangMapForProducts["tr"] = TRProducts
}
