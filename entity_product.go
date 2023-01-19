package yext

import (
	"encoding/json"
)

const ENTITYTYPE_PRODUCT EntityType = "product"

type ProductEntity struct {
	BaseEntity
	AvailabilityDate             *string           `json:"availabilityDate,omitempty"`
	Brand                        *string           `json:"brand,omitempty"`
	Bundle                       **bool            `json:"bundle,omitempty"`
	Color                        *string           `json:"color,omitempty"`
	Condition                    *string           `json:"condition,omitempty"`
	EnergyEfficiencyClass        *string           `json:"energyEfficiencyClass,omitempty"`
	ExpirationDate               *string           `json:"expirationDate,omitempty"`
	GTIN                         *string           `json:"gtin,omitempty"`
	IncludesAdultContent         **bool            `json:"includesAdultContent,omitempty"`
	InstallmentPlan              **Plan            `json:"installmentPlan,omitempty"`
	InventoryQuantity            *int              `json:"inventoryQuantity,omitempty"`
	LoyaltyPoints                **LoyaltyPoints   `json:"loyaltyPoints,omitempty"`
	Material                     *string           `json:"material,omitempty"`
	MaximumEnergyEfficiencyClass *string           `json:"maximumEnergyEfficiencyClass,omitempty"`
	MinimumEnergyEfficiencyClass *string           `json:"minimumEnergyEfficiencyClass,omitempty"`
	MPN                          *string           `json:"mpn,omitempty"`
	NumberOfItemsInPack          *int              `json:"numberOfItemsInPack,omitempty"`
	Pattern                      *string           `json:"pattern,omitempty"`
	Price                        **Price           `json:"price,omitempty"`
	PrimaryPhoto                 **Photo           `json:"primaryPhoto,omitempty"`
	ProductHighlights            *UnorderedStrings `json:"productHighlights,omitempty"`
	ProductionCost               **Price           `json:"productionCost,omitempty"`
	RichTextDescription          *string           `json:"richTextDescription,omitempty"`
	Sales                        *[]**Sale         `json:"sales,omitempty"`
	SalesChannel                 *string           `json:"salesChannel,omitempty"`
	ShippingHeight               **ProductUnit     `json:"shippingHeight,omitempty"`
	ShippingLength               **ProductUnit     `json:"shippingLength,omitempty"`
	ShippingWeight               **ProductUnit     `json:"shippingWeight,omitempty"`
	ShippingWidth                **ProductUnit     `json:"shippingWidth,omitempty"`
	Size                         *string           `json:"size,omitempty"`
	SizeSystem                   *string           `json:"sizeSystem,omitempty"`
	SizeType                     *[]string         `json:"sizeType,omitempty"`
	SKU                          *string           `json:"sku,omitempty"`
	StockStatus                  *string           `json:"stockStatus,omitempty"`
	SubscriptionPlan             Plan              `json:"subscriptionPlan,omitempty"`
	TargetAgeGroup               *string           `json:"targetAgeGroup,omitempty"`
	TargetGender                 *string           `json:"targetGender,omitempty"`
	UnitPricingBaseMeasure       **ProductUnit     `json:"unitPricingBaseMeasure,omitempty"`
	UnitPricingMeasure           **ProductUnit     `json:"unitPricingMeasure,omitempty"`
	Name                         *string           `json:"name,omitempty"`
	Keywords                     *[]string         `json:"keywords,omitempty"`
	PhotoGallery                 *[]Photo          `json:"photoGallery,omitempty"`
	Videos                       *[]Video          `json:"videos,omitempty"`
	Timezone                     *string           `json:"timezone,omitempty"`
}

func (l *ProductEntity) UnmarshalJSON(data []byte) error {
	type Alias ProductEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(l),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(l, data)
}

func (y ProductEntity) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
	}
	return ""
}

func (y ProductEntity) GetRichTextDescription() string {
	if y.RichTextDescription != nil {
		return *y.RichTextDescription
	}
	return ""
}

func (y ProductEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y ProductEntity) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (y ProductEntity) GetVideos() (v []Video) {
	if y.Videos != nil {
		v = *y.Videos
	}
	return v
}

type Plan struct {
	NumberOfPeriods *int    `json:"numberOfPeriods,omitempty"`
	Price           **Price `json:"price,omitempty"`
	Period          *string `json:"period,omitempty"`
}

func NullablePlan(p *Plan) **Plan {
	return &p
}

func NullPlan() **Plan {
	var p *Plan
	return &p
}

type LoyaltyPoints struct {
	LoyaltyProgramName *string  `json:"loyaltyProgramName,omitempty"`
	NumberOfPoints     *int     `json:"numberOfPoints,omitempty"`
	RatioToCurrency    *float64 `json:"ratioToCurrency,omitempty"`
}

func NullableLoyaltyPoints(p *LoyaltyPoints) **LoyaltyPoints {
	return &p
}

func NullLoyaltyPoints() **LoyaltyPoints {
	var p *LoyaltyPoints
	return &p
}

type Sale struct {
	SalePrice     **Price `json:"salePrice,omitempty"`
	SaleStartDate *string `json:"saleStartDate,omitempty"`
	SaleEndDate   *string `json:"saleEndDate,omitempty"`
}

func NullableSale(s *Sale) **Sale {
	return &s
}

func NullSale() **Sale {
	var s *Sale
	return &s
}

type ProductUnit struct {
	Unit  *string `json:"unit,omitempty"`
	Value *string `json:"value,omitempty"`
}

func NullableProductUnit(p *ProductUnit) **ProductUnit {
	return &p
}

func NullProductUnit() **ProductUnit {
	var p *ProductUnit
	return &p
}
