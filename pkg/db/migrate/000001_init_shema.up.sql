CREATE TABLE students (
                                        id SERIAL PRIMARY KEY,
                                        student_id TEXT NOT NULL UNIQUE,
                                        first_name TEXT NOT NULL,
                                        last_name TEXT NOT NULL,
                                        age BIGINT,
                                        email TEXT NOT NULL UNIQUE,
                                        gender TEXT,
                                        favourite_color TEXT,
                                        student_address TEXT NOT NULL DEFAULT '{}',
                                        created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                                        updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
                                        deleted BOOLEAN DEFAULT FALSE
);
CREATE TABLE address (
                                       address_id TEXT NOT NULL,
                                       street TEXT,
                                       city TEXT,
                                       planet TEXT,
                                       phone TEXT NOT NULL
);

