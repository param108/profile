--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)

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
-- Name: onetime; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.onetime (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    data text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    writer uuid DEFAULT '6bd8d353-927e-463b-abb4-4b3c08c7c3af'::uuid NOT NULL
);


--
-- Name: resources; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.resources (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    t character varying(50) NOT NULL,
    value integer DEFAULT 0,
    max integer DEFAULT 0,
    writer uuid NOT NULL,
    CONSTRAINT resources_check CHECK (((value >= 0) AND (value <= max)))
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: sp_group_messages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_group_messages (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    sp_group_id uuid NOT NULL,
    sp_user_id uuid NOT NULL,
    sp_message_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    writer uuid NOT NULL
);


--
-- Name: sp_group_users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_group_users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    sp_group_id uuid NOT NULL,
    sp_user_id uuid NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    writer uuid NOT NULL
);


--
-- Name: sp_groups; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_groups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(100) NOT NULL,
    parent uuid,
    deleted boolean DEFAULT false NOT NULL,
    description text DEFAULT ''::text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    writer uuid NOT NULL
);


--
-- Name: sp_message_comments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_message_comments (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    sp_user_id uuid NOT NULL,
    sp_message_id uuid NOT NULL,
    msg_text text DEFAULT ''::text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    writer uuid NOT NULL
);


--
-- Name: sp_messages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_messages (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    sp_user_id uuid NOT NULL,
    msg_type character varying(100) NOT NULL,
    msg_value integer DEFAULT 0 NOT NULL,
    msg_text text DEFAULT ''::text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    writer uuid NOT NULL,
    sp_user_photo_url character varying(500) DEFAULT ''::character varying NOT NULL
);


--
-- Name: sp_otps; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_otps (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    phone character varying(20) NOT NULL,
    code character varying(10) NOT NULL,
    expiry timestamp with time zone NOT NULL,
    writer uuid NOT NULL,
    retries integer DEFAULT 0 NOT NULL
);


--
-- Name: sp_services; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_services (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    code character varying(100) NOT NULL,
    name text NOT NULL,
    category character varying(100) NOT NULL,
    unit character varying(50) NOT NULL,
    description text NOT NULL,
    short_description text NOT NULL,
    question text NOT NULL
);


--
-- Name: sp_users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sp_users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    phone character varying(20) NOT NULL,
    name character varying(100) NOT NULL,
    photo_url text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted boolean DEFAULT false NOT NULL,
    writer uuid NOT NULL,
    deleted_at timestamp with time zone,
    profile_updated boolean DEFAULT false NOT NULL
);


