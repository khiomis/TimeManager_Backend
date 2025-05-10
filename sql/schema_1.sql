-- Already executed: false

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE TBL_USERS
(
    ID_USER               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    DT_CREATED_AT         timestamptz  not null default current_timestamp,
    DT_UPDATED_AT         timestamptz  not null,
    NM_USER               varchar(128) not null,
    DS_EMAIL              varchar(128) not null,
    DS_PASSWORD           varchar(128) not null,
    TP_STATUS             integer      not null default 0,
    DS_TEMPORARY_PASSWORD varchar(128) not null default ''
);

CREATE TABLE TBL_VALIDATION_TOKENS
(
    ID_VALIDATION_TOKEN BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    DT_CREATED_AT       timestamptz                           not null default current_timestamp,
    DT_UPDATED_AT       timestamptz                           not null,
    CD_VALIDATION_TOKEN varchar(10)                           not null,
    DT_EXPIRE_AT        timestamptz                           NOT NULL,
    TP_VALIDATION_TOKEN INT2                                  NOT NULL,
    ID_USER             BIGINT REFERENCES TBL_USERS (ID_USER) NOT NULL
);

CREATE TABLE TBL_SESSIONS
(
    ID_SESSION    UUID PRIMARY KEY                               DEFAULT uuid_generate_v4(),
    DT_CREATED_AT timestamptz                           not null default current_timestamp,
    ID_USER       BIGINT REFERENCES TBL_USERS (ID_USER) NOT NULL,
    expires_at    TIMESTAMPTZ
)