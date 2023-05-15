--
-- PostgreSQL database dump
--

-- Dumped from database version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)

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
-- Name: tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tags (
    id bigint NOT NULL,
    user_id uuid NOT NULL,
    tag character varying(50) NOT NULL,
    created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL
);


--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: tweet_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tweet_tags (
    id bigint NOT NULL,
    tag character varying(50) NOT NULL,
    tweet_id uuid NOT NULL,
    user_id uuid NOT NULL
);


--
-- Name: tweet_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.tweet_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tweet_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.tweet_tags_id_seq OWNED BY public.tweet_tags.id;


--
-- Name: tweets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tweets (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    tweet character varying(300) NOT NULL,
    flags character varying(100) DEFAULT ''::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL
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
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: tweet_tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tweet_tags ALTER COLUMN id SET DEFAULT nextval('public.tweet_tags_id_seq'::regclass);


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
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: tweet_tags tweet_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tweet_tags
    ADD CONSTRAINT tweet_tags_pkey PRIMARY KEY (id);


--
-- Name: tweets tweets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tweets
    ADD CONSTRAINT tweets_pkey PRIMARY KEY (id);


--
-- Name: twitter_challenges twitter_challenges_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.twitter_challenges
    ADD CONSTRAINT twitter_challenges_pkey PRIMARY KEY (id);


--
-- Name: tags uniq_user_id_tag; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT uniq_user_id_tag UNIQUE (user_id, tag);


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
-- Name: idx_index_user_id_tag; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_index_user_id_tag ON public.tags USING btree (user_id, tag);


--
-- Name: idx_tags_tag; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tags_tag ON public.tags USING btree (tag);


--
-- Name: idx_tweet_tags_user_id_tag; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweet_tags_user_id_tag ON public.tweet_tags USING btree (user_id, tag);


--
-- Name: idx_tweets_timestamp; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweets_timestamp ON public.tweets USING btree (created_at);


--
-- Name: idx_tweets_user_timestamp; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweets_user_timestamp ON public.tweets USING btree (user_id, created_at);


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

