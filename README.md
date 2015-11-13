# goYML
Yandex Market YML (XML) generator for Go (golang)

File format description: https://yandex.ru/support/partnermarket/yml/about-yml.xml

Usage:
```go

	// Create YML
	ymlCat := NewYML("BestShop", "Best online seller Inc.", "http://best.seller.ru/")

	// Additional info
	ymlCat.Shop.Platform = "CMS"
	ymlCat.Shop.Version = "2.3"
	ymlCat.Shop.Agency = "Agency"
	ymlCat.Shop.Email = "CMS@CMS.ru"

	// id, rate, plus
	ymlCat.AddCurrency("RUR", 1, 0)

	// Categories
	ymlCat.AddCategory(1, 0, "Книги")
	ymlCat.AddCategory(2, 1, "Детективы")
	ymlCat.AddCategory(3, 1, "Боевики")
	ymlCat.AddCategory(4, 0, "Видео")
	ymlCat.AddCategory(5, 4, "Комедии")
	ymlCat.AddCategory(6, 0, "Принтеры")
	ymlCat.AddCategory(7, 0, "Оргтехника")

	// Delivery
	// cost, daysFrom, daysTo (if 0 - omitted), orderBefore
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
		Picture:              "http://best.seller.ru/img/device12345.jpg",
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
		Picture:              "http://best.seller.ru/img/device12345.jpg",
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
		fmt.Println(err.Error())
	}
	err = offer2.Validate()
	if err != nil {
		fmt.Println(err.Error())
	}

	ymlCat.AddOffer(offer)
	ymlCat.AddOffer(offer2)

	ExportToFile(ymlCat, "/path/to/file/yml.xml", true)

```