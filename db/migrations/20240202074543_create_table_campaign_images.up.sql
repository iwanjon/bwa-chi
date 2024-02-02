


create table campaign_images (
        id integer primary key generated always as identity, file_name varchar(50) NOT NULL, is_primary INTEGER DEFAULT 0 , campaign_id INTEGER  ,created_at TIMESTAMP default Now(), updated_at TIMESTAMP DEFAULT now(),     
        CONSTRAINT fk_campaign
        FOREIGN KEY(campaign_id) 
        REFERENCES campaigns(id)
);