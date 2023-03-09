package dpfm_api_output_formatter

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
	"golang.org/x/xerrors"
)

func ConvertToConcatMessage(filePath string) (*[]DataConcatenation, error) {
	data := make([]DataConcatenation, 0)
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return nil, xerrors.Errorf("file open error: %w", err)
	}
	// defer f.Close()

	allHeaders := make([]OrdersHeader, 0)
	allItems := make([]OrdersItem, 0)
	allItemPricingElements := make([]OrdersItemPricingElement, 0)
	allItemScheduleLines := make([]OrdersItemScheduleLine, 0)
	allAddresses := make([]OrdersAddress, 0)
	allPartners := make([]OrdersPartner, 0)

	err = gocsv.UnmarshalFile(f, &allHeaders)
	if err != nil {
		return nil, xerrors.Errorf("read order header error: %w", err)
	}
	f.Close()
	f, _ = os.OpenFile(filePath, os.O_RDONLY, 0)
	err = gocsv.UnmarshalFile(f, &allItems)
	if err != nil {
		return nil, xerrors.Errorf("read order items error: %w", err)
	}
	f.Seek(0, 0)
	err = gocsv.UnmarshalFile(f, &allItemPricingElements)
	if err != nil {
		return nil, xerrors.Errorf("read order itemPricingElements error: %w", err)
	}
	f.Seek(0, 0)
	err = gocsv.UnmarshalFile(f, &allItemScheduleLines)
	if err != nil {
		return nil, xerrors.Errorf("read order itemScheduleLines error: %w", err)
	}
	f.Seek(0, 0)
	err = gocsv.UnmarshalFile(f, &allAddresses)
	if err != nil {
		return nil, xerrors.Errorf("read order addresses error: %w", err)
	}
	f.Seek(0, 0)
	err = gocsv.UnmarshalFile(f, &allPartners)
	if err != nil {
		return nil, xerrors.Errorf("read order partners error: %w", err)
	}

	allHeaders = toNull(allHeaders)
	allItems = toNull(allItems)
	allItemPricingElements = toNull(allItemPricingElements)
	allItemScheduleLines = toNull(allItemScheduleLines)
	allAddresses = toNull(allAddresses)
	allPartners = toNull(allPartners)
	allHeaders = getUniqueHeaders(&allHeaders)

	for headerIdx, v := range allHeaders {
		oID := v.OrderID
		items := getUniqueItem(&allItems, oID)
		if len(items) == 0 {
			return nil, xerrors.Errorf("items length is 0")
		}
		itemIDs := getItemIDs(items, oID)
		itemPricingElements := getUniquePricing(&allItemPricingElements, oID, itemIDs)
		itemScheduleLines := getUniqueSchedules(&allItemScheduleLines, oID, itemIDs)
		addresses := getUniqueAddresses(&allAddresses, oID, itemIDs)
		partners := getUniquePartners(&allPartners, oID)
		orderAll := DataConcatenation{
			Header:             allHeaders[headerIdx],
			Item:               items,
			ItemPricingElement: itemPricingElements,
			ItemScheduleLine:   itemScheduleLines,
			Address:            addresses,
			Partner:            partners,
		}
		data = append(data, orderAll)
	}

	return &data, nil
}

func getOrderIDs(headers []OrdersHeader) []int {
	idCheck := make(map[int]struct{})
	for _, v := range headers {
		idCheck[v.OrderID] = struct{}{}
	}
	ids := make([]int, 0)
	for k := range idCheck {
		ids = append(ids, k)
	}
	return ids
}

func getUniqueHeaders(headers *[]OrdersHeader) []OrdersHeader {
	idCheck := make(map[int]struct{})
	unique := make([]OrdersHeader, 0, len(*headers))
	for i, v := range *headers {
		_, ok := idCheck[v.OrderID]
		if ok {
			continue
		}
		idCheck[v.OrderID] = struct{}{}
		unique = append(unique, (*headers)[i])
	}
	return unique
}

func getUniqueItem(items *[]OrdersItem, oID int) []OrdersItem {
	itemIDs := make([]int, 0)
	unique := make([]OrdersItem, 0, len(*items))
	for i, v := range *items {
		if v.OrderID != oID || isContain(itemIDs, v.OrderItem) {
			continue
		}
		itemIDs = append(itemIDs, v.OrderItem)
		unique = append(unique, (*items)[i])
	}
	return unique
}

