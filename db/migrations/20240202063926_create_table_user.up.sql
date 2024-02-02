create table users (
    id integer primary key generated always as identity, name varchar(50) NOT NULL, Occupation varchar(50), email varCHAR(50) not null, password_hash varchar(256) not null, avatar_file_name varchar(256), role varchar(50), token varchar(256), created_at TIMESTAMP default Now(), updated_at TIMESTAMP DEFAULT now() 
);





