CREATE TABLE orders_data (order_uid VARCHAR(50) UNIQUE PRIMARY KEY, track_number VARCHAR(50), entry VARCHAR(50), locale VARCHAR(50), internal_signature VARCHAR(50), customer_id VARCHAR(50), delivery_service VARCHAR(50), shardkey VARCHAR(50), sm_id INTEGER, date_created VARCHAR(50), oof_shard VARCHAR(50));

CREATE TABLE delivery (name VARCHAR(50), phone VARCHAR(50), zip VARCHAR(50), city VARCHAR(50), address VARCHAR(50), region VARCHAR(50), email VARCHAR(50), order_uid VARCHAR(50) UNIQUE PRIMARY KEY);

CREATE TABLE payment (transaction VARCHAR(50), request_id VARCHAR(50), currency VARCHAR(50), provider VARCHAR(50), amount INTEGER, payment_dt INTEGER, bank VARCHAR(50), delivery_cost INTEGER, goods_total INTEGER, custom_fee INTEGER, order_uid VARCHAR(50) UNIQUE PRIMARY KEY);

CREATE TABLE item (chrt_id INTEGER, track_number VARCHAR(50), price INTEGER, rid VARCHAR(50), name VARCHAR(50), sale INTEGER, size VARCHAR(50), total_price INTEGER, nm_id INTEGER, brand VARCHAR(50), status INTEGER, order_uid VARCHAR(50) UNIQUE PRIMARY KEY);
