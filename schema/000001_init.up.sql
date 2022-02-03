CREATE TABLE users ( 
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
    telegram_id INTEGER NOT NULL UNIQUE, 
    name VARCHAR(255) DEFAULT "", 
    number VARCHAR(255) DEFAULT "", 
    dialogue_status VARCHAR(255) DEFAULT "pre_registration" 
);

CREATE TABLE couriers (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
    telegram_id INTEGER NOT NULL UNIQUE, 
    name VARCHAR(255) DEFAULT "", 
    number VARCHAR(255) DEFAULT "", 
    dialogue_status VARCHAR(255) DEFAULT "pre_registration" 
);

CREATE TABLE products ( 
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
    title VARCHAR(255) DEFAULT "", 
    price_kg INTEGER DEFAULT 0,
    price_bag INTEGER DEFAULT 0,
    description TEXT DEFAULT "",
    image_name TEXT DEFAULT ""
);    
      
CREATE TABLE shoping_cart ( 
    order_id INTEGER REFERENCES orders (id) ON DELETE CASCADE NOT NULL, 
    product_id INTEGER REFERENCES products (id) ON DELETE CASCADE NOT NULL,
    price INTEGER DEFAULT 0,
    unit_price INTEGER DEFAULT 0,
    delivery_format VARCHAR(5) DEFAULT "",
    quantity INTEGER DEFAULT 1 
);    
      
CREATE TABLE orders (
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	user_id INTEGER REFERENCES users (id),
 	user_name VARCHAR(255) DEFAULT "",
 	user_number VARCHAR(255) DEFAULT "",
 	delivery_adress VARCHAR(255) DEFAULT "",
 	order_date datetime DEFAULT '',
    order_status VARCHAR(255) DEFAULT "in_progress"
);

CREATE TABLE couriers_orders (
    courier_id INTEGER REFERENCES couriers (id) ON DELETE CASCADE NOT NULL,
    order_id INTEGER REFERENCES orders (id) ON DELETE CASCADE NOT NULL,  
    status VARCHAR(255) DEFAULT "active"
);