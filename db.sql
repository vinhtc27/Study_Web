CREATE TABLE users (
	id bigserial,
	"name" varchar,
	dob varchar,
	sex varchar,
	avatar varchar,
	email varchar,
	address varchar,
	phone varchar,
	idcard varchar,
	"national" varchar,
	channels jsonb DEFAULT '[]'::jsonb,
	createddate timestamp NOT NULL,
	updateddate timestamp NOT NULL
);

CREATE TABLE accounts (
	id bigserial,
	userid bigint,
	email varchar,
	"password" varchar,
	accounttype varchar,
	accountstatus varchar,
	createddate timestamp NOT NULL,
	updateddate timestamp NOT NULL
);

CREATE TABLE channels (
    id bigserial,
    name varchar,
    avatar varchar,
	members jsonb DEFAULT '[]'::jsonb,
	messages jsonb DEFAULT '[]'::jsonb,
    tasks jsonb DEFAULT '[]'::jsonb,
    createddate timestamp NOT NULL,
    updateddate timestamp NOT NULL
);
