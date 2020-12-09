package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)
type ZwdProduct struct {
	Name string `json:"name"`
	OriginalPrice string `json:"original_price"`
	WholesalePrice string `json:"wholesale_price"`
	GoodsNo string `json:"goods_no"`
	Colors []Color
	Sizes []Size
	Images []Image
	OnSaleTime string `json:"on_sale_time"`
	StoreName string `json:"store_name"`
	Phones []Phone
	Address string `json:"address"`
	Details []Detail
}
type Color struct {
	ColorName string `json:"color_name"`
}
type Size struct {
	SizeName string `json:"size_name"`
	SizePrice string `json:"size_price"`
}
type Phone struct {
	PhoneNumber string `json:"phone_number"`
}
type Image struct {
	ImageUrl string `json:"image_url"`
}
type Detail struct {
	Label string `json:"label"`
}

func (p ZwdProduct) Scrape(url string) ZwdProduct{

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			p.Name = r.HTMLDoc.Find("#J_goodsForm > div.goods-item-title > span").Text()
			p.OriginalPrice = r.HTMLDoc.Find("#J_goodsForm > div.goods-parameter-outside-container > div.goods-page-server-container > div.goods-page-top > div > span.goods-value").Text()
			p.WholesalePrice = r.HTMLDoc.Find("#goods-pifa-price").Text()
			p.GoodsNo = r.HTMLDoc.Find("#J_goodsForm > div.goods-parameter-outside-container > div.goods-all-parameter-container > div:nth-child(1) > div.parameter-right > a").Text()
			p.OnSaleTime = r.HTMLDoc.Find("#J_goodsForm > div.goods-parameter-outside-container > div.goods-all-parameter-container > div:nth-child(2) > div.parameter-right > a").Text()
			p.OnSaleTime = TrimHtml(p.OnSaleTime)
			r.HTMLDoc.Find("#sku-color-selector > span").Each(func(_ int, s *goquery.Selection) {
				color :=  new(Color)
				color.ColorName, _ = s.Attr("title")
				p.Colors = append(p.Colors, *color)
			})
			r.HTMLDoc.Find("#sku-size-selector > ul > li").Each(func(_ int, s *goquery.Selection) {
				size :=  new(Size)
				size.SizeName = TrimHtml(s.Find("div:nth-child(1)").Text())
				size.SizePrice = TrimHtml(s.Find("div:nth-child(2)").Text())
				p.Sizes = append(p.Sizes, *size)
			})
			r.HTMLDoc.Find("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.view-BigPicture > div.BigPicture-content > div.left-area > div.imgWrap").Each(func(_ int, s *goquery.Selection) {
				image := new(Image)
				image.ImageUrl, _ = s.Find("a").Attr("href")
				p.Images = append(p.Images, *image)
			})
			p.StoreName = r.HTMLDoc.Find("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.goods-shop-info-wrap > div > div.upper-part > div.new-shop-panel-header > div.left").Text()
			p.StoreName = TrimHtml(p.StoreName)
			r.HTMLDoc.Find("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.goods-shop-info-wrap > div > div.second-part > ul > li:nth-child(2) > div.new-shop-panel-item-value").Each(func(_ int, s *goquery.Selection) {
				phone :=  new(Phone)
				phone.PhoneNumber= TrimHtml(s.Find("a").Text())
				p.Phones = append(p.Phones, *phone)
			})

			p.Address = r.HTMLDoc.Find("body > div.web-container > div.item-clear-container.main-function > div.promote-goods-page-container > div.goods-shop-info-wrap > div > div.second-part > ul > li:nth-child(6) > div.new-shop-panel-item-value.new-shop-panel-item-value-showall").Text()
			p.Address = TrimHtml(p.Address)

			r.HTMLDoc.Find("body > div.web-container > div:nth-child(12) > div.detail-left > div.promote-goods-details-container.vip4-detail > div.promote-goods-details-right.pull-right > div.details-right-content > div").Each(func(_ int, s *goquery.Selection) {
				detail := new(Detail)
				detail.Label = TrimHtml(s.Text())
				p.Details = append(p.Details, *detail)
			})

		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
	return p
}