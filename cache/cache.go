package cache

import (
	"ServerWithNatsL0/database"
	"ServerWithNatsL0/model"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"database/sql"
)

func CacheRecovery(db *sql.DB, CacheMap map[string]model.OrdersData) error {
	//Getting the OrderUid slice from DB
	SliceOrderUid, err := database.GetOrderUidFromDB(db)
	if err != nil {
		return errors.Wrap(err, "cache, func database.GetOrderUidFromDB")
	}
	//Iterate the array and get the OrderUid for calling functions...
	for _, OrderUid := range SliceOrderUid {
		//Getting data for ggg struct from DB
		OrdersData, err := database.GetOrdersDataFromDB(db, OrderUid)
		if err != nil {
			return errors.Wrap(err, "cache, func database.GetOrdersDataFromDB")
		}
		//Getting data for Delivery struct OrdersData struct from DB
		Delivery, err := database.GetDeliveryFromDB(db, OrderUid)
		if err != nil {
			return errors.Wrap(err, "cache, func database.GetDeliveryFromDB")
		}
		//Getting data for Payment struct in OrdersData struct from DB
		Payment, err := database.GetPaymentFromDB(db, OrderUid)
		if err != nil {
			return errors.Wrap(err, "cache, func database.GetPaymentFromDB")
		}
		//Getting data for Items slice in OrdersData struct from DB
		OrdersData.Items, err = database.GetItemsFromDB(db, OrderUid)
		if err != nil {
			return errors.Wrap(err, "cache, func database.GetItemsFromDB")
		}
		//Combining data structures
		OrdersData.Delivery = *Delivery
		OrdersData.Payment = *Payment
		//Save the data structure, using the OrderUid as the key for the CacheMap
		CacheMap[OrderUid] = *OrdersData
	}

	return nil
}
