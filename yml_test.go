package goyml

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestYml(t *testing.T) {

	// Create YML
	ymlCat := NewYML("BestShop", "Best online seller Inc.", "http://best.seller.ru/")

	// Additional info
	ymlCat.Shop.Platform = "CMS"
	ymlCat.Shop.Version = "2.3"
	ymlCat.Shop.Agency = "Agency"
	ymlCat.Shop.Email = "CMS@CMS.ru"

	ymlCat.AddCurrency("RUR", "1", 0)

	someTime, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	ymlCat.SetDate(someTime)

	// Categories
	ymlCat.AddCategory(1, 0, "Книги")
	ymlCat.AddCategory(2, 1, "Детективы")
	ymlCat.AddCategory(3, 1, "Боевики")
	ymlCat.AddCategory(4, 0, "Видео")
	ymlCat.AddCategory(5, 4, "Комедии")
	ymlCat.AddCategory(6, 0, "Принтеры")
	ymlCat.AddCategory(7, 0, "Оргтехника")

	// Delivery
	ymlCat.AddDeliveryOption(0, 0, 0, 10)
	ymlCat.AddDeliveryOption(0, 1, 0, 0)

	// Simple Offer
	offer := Offer{
		Id:                   "123",
		Available:            true,
		Bid:                  21,
		Url:                  "http://best.seller.ru/product_page.asp?pid=12348",
		Price:                600,
		OldPrice:             800,
		CurrencyId:           "USD",
		CategoryId:           6,
		Picture:              []string{"http://best.seller.ru/img/device12345.jpg"},
		Store:                false,
		Pickup:               true,
		Delivery:             false,
		Name:                 "Наручные часы Casio A1234567B",
		Vendor:               "Casio",
		VendorCode:           "A1234567B",
		Description:          "Изящные наручные часы.",
		SalesNotes:           "Необходима предоплата.",
		ManufacturerWarranty: true,
		CountryOfOrigin:      "Япония",
		Cpa:                  1,
	}
	offer.AddBarcode("0123456789012")
	offer.AddAge("year", "18")

	// Offer vendor.model
	offer2 := Offer{
		Id:                   "12341",
		Available:            true,
		Type:                 TypeVendorModel,
		Bid:                  13,
		Url:                  "http://best.seller.ru/product_page.asp?pid=12344",
		Price:                16800,
		OldPrice:             17000,
		CurrencyId:           "USD",
		CategoryId:           6,
		Picture:              []string{"http://best.seller.ru/img/device12345.jpg"},
		Store:                false,
		Pickup:               false,
		Delivery:             true,
		TypePrefix:           "Принтер",
		Vendor:               "HP",
		Model:                "Deskjet D2663",
		Description:          "Серия принтеров для людей, которым нужен надежный, простой в использовании цветной принтер для повседневной печати...",
		SalesNotes:           "Необходима предоплата.",
		ManufacturerWarranty: true,
		CountryOfOrigin:      "Япония",
		Cpa:                  1,
		Rec:                  "123123,1214,243",
		Expiry:               "P5Y",
		Weight:               2.07,
		Dimensions:           "100/25.45/11.112",
	}
	offer2.AddBarcode("1234567890120")
	offer2.AddParam("Максимальный формат", "", "А4")
	offer2.AddParam("Технология печати", "", "термическая струйная")
	offer2.AddParam("Тип печати", "", "Цветная")
	offer2.AddParam("Количество страниц в месяц", "стр", "1000")
	offer2.AddParam("Потребляемая мощность", "Вт", "20")
	offer2.AddParam("Вес", "кг", "2.73")

	err := offer.Validate()
	if err != nil {
		t.Error(err)
	}
	err = offer2.Validate()
	if err != nil {
		t.Error(err)
	}

	ymlCat.AddOffer(offer)
	ymlCat.AddOffer(offer2)

	data, err := xml.Marshal(ymlCat)
	if err != nil {
		t.Error(err)
	}

	var expected = `<yml_catalog date="2013-02-03 00:00"><shop><name>BestShop</name><company>Best online seller Inc.</company><url>http://best.seller.ru/</url><platform>CMS</platform><version>2.3</version><agency>Agency</agency><email>CMS@CMS.ru</email><currencies><currency id="RUR" rate="1" plus="0"></currency></currencies><categories><category id="1">Книги</category><category id="2" parentId="1">Детективы</category><category id="3" parentId="1">Боевики</category><category id="4">Видео</category><category id="5" parentId="4">Комедии</category><category id="6">Принтеры</category><category id="7">Оргтехника</category></categories><delivery-options><option cost="0" days="0" order-before="10"></option><option cost="0" days="1"></option></delivery-options><offers><offer id="123" bid="21" available="true"><url>http://best.seller.ru/product_page.asp?pid=12348</url><price>600</price><oldprice>800</oldprice><currencyId>USD</currencyId><categoryId>6</categoryId><picture>http://best.seller.ru/img/device12345.jpg</picture><pickup>true</pickup><name>Наручные часы Casio A1234567B</name><vendor>Casio</vendor><vendorCode>A1234567B</vendorCode><description>Изящные наручные часы.</description><sales_notes>Необходима предоплата.</sales_notes><manufacturer_warranty>true</manufacturer_warranty><country_of_origin>Япония</country_of_origin><age unit="year">18</age><barcode>0123456789012</barcode><cpa>1</cpa></offer><offer id="12341" bid="13" type="vendor.model" available="true"><url>http://best.seller.ru/product_page.asp?pid=12344</url><price>16800</price><oldprice>17000</oldprice><currencyId>USD</currencyId><categoryId>6</categoryId><picture>http://best.seller.ru/img/device12345.jpg</picture><delivery>true</delivery><typePrefix>Принтер</typePrefix><vendor>HP</vendor><model>Deskjet D2663</model><description>Серия принтеров для людей, которым нужен надежный, простой в использовании цветной принтер для повседневной печати...</description><sales_notes>Необходима предоплата.</sales_notes><manufacturer_warranty>true</manufacturer_warranty><country_of_origin>Япония</country_of_origin><barcode>1234567890120</barcode><cpa>1</cpa><rec>123123,1214,243</rec><expiry>P5Y</expiry><weight>2.07</weight><dimensions>100/25.45/11.112</dimensions><param name="Максимальный формат">А4</param><param name="Технология печати">термическая струйная</param><param name="Тип печати">Цветная</param><param name="Количество страниц в месяц" unit="стр">1000</param><param name="Потребляемая мощность" unit="Вт">20</param><param name="Вес" unit="кг">2.73</param></offer></offers></shop></yml_catalog>`
	if string(data) != expected {
		t.Log(string(data))
		t.Error("YML incorrect")
	}
}
