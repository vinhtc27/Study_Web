CREATE TABLE public.users (
	id bigserial NOT NULL,
	"name" varchar(255) NULL,
	dob varchar(255) NULL,
	sex varchar(255) NULL,
	avartar varchar(255) NULL,
	email varchar(255) NULL,
	address varchar(255) NULL,
	phone varchar(255) NULL,
	idcard varchar(255) NULL,
	"national" varchar(255) NULL,
	createddate timestamp NOT NULL,
	updateddate timestamp NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE public.accounts (
	id bigserial NOT NULL,
	userid int8 NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	accounttype varchar(255) NOT NULL,
	accountstatus varchar(255) NOT NULL,
	createddate timestamp NOT NULL,
	updateddate timestamp NOT NULL,
	CONSTRAINT accounts_pk PRIMARY KEY (id),
	CONSTRAINT accounts_fk FOREIGN KEY (id) REFERENCES public.users(id)
);

CREATE TABLE channels (
      id bigserial NOT NULL,
      name varchar(255) not null,
      member_id int8 NOT NULL,
      member_name varchar not null,
      role smallint not null,
      created_date timestamp NOT NULL,
      CONSTRAINT channels_pk PRIMARY KEY (id),
      CONSTRAINT channels_fk FOREIGN KEY (member_id) REFERENCES public.users(id)
);


CREATE TABLE message (
     id bigserial NOT NULL,
     sender_id int8 NOT NULL,
     sender_name varchar not null,
     content text not null,
     channel_id int8 NOT null,
     created_date timestamp NOT NULL,
     last_modified timestamp,
     type int not null,
     CONSTRAINT message_pk PRIMARY KEY (id),
     CONSTRAINT message_fk FOREIGN KEY (channel_id) REFERENCES public.channels(id)
);


