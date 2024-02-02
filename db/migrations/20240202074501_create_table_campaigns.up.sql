
create table campaigns (
        id integer primary key generated always as identity, name varchar(50) NOT NULL, short_description text, description text, perks varchar(50), backer_count INTEGER DEFAULT 0 , goal_amount INTEGER DEFAULT 0 , current_amount INTEGER DEFAULT 0 , slug VARCHAR(50), user_id INTEGER  ,created_at TIMESTAMP default Now(), updated_at TIMESTAMP DEFAULT now(),     
        CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
);