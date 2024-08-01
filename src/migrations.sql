CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE passwd (
    id INT PRIMARY KEY REFERENCES users(id),
    passwd VARCHAR(255) NOT NULL
);

CREATE TABLE courses (
    course_name VARCHAR(255) UNIQUE NOT NULL,
    available_spots INT NOT NULL
);
