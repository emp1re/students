CREATE TABLE students (
                                        id SERIAL PRIMARY KEY,
                                        first_name TEXT NOT NULL,
                                        last_name TEXT NOT NULL,
                                        age BIGINT,
                                        email TEXT NOT NULL UNIQUE,
                                        gender TEXT,
                                        favourite_color TEXT,
                                        addresses VARCHAR,
                                        created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                                        updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                                        deleted BOOLEAN DEFAULT FALSE
);
CREATE TABLE address (
                                       id SERIAL PRIMARY KEY,
                                       street TEXT ,
                                       city TEXT,
                                       planet TEXT,
                                       phone TEXT NOT NULL
);
