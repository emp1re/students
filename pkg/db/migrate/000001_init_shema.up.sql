CREATE TABLE students (
                                        id SERIAL PRIMARY KEY,
                                        student_id BIGINT NOT NULL UNIQUE,
                                        first_name TEXT NOT NULL,
                                        last_name TEXT NOT NULL,
                                        age SERIAL,
                                        email TEXT NOT NULL UNIQUE,
                                        gender TEXT,
                                        favourite_color TEXT,
                                        student_address BIGINT NOT NULL,
                                        created_at BIGINT NOT NULL DEFAULT 0,
                                        updated_at BIGINT NOT NULL DEFAULT 0,
                                        deleted BOOLEAN DEFAULT FALSE
);
CREATE TABLE address (
                                       address_id BIGINT NOT NULL,
                                       street TEXT,
                                       city TEXT,
                                       planet TEXT,
                                       phone TEXT NOT NULL
);
CREATE TABLE index (
                       index_id BIGINT NOT NULL

);
