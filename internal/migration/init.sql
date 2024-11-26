CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255),
    song_name VARCHAR(255),
    lyrics TEXT,
    link VARCHAR(255)
);
