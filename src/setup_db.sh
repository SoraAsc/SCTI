#!/usr/bin/env bash

# Create the database
psql -U postgres -c 'CREATE DATABASE "scti-db";'

# Run migrations or create tables
psql -U postgres -d scti-db -f migrations.sql
