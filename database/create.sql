-- CREATE DATABASE test_ozon;

CREATE TABLE IF NOT EXISTS public.addresses(
    id SERIAL,
    long_url text UNIQUE,
    short_url text UNIQUE
);


DO $$
BEGIN
    IF (SELECT count(id) FROM public.addresses) > 0 THEN
        DELETE FROM public.addresses;
    END IF;
END $$;