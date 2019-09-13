-- database
  drop database if exists noonde_api;
  create database         noonde_api charset utf8mb4 collate utf8mb4_bin;

-- user writer
create user if not exists 'noonde_w'@'localhost' identified by 'noonde_dev';
grant all privileges on noonde_api.* to 'noonde_w'@'localhost';


-- user reader
create user if not exists 'noonde_r'@'localhost' identified by 'noonde_dev';
grant all privileges on noonde_api.* to 'noonde_r'@'localhost';