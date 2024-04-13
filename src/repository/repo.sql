
create database book_service;
create user tisan with encrypted password 'tisan';
grant all privileges on database book_service to tisan;



create schema if not exists book_service;
