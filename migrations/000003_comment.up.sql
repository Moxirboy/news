CREATE TABLE comment
(
    id      SERIAL PRIMARY KEY,
    blog_id int NOT NULL,
    body    VARCHAR(255)
);
