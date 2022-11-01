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
create index idx_users_id on users(id);

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
create index idx_accounts_id on accounts(id);

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
create index idx_channels_id on channels(id);
              








