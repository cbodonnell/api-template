-- Initialization Script

-- Create tables

-- public.users definition

DO $$
BEGIN

	CREATE TABLE IF NOT EXISTS public.users (
		id serial NOT NULL,
		"name" text NOT NULL,
		discoverable bool NOT NULL,
		iri text NOT NULL,
		CONSTRAINT users_pkey PRIMARY KEY (id)
	);

END
$$

