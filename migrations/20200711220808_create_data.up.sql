CREATE TABLE data (
    id bigserial not null primary key,
    country varchar not null unique,
    cases bigserial not null,
    deaths bigserial not null,
    recovered bigserial not null
);