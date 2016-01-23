// Yandex YML generator: https://yandex.ru/support/partnermarket/yml/about-yml.xml
package goyml

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const Header = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE yml_catalog SYSTEM "shops.dtd">
`

func NewYML(name, company, url string) Catalog {
	yml := Catalog{}
	yml.SetDate(time.Now())
	yml.Shop.Name = name
	yml.Shop.Company = company
	yml.Shop.Url = url
	return yml
}

type Catalog struct {
	XMLName struct{} `xml:"yml_catalog"`
	Shop    Shop     `xml:"shop"`
	Date    string   `xml:"date,attr"`
}

func (c *Catalog) SetDate(t time.Time) {
	c.Date = t.Format("2006-01-02 15:04")
}

func (c *Catalog) AddCurrency(id, rate string, plus float64) {
	c.Shop.Currencies.add(id, rate, plus)
}

func (c *Catalog) AddCategory(id, parentId int, name string) {
	c.Shop.Categories.add(id, parentId, name)
}

func (c *Catalog) AddDeliveryOption(cost int, daysFrom, daysTo int, orderBefore int) {
	if c.Shop.DeliveryOptions == nil {
		c.Shop.DeliveryOptions = &DeliveryOptions{}
	}
	c.Shop.DeliveryOptions.add(cost, daysFrom, daysTo, orderBefore)
}

func (c *Catalog) AddOffer(offer Offer) {
	c.Shop.Offers.add(offer)
}

type Shop struct {
	Name            string           `xml:"name"`
	Company         string           `xml:"company"`
	Url             string           `xml:"url"`
	Platform        string           `xml:"platform,omitempty"`
	Version         string           `xml:"version,omitempty"`
	Agency          string           `xml:"agency,omitempty"`
	Email           string           `xml:"email,omitempty"`
	Currencies      Currencies       `xml:"currencies"`
	Categories      Categories       `xml:"categories"`
	DeliveryOptions *DeliveryOptions `xml:"delivery-options"`
	Cpa             int              `xml:"cpa,omitempty"`
	Offers          Offers           `xml:"offers"`
}

type Currencies struct {
	Currency []Currency `xml:"currency"`
}

func (cur *Currencies) add(id, rate string, plus float64) {
	cur.Currency = append(cur.Currency, Currency{Id: id, Rate: rate, Plus: plus})
}

type Currency struct {
	Id   string  `xml:"id,attr"`
	Rate string  `xml:"rate,attr"`
	Plus float64 `xml:"plus,attr"`
}

type Categories struct {
	Category []Category `xml:"category"`
}

func (cat *Categories) add(id, parentId int, name string) {
	cat.Category = append(cat.Category, Category{Id: id,
		ParentId: parentId, Name: name})
}

type Category struct {
	Id       int    `xml:"id,attr"`
	ParentId int    `xml:"parentId,attr,omitempty"`
	Name     string `xml:",innerxml"`
}

type DeliveryOptions struct {
	Options []DeliveryOption `xml:"option"`
}

//// cost 0 - free of charge delivery
func (do *DeliveryOptions) add(cost int, daysFrom, daysTo int, orderBefore int) {
	daysStr := ""
	if daysFrom > 255 {
		daysFrom = 255
	}
	if daysTo > 255 {
		daysTo = 255
	}
	if daysTo < daysFrom {
		daysTo = daysFrom
	}
	if daysFrom == daysTo || daysTo == 0 {
		daysStr = strconv.Itoa(daysFrom)
	} else {
		daysStr = fmt.Sprintf("%d-%d", daysFrom, daysTo)
	}
	option := DeliveryOption{
		Cost:        cost,
		Days:        daysStr,
		OrderBefore: orderBefore,
	}
	do.Options = append(do.Options, option)
}

type DeliveryOption struct {
	Cost        int    `xml:"cost,attr"`
	Days        string `xml:"days,attr"`
	OrderBefore int    `xml:"order-before,attr,omitempty"`
}

type Offers struct {
	Offers []Offer `xml:"offer"`
}

func (o *Offers) add(offer Offer) {
	o.Offers = append(o.Offers, offer)
}

type OfferType string

const (
	TypeVendorModel OfferType = "vendor.model"
)

// Available = true - delivery in 2 days
// Available = false - delivery from 3 days to 2 months
type Offer struct {
	Id                   string           `xml:"id,attr"`
	Bid                  uint             `xml:"bid,attr,omitempty"`
	CBid                 uint             `xml:"cid,attr,omitempty"`
	Type                 OfferType        `xml:"type,attr,omitempty"`
	Available            bool             `xml:"available,attr"`
	Url                  string           `xml:"url,omitempty"`
	Price                float64          `xml:"price"`
	OldPrice             float64          `xml:"oldprice,omitempty"`
	CurrencyId           string           `xml:"currencyId"`
	CategoryId           int              `xml:"categoryId"`
	MarketCategory       string           `xml:"market_category,omitempty"`
	Picture              []string         `xml:"picture,omitempty"`
	Store                bool             `xml:"store,omitempty"`
	Pickup               bool             `xml:"pickup,omitempty"`
	Delivery             bool             `xml:"delivery,omitempty"`
	DeliveryOptions      *DeliveryOptions `xml:"delivery-options,omitempty"`
	Name                 string           `xml:"name,omitempty"`
	TypePrefix           string           `xml:"typePrefix,omitempty"`
	Vendor               string           `xml:"vendor,omitempty"`
	VendorCode           string           `xml:"vendorCode,omitempty"`
	Model                string           `xml:"model,omitempty"`
	Description          string           `xml:"description,omitempty"`
	SalesNotes           string           `xml:"sales_notes,omitempty"`
	ManufacturerWarranty bool             `xml:"manufacturer_warranty,omitempty"`
	CountryOfOrigin      string           `xml:"country_of_origin,omitempty"`
	Downloadable         bool             `xml:"downloadable,omitempty"`
	Adult                bool             `xml:"adult,omitempty"`
	Age                  *Age             `xml:"age,omitempty"`
	Barcode              []string         `xml:"barcode,omitempty"`
	Cpa                  int              `xml:"cpa,omitempty"`
	Rec                  string           `xml:"rec,omitempty"`
	Expiry               string           `xml:"expiry,omitempty"`
	Weight               float64          `xml:"weight,omitempty"`
	Dimensions           string           `xml:"dimensions,omitempty"`
	Params               []Param          `xml:"param,omitempty"`
}

func (o *Offer) AddPicture(pic string) {
	o.Picture = append(o.Picture, pic)
}

func (o *Offer) AddDeliveryOption(cost int, daysFrom, daysTo int, orderBefore int) {
	if o.DeliveryOptions == nil {
		o.DeliveryOptions = &DeliveryOptions{}
	}
	o.DeliveryOptions.add(cost, daysFrom, daysTo, orderBefore)
}

func (o *Offer) AddBarcode(barcode string) {
	o.Barcode = append(o.Barcode, barcode)
}

func (o *Offer) AddAge(unit string, value string) {
	if o.Age == nil {
		o.Age = &Age{Unit: unit, Value: value}
	}
}

func (o *Offer) AddParam(name, unit, value string) {
	o.Params = append(o.Params, Param{Name: name, Unit: unit, Value: value})
}

type Param struct {
	Name  string `xml:"name,attr"`
	Unit  string `xml:"unit,attr,omitempty"`
	Value string `xml:",innerxml"`
}

// Validate offer
func (o Offer) Validate() error {
	if utf8.RuneCountInString(o.Id) > 20 {
		return errors.New("Id more than 20 cahrs")
	}
	if o.Type == TypeVendorModel && (o.Vendor == "" || o.Model == "") {
		return errors.New("Vendor or Model is empty")
	}
	if o.Price == 0 {
		return errors.New("Price is zero")
	}
	if o.OldPrice > 0 && o.OldPrice <= o.Price {
		return errors.New("OldPrice less than Price")
	}
	if utf8.RuneCountInString(o.CurrencyId) != 3 {
		return errors.New("CurrencyId less than 3 chars")
	}
	if o.CategoryId > 999999999999999999 {
		return errors.New("CategoryId more than 18 cahrs")
	}
	for _, pic := range o.Picture {
		if utf8.RuneCountInString(pic) > 512 {
			return errors.New("Picture more than 512 cahrs")
		}
	}

	descrTrim := strings.Replace(o.Description, ",", "", -1)
	descrTrim = strings.Replace(descrTrim, ".", "", -1)
	if utf8.RuneCountInString(descrTrim) > 175 {
		return errors.New("Description more than 175 cahrs")
	}

	if utf8.RuneCountInString(o.SalesNotes) > 50 {
		log.Println(o.SalesNotes)
		return errors.New("SalesNotes more than 50 cahrs")
	}

	if o.CountryOfOrigin != "" {
		if _, ok := Countries[o.CountryOfOrigin]; !ok {
			return errors.New("CountryOfOrigin not valid")
		}
	}
	if o.Age != nil && o.Age.Unit != "" {
		switch o.Age.Unit {
		case "year":
			if !isInSlice([]string{"0", "6", "12", "16", "18"}, o.Age.Value) {
				return errors.New("Age.Value is incorrect")
			}
		case "month":
			if !isInSlice([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}, o.Age.Value) {
				return errors.New("Age.Value is incorrect")
			}
		default:
			return errors.New("Age.Unit is incorrect")
		}
	}
	return nil
}

type Age struct {
	Unit  string `xml:"unit,attr"`
	Value string `xml:",innerxml"`
}

func isInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