func getUniquePricing(pricings *[]OrdersItemPricingElement, oID int, itemIDs []int) []OrdersItemPricingElement {
	unique := make([]OrdersItemPricingElement, 0, len(*pricings))

	for _, oiID := range itemIDs {
		counters := make([]int, 0)
		for i, v := range *pricings {
			if v.OrderID != oID || oiID != v.OrderItem {
				continue
			}
			if isContain(counters, v.PricingProcedureCounter) {
				continue
			}
			counters = append(counters, v.PricingProcedureCounter)
			unique = append(unique, (*pricings)[i])
		}
	}
	return unique
}

func getUniqueSchedules(schedules *[]OrdersItemScheduleLine, oID int, itemIDs []int) []OrdersItemScheduleLine {
	unique := make([]OrdersItemScheduleLine, 0, len(*schedules))

	for _, oiID := range itemIDs {
		schedulelines := make([]int, 0)
		for i, v := range *schedules {
			if v.OrderID != oID || !isContain(itemIDs, v.OrderItem) || v.OrderItem != oiID {
				continue
			}

			if isContain(schedulelines, v.ScheduleLine) {
				continue
			}

			schedulelines = append(schedulelines, v.ScheduleLine)
			unique = append(unique, (*schedules)[i])
		}
	}
	return unique
}
func getUniqueAddresses(addresses *[]OrdersAddress, oID int, itemID []int) []OrdersAddress {
	addIDs := make([]int, 0)
	unique := make([]OrdersAddress, 0, len(*addresses))
	for i, v := range *addresses {
		if v.OrderID != oID || v.AddressID == nil || isContain(addIDs, *v.AddressID) {
			continue
		}
		addIDs = append(addIDs, *v.AddressID)
		unique = append(unique, (*addresses)[i])
	}
	return unique
}

func getUniquePartners(partners *[]OrdersPartner, oID int) []OrdersPartner {
	bpIDs := getBusinessPartnerIDs(partners, oID)
	unique := make([]OrdersPartner, 0, len(*partners))
	for _, bpID := range bpIDs {
		partnerFuncs := make([]string, 0)
		for i, v := range *partners {
			if v.OrderID != oID || v.PartnerFunction == nil || v.BusinessPartner == nil || *v.BusinessPartner != bpID {
				continue
			}

			if isContain(partnerFuncs, *v.PartnerFunction) {
				continue
			}
			partnerFuncs = append(partnerFuncs, *v.PartnerFunction)
			unique = append(unique, (*partners)[i])
		}
	}
	return unique
}

func getItemIDs(csv []OrdersItem, oID int) []int {
	idCheck := make(map[int]struct{})
	for _, v := range csv {
		if v.OrderID != oID {
			continue
		}
		idCheck[v.OrderItem] = struct{}{}
	}
	ids := make([]int, 0)
	for k := range idCheck {
		ids = append(ids, k)
	}
	return ids
}

func getBusinessPartnerIDs(partners *[]OrdersPartner, oID int) []int {
	idCheck := make(map[int]struct{})
	for _, v := range *partners {
		if v.OrderID != oID {
			continue
		}

		if v.BusinessPartner == nil {
			continue
		}
		idCheck[*v.BusinessPartner] = struct{}{}
	}
	ids := make([]int, 0)
	for k := range idCheck {
		ids = append(ids, k)
	}
	return ids
}

func isContain[T comparable](obj []T, targ T) bool {
	for _, v := range obj {
		if v == targ {
			return true
		}
	}
	return false
}

func toNull[T any](obj []T) []T {
	b, _ := json.Marshal(obj)
	j := make([]map[string]interface{}, 0)
	json.Unmarshal(b, &j)
	for i := range j {
		for k, v := range j[i] {
			switch typedV := v.(type) {
			case int:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case int32:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case int64:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case float32:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case float64:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case string:
				if isZero(typedV) {
					j[i][k] = nil
				}
			}
		}
	}
	b, _ = json.Marshal(j)
	json.Unmarshal(b, &obj)
	return obj
}
func isZero[T comparable](obj T) bool {
	var zero T
	return obj == zero
}
