package database

import (
	"ServerWithNatsL0/model"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"database/sql"
)

func SaveDataInDB(db *sql.DB, od model.OrdersData) error {
	//Save the OrdersData structure in DB
	err := SaveOrdersDataInDB(db, od)
	if err != nil {
		return errors.Wrap(err, "func SaveOrdersDataInDB")
	}
	//Save the Delivery structure in DB
	err = SaveDeliveryInDB(db, od)
	if err != nil {
		return errors.Wrap(err, "func SaveDeliveryInDB")
	}
	//Save the Payment structure in DB
	err = SavePaymentInDB(db, od)
	if err != nil {
		return errors.Wrap(err, "func SavePaymentInDb")
	}
	////Save the Item structures in DB
	err = SaveItemsInDB(db, od)
	if err != nil {
		return errors.Wrap(err, "func SaveItemsInDb")
	}

	return nil
}

func SaveOrdersDataInDB(db *sql.DB, od model.OrdersData) error {
	//Request to the DB to add data
	_, err := db.Exec(`INSERT INTO orders_data (
				order_uid, track_number, entry, 
				locale, internal_signature, 
				customer_id, delivery_service, 
				shardkey, sm_id, 
				date_created, oof_shard
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		od.OrderUid, od.TrackNumber, od.Entry,
		od.Locale, od.InternalSignature,
		od.CustomerId, od.DeliveryService,
		od.Shardkey, od.SmId,
		od.DateCreated, od.OofShard,
	)
	if err != nil {
		return err
	}

	return nil
}

func SaveDeliveryInDB(db *sql.DB, od model.OrdersData) error {
	//Request to the DB to add data
	_, err := db.Exec(`INSERT INTO delivery (
                                order_uid, name, phone, zip,
                                city, address, region, email
                        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		od.OrderUid, od.Delivery.Name, od.Delivery.Phone, od.Delivery.Zip,
		od.Delivery.City, od.Delivery.Address, od.Delivery.Region, od.Delivery.Email,
	)
	if err != nil {
		return err
	}

	return nil
}

func SavePaymentInDB(db *sql.DB, od model.OrdersData) error {
	//Request to the DB to add data
	_, err := db.Exec(`INSERT INTO payment (
                                order_uid, transaction, request_id, 
				currency, provider, amount, payment_dt,
				bank, delivery_cost, goods_total, custom_fee
                        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		od.OrderUid, od.Payment.Transaction, od.Payment.RequestId, od.Payment.Currency,
		od.Payment.Provider, od.Payment.Amount, od.Payment.PaymentDt, od.Payment.Bank,
		od.Payment.DeliveryCost, od.Payment.GoodsTotal, od.Payment.CustomFee,
	)
	if err != nil {
		return err
	}

	return nil
}

func SaveItemsInDB(db *sql.DB, od model.OrdersData) error {
	//Iterate over the slice Items and add the Item structure to the DB
	for _, Item := range od.Items {
		//Request to the DB to add data
		_, err := db.Exec(`INSERT INTO item (
					order_uid, chrt_id, track_number,
					price, rid, name, sale, size,
					total_price, nm_id, brand, status
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
			od.OrderUid, Item.ChrtId, Item.TrackNumber, Item.Price,
			Item.Rid, Item.Name, Item.Sale, Item.Size, Item.TotalPrice,
			Item.NmId, Item.Brand, Item.Status,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
