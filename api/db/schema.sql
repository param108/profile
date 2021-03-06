--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.6 (Ubuntu 13.6-0ubuntu0.21.10.1)

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

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: user_role; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.user_role AS ENUM (
    'user',
    'admin'
);


SET default_table_access_method = heap;

--
-- Name: invalid_tokens; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.invalid_tokens (
    token text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    writer uuid NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: twitter_challenges; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.twitter_challenges (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    challenge text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    writer uuid NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    handle text NOT NULL,
    profile text,
    role public.user_role,
    writer uuid NOT NULL
);


--
-- Name: invalid_tokens invalid_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invalid_tokens
    ADD CONSTRAINT invalid_tokens_pkey PRIMARY KEY (token);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: twitter_challenges twitter_challenges_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.twitter_challenges
    ADD CONSTRAINT twitter_challenges_pkey PRIMARY KEY (id);


--
-- Name: users users_handle_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_handle_key UNIQUE (handle);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_twitter_challenges_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_twitter_challenges_created_at ON public.twitter_challenges USING btree (created_at);


--
-- Name: users_on_handle; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX users_on_handle ON public.users USING btree (handle);


--
-- PostgreSQL database dump complete
--

