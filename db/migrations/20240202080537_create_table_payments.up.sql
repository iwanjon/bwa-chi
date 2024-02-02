
create table payments (
    id integer primary key generated always as identity, name varchar(50) NOT NULl, amount INTEGER DEFAULT 0, created_at TIMESTAMP default Now(), updated_at TIMESTAMP DEFAULT now() 
);




