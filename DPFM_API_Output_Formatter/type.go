package dpfm_api_output_formatter

type OrdersSDC struct {
	ConnectionKey       string             `json:"connection_key"`
	Result              bool               `json:"result"`
	RedisKey            string             `json:"redis_key"`
	Filepath            string             `json:"filepath"`
	APIStatusCode       int                `json:"api_status_code"`
	RuntimeSessionID    string             `json:"runtime_session_id"`
	BusinessPartnerID   *int               `json:"business_partner"`
	ServiceLabel        string             `json:"service_label"`
	APIType             string             `json:"api_type"`
	DataConcatenation   *DataConcatenation `json:"DataConcatenation"`
	APISchema           string             `json:"api_schema"`
	Accepter            []string           `json:"accepter"`
	Deleted             bool               `json:"deleted"`
	APIProcessingResult bool               `json:"api_processing_result"`
	APIProcessingError  *string            `json:"api_processing_error"`
}

type DataConcatenation struct {
	Header             OrdersHeader               `json:"OrdersHeader"`
	Item               []OrdersItem               `json:"OrdersItem"`
	ItemPricingElement []OrdersItemPricingElement `json:"OrdersItemPricingElement"`
	ItemScheduleLine   []OrdersItemScheduleLine   `json:"OrdersItemScheduleLine"`
	Address            []OrdersAddress            `json:"OrdersAddress"`
	Partner            []OrdersPartner            `json:"OrdersPartner"`
}

type OrdersHeader struct {
	OrderID                          int      `json:"OrderID" csv:"1"`
	OrderDate                        *string  `json:"OrderDate" csv:"2"`
	OrderType                        *string  `json:"OrderType" csv:"3"`
	SupplyChainRelationshipID        *int     `json:"SupplyChainRelationshipID" csv:"4"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID" csv:"5"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID" csv:"6"`
	Buyer                            *int     `json:"Buyer" csv:"7"`
	Seller                           *int     `json:"Seller" csv:"8"`
	BillToParty                      *int     `json:"BillToParty" csv:"9"`
	BillFromParty                    *int     `json:"BillFromParty" csv:"10"`
	BillToCountry                    *string  `json:"BillToCountry" csv:"11"`
	BillFromCountry                  *string  `json:"BillFromCountry" csv:"12"`
	Payer                            *int     `json:"Payer" csv:"13"`
	Payee                            *int     `json:"Payee" csv:"14"`
	CreationDate                     *string  `json:"CreationDate" csv:"15"`
	LastChangeDate                   *string  `json:"LastChangeDate" csv:"16"`
	ContractType                     *string  `json:"ContractType" csv:"17"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate" csv:"18"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate" csv:"19"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate" csv:"20"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate" csv:"21"`
	TotalNetAmount                   *float32 `json:"TotalNetAmount" csv:"22"`
	TotalTaxAmount                   *float32 `json:"TotalTaxAmount" csv:"23"`
	TotalGrossAmount                 *float32 `json:"TotalGrossAmount" csv:"24"`
	HeaderDeliveryStatus             *string  `json:"HeaderDeliveryStatus" csv:"25"`
	HeaderBillingStatus              *string  `json:"HeaderBillingStatus" csv:"26"`
	HeaderDocReferenceStatus         *string  `json:"HeaderDocReferenceStatus" csv:"27"`
	TransactionCurrency              *string  `json:"TransactionCurrency" csv:"28"`
	PricingDate                      *string  `json:"PricingDate" csv:"29"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate" csv:"30"`
	RequestedDeliveryDate            *string  `json:"RequestedDeliveryDate" csv:"31"`
	RequestedDeliveryTime            *string  `json:"RequestedDeliveryTime" csv:"32"`
	HeaderCompleteDeliveryIsDefined  *bool    `json:"HeaderCompleteDeliveryIsDefined" csv:"33"`
	Incoterms                        *string  `json:"Incoterms" csv:"34"`
	PaymentTerms                     *string  `json:"PaymentTerms" csv:"35"`
	PaymentMethod                    *string  `json:"PaymentMethod" csv:"36"`
	ReferenceDocument                *int     `json:"ReferenceDocument" csv:"37"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem" csv:"38"`
	AccountAssignmentGroup           *string  `json:"AccountAssignmentGroup" csv:"39"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate" csv:"40"`
	InvoiceDocumentDate              *string  `json:"InvoiceDocumentDate" csv:"41"`
	IsExportImport                   *bool    `json:"IsExportImport" csv:"42"`
	HeaderText                       *string  `json:"HeaderText" csv:"43"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus" csv:"44"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus" csv:"45"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus" csv:"46"`
	IsCancelled                      *bool    `json:"IsCancelled" csv:"47"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion" csv:"48"`
}

