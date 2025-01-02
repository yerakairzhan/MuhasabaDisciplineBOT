-- Create the users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,              -- Unique identifier for each user
                       user_id bigint NOT NULL UNIQUE,        -- Telegram user ID (must be unique)
                       username VARCHAR(100) NOT NULL,     -- Telegram username
                       created_at TIMESTAMP DEFAULT NOW()  -- Timestamp of user creation
);

-- Create the fajr table
CREATE TABLE fajr (
                      id SERIAL PRIMARY KEY,               -- Unique identifier for each record
                      user_id bigint NOT NULL UNIQUE,        -- Telegram user ID (must be unique)
                      done BOOLEAN DEFAULT FALSE,          -- Task completion status
                      created_at TIMESTAMP DEFAULT NOW(),  -- Timestamp of task creation
                      CONSTRAINT fk_user_fajr
                          FOREIGN KEY (user_id)
                              REFERENCES users(user_id)        -- Link to users.user_id
                              ON DELETE CASCADE,               -- Delete tasks if the user is deleted
                      CONSTRAINT unique_user_fajr UNIQUE (user_id) -- Ensure one record per user
);

-- Create the duha table
CREATE TABLE duha (
                      id SERIAL PRIMARY KEY,               -- Unique identifier for each record
                      user_id bigint NOT NULL UNIQUE,        -- Telegram user ID (must be unique)
                      done BOOLEAN DEFAULT FALSE,          -- Task completion status
                      created_at TIMESTAMP DEFAULT NOW(),  -- Timestamp of task creation
                      CONSTRAINT fk_user_duha
                          FOREIGN KEY (user_id)
                              REFERENCES users(user_id)        -- Link to users.user_id
                              ON DELETE CASCADE,               -- Delete tasks if the user is deleted
                      CONSTRAINT unique_user_duha UNIQUE (user_id) -- Ensure one record per user
);

-- Create the tafsir table
CREATE TABLE tafsir (
                        id SERIAL PRIMARY KEY,               -- Unique identifier for each record
                        user_id bigint NOT NULL UNIQUE,        -- Telegram user ID (must be unique)
                        done BOOLEAN DEFAULT FALSE,          -- Task completion status
                        created_at TIMESTAMP DEFAULT NOW(),  -- Timestamp of task creation
                        CONSTRAINT fk_user_tafsir
                            FOREIGN KEY (user_id)
                                REFERENCES users(user_id)        -- Link to users.user_id
                                ON DELETE CASCADE,               -- Delete tasks if the user is deleted
                        CONSTRAINT unique_user_tafsir UNIQUE (user_id) -- Ensure one record per user
);
