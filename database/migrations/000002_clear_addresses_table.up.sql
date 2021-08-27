DO $$
BEGIN
    IF (SELECT count(id) FROM public.addresses) > 0 THEN
        DELETE FROM public.addresses;
    END IF;
END $$;