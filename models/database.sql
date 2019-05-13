BEGIN TRANSACTION;
PRAGMA foreign_keys=ON;
CREATE TABLE detectives (
    id text primary key,
    firstname text,
    lastname text,
    city text,
    postalcode text,
    phone text,
    device text,
    position integer,
    FOREIGN KEY (position) REFERENCES positions(position) ON DELETE CASCADE
);

CREATE TABLE requests (
    id integer primary key,
    detective_id text,
    created_at text,
    info text,
    curriculum text,
    req_status integer,
    FOREIGN KEY (detective_id) REFERENCES detectives(id) ON DELETE CASCADE
);

CREATE TABLE positions (
    position integer primary key,
    price integer 
);

COMMIT;
