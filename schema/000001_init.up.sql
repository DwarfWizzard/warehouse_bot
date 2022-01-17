CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	telegram_id INTEGER NOT NULL UNIQUE,
	name VARCHAR(255) DEFAULT "",
	number VARCHAR(255) DEFAULT "",
    dialogue_status VARCHAR(255) DEFAULT "pre_registration"
);

CREATE TABLE products (
	id INTEGER PRIMARY KEY AUTOINCREMENT  NOT NULL,
	title VARCHAR(255) DEFAULT "",
	price VARCHAR(255) DEFAULT "",
	description TEXT DEFAULT ""
);

CREATE TABLE shoping_cart ( 
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
    user_id INTEGER REFERENCES users(id) NOT NULL, 
    adress TEXT DEFAULT "", 
    delivery_date TEXT DEFAULT ""
);

CREATE TABLE product_lists ( 
    cart_id INTEGER REFERENCES shoping_cart (id) NOT NULL, 
    product_id INTEGER REFERENCES products (id) NOT NULL, 
    quantity INTEGER DEFAULT 1 
);