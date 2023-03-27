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
	QuantityPerPackage                            *float32 `json:"QuantityPerPackage" csv:"124"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy" csv:"125"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus" csv:"126"`
	ConfirmedOrderQuantityInBaseUnit              *float32 `json:"ConfirmedOrderQuantityInBaseUnit" csv:"127"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit" csv:"128"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight" csv:"129"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight" csv:"130"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight" csv:"131"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight" csv:"132"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity" csv:"133"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit" csv:"134"`
	NetAmount                                     *float32 `json:"NetAmount" csv:"135"`
	TaxAmount                                     *float32 `json:"TaxAmount" csv:"136"`
	GrossAmount                                   *float32 `json:"GrossAmount" csv:"137"`
	InvoiceDocumentDate                           *string  `json:"InvoiceDocumentDate" csv:"138"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner" csv:"139"`
	ProductionPlant                               *string  `json:"ProductionPlant" csv:"140"`
	ProductionPlantTimeZone                       *string  `json:"ProductionPlantTimeZone" csv:"141"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation" csv:"142"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant" csv:"143"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant" csv:"144"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch" csv:"145"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate" csv:"146"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime" csv:"147"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate" csv:"148"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime" csv:"149"`
	InspectionPlan                                *int     `json:"InspectionPlan" csv:"150"`
	InspectionPlant                               *string  `json:"InspectionPlant" csv:"151"`
	InspectionOrder                               *int     `json:"InspectionOrder" csv:"152"`
	Incoterms                                     *string  `json:"Incoterms" csv:"153"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification" csv:"154"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry" csv:"155"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry" csv:"156"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification" csv:"157"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup" csv:"158"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup" csv:"159"`
	PaymentTerms                                  *string  `json:"PaymentTerms" csv:"160"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate" csv:"161"`
	PaymentDueDate                                *string  `json:"PaymentDueDate" csv:"162"`
	NetPaymentDays                                *int     `json:"NetPaymentDays" csv:"163"`
	PaymentMethod                                 *string  `json:"PaymentMethod" csv:"164"`
	Project                                       *string  `json:"Project" csv:"165"`
	AccountingExchangeRate                        *float32 `json:"AccountingExchangeRate" csv:"166"`
	ReferenceDocument                             *int     `json:"ReferenceDocument" csv:"167"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem" csv:"168"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined" csv:"169"`
	ItemDeliveryStatus                            *string  `json:"ItemDeliveryStatus" csv:"170"`
	IssuingStatus                                 *string  `json:"IssuingStatus" csv:"171"`
	ReceivingStatus                               *string  `json:"ReceivingStatus" csv:"172"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus" csv:"173"`
	TaxCode                                       *string  `json:"TaxCode" csv:"174"`
	TaxRate                                       *float32 `json:"TaxRate" csv:"175"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin" csv:"176"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage" csv:"177"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus" csv:"178"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus" csv:"179"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus" csv:"180"`
	IsCancelled                                   *bool    `json:"IsCancelled" csv:"181"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion" csv:"182"`
}

type OrdersItemPricingElement struct {
	OrderID                    int      `json:"OrderID" csv:"1"`
	OrderItem                  int      `json:"OrderItem" csv:"66"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID" csv:"183"`
	Buyer                      int      `json:"Buyer" csv:"184"`
	Seller                     int      `json:"Seller" csv:"185"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter" csv:"186"`
	ConditionRecord            *int     `json:"ConditionRecord" csv:"187"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber" csv:"188"`
	ConditionType              *string  `json:"ConditionType" csv:"189"`
	PricingDate                *string  `json:"PricingDate" csv:"190"`
	ConditionRateValue         *float32 `json:"ConditionRateValue" csv:"191"`
	ConditionCurrency          *string  `json:"ConditionCurrency" csv:"192"`
	ConditionQuantity          *float32 `json:"ConditionQuantity" csv:"193"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit" csv:"194"`
	TaxCode                    *string  `json:"TaxCode" csv:"195"`
	ConditionAmount            *float32 `json:"ConditionAmount" csv:"196"`
	TransactionCurrency        *string  `json:"TransactionCurrency" csv:"197"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged" csv:"198"`
}

type OrdersItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID" csv:"1"`
	OrderItem                                       int      `json:"OrderItem" csv:"66"`
	ScheduleLine                                    int      `json:"ScheduleLine" csv:"199"`
	SupplyChainRelationshipID                       *int     `json:"SupplyChainRelationshipID" csv:"200"`
	SupplyChainRelationshipStockConfPlantID         *int     `json:"SupplyChainRelationshipStockConfPlantID" csv:"201"`
	Product                                         *string  `json:"Product" csv:"202"`
	StockConfirmationBussinessPartner               *int     `json:"StockConfirmationBussinessPartner" csv:"203"`
	StockConfirmationPlant                          *string  `json:"StockConfirmationPlant" csv:"204"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone" csv:"205"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch" csv:"206"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate" csv:"207"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate" csv:"208"`
	RequestedDeliveryDate                           *string  `json:"RequestedDeliveryDate" csv:"209"`
	RequestedDeliveryTime                           *string  `json:"RequestedDeliveryTime" csv:"210"`
	ConfirmedDeliveryDate                           *string  `json:"ConfirmedDeliveryDate" csv:"211"`
	ScheduleLineOrderQuantity                       *float32 `json:"ScheduleLineOrderQuantity" csv:"212"`
	OriginalOrderQuantityInBaseUnit                 *float32 `json:"OriginalOrderQuantityInBaseUnit" csv:"213"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit" csv:"214"`
	ConfirmedOrderQuantityByPDTAvailCheck           *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheck" csv:"215"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit" csv:"216"`
	UndeliveredQuantityInBaseUnit                   *float32 `json:"UndeliveredQuantityInBaseUnit" csv:"217"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit" csv:"218"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed" csv:"219"`
	PlusMinusFlag                                   *string  `json:"PlusMinusFlag" csv:"220"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus" csv:"221"`
	IsCancelled                                     *bool    `json:"IsCancelled" csv:"222"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion" csv:"223"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID" csv:"1"`
	PartnerFunction         *string `json:"PartnerFunction" csv:"224"`
	BusinessPartner         *int    `json:"BusinessPartner" csv:"225"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName" csv:"226"`
	BusinessPartnerName     *string `json:"BusinessPartnerName" csv:"227"`
	Organization            *string `json:"Organization" csv:"228"`
	Country                 *string `json:"Country" csv:"229"`
	Language                *string `json:"Language" csv:"230"`
	Currency                *string `json:"Currency" csv:"231"`
	ExternalDocumentID      *string `json:"ExternalDocumentID" csv:"232"`
	AddressID               *int    `json:"AddressID" csv:"233"`
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
