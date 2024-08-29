#!/usr/bin/env bash

psql -U postgres -c 'DROP DATABASE IF EXISTS "scti-db";'

# Create the database
psql -U postgres -c 'CREATE DATABASE "scti-db";'

# Run migrations or create tables
psql -U postgres -d scti-db -f migrations.sql

# Populate the database
psql -U postgres -d scti-db -f populate.sql

# Create your account
echo "Qual o seu email: "
read -r Email
Email=${Email:-user@test.com}
cmd="INSERT INTO users (email, name, uuid, verificationCode, isVerified, isAdmin, isPaid) VALUES ('$Email', 'Voce', '623e4567-e89b-12d3-a456-426614174000', '623e4', TRUE, TRUE, TRUE);"

psql -U postgres -d scti-db -c "$cmd"

pass="INSERT INTO passwd (id, passwd) SELECT id, CONCAT('senha_segura_', id) FROM users WHERE email = '$Email';"

psql -U postgres -d scti-db -c "$pass"
