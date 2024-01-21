CREATE TABLE blog
(
    id               SERIAL PRIMARY KEY,
    title            VARCHAR(255) NOT NULL,
    content          VARCHAR(255) NOT NULL,
    author           VARCHAR(100) NOT NULL,
    publication_date DATE         NOT NULL,
    updatedAt        DATE,
    DeletedAt        Date

);
