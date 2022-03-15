package database

import (
	"ServerWithNatsL0/model"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"database/sql"
)

func GetOrderUidFromDB(db *sql.DB) ([]string, error) {

	var OrderUid string
	var SliceOrderUid []string
	//Executing a query to the DB
	rows, err := db.Query("SELECT order_uid FROM orders_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Iterate over the returned rows using rows.Next.
	for rows.Next() {
		err := rows.Scan(&OrderUid) //Copy the return value to the variable OrderUid
		if err != nil {
			return nil, err
		}
		//Filling the slice with the received values
		SliceOrderUid = append(SliceOrderUid, OrderUid)
	}

	return SliceOrderUid, nil

}

func GetOrdersDataFromDB(db *sql.DB, OrderUid string) (*model.OrdersData, error) {

	OrdersData := model.OrdersData{}
	//Executing a query to the DB
	err := db.QueryRow(`SELECT * FROM orders_data WHERE order_uid=$1`, OrderUid).Scan(&OrdersData.OrderUid,
		&OrdersData.TrackNumber, &OrdersData.Entry, &OrdersData.Locale, &OrdersData.InternalSignature,
		&OrdersData.CustomerId, &OrdersData.DeliveryService, &OrdersData.Shardkey, &OrdersData.SmId,
		&OrdersData.DateCreated, &OrdersData.OofShard,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows: //no rows in result set
			return nil, errors.Wrap(err, "func GetOrdersDataFromDB, ROWS NOT FOUND")
		default:
			return nil, errors.Wrap(err, "func GetOrdersDataFromDB")
		}
	}
	return &OrdersData, nil
}

func GetDeliveryFromDB(db *sql.DB, OrderUid string) (*model.Delivery, error) {

	Delivery := model.Delivery{}
	//Executing a query to the DB
	err := db.QueryRow(`SELECT name, phone, zip, city, address, region, email 
				FROM delivery 
				WHERE order_uid=$1`, OrderUid).Scan(&Delivery.Name,
		&Delivery.Phone, &Delivery.Zip, &Delivery.City, &Delivery.Address,
		&Delivery.Region, &Delivery.Email,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows: //no rows in result set
			return nil, errors.Wrap(err, "func GetDeliveryFromDB, ROWS NOT FOUND")
		default:
			return nil, errors.Wrap(err, "func GetDeliveryFromDB")
		}
	}
	return &Delivery, nil
}

func GetPaymentFromDB(db *sql.DB, OrderUid string) (*model.Payment, error) {

	Payment := model.Payment{}
	//Executing a query to the DB
	err := db.QueryRow(`SELECT transaction, request_id, currency, 
				provider, amount, payment_dt, bank,
				delivery_cost, goods_total, custom_fee
				FROM payment 
				WHERE order_uid=$1`, OrderUid).Scan(&Payment.Transaction,
		&Payment.RequestId, &Payment.Currency, &Payment.Provider,
		&Payment.Amount, &Payment.PaymentDt, &Payment.Bank,
		&Payment.DeliveryCost, &Payment.GoodsTotal, &Payment.CustomFee,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows: //no rows in result set
			return nil, errors.Wrap(err, "func GetPaymentFromDB ROWS NOT FOUND")
		default:
			return nil, errors.Wrap(err, "func GetPaymentFromDB")
		}
	}
	return &Payment, nil
}

func GetItemsFromDB(db *sql.DB, OrderUid string) ([]model.Item, error) {

	var Item model.Item

	var Items []model.Item
	//Executing a query to the DB
	rows, err := db.Query(`SELECT chrt_id, track_number, price,
				rid, name, sale, size, total_price,
				nm_id, brand, status 
				FROM item 
				WHERE order_uid=$1`, OrderUid)
	if err != nil {
		return nil, errors.Wrap(err, "func GetItemsFromDB")
	}
	defer rows.Close()
	//Iterate over the returned rows using rows.Next.
	for rows.Next() {
		//Copy the return value to the variable OrderUid
		err := rows.Scan(&Item.ChrtId, &Item.TrackNumber, &Item.Price,
			&Item.Rid, &Item.Name, &Item.Sale, &Item.Size, &Item.TotalPrice,
			&Item.NmId, &Item.Brand, &Item.Status,
		)
		if err != nil {
			return nil, errors.Wrap(err, "func GetItemsFromDB")
		}
		//Filling the slice with the received values
		Items = append(Items, Item)
	}

	return Items, nil
}