--
-- Name: tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tags (
    id bigint NOT NULL,
    user_id uuid NOT NULL,
    tag character varying(50) NOT NULL,
    created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    writer uuid NOT NULL
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
-- Name: thread_tweets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.thread_tweets (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    tweet_id uuid NOT NULL,
    thread_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    writer uuid NOT NULL
);


--
-- Name: threads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.threads (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    writer uuid NOT NULL,
    name character varying(50) DEFAULT ''::character varying NOT NULL
);


--
-- Name: tweet_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tweet_tags (
    id bigint NOT NULL,
    tag character varying(50) NOT NULL,
    tweet_id uuid NOT NULL,
    user_id uuid NOT NULL,
    writer uuid NOT NULL
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
    tweet character varying(500) NOT NULL,
    flags text DEFAULT ''::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    writer uuid NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    image character varying(150) DEFAULT ''::character varying NOT NULL,
    image_compressed boolean DEFAULT false,
    image_compressed_failed boolean DEFAULT false
);


--
-- Name: twitter_challenges; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.twitter_challenges (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    challenge text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    writer uuid NOT NULL,
    redirect_uri text DEFAULT ''::text NOT NULL
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
-- Name: onetime onetime_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.onetime
    ADD CONSTRAINT onetime_pkey PRIMARY KEY (id);


--
-- Name: resources resources_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resources
    ADD CONSTRAINT resources_pkey PRIMARY KEY (id);


--
-- Name: resources resources_user_id_writer_t_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.resources
    ADD CONSTRAINT resources_user_id_writer_t_key UNIQUE (user_id, writer, t);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sp_group_messages sp_group_messages_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_group_messages
    ADD CONSTRAINT sp_group_messages_pkey PRIMARY KEY (id);


--
-- Name: sp_group_users sp_group_users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_group_users
    ADD CONSTRAINT sp_group_users_pkey PRIMARY KEY (id);


--
-- Name: sp_groups sp_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_groups
    ADD CONSTRAINT sp_groups_pkey PRIMARY KEY (id);


--
-- Name: sp_message_comments sp_message_comments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_message_comments
    ADD CONSTRAINT sp_message_comments_pkey PRIMARY KEY (id);


--
-- Name: sp_messages sp_messages_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_messages
    ADD CONSTRAINT sp_messages_pkey PRIMARY KEY (id);


--
-- Name: sp_otps sp_otps_phone_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_otps
    ADD CONSTRAINT sp_otps_phone_key UNIQUE (phone);


--
-- Name: sp_otps sp_otps_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_otps
    ADD CONSTRAINT sp_otps_pkey PRIMARY KEY (id);


--
-- Name: sp_services sp_services_code_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_services
    ADD CONSTRAINT sp_services_code_key UNIQUE (code);


--
-- Name: sp_users sp_users_phone_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_users
    ADD CONSTRAINT sp_users_phone_key UNIQUE (phone);


--
-- Name: sp_users sp_users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sp_users
    ADD CONSTRAINT sp_users_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: thread_tweets thread_tweets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.thread_tweets
    ADD CONSTRAINT thread_tweets_pkey PRIMARY KEY (id);


--
-- Name: threads threads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.threads
    ADD CONSTRAINT threads_pkey PRIMARY KEY (id);


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
-- Name: idx_image_compressed_failed; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_image_compressed_failed ON public.tweets USING btree (image_compressed, image_compressed_failed);


--
-- Name: idx_index_user_id_tag; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_index_user_id_tag ON public.tags USING btree (user_id, tag);


--
-- Name: idx_resources_user_id_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_resources_user_id_type ON public.resources USING btree (user_id, writer, t);


--
-- Name: idx_sp_group_messages_sp_group_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_group_messages_sp_group_id ON public.sp_group_messages USING btree (sp_group_id, created_at, writer);


--
-- Name: idx_sp_group_users_sp_group_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_group_users_sp_group_id ON public.sp_group_users USING btree (sp_group_id, writer);


--
-- Name: idx_sp_group_users_sp_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_group_users_sp_user_id ON public.sp_group_users USING btree (sp_user_id, writer);


--
-- Name: idx_sp_groups_name; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_groups_name ON public.sp_groups USING btree (name, writer);


--
-- Name: idx_sp_message_comments_sp_message_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_message_comments_sp_message_id ON public.sp_message_comments USING btree (sp_message_id, writer);


--
-- Name: idx_sp_messages_msg_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_messages_msg_type ON public.sp_messages USING btree (msg_type, created_at, writer);


--
-- Name: idx_sp_messages_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_messages_user_id ON public.sp_messages USING btree (sp_user_id, created_at, writer);


--
-- Name: idx_sp_services_category; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_services_category ON public.sp_services USING btree (category);


--
-- Name: idx_sp_services_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sp_services_code ON public.sp_services USING btree (code);


--
-- Name: idx_tags_tag; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tags_tag ON public.tags USING btree (tag);


--
-- Name: idx_tags_writer; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tags_writer ON public.tags USING btree (writer);


--
-- Name: idx_thread_tweets_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_thread_tweets_user_id ON public.thread_tweets USING btree (user_id, thread_id, writer);


--
-- Name: idx_threads_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_threads_user_id ON public.threads USING btree (user_id, writer);


--
-- Name: idx_tweet_tags_user_id_tag; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweet_tags_user_id_tag ON public.tweet_tags USING btree (user_id, tag);


--
-- Name: idx_tweet_tags_writer; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweet_tags_writer ON public.tweet_tags USING btree (writer);


--
-- Name: idx_tweets_timestamp; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweets_timestamp ON public.tweets USING btree (created_at);


--
-- Name: idx_tweets_user_timestamp; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweets_user_timestamp ON public.tweets USING btree (user_id, created_at);


--
-- Name: idx_tweets_writer; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_tweets_writer ON public.tweets USING btree (writer);


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

