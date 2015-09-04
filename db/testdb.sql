DROP TABLE IF EXISTS users;
CREATE TABLE users 
(
/*id uniqueidentifier,*/
first_name varchar(255) NOT NULL,
last_name varchar(255) NOT NULL,
email varchar(255) NOT NULL UNIQUE,
username varchar(255) NOT NULL UNIQUE,
password varchar(255) NOT NULL

/*LastLoginDate int,*/
);

INSERT INTO users (first_name, last_name, email, username, password)
VALUES ('David', 'Vydra', 'david@vydra.net', 'dvydra', 'password');

INSERT INTO users (first_name, last_name, email, username, password)
VALUES ('Maxim', 'Veligan', 'maximveligan@gmail.com', 'grodion', 'password');

INSERT INTO users (first_name, last_name, email, username, password)
VALUES ('Shawn', 'Shacterman', 'shawnshact@msea.us', 'doozleberry', 'password');