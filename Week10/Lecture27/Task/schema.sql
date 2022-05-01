CREATE TABLE stories (
    story_id int not null primary key,
    story_title text not null,
    story_score int not null, 
    created_at datetime not null
);