package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type VvicProduct struct {
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

func (p VvicProduct) Scrape(url string) VvicProduct{

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			p.Name = r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fl.item-left.mt20 > div.product-detail > div.d-name > strong").Text()
			p.OriginalPrice = r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fl.item-left.mt20 > div.product-detail > div.price-time-buyer > div:nth-child(2) > div.p-value > span.d-sale").Text()
			p.WholesalePrice = r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fl.item-left.mt20 > div.product-detail > div.price-time-buyer > div.v-price.d-p > div.p-value > span > strong.d-sale").Text()
			p.GoodsNo = r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fl.item-left.mt20 > div.product-detail > dl:nth-child(6) > dd:nth-child(1) > div.value.ff-arial").Text()
			p.GoodsNo = TrimHtml(p.GoodsNo)
			p.OnSaleTime = r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fl.item-left.mt20 > div.product-detail > dl:nth-child(6) > dd:nth-child(2) > div.value.ff-arial").Text()
			p.OnSaleTime = TrimHtml(p.OnSaleTime)
			r.HTMLDoc.Find("#j-buy > dd:nth-child(2) > div.value.color-choice > ul > li.selectColorPic.selectColor").Each(func(_ int, s *goquery.Selection) {
				color :=  new(Color)
				color.ColorName, _ = s.Attr("data-color")
				p.Colors = append(p.Colors, *color)
			})
			r.HTMLDoc.Find("#size-container > div.value.goods-choice.goods-choice-size.goods-choice-size__height > ul > div > li").Each(func(_ int, s *goquery.Selection) {
				size :=  new(Size)
				size.SizeName = TrimHtml(s.Find("span.selectSize").Text())
				size.SizePrice = s.Find("span.skuPrice.skuPrice").Text()
				p.Sizes = append(p.Sizes, *size)
			})
			r.HTMLDoc.Find("#thumblist > div.owl-stage-outer > div > div").Each(func(_ int, s *goquery.Selection) {
				image := new(Image)
				image.ImageUrl, _ = s.Find("div > a > img").Attr("big")
				p.Images = append(p.Images, *image)
			})
			p.StoreName = r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fr.item-right.mt20 > div.shop-info.shop-gold.shop-strength > div > h2 > span").Text()
			p.StoreName = TrimHtml(p.StoreName)
			r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fr.item-right.mt20 > div.shop-info.shop-gold.shop-strength > div > ul > li.tel-list > div.text").Each(func(_ int, s *goquery.Selection) {
				phone :=  new(Phone)
				phone.PhoneNumber = s.Find("p").Text()
				phone.PhoneNumber = TrimHtml(phone.PhoneNumber)
				p.Phones = append(p.Phones, *phone)
			})

			p.Address = r.HTMLDoc.Find("body > div.w.clearfix > div.item-content.clearfix > div.fr.item-right.mt20 > div.shop-info.shop-gold.shop-strength > div > ul > li:nth-child(8) > div.text").Text()
			p.Address = TrimHtml(p.Address)

			r.HTMLDoc.Find("#info > div:nth-child(1) > div.d-attr.clearfix > ul > li").Each(func(_ int, s *goquery.Selection) {
				detail := new(Detail)
				detail.Label = TrimHtml(s.Text())
				p.Details = append(p.Details, *detail)
			})
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
	return p
}