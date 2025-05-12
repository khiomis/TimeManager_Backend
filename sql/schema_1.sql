-- Already executed: false

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS TBL_USERS
(
    ID_USER               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ID_USER_UUID          UUID                  DEFAULT uuid_generate_v4(),
    DT_CREATED_AT         timestamptz  not null default current_timestamp,
    DT_UPDATED_AT         timestamptz  not null,
    NM_USER               varchar(128) not null,
    DS_EMAIL              varchar(128) not null,
    DS_PASSWORD           varchar(128) not null,
    TP_STATUS             integer      not null default 0,
    DS_TEMPORARY_PASSWORD varchar(128) not null default ''
);

CREATE TABLE IF NOT EXISTS TBL_VALIDATION_TOKENS
(
    ID_VALIDATION_TOKEN UUID PRIMARY KEY                               DEFAULT uuid_generate_v4(),
    DT_CREATED_AT       timestamptz                           not null default current_timestamp,
    CD_VALIDATION_TOKEN varchar(10)                           not null,
    DT_EXPIRE_AT        timestamptz                           NOT NULL,
    TP_VALIDATION_TOKEN INT2                                  NOT NULL,
    ID_USER             BIGINT REFERENCES TBL_USERS (ID_USER) NOT NULL
);

CREATE TABLE IF NOT EXISTS TBL_SESSIONS
(
    ID_SESSION    UUID PRIMARY KEY                               DEFAULT uuid_generate_v4(),
    DT_CREATED_AT timestamptz                           not null default current_timestamp,
    ID_USER       BIGINT REFERENCES TBL_USERS (ID_USER) NOT NULL,
    DT_EXPIRES_AT TIMESTAMPTZ                           NOT NULL
);

CREATE TABLE IF NOT EXISTS TBL_PROJECTS
(
    ID_PROJECT      BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ID_PROJECT_UUID UUID                                           DEFAULT uuid_generate_v4(),
    DT_CREATED_AT   timestamptz                           not null default current_timestamp,
    DT_UPDATED_AT   timestamptz                           not null,
    NM_PROJECT      varchar(128)                          not null,
    VL_COLOR        integer,
    ID_OWNER        BIGINT REFERENCES TBL_USERS (ID_USER) NOT NULL
);

CREATE TABLE IF NOT EXISTS TBL_TAGS
(
    ID_TAG        BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    DT_CREATED_AT timestamptz                           not null default current_timestamp,
    DT_UPDATED_AT timestamptz                           not null,
    NM_TAG        varchar(128)                          not null,
    VL_COLOR      integer,
    ID_OWNER      BIGINT REFERENCES TBL_USERS (ID_USER) NOT NULL,
    ID_PROJECT    BIGINT REFERENCES TBL_PROJECTS (ID_PROJECT)
);

CREATE TABLE IF NOT EXISTS TBL_TASKS
(
    ID_TASK       BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ID_TASK_UUID  UUID                                                 DEFAULT uuid_generate_v4(),
    DT_CREATED_AT timestamptz                                 not null default current_timestamp,
    DT_UPDATED_AT timestamptz                                 not null,
    NM_TASK       varchar(128)                                not null,
    ID_PROJECT    BIGINT REFERENCES TBL_PROJECTS (ID_PROJECT) not null,
    ID_OWNER      BIGINT REFERENCES TBL_USERS (ID_USER)       NOT NULL
);

CREATE TABLE IF NOT EXISTS TBL_ENTRIES
(
    ID_ENTRY      BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ID_ENTRY_UUID UUID                                           DEFAULT uuid_generate_v4(),
    DT_CREATED_AT timestamptz                           not null default current_timestamp,
    DT_UPDATED_AT timestamptz                           not null,
    NM_ENTRY      varchar(128)                          not null,
    DS_ENTRY      varchar(4096)                         not null,
    DT_START_AT   timestamptz                           not null,
    DT_FINISH_AT  timestamptz                           not null,
    ID_OWNER      BIGINT REFERENCES TBL_USERS (ID_USER) NOT NULL,
    ID_PROJECT    BIGINT REFERENCES TBL_PROJECTS (ID_PROJECT),
    ID_TASK       BIGINT REFERENCES TBL_TASKS (ID_TASK)
);

CREATE TABLE IF NOT EXISTS CRZ_PROJECTS_X_TAGS
(
    ID_PROJECT BIGINT REFERENCES TBL_PROJECTS (ID_PROJECT),
    ID_TAG     BIGINT REFERENCES TBL_TAGS (ID_TAG),
    PRIMARY KEY (ID_PROJECT, ID_TAG)
);

CREATE TABLE IF NOT EXISTS CRZ_TASKS_X_TAGS
(
    ID_TASK BIGINT REFERENCES TBL_TASKS (ID_TASK) NOT NULL,
    ID_TAG  BIGINT REFERENCES TBL_TAGS (ID_TAG)   NOT NULL,
    PRIMARY KEY (ID_TASK, ID_TAG)
);