CREATE TABLE users (
    id            VARCHAR(255)    NOT NULL  PRIMARY KEY,
    mail_address  VARCHAR(255)    NOT NULL,
    password      VARCHAR(50)     NOT NULL,
    unique (mail_address)
) DEFAULT CHARSET=utf8mb4 ENGINE=InnoDB;
