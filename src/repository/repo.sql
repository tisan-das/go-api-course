
------- Schema creation for book service -------
create database book_service;
create user tisan with encrypted password 'tisan';
grant all privileges on database book_service to tisan;
create schema if not exists book_service;


------- Create another user to check rotation of DB credential -------
create user tisan2 with encrypted password 'tisan2';
grant all privileges on database book_service to tisan2;
grant all privileges on schema book_service to tisan2;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA book_service TO tisan2;

alter user tisan with encrypted password 'das';


------- Remove the extra user created to verify DB credential rotation -------
revoke all privileges on all tables in schema book_service from tisan2;
revoke all privileges on schema book_service from tisan2;
revoke all privileges on database book_service from tisan2;
drop user tisan2;
