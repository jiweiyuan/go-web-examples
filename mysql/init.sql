DROP TABLE IF EXISTS album;

CREATE TABLE album(
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(100) NOT NULL,
    artist VARCHAR(100) NOT NULL,
    price  DECIMAL(5, 2) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO album
(title, artist, price)
VALUES
( "万马奔腾", "徐悲鸿", 89.9),
("sky", "vongao", 23.89)