type OrdersItem struct {
	OrderID                                       int      `json:"OrderID" csv:"1"`
	OrderItem                                     int      `json:"OrderItem" csv:"66"`
	OrderItemCategory                             *string  `json:"OrderItemCategory" csv:"67"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID" csv:"68"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID" csv:"69"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID" csv:"70"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID" csv:"71"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID" csv:"72"`
	OrderItemText                                 *string  `json:"OrderItemText" csv:"73"`
	OrderItemTextByBuyer                          *string  `json:"OrderItemTextByBuyer" csv:"74"`
	OrderItemTextBySeller                         *string  `json:"OrderItemTextBySeller" csv:"75"`
	Product                                       *string  `json:"Product" csv:"76"`
	ProductStandardID                             *string  `json:"ProductStandardID" csv:"77"`
	ProductGroup                                  *string  `json:"ProductGroup" csv:"78"`
	BaseUnit                                      *string  `json:"BaseUnit" csv:"79"`
	PricingDate                                   *string  `json:"PricingDate" csv:"80"`
	PriceDetnExchangeRate                         *float32 `json:"PriceDetnExchangeRate" csv:"81"`
	RequestedDeliveryDate                         *string  `json:"RequestedDeliveryDate" csv:"82"`
	RequestedDeliveryTime                         *string  `json:"RequestedDeliveryTime" csv:"83"`
	DeliverToParty                                *int     `json:"DeliverToParty" csv:"84"`
	DeliverFromParty                              *int     `json:"DeliverFromParty" csv:"85"`
	CreationDate                                  *string  `json:"CreationDate" csv:"86"`
	CreationTime                                  *string  `json:"CreationTime" csv:"87"`
	LastChangeDate                                *string  `json:"LastChangeDate" csv:"88"`
	LastChangeTime                                *string  `json:"LastChangeTime" csv:"89"`
	DeliverToPlant                                *string  `json:"DeliverToPlant" csv:"90"`
	DeliverToPlantTimeZone                        *string  `json:"DeliverToPlantTimeZone" csv:"91"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation" csv:"92"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant" csv:"93"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant" csv:"94"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch" csv:"95"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate" csv:"96"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime" csv:"97"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate" csv:"98"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime" csv:"99"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant" csv:"100"`
	DeliverFromPlantTimeZone                      *string  `json:"DeliverFromPlantTimeZone" csv:"101"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation" csv:"102"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant" csv:"103"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant" csv:"104"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch" csv:"105"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate" csv:"106"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime" csv:"107"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate" csv:"108"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime" csv:"109"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit" csv:"110"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner" csv:"111"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant" csv:"112"`
	StockConfirmationPlantTimeZone                *string  `json:"StockConfirmationPlantTimeZone" csv:"113"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant" csv:"114"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant" csv:"115"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch" csv:"116"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate" csv:"117"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime" csv:"118"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate" csv:"119"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime" csv:"120"`
	ServicesRenderingDate                         *string  `json:"ServicesRenderingDate" csv:"121"`
	OrderQuantityInBaseUnit                       *float32 `json:"OrderQuantityInBaseUnit" csv:"122"`
	OrderQuantityInDeliveryUnit                   *float32 `json:"OrderQuantityInDeliveryUnit" csv:"123"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy" csv:"124"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus" csv:"125"`
	ConfirmedOrderQuantityInBaseUnit              *float32 `json:"ConfirmedOrderQuantityInBaseUnit" csv:"126"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit" csv:"127"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight" csv:"128"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight" csv:"129"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight" csv:"130"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight" csv:"131"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity" csv:"132"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit" csv:"133"`
	NetAmount                                     *float32 `json:"NetAmount" csv:"134"`
	TaxAmount                                     *float32 `json:"TaxAmount" csv:"135"`
	GrossAmount                                   *float32 `json:"GrossAmount" csv:"136"`
	InvoiceDocumentDate                           *string  `json:"InvoiceDocumentDate" csv:"137"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner" csv:"138"`
	ProductionPlant                               *string  `json:"ProductionPlant" csv:"139"`
	ProductionPlantTimeZone                       *string  `json:"ProductionPlantTimeZone" csv:"140"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation" csv:"141"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant" csv:"142"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant" csv:"143"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch" csv:"144"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate" csv:"145"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime" csv:"146"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate" csv:"147"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime" csv:"148"`
	Incoterms                                     *string  `json:"Incoterms" csv:"149"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification" csv:"150"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry" csv:"151"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry" csv:"152"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification" csv:"153"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup" csv:"154"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup" csv:"155"`
	PaymentTerms                                  *string  `json:"PaymentTerms" csv:"156"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate" csv:"157"`
	PaymentDueDate                                *string  `json:"PaymentDueDate" csv:"158"`
	NetPaymentDays                                *int     `json:"NetPaymentDays" csv:"159"`
	PaymentMethod                                 *string  `json:"PaymentMethod" csv:"160"`
	Project                                       *string  `json:"Project" csv:"161"`
	AccountingExchangeRate                        *float32 `json:"AccountingExchangeRate" csv:"162"`
	ReferenceDocument                             *int     `json:"ReferenceDocument" csv:"163"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem" csv:"164"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined" csv:"165"`
	ItemDeliveryStatus                            *string  `json:"ItemDeliveryStatus" csv:"166"`
	IssuingStatus                                 *string  `json:"IssuingStatus" csv:"167"`
	ReceivingStatus                               *string  `json:"ReceivingStatus" csv:"168"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus" csv:"169"`
	TaxCode                                       *string  `json:"TaxCode" csv:"170"`
	TaxRate                                       *float32 `json:"TaxRate" csv:"171"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin" csv:"172"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage" csv:"173"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus" csv:"174"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus" csv:"175"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus" csv:"176"`
	IsCancelled                                   *bool    `json:"IsCancelled" csv:"177"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion" csv:"178"`
}

type OrdersItemPricingElement struct {
	OrderID                    int      `json:"OrderID" csv:"1"`
	OrderItem                  int      `json:"OrderItem" csv:"66"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID" csv:"179"`
	Buyer                      int      `json:"Buyer" csv:"180"`
	Seller                     int      `json:"Seller" csv:"181"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter" csv:"182"`
	ConditionRecord            *int     `json:"ConditionRecord" csv:"183"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber" csv:"184"`
	ConditionType              *string  `json:"ConditionType" csv:"185"`
	PricingDate                *string  `json:"PricingDate" csv:"186"`
	ConditionRateValue         *float32 `json:"ConditionRateValue" csv:"187"`
	ConditionCurrency          *string  `json:"ConditionCurrency" csv:"188"`
	ConditionQuantity          *float32 `json:"ConditionQuantity" csv:"189"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit" csv:"190"`
	TaxCode                    *string  `json:"TaxCode" csv:"191"`
	ConditionAmount            *float32 `json:"ConditionAmount" csv:"192"`
	TransactionCurrency        *string  `json:"TransactionCurrency" csv:"193"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged" csv:"194"`
}

type OrdersItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID" csv:"1"`
	OrderItem                                       int      `json:"OrderItem" csv:"66"`
	ScheduleLine                                    int      `json:"ScheduleLine" csv:"195"`
	SupplyChainRelationshipID                       *int     `json:"SupplyChainRelationshipID" csv:"196"`
	SupplyChainRelationshipStockConfPlantID         *int     `json:"SupplyChainRelationshipStockConfPlantID" csv:"197"`
	Product                                         *string  `json:"Product" csv:"198"`
	StockConfirmationBussinessPartner               *int     `json:"StockConfirmationBussinessPartner" csv:"199"`
	StockConfirmationPlant                          *string  `json:"StockConfirmationPlant" csv:"200"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone" csv:"201"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch" csv:"202"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate" csv:"203"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate" csv:"204"`
	RequestedDeliveryDate                           *string  `json:"RequestedDeliveryDate" csv:"205"`
	RequestedDeliveryTime                           *string  `json:"RequestedDeliveryTime" csv:"206"`
	ConfirmedDeliveryDate                           *string  `json:"ConfirmedDeliveryDate" csv:"207"`
	ScheduleLineOrderQuantity                       *float32 `json:"ScheduleLineOrderQuantity" csv:"208"`
	OriginalOrderQuantityInBaseUnit                 *float32 `json:"OriginalOrderQuantityInBaseUnit" csv:"209"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit" csv:"210"`
	ConfirmedOrderQuantityByPDTAvailCheck           *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheck" csv:"211"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit" csv:"212"`
	UndeliveredQuantityInBaseUnit                   *float32 `json:"UndeliveredQuantityInBaseUnit" csv:"213"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit" csv:"214"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed" csv:"215"`
	PlusMinusFlag                                   *string  `json:"PlusMinusFlag" csv:"216"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus" csv:"217"`
	IsCancelled                                     *bool    `json:"IsCancelled" csv:"218"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion" csv:"219"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID" csv:"1"`
	PartnerFunction         *string `json:"PartnerFunction" csv:"220"`
	BusinessPartner         *int    `json:"BusinessPartner" csv:"221"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName" csv:"222"`
	BusinessPartnerName     *string `json:"BusinessPartnerName" csv:"223"`
	Organization            *string `json:"Organization" csv:"224"`
	Country                 *string `json:"Country" csv:"225"`
	Language                *string `json:"Language" csv:"226"`
	Currency                *string `json:"Currency" csv:"227"`
	ExternalDocumentID      *string `json:"ExternalDocumentID" csv:"228"`
	AddressID               *int    `json:"AddressID" csv:"229"`
}

