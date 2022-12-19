create table books(
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) not null,
    last_date_read DATE not null,
    chapter INTEGER not null,
    page INTEGER not null
);