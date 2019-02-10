-- [schema.sql]
-- PostgreSQL 11: initial database schema for gogogadget.live.

-- Drop, create, and use a database named "gogogadget."
DROP DATABASE IF EXISTS "gogogadget";
CREATE DATABASE "gogogadget";
USE "gogogadget";

-- Create a special user for the database and grant all privileges.
-- This is a special user created ONLY to connect to this DB, in this specific Go application.
CREATE USER "gorm" WITH ENCRYPTED PASSWORD 'gorm';
GRANT ALL PRIVILEGES ON DATABASE "gogogadget" to "gorm";