type OrdersAddress struct {
	OrderID     int     `json:"OrderID" csv:"1"`
	AddressID   *int    `json:"AddressID" csv:"56"`
	PostalCode  *string `json:"PostalCode" csv:"57"`
	LocalRegion *string `json:"LocalRegion" csv:"58"`
	Country     *string `json:"Country" csv:"59"`
	District    *string `json:"District" csv:"60"`
	StreetName  *string `json:"StreetName" csv:"61"`
	CityName    *string `json:"CityName" csv:"62"`
	Building    *string `json:"Building" csv:"63"`
	Floor       *int    `json:"Floor" csv:"64"`
	Room        *int    `json:"Room" csv:"65"`
}

// func (nb *OrdersAddress) UnmarshalCSV(val string) error {
// 	// Workaround for https://github.com/gocarina/gocsv/issues/202
// 	// gocsv unmarshals "" into &false, not nil for a *bool
// 	// This is a gocsv bug, but to workaround it, create our own type to emulate
// 	// the same behavior
// 	if val == "" {
// 		return nil
// 	}

// 	var b = []struct {
// 		B *bool `csv:"b"`
// 	}{}
// 	csv := fmt.Sprintf("b\n%s", val)
// 	err := gocsv.UnmarshalString(csv, &b)
// 	if err != nil {
// 		return err
// 	}
// 	nb.B = b[0].B
// 	return nil
// }
