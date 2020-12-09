package main

import (
	"fmt"
	"log"
)

func init()  {
	log.SetFlags(log.Llongfile|log.LstdFlags)
}
func main()  {

 	var zpwdp ZwdProduct
	url := "https://gz.17zwd.com/item.htm?GID=117517428&spm=0.42.0.378.117517428.0.1"
	zpwdp = zpwdp.Scrape(url)
	fmt.Println(zpwdp)

	var vvicp VvicProduct
	vurl := "https://www.vvic.com/item/5fb9cd69d932b600016f1337"
	vvicp = vvicp.Scrape(vurl)
	fmt.Println(vvicp)
}
