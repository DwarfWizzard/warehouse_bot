CREATE TABLE users ( 
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
    telegram_id INTEGER NOT NULL UNIQUE, name VARCHAR(255) DEFAULT "", 
    number VARCHAR(255) DEFAULT "", 
    dialogue_status VARCHAR(255) DEFAULT "pre_registration" 
);    
      
CREATE TABLE products ( 
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
    title VARCHAR(255) DEFAULT "", 
    price INTEGER DEFAULT 0, 
    description TEXT DEFAULT "" 
);    
      
CREATE TABLE shoping_cart ( 
    order_id INTEGER REFERENCES orders (id) NOT NULL, 
    product_id INTEGER REFERENCES products (id) NOT NULL,
    price INTEGER DEFAULT 0,
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