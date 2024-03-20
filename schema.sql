CREATE TABLE Players (
    id INTEGER NOT NULL AUTOINCREMENT UNIQUE,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    PRIMARY KEY (id, email)
);

CREATE TABLE Scores (
    id INTEGER PRIMARY KEY,
    score INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    FOREIGN KEY (player_id) REFERENCES Players(id)
);

CREATE TABLE Auth (
    id INTEGER PRIMARY KEY,
    hash TEXT NOT NULL,
    player_id INTEGER NOT NULL,
    player_email TEXT NOT NULL,
    FOREIGN KEY (player_id) REFERENCES Players(id)
    FOREIGN KEY (player_email) REFERENCES Players(email)
);
