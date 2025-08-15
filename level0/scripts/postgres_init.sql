CREATE TABLE Orders (
  order_uid VARCHAR(50) PRIMARY KEY,
  track_number VARCHAR(50) NOT NULL,
  entry VARCHAR(10) NOT NULL,
  locale VARCHAR(10) NOT NULL,
  customer_id VARCHAR(50) NOT NULL,
  delivery_service VARCHAR(50) NOT NULL,
  shardkey VARCHAR(10) NOT NULL,
  sm_id INTEGER NOT NULL,
  date_created TIMESTAMP WITH TIME ZONE NOT NULL,
  oof_shard VARCHAR(10) NOT NULL,
  internal_signature VARCHAR(255)
);

CREATE TABLE Delivery (
  order_uid VARCHAR(50) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  zip VARCHAR(20) NOT NULL,
  city VARCHAR(100) NOT NULL,
  address VARCHAR(255) NOT NULL,
  region VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  FOREIGN KEY (order_uid) REFERENCES Orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE Payment (
  transaction VARCHAR(50) PRIMARY KEY,
  order_uid VARCHAR(50) NOT NULL,
  request_id VARCHAR(50),
  currency VARCHAR(10) NOT NULL,
  provider VARCHAR(50) NOT NULL,
  amount INTEGER NOT NULL,
  payment_dt BIGINT NOT NULL,
  bank VARCHAR(50) NOT NULL,
  delivery_cost INTEGER NOT NULL,
  goods_total INTEGER NOT NULL,
  custom_fee INTEGER NOT NULL,
  FOREIGN KEY (order_uid) REFERENCES Orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE Items (
  chrt_id INTEGER PRIMARY KEY,
  order_uid VARCHAR(50) NOT NULL,
  track_number VARCHAR(50) NOT NULL,
  price INTEGER NOT NULL,
  rid VARCHAR(50) NOT NULL,
  name VARCHAR(255) NOT NULL,
  sale INTEGER NOT NULL,
  size VARCHAR(10) NOT NULL,
  total_price INTEGER NOT NULL,
  nm_id INTEGER NOT NULL,
  brand VARCHAR(100) NOT NULL,
  status INTEGER NOT NULL,
  FOREIGN KEY (order_uid) REFERENCES Orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE test (
  a VARCHAR(50),
  b INTEGER
);