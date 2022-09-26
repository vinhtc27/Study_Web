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
