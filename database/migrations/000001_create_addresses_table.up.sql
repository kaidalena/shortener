CREATE TABLE IF NOT EXISTS public.addresses(
       id SERIAL,
       long_url text UNIQUE,
       short_url text UNIQUE
);