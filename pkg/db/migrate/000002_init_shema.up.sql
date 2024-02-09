CREATE TABLE students (
                                        id SERIAL PRIMARY KEY,
                                        student_id INT NOT NULL UNIQUE,
                                        first_name TEXT NOT NULL,
                                        last_name TEXT NOT NULL,
                                        age BIGINT,
                                        email TEXT NOT NULL UNIQUE,
                                        gender TEXT,
                                        favourite_color TEXT,
                                        student_address INT NOT NULL,
                                        created_at timestamptz,
                                        updated_at timestamptz,
                                        deleted BOOLEAN DEFAULT FALSE
);
CREATE TABLE address (
                                       address_id INT NOT NULL,
                                       street TEXT,
                                       city TEXT,
                                       planet TEXT,
                                       phone TEXT NOT NULL
);
CREATE TABLE index (
                       index_id INT NOT NULL

);
