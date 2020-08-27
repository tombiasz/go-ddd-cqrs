--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3 (Debian 12.3-1.pgdg100+1)
-- Dumped by pg_dump version 12.3 (Ubuntu 12.3-1.pgdg18.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: coupons; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.coupons (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    code character varying(11) NOT NULL,
    description character varying(200),
    status character varying(20) NOT NULL,
    expdays integer NOT NULL,
    activatedat timestamp with time zone DEFAULT now() NOT NULL,
    expiredat timestamp with time zone,
    usedat timestamp with time zone
);


ALTER TABLE public.coupons OWNER TO postgres;

--
-- Data for Name: coupons; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.coupons (id, email, code, description, status, expdays, activatedat, expiredat, usedat) FROM stdin;
831aa974-545f-4fa6-820f-ba861088c831	test@expired.at	expiredcode	desc of expired code	Expired	7	2020-07-26 17:18:56.323176+00	2020-08-02 17:18:56.323176+00	\N
d831cdb4-0f79-4acf-bf05-ef3ad98896f8	test@used.at	usedcode	desc of used code	Used	7	2020-08-05 17:19:43.717418+00	\N	2020-08-09 17:19:43.717418+00
1219b7dc-538f-44f6-8976-0d609ac71ea4	test@test.com	gdm_q3sgg	some desc	Active	7	2020-08-16 16:27:29.984462+00	\N	\N
8d2a7a5d-d8a1-4d11-9953-4c6bf3638d42	test@test.com	gweycgngg	some desc	Active	7	2020-08-18 20:33:29.380095+00	\N	\N
123e4567-e89b-12d3-a456-426614174001	fizz@buzz.com	qnothercode	couponzzzz	Expired	7	2020-08-09 16:58:17.784432+00	2020-08-20 20:18:16.601934+00	\N
123e4567-e89b-12d3-a456-426614174000	foo@bar.com	couponcode	test coupon	Expired	7	2020-08-09 16:58:17.784432+00	2020-08-20 20:18:16.611889+00	\N
7f13a4f3-9e00-4f0d-8b5d-eca0eb4e36a1	test@test.com	tyly5mhgr	some desc	Used	7	2020-08-18 20:33:19.02338+00	\N	2020-08-20 20:42:18.939443+00


--
-- Name: coupons coupons_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.coupons
    ADD CONSTRAINT coupons_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

