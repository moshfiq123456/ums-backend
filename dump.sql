--
-- PostgreSQL database dump
--

\restrict CgAcI4osYmb8GbTMQez4NCENw49tWbBXWc0bV12RepDEOmuBCXYW5oM03OVkBpv

-- Dumped from database version 16.11
-- Dumped by pg_dump version 16.11

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
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: audit_logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.audit_logs (
    id bigint NOT NULL,
    user_id uuid,
    action character varying(100) NOT NULL,
    entity character varying(50) NOT NULL,
    entity_id character varying(50),
    old_value jsonb,
    new_value jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    organization_id uuid,
    service character varying(50) DEFAULT 'ums'::character varying,
    ip_address character varying(45)
);


--
-- Name: audit_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.audit_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: audit_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.audit_logs_id_seq OWNED BY public.audit_logs.id;


--
-- Name: login_sessions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.login_sessions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    refresh_token_hash character varying(255) NOT NULL,
    refresh_expires_at timestamp without time zone NOT NULL,
    ip_address character varying(45),
    user_agent text,
    logged_in_at timestamp without time zone DEFAULT now() NOT NULL,
    logged_out_at timestamp without time zone,
    organization_id uuid,
    device_info jsonb DEFAULT '{}'::jsonb
);


--
-- Name: organizations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.organizations (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(200) NOT NULL,
    slug character varying(100) NOT NULL,
    domain character varying(200),
    logo_url character varying(500),
    is_active boolean DEFAULT true NOT NULL,
    plan character varying(50) DEFAULT 'free'::character varying NOT NULL,
    settings jsonb DEFAULT '{}'::jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.permissions (
    id bigint NOT NULL,
    code character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    organization_id uuid NOT NULL
);


--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: role_permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.role_permissions (
    id bigint NOT NULL,
    role_id bigint NOT NULL,
    permission_id bigint NOT NULL
);


--
-- Name: role_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.role_permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: role_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.role_permissions_id_seq OWNED BY public.role_permissions.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    code character varying(100) NOT NULL,
    description text,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    organization_id uuid NOT NULL
);


--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: service_clients; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.service_clients (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    organization_id uuid NOT NULL,
    name character varying(100) NOT NULL,
    client_id character varying(100) NOT NULL,
    client_secret_hash text NOT NULL,
    allowed_scopes text[] DEFAULT '{}'::text[],
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: user_hierarchy; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_hierarchy (
    id bigint NOT NULL,
    parent_user_id uuid NOT NULL,
    child_user_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: user_hierarchy_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_hierarchy_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_hierarchy_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_hierarchy_id_seq OWNED BY public.user_hierarchy.id;


--
-- Name: user_permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_permissions (
    id bigint NOT NULL,
    user_id uuid NOT NULL,
    permission_id bigint NOT NULL,
    allow boolean DEFAULT true NOT NULL
);


--
-- Name: user_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_permissions_id_seq OWNED BY public.user_permissions.id;


--
-- Name: user_roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_roles (
    id bigint NOT NULL,
    user_id uuid NOT NULL,
    role_id bigint NOT NULL,
    assigned_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: user_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_roles_id_seq OWNED BY public.user_roles.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(150),
    password_hash text,
    phone character varying(20),
    status character varying(20) DEFAULT 'active'::character varying NOT NULL,
    email_verified_at timestamp without time zone,
    last_login_at timestamp without time zone,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    avatar_url character varying(500),
    organization_id uuid NOT NULL,
    user_type character varying(20) DEFAULT 'member'::character varying NOT NULL,
    metadata jsonb DEFAULT '{}'::jsonb
);


--
-- Name: audit_logs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs ALTER COLUMN id SET DEFAULT nextval('public.audit_logs_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: role_permissions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_permissions ALTER COLUMN id SET DEFAULT nextval('public.role_permissions_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: user_hierarchy id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_hierarchy ALTER COLUMN id SET DEFAULT nextval('public.user_hierarchy_id_seq'::regclass);


--
-- Name: user_permissions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_permissions ALTER COLUMN id SET DEFAULT nextval('public.user_permissions_id_seq'::regclass);


--
-- Name: user_roles id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_roles ALTER COLUMN id SET DEFAULT nextval('public.user_roles_id_seq'::regclass);


--
-- Data for Name: audit_logs; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.audit_logs (id, user_id, action, entity, entity_id, old_value, new_value, created_at, organization_id, service, ip_address) FROM stdin;
1	7bd17c85-5ded-4feb-a692-76422bc5dab2	user_created	user	alice@ums.com	\N	\N	2026-02-13 00:20:50.528756	\N	ums	\N
2	7bd17c85-5ded-4feb-a692-76422bc5dab2	role_assigned	user	alice@ums.com	\N	\N	2026-02-13 00:20:50.534939	\N	ums	\N
3	7bd17c85-5ded-4feb-a692-76422bc5dab2	user_created	user	david@ums.com	\N	\N	2026-02-18 00:20:50.536273	\N	ums	\N
4	7bd17c85-5ded-4feb-a692-76422bc5dab2	role_assigned	user	david@ums.com	\N	\N	2026-02-18 00:20:50.537375	\N	ums	\N
5	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	user_created	user	priya@ums.com	\N	\N	2026-03-05 00:20:50.538536	\N	ums	\N
6	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	role_assigned	user	priya@ums.com	\N	\N	2026-03-05 00:20:50.539741	\N	ums	\N
7	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	user_created	user	sophie@ums.com	\N	\N	2026-03-20 00:20:50.541077	\N	ums	\N
8	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	user_created	user	bob@ums.com	\N	\N	2026-02-23 00:20:50.542256	\N	ums	\N
9	7dbbc23d-197f-464b-9288-5eb6dfde8e61	user_created	user	nina@ums.com	\N	\N	2026-04-04 00:20:50.543227	\N	ums	\N
10	7dbbc23d-197f-464b-9288-5eb6dfde8e61	user_status_changed	user	eve@ums.com	\N	\N	2026-04-29 00:20:50.544192	\N	ums	\N
11	7bd17c85-5ded-4feb-a692-76422bc5dab2	user_status_changed	user	frank@ums.com	\N	\N	2026-05-02 00:20:50.54541	\N	ums	\N
12	7bd17c85-5ded-4feb-a692-76422bc5dab2	role_created	role	auditor	\N	\N	2026-03-30 00:20:50.54668	\N	ums	\N
13	7bd17c85-5ded-4feb-a692-76422bc5dab2	role_created	role	hr_manager	\N	\N	2026-03-03 00:20:50.548527	\N	ums	\N
14	7bd17c85-5ded-4feb-a692-76422bc5dab2	permission_created	permission	audit:read	\N	\N	2026-03-30 00:20:50.550399	\N	ums	\N
15	7bd17c85-5ded-4feb-a692-76422bc5dab2	permission_created	permission	report:generate	\N	\N	2026-03-30 00:20:50.552757	\N	ums	\N
16	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	user_permission_assigned	user	bob@ums.com	\N	\N	2026-04-09 00:20:50.555088	\N	ums	\N
17	7dbbc23d-197f-464b-9288-5eb6dfde8e61	hierarchy_assigned	user	nina@ums.com	\N	\N	2026-04-05 00:20:50.557005	\N	ums	\N
18	7bd17c85-5ded-4feb-a692-76422bc5dab2	hierarchy_assigned	user	alice@ums.com	\N	\N	2026-02-14 00:20:50.558957	\N	ums	\N
\.


--
-- Data for Name: login_sessions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.login_sessions (id, user_id, refresh_token_hash, refresh_expires_at, ip_address, user_agent, logged_in_at, logged_out_at, organization_id, device_info) FROM stdin;
fd05bc3c-426b-47e6-80c7-a5c5248fe2be	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-11 00:00:41.658195	\N	\N	2026-05-04 00:00:41.65865	\N	\N	{}
ebc5eb7b-d8c0-4916-87b4-b8f72294e328	7bd17c85-5ded-4feb-a692-76422bc5dab2	historical-session-hash	2026-04-30 16:20:50.560834	192.168.0.101	Mozilla/5.0 Chrome/123 Linux	2026-04-23 16:20:50.560834	2026-04-23 18:20:50.560834	\N	{}
43f78e1b-ef6e-48db-8fc4-865f5866e41e	7bd17c85-5ded-4feb-a692-76422bc5dab2	historical-session-hash	2026-05-05 16:20:50.567152	192.168.0.101	Mozilla/5.0 Chrome/123 Linux	2026-04-28 16:20:50.567152	2026-04-28 18:20:50.567152	\N	{}
ae196a6e-3d9f-4012-9495-9aa13faf70bb	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	historical-session-hash	2026-05-02 16:20:50.569708	192.168.0.102	Mozilla/5.0 Firefox/124 Linux	2026-04-25 16:20:50.569708	2026-04-25 18:20:50.569708	\N	{}
94f11bb1-5b2c-4ecd-9874-6c1fd03584e5	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	historical-session-hash	2026-05-07 16:20:50.572212	192.168.0.102	Mozilla/5.0 Firefox/124 Linux	2026-04-30 16:20:50.572212	2026-04-30 18:20:50.572212	\N	{}
11af0324-95db-4b01-b06e-7f2cf4817576	7dbbc23d-197f-464b-9288-5eb6dfde8e61	historical-session-hash	2026-05-03 16:20:50.574653	192.168.0.105	Mozilla/5.0 Chrome/123 Windows	2026-04-26 16:20:50.574653	2026-04-26 18:20:50.574653	\N	{}
a4cb8da2-1404-490c-9858-887cd36bce1e	2c05e61c-c6e0-4de3-8a5c-44a8e22ac4a8	historical-session-hash	2026-05-04 16:20:50.577187	192.168.0.103	Mozilla/5.0 Safari/17 macOS	2026-04-27 16:20:50.577187	2026-04-27 18:20:50.577187	\N	{}
f2ce6057-785b-4f43-9e05-150e50170e10	14f9cca5-62d3-42ce-b47f-5cedea7c1622	historical-session-hash	2026-05-06 16:20:50.579537	192.168.0.106	Mozilla/5.0 Chrome/123 Linux	2026-04-29 16:20:50.579537	2026-04-29 18:20:50.579537	\N	{}
ef83cdd5-984e-4b43-9cb7-3885562ece84	f1071943-d948-458e-97bb-05caa03c447b	historical-session-hash	2026-05-08 16:20:50.581422	192.168.0.107	Mozilla/5.0 Edge/123 Windows	2026-05-01 16:20:50.581422	2026-05-01 18:20:50.581422	\N	{}
7d661d82-9452-4165-b0b5-3fc36b07bdab	9c542918-66cb-409d-bf53-22a8c95bb608	historical-session-hash	2026-05-09 16:20:50.583478	192.168.0.109	Mozilla/5.0 Chrome/123 Linux	2026-05-02 16:20:50.583478	2026-05-02 18:20:50.583478	\N	{}
609653b7-6e64-425d-82cc-3b0a0496b3b7	69f8f2d7-8ed4-4fdc-b7e9-715dac5f54aa	historical-session-hash	2026-04-25 16:20:50.585281	10.0.0.55	curl/7.88.0	2026-04-18 16:20:50.585281	2026-04-18 18:20:50.585281	\N	{}
f1fe141b-c8fc-4f9d-933a-2b308cf72efd	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-14 17:30:00.006753	\N	\N	2026-05-07 17:30:00.010962	\N	\N	{}
f4b492c3-0562-48e4-afd4-2c4b134e2393	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-16 21:28:43.326228	\N	\N	2026-05-09 21:28:43.331134	\N	\N	{}
e08d96c0-c515-4624-9ba0-9983dab1f91c	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-16 22:49:10.24823	\N	\N	2026-05-09 22:49:10.248704	\N	\N	{}
f5447989-a7da-4eee-975e-394cd445cfa8	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-17 10:44:26.088912	\N	\N	2026-05-10 10:44:26.089304	\N	\N	{}
40118d52-7c44-42ca-a491-fda277100680	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-22 18:14:22.116235	\N	\N	2026-05-15 18:14:22.1171	\N	\N	{}
cdd5223d-b4b4-47fe-85c8-d8ce81f7cbfe	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-22 18:20:30.75227	\N	\N	2026-05-15 18:20:30.752747	\N	\N	{}
41631906-f417-4246-b87a-ebfbf1fd1673	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-22 18:21:07.657807	\N	\N	2026-05-15 18:21:07.658193	\N	\N	{}
1db640aa-cf73-4700-9ffd-e8c664020db4	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-22 18:29:57.489186	\N	\N	2026-05-15 18:29:57.489463	\N	\N	{}
481a38e2-c2ff-4516-acf3-81862908c2c2	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-22 18:30:26.702382	\N	\N	2026-05-15 18:30:26.702669	\N	\N	{}
08ae006a-7a0f-4110-b450-d8e67ff8b570	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-22 18:33:03.46146	\N	\N	2026-05-15 18:33:03.46181	\N	\N	{}
4e1f4e2e-8e3d-4649-8202-7a9f35af73ea	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-23 00:39:41.21819	\N	\N	2026-05-16 00:39:41.220307	\N	\N	{}
bf5d3224-a963-4921-8562-8e883d75d81e	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-21 17:09:01.847029	\N	\N	2026-05-14 17:09:01.853591	2026-05-17 22:35:02.956792	\N	{}
d6c72f02-8585-4197-ac3e-843e5a7c84d1	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-27 13:46:06.595292	\N	\N	2026-01-20 13:46:06.59574	\N	00000000-0000-0000-0000-000000000001	{}
fe234fd1-5968-4405-8475-e935120828ee	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-27 13:52:48.581083	\N	\N	2026-01-20 13:52:48.581566	\N	00000000-0000-0000-0000-000000000001	{}
ddcc9843-ca27-4511-af30-043b14c579a9	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-20 14:13:34.024432	\N	\N	2026-01-20 14:13:34.024837	\N	00000000-0000-0000-0000-000000000001	{}
ab82aa98-f441-4fba-88e6-5bcc20d91b66	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-20 14:13:38.404438	\N	\N	2026-01-20 14:13:38.404747	\N	00000000-0000-0000-0000-000000000001	{}
3583d96b-9c95-4b26-a13f-95851f6718e0	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-27 14:14:01.005557	\N	\N	2026-01-20 14:14:01.006142	\N	00000000-0000-0000-0000-000000000001	{}
e6b34d8f-c6a3-4d10-b50a-1a97501a13c0	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-28 13:21:48.206686	\N	\N	2026-01-21 13:21:48.20721	\N	00000000-0000-0000-0000-000000000001	{}
dc5cc4f4-607f-44e2-835e-257d0f11fd06	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-28 13:30:23.127157	\N	\N	2026-01-21 13:30:23.127504	\N	00000000-0000-0000-0000-000000000001	{}
8653eb78-1273-4c08-9624-cc5fc7906b9d	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-28 17:41:15.519216	\N	\N	2026-01-21 17:41:15.530136	\N	00000000-0000-0000-0000-000000000001	{}
1f5b216d-6ebc-4284-8c53-34b4ba1aad3a	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-28 18:02:29.057991	\N	\N	2026-01-21 18:02:29.058403	\N	00000000-0000-0000-0000-000000000001	{}
726cb7f9-ddc3-4d18-bfab-191fb254bb4b	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-28 18:10:15.255807	\N	\N	2026-01-21 18:10:15.256206	\N	00000000-0000-0000-0000-000000000001	{}
268057f0-3dd6-4592-8227-89246b1669e9	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 17:01:13.610022	\N	\N	2026-01-23 17:01:13.610511	\N	00000000-0000-0000-0000-000000000001	{}
f5a0521b-c97f-416b-aaa7-ef06dc3b11c2	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 17:56:58.777687	\N	\N	2026-01-23 17:56:58.778106	\N	00000000-0000-0000-0000-000000000001	{}
22fb870b-4a55-4380-83d6-34a1bbbf1554	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 18:04:28.080283	\N	\N	2026-01-23 18:04:28.080681	\N	00000000-0000-0000-0000-000000000001	{}
100aff0a-3cff-45fb-a523-34a6e449a8a6	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 18:09:15.462711	\N	\N	2026-01-23 18:09:15.463591	\N	00000000-0000-0000-0000-000000000001	{}
dda9e1ba-3408-4946-be6d-52215501b865	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 18:09:36.923882	\N	\N	2026-01-23 18:09:36.924279	\N	00000000-0000-0000-0000-000000000001	{}
aedaba27-4320-47ee-ad10-293f30d78f20	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 18:13:24.093983	\N	\N	2026-01-23 18:13:24.094387	\N	00000000-0000-0000-0000-000000000001	{}
aacc4b4b-1ed5-4975-a0e1-7193f07f0889	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 18:16:41.102772	\N	\N	2026-01-23 18:16:41.103129	\N	00000000-0000-0000-0000-000000000001	{}
80227b5f-4499-479b-b404-de2b5b810c1f	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 18:45:26.142206	\N	\N	2026-01-23 18:45:26.142622	\N	00000000-0000-0000-0000-000000000001	{}
8c15455f-b208-46fd-99ac-30491a730db0	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 13:23:19.992694	\N	\N	2026-01-26 13:23:19.997256	\N	00000000-0000-0000-0000-000000000001	{}
199069a9-d12a-4fe6-b805-4721f32b527e	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 13:24:46.149657	\N	\N	2026-01-26 13:24:46.149965	\N	00000000-0000-0000-0000-000000000001	{}
7ce95592-ba56-4d52-ada5-f4fd9dd912b4	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 13:27:10.897899	\N	\N	2026-01-26 13:27:10.898316	\N	00000000-0000-0000-0000-000000000001	{}
cab05b7d-d3a2-4d32-b74e-2a8f7d9c1e84	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 13:34:30.482592	\N	\N	2026-01-26 13:34:30.482972	\N	00000000-0000-0000-0000-000000000001	{}
9378bcc8-ef0b-432d-b767-7fdfe33bcc7c	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 13:37:01.014424	\N	\N	2026-01-26 13:37:01.014729	\N	00000000-0000-0000-0000-000000000001	{}
a003a75e-d216-4a6e-a8cd-aa184a9e6e7b	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 17:12:03.824417	\N	\N	2026-01-26 17:12:03.824788	\N	00000000-0000-0000-0000-000000000001	{}
be6053ec-3dfd-4bd8-ad2e-af65469ec363	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 17:41:33.605168	\N	\N	2026-01-26 17:41:33.605609	\N	00000000-0000-0000-0000-000000000001	{}
22497381-f793-4d18-8c6f-b01932d3e23b	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 18:01:06.085697	\N	\N	2026-01-26 18:01:06.08595	\N	00000000-0000-0000-0000-000000000001	{}
283e3d57-f9dc-4f0b-905c-f9f0e9f157c0	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 18:30:53.366037	\N	\N	2026-01-26 18:30:53.366417	\N	00000000-0000-0000-0000-000000000001	{}
6b641ab9-3dcb-4ac3-9cde-950b4a2f264a	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:00:31.844353	\N	\N	2026-01-26 19:00:31.844671	\N	00000000-0000-0000-0000-000000000001	{}
95327f7f-a874-4359-bcbb-d9176b99c6f7	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:01:38.767599	\N	\N	2026-01-26 19:01:38.768139	\N	00000000-0000-0000-0000-000000000001	{}
6d39aa71-c6e9-4ecb-b3b4-d6cbd58ea948	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:02:08.291831	\N	\N	2026-01-26 19:02:08.292398	\N	00000000-0000-0000-0000-000000000001	{}
583e27b4-8587-4d5e-93bc-cd36047b0c44	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:02:21.299464	\N	\N	2026-01-26 19:02:21.299761	\N	00000000-0000-0000-0000-000000000001	{}
1fe8aec5-1d51-4e79-85f6-212c303b7686	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:05:24.094271	\N	\N	2026-01-26 19:05:24.094637	\N	00000000-0000-0000-0000-000000000001	{}
f3cfa565-4516-40ef-84b9-272290f4ef46	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:35:05.616599	\N	\N	2026-01-26 19:35:05.617025	\N	00000000-0000-0000-0000-000000000001	{}
0c4a720d-c350-483b-876a-abc5bc5b580f	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:35:55.991319	\N	\N	2026-01-26 19:35:55.991737	\N	00000000-0000-0000-0000-000000000001	{}
e746a089-9fe4-4b4e-bd35-42368ef0acc7	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:37:06.361986	\N	\N	2026-01-26 19:37:06.362652	\N	00000000-0000-0000-0000-000000000001	{}
2dcbda05-1dd8-4693-8110-45c1a6aa9a34	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:50:05.795634	\N	\N	2026-01-26 19:50:05.796615	\N	00000000-0000-0000-0000-000000000001	{}
185db010-7c76-4dc0-bd6d-f99c3243b92b	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 19:59:12.245672	\N	\N	2026-01-26 19:59:12.246345	\N	00000000-0000-0000-0000-000000000001	{}
8829d4b7-01d2-4900-af1d-2a092f8036e3	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 20:04:07.00916	\N	\N	2026-01-26 20:04:07.009575	\N	00000000-0000-0000-0000-000000000001	{}
7cf00c9d-c9ec-45af-ac66-9d2e3b2ac6cc	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-01-30 18:19:01.916726	\N	\N	2026-01-23 18:19:01.916986	2026-01-26 20:07:45.727266	00000000-0000-0000-0000-000000000001	{}
d62925be-7a22-4312-b9ec-cee38fb25f00	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 20:11:07.810003	\N	\N	2026-01-26 20:11:07.810573	\N	00000000-0000-0000-0000-000000000001	{}
56628628-a2ab-4e4e-8c77-70b8ce5f139e	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 20:12:16.27374	\N	\N	2026-01-26 20:12:16.274135	\N	00000000-0000-0000-0000-000000000001	{}
2d64a809-885f-412a-94b9-cf24499ef95d	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 20:13:16.766375	\N	\N	2026-01-26 20:13:16.766729	\N	00000000-0000-0000-0000-000000000001	{}
ecaa27c8-4da8-4511-98ca-e3a3cb994359	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 20:13:53.56475	\N	\N	2026-01-26 20:13:53.565443	\N	00000000-0000-0000-0000-000000000001	{}
91733449-c8f6-43c0-aece-5f9ad16c81ec	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-02 20:20:05.215557	\N	\N	2026-01-26 20:20:05.215989	\N	00000000-0000-0000-0000-000000000001	{}
debe8fd7-911a-411b-99ae-77c7bd6a8fba	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:05:06.118316	\N	\N	2026-01-27 13:05:06.118959	\N	00000000-0000-0000-0000-000000000001	{}
61368b00-ff94-4ea1-8db0-98fe20b5421e	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:07:22.108152	\N	\N	2026-01-27 13:07:22.108531	\N	00000000-0000-0000-0000-000000000001	{}
f479c11a-ac9a-499c-b1ab-209573f4ef0e	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:13:39.474145	\N	\N	2026-01-27 13:13:39.474611	\N	00000000-0000-0000-0000-000000000001	{}
be2d2921-3d0c-4dce-a470-789090c9837e	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:14:11.478128	\N	\N	2026-01-27 13:14:11.478527	\N	00000000-0000-0000-0000-000000000001	{}
797d1839-0350-4661-995b-a6760409fca1	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:16:13.689623	\N	\N	2026-01-27 13:16:13.689969	2026-01-27 13:36:12.536401	00000000-0000-0000-0000-000000000001	{}
d8a1b4f3-3db7-4633-aeb3-764292189c76	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:36:39.278281	\N	\N	2026-01-27 13:36:39.278763	2026-01-27 13:36:42.87306	00000000-0000-0000-0000-000000000001	{}
ceb66533-a544-4a6f-8edb-240dd73c9de5	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:38:40.51771	\N	\N	2026-01-27 13:38:40.518084	2026-01-27 13:49:34.582002	00000000-0000-0000-0000-000000000001	{}
a0030dec-c43b-4ba6-8d4d-fc990617c99f	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 13:55:34.955614	\N	\N	2026-01-27 13:55:34.956016	2026-01-27 14:05:53.674078	00000000-0000-0000-0000-000000000001	{}
2707ab8d-ccf8-44f0-aa1e-53e89796e531	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-03 14:40:50.517761	\N	\N	2026-01-27 14:40:50.518212	2026-01-27 14:46:56.101656	00000000-0000-0000-0000-000000000001	{}
87682904-63d2-4740-976f-3f53f4f40fe2	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-04 13:59:09.339073	\N	\N	2026-01-28 13:59:09.343622	\N	00000000-0000-0000-0000-000000000001	{}
1978fff2-4dcf-4fe7-89f3-17da8451f5bd	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-16 13:35:32.543228	\N	\N	2026-02-09 13:35:32.547526	\N	00000000-0000-0000-0000-000000000001	{}
824a6906-4351-4a76-a2fa-a1aaa515d13b	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-02-24 17:50:44.644747	\N	\N	2026-02-17 17:50:44.645816	\N	00000000-0000-0000-0000-000000000001	{}
a730c079-3f16-47b6-804b-db336d98e5ee	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-03-12 16:55:48.692082	\N	\N	2026-03-05 16:55:48.692477	\N	00000000-0000-0000-0000-000000000001	{}
48662335-6cd6-4cb0-ae3e-e7a01ca260cd	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-03-12 17:39:47.704646	\N	\N	2026-03-05 17:39:47.705044	\N	00000000-0000-0000-0000-000000000001	{}
27ff8dac-8a53-4cd5-9874-8ae4d1704511	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-04-08 13:38:01.194335	\N	\N	2026-04-01 13:38:01.194754	2026-04-10 13:59:41.570521	00000000-0000-0000-0000-000000000001	{}
15055b0a-1bdc-4a1c-a10c-ce5514fb1f43	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-04-17 14:07:42.884401	\N	\N	2026-04-10 14:07:42.885429	2026-04-10 15:16:18.707951	00000000-0000-0000-0000-000000000001	{}
a9845298-2073-41c1-a630-27bf5693c1aa	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-04-17 15:18:12.136288	\N	\N	2026-04-10 15:18:12.136712	2026-04-10 15:29:12.389916	00000000-0000-0000-0000-000000000001	{}
61fa7e8e-5456-453e-8a0a-0be6c65d575d	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-04-17 15:29:22.43469	\N	\N	2026-04-10 15:29:22.435093	2026-04-10 15:35:02.635328	00000000-0000-0000-0000-000000000001	{}
122bc0f1-d22b-4339-b52c-f2180f7ad266	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-04-17 16:35:46.735264	\N	\N	2026-04-10 16:35:46.73563	\N	00000000-0000-0000-0000-000000000001	{}
a70477cd-9131-4039-8900-517732e03fd5	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-04-28 18:08:13.48002	\N	\N	2026-04-21 18:08:13.48053	2026-04-22 00:10:55.371944	00000000-0000-0000-0000-000000000001	{}
f8935851-0030-49f6-8720-8b88165f230e	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-04-29 00:44:01.669324	\N	\N	2026-04-22 00:44:01.670868	\N	00000000-0000-0000-0000-000000000001	{}
2aa1adac-577f-4889-a4d0-f380a93b3849	3c77374a-29b2-4bc3-93dd-011277c861bd		2026-05-05 13:55:16.479268	\N	\N	2026-04-28 13:55:16.480052	\N	00000000-0000-0000-0000-000000000001	{}
8b64633a-79d7-4b5b-967f-d539faac5230	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-11 00:08:48.851783	\N	\N	2026-05-04 00:08:48.852285	\N	\N	{}
35a796a2-fdce-45ad-a7aa-9e1932ecdd5a	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-11 00:09:09.135717	\N	\N	2026-05-04 00:09:09.13606	\N	\N	{}
81fb53c5-f4aa-4dbc-8ca8-45422e9aafea	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-11 00:23:03.855621	\N	\N	2026-05-04 00:23:03.856104	\N	\N	{}
fcf1f162-47b7-4442-bbba-a946caeb1aea	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-11 00:33:43.869924	\N	\N	2026-05-04 00:33:43.870269	\N	\N	{}
88f840a8-00e2-4f56-b3e2-f585c71bda21	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-11 00:33:51.384613	\N	\N	2026-05-04 00:33:51.385032	\N	\N	{}
af82b2a6-3690-42eb-a3d3-66bf4e07b394	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 10:58:43.97676	\N	\N	2026-05-08 10:58:43.981027	\N	\N	{}
d1c07f60-73d0-48e9-8118-3b021bb824f6	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 11:02:05.009251	\N	\N	2026-05-08 11:02:05.009672	\N	\N	{}
2d60e8ca-e285-4d41-ae55-5ba3cc650222	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 18:04:02.627649	\N	\N	2026-05-08 18:04:02.633522	\N	\N	{}
9c044335-74c4-47ef-96bf-d1e9a7eaedea	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 18:04:20.66815	\N	\N	2026-05-08 18:04:20.671327	\N	\N	{}
b9bf3670-a94a-4f3b-9785-c39e77d20242	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 18:04:45.116674	\N	\N	2026-05-08 18:04:45.117445	\N	\N	{}
89c249f6-9a31-469e-8acf-923ddebc12da	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 18:05:06.975585	\N	\N	2026-05-08 18:05:06.976387	\N	\N	{}
04efce1c-6b3a-4190-b2f8-794de1d4715c	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 18:05:31.241409	\N	\N	2026-05-08 18:05:31.242549	\N	\N	{}
a79d2b12-b30e-447f-b763-c679dec39f85	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-15 18:06:05.47719	\N	\N	2026-05-08 18:06:05.478307	\N	\N	{}
7ea2bf40-19a5-4fbc-9c7b-8d9e611b0f32	7bd17c85-5ded-4feb-a692-76422bc5dab2		2026-05-16 11:15:54.016895	\N	\N	2026-05-09 11:15:54.019243	\N	\N	{}
35f9eb94-4a13-42b7-b51a-9ac8badb7652	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-16 21:58:38.197399	\N	\N	2026-05-09 21:58:38.202359	\N	\N	{}
39bce813-a1fd-47fe-a399-c6c31b89e9d8	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-16 22:48:47.023098	\N	\N	2026-05-09 22:48:47.023746	\N	\N	{}
6bdc03c8-a593-45cf-81de-524c3cf7566d	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c		2026-05-17 02:53:49.971847	\N	\N	2026-05-10 02:53:49.976132	2026-05-10 09:11:36.43744	\N	{}
\.


--
-- Data for Name: organizations; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.organizations (id, name, slug, domain, logo_url, is_active, plan, settings, created_at, updated_at, deleted_at) FROM stdin;
00000000-0000-0000-0000-000000000001	Default Organization	default	\N	\N	t	free	{}	2026-04-28 08:08:57.218601	2026-04-28 08:08:57.218601	\N
\.


--
-- Data for Name: permissions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.permissions (id, code, name, description, created_at, updated_at, organization_id) FROM stdin;
5	user.create	Create User	Allows creating a new user in the system	2026-01-19 17:28:56.081684	2026-01-19 17:28:56.081684	00000000-0000-0000-0000-000000000001
6	user.update	Update User	Allows updating a user in the system	2026-01-19 17:29:47.059285	2026-01-19 17:29:47.059285	00000000-0000-0000-0000-000000000001
7	user.read	Read User	Allows reading a user in the system	2026-01-19 17:30:55.277256	2026-01-19 17:30:55.277256	00000000-0000-0000-0000-000000000001
8	user.delete	Delete User	Allows deleting  a user in the system	2026-01-19 17:34:15.030198	2026-01-19 17:34:15.030198	00000000-0000-0000-0000-000000000001
9	order:read	Read Orders	View customer orders	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
10	order:update	Update Orders	Update order status	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
11	order:delete	Delete Orders	Cancel/delete orders	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
12	product:create	Create Products	Add new products	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
13	product:read	Read Products	View products	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
14	product:update	Update Products	Edit existing products	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
15	product:delete	Delete Products	Remove products	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
16	category:manage	Manage Categories	Create/edit/delete categories	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
17	inventory:manage	Manage Inventory	Update stock levels	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
18	payment:read	Read Payments	View payment transactions	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
19	discount:manage	Manage Discounts	Create/edit discount codes	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
20	customer:read	Read Customers	View customer profiles	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
21	customer:update	Update Customers	Edit customer information	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
22	customer:delete	Delete Customers	Remove customer accounts	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
23	analytics:read	Read Analytics	View reports and dashboards	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
24	settings:manage	Manage Settings	Configure store settings	2026-04-28 08:08:57.548765	2026-04-28 08:08:57.548765	00000000-0000-0000-0000-000000000001
25	user:create	Create User	Create a new user account	2026-05-03 18:14:46.225389	2026-05-03 18:14:46.225389	00000000-0000-0000-0000-000000000001
26	user:read	Read User	View user profile and details	2026-05-03 18:14:46.230131	2026-05-03 18:14:46.230131	00000000-0000-0000-0000-000000000001
27	user:update	Update User	Edit user information	2026-05-03 18:14:46.232267	2026-05-03 18:14:46.232267	00000000-0000-0000-0000-000000000001
28	user:delete	Delete User	Permanently delete a user	2026-05-03 18:14:46.234477	2026-05-03 18:14:46.234477	00000000-0000-0000-0000-000000000001
29	user:status	Update User Status	Change status: active / inactive / blocked	2026-05-03 18:14:46.236706	2026-05-03 18:14:46.236706	00000000-0000-0000-0000-000000000001
30	role:create	Create Role	Create a new role	2026-05-03 18:14:46.239039	2026-05-03 18:14:46.239039	00000000-0000-0000-0000-000000000001
31	role:read	Read Role	View role details	2026-05-03 18:14:46.24123	2026-05-03 18:14:46.24123	00000000-0000-0000-0000-000000000001
32	role:update	Update Role	Edit role name and description	2026-05-03 18:14:46.243361	2026-05-03 18:14:46.243361	00000000-0000-0000-0000-000000000001
33	role:delete	Delete Role	Delete a role	2026-05-03 18:14:46.245609	2026-05-03 18:14:46.245609	00000000-0000-0000-0000-000000000001
34	role:status	Update Role Status	Activate or deactivate a role	2026-05-03 18:14:46.247906	2026-05-03 18:14:46.247906	00000000-0000-0000-0000-000000000001
35	permission:create	Create Permission	Create a new permission	2026-05-03 18:14:46.250023	2026-05-03 18:14:46.250023	00000000-0000-0000-0000-000000000001
36	permission:read	Read Permission	View permission details	2026-05-03 18:14:46.252118	2026-05-03 18:14:46.252118	00000000-0000-0000-0000-000000000001
37	permission:update	Update Permission	Edit permission information	2026-05-03 18:14:46.254154	2026-05-03 18:14:46.254154	00000000-0000-0000-0000-000000000001
38	permission:delete	Delete Permission	Delete a permission	2026-05-03 18:14:46.256042	2026-05-03 18:14:46.256042	00000000-0000-0000-0000-000000000001
39	user_role:assign	Assign Role to User	Assign one or more roles to a user	2026-05-03 18:14:46.258071	2026-05-03 18:14:46.258071	00000000-0000-0000-0000-000000000001
40	user_role:remove	Remove Role from User	Remove roles from a user	2026-05-03 18:14:46.26024	2026-05-03 18:14:46.26024	00000000-0000-0000-0000-000000000001
41	user_role:read	Read User Roles	View roles assigned to a user	2026-05-03 18:14:46.262533	2026-05-03 18:14:46.262533	00000000-0000-0000-0000-000000000001
42	user_permission:assign	Assign Permission to User	Grant a permission directly to a user	2026-05-03 18:14:46.264438	2026-05-03 18:14:46.264438	00000000-0000-0000-0000-000000000001
43	user_permission:remove	Remove Permission from User	Revoke a direct permission from a user	2026-05-03 18:14:46.266497	2026-05-03 18:14:46.266497	00000000-0000-0000-0000-000000000001
44	user_permission:read	Read User Permissions	View direct permissions of a user	2026-05-03 18:14:46.268582	2026-05-03 18:14:46.268582	00000000-0000-0000-0000-000000000001
45	hierarchy:manage	Manage Hierarchy	Assign/remove parent-child user relationships	2026-05-03 18:14:46.270548	2026-05-03 18:14:46.270548	00000000-0000-0000-0000-000000000001
46	hierarchy:read	Read Hierarchy	View the user org hierarchy	2026-05-03 18:14:46.272371	2026-05-03 18:14:46.272371	00000000-0000-0000-0000-000000000001
47	report:generate	Generate Reports	Create and export system reports	2026-05-03 18:14:46.274514	2026-05-03 18:14:46.274514	00000000-0000-0000-0000-000000000001
48	audit:read	Read Audit Logs	View the system audit trail	2026-05-03 18:14:46.276538	2026-05-03 18:14:46.276538	00000000-0000-0000-0000-000000000001
49	audit:export	Export Audit Logs	Download audit log exports	2026-05-03 18:14:46.278503	2026-05-03 18:14:46.278503	00000000-0000-0000-0000-000000000001
\.


--
-- Data for Name: role_permissions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.role_permissions (id, role_id, permission_id) FROM stdin;
2	1	5
3	1	6
4	1	7
5	3	25
6	3	26
7	3	27
8	3	28
9	3	29
10	3	30
11	3	31
12	3	32
13	3	33
14	3	34
15	3	35
16	3	36
17	3	37
18	3	38
19	3	39
20	3	40
21	3	41
22	3	42
23	3	43
24	3	44
25	3	45
26	3	46
27	3	47
28	3	48
29	3	49
30	4	25
31	4	26
32	4	27
33	4	29
34	4	31
35	4	36
36	4	39
37	4	40
38	4	41
39	4	42
40	4	43
41	4	44
42	4	45
43	4	46
44	4	47
45	5	25
46	5	26
47	5	27
48	5	29
49	5	31
50	5	39
51	5	41
52	5	46
53	6	26
54	6	31
55	6	36
56	6	41
57	6	46
58	7	26
59	7	31
60	7	36
61	7	41
62	7	44
63	7	46
64	8	26
65	8	31
66	8	36
67	8	48
68	8	49
69	8	47
70	9	26
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.roles (id, name, code, description, is_active, created_at, updated_at, organization_id) FROM stdin;
1	Admin	ADMIN	System administrator role	t	2026-01-19 17:22:52.259202	2026-01-19 17:22:52.259202	00000000-0000-0000-0000-000000000001
2	Sales Manager	SALES_MANAGER	Manages sales team	t	2026-01-19 17:23:52.420361	2026-01-19 17:23:52.420361	00000000-0000-0000-0000-000000000001
3	Super Admin	super_admin	Full unrestricted system access	t	2026-05-03 18:14:46.210368	2026-05-03 18:14:46.210368	00000000-0000-0000-0000-000000000001
4	Manager	manager	Manages users, roles and assignments	t	2026-05-03 18:14:46.214669	2026-05-03 18:14:46.214669	00000000-0000-0000-0000-000000000001
5	HR Manager	hr_manager	Handles user onboarding and status changes	t	2026-05-03 18:14:46.216162	2026-05-03 18:14:46.216162	00000000-0000-0000-0000-000000000001
6	Support Agent	support_agent	Read access to users and tickets	t	2026-05-03 18:14:46.217773	2026-05-03 18:14:46.217773	00000000-0000-0000-0000-000000000001
7	Viewer	viewer	Read-only access across the system	t	2026-05-03 18:14:46.219146	2026-05-03 18:14:46.219146	00000000-0000-0000-0000-000000000001
8	Auditor	auditor	Access to audit logs and reports	t	2026-05-03 18:14:46.220492	2026-05-03 18:14:46.220492	00000000-0000-0000-0000-000000000001
9	Guest	guest	Minimal temporary access	f	2026-05-03 18:14:46.222083	2026-05-03 18:14:46.222083	00000000-0000-0000-0000-000000000001
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.schema_migrations (version, dirty) FROM stdin;
16	f
\.


--
-- Data for Name: service_clients; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.service_clients (id, organization_id, name, client_id, client_secret_hash, allowed_scopes, is_active, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: user_hierarchy; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_hierarchy (id, parent_user_id, child_user_id, created_at) FROM stdin;
1	c2bfb817-613d-4ec8-be89-52d7367171e6	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c	2026-01-19 18:30:41.21042
2	7bd17c85-5ded-4feb-a692-76422bc5dab2	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	2026-05-03 18:20:50.510208
3	7bd17c85-5ded-4feb-a692-76422bc5dab2	7dbbc23d-197f-464b-9288-5eb6dfde8e61	2026-05-03 18:20:50.512723
4	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	14f9cca5-62d3-42ce-b47f-5cedea7c1622	2026-05-03 18:20:50.51416
5	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	f1071943-d948-458e-97bb-05caa03c447b	2026-05-03 18:20:50.51568
6	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	2c05e61c-c6e0-4de3-8a5c-44a8e22ac4a8	2026-05-03 18:20:50.517674
7	14f9cca5-62d3-42ce-b47f-5cedea7c1622	ea2b32dd-0c0b-4278-9002-260fa8548de3	2026-05-03 18:20:50.519258
8	f1071943-d948-458e-97bb-05caa03c447b	c37f81aa-79bf-4541-a36a-97cc94d92d8e	2026-05-03 18:20:50.520727
9	2c05e61c-c6e0-4de3-8a5c-44a8e22ac4a8	94fbc977-f19c-4723-9793-44975d51e237	2026-05-03 18:20:50.522376
10	7dbbc23d-197f-464b-9288-5eb6dfde8e61	9c542918-66cb-409d-bf53-22a8c95bb608	2026-05-03 18:20:50.523823
11	7dbbc23d-197f-464b-9288-5eb6dfde8e61	63ac4ac0-abd4-4f64-81d3-8c6e0a034539	2026-05-03 18:20:50.525021
12	7dbbc23d-197f-464b-9288-5eb6dfde8e61	3ad648cb-1cf9-452f-8eeb-7aa38ebf5eb5	2026-05-03 18:20:50.526291
13	7dbbc23d-197f-464b-9288-5eb6dfde8e61	4e35d96d-0cac-4004-95f3-0a22f2011c15	2026-05-03 18:20:50.527613
\.


--
-- Data for Name: user_permissions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_permissions (id, user_id, permission_id, allow) FROM stdin;
1	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c	5	t
2	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c	6	t
3	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c	7	t
4	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c	8	t
5	2c05e61c-c6e0-4de3-8a5c-44a8e22ac4a8	48	t
6	94fbc977-f19c-4723-9793-44975d51e237	47	t
7	f1071943-d948-458e-97bb-05caa03c447b	27	t
8	3ad648cb-1cf9-452f-8eeb-7aa38ebf5eb5	26	t
9	3ad648cb-1cf9-452f-8eeb-7aa38ebf5eb5	31	t
\.


--
-- Data for Name: user_roles; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_roles (id, user_id, role_id, assigned_at) FROM stdin;
1	b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c	1	2026-01-19 12:07:01.905637
2	3c77374a-29b2-4bc3-93dd-011277c861bd	1	2026-04-28 07:54:06.171479
3	d1a7e67d-90ff-42b6-abe5-9e54bbc4624f	2	2026-04-28 07:54:06.550499
4	7bd17c85-5ded-4feb-a692-76422bc5dab2	3	2026-02-03 01:20:49.612293
5	0ae42eb2-8765-4fe6-8b09-2273a7736fe4	4	2026-02-13 01:20:49.683264
6	7dbbc23d-197f-464b-9288-5eb6dfde8e61	4	2026-02-18 01:20:49.752948
7	7dbbc23d-197f-464b-9288-5eb6dfde8e61	5	2026-02-18 01:20:49.752948
8	14f9cca5-62d3-42ce-b47f-5cedea7c1622	5	2026-03-05 01:20:49.817283
9	ea2b32dd-0c0b-4278-9002-260fa8548de3	5	2026-03-10 01:20:49.879795
10	f1071943-d948-458e-97bb-05caa03c447b	6	2026-03-20 01:20:49.944328
11	c37f81aa-79bf-4541-a36a-97cc94d92d8e	6	2026-03-25 01:20:50.007963
12	9c542918-66cb-409d-bf53-22a8c95bb608	8	2026-04-04 01:20:50.070507
13	2c05e61c-c6e0-4de3-8a5c-44a8e22ac4a8	7	2026-02-23 01:20:50.133182
14	94fbc977-f19c-4723-9793-44975d51e237	7	2026-02-28 01:20:50.196339
15	69f8f2d7-8ed4-4fdc-b7e9-715dac5f54aa	9	2026-04-24 01:20:50.318935
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, name, email, password_hash, phone, status, email_verified_at, last_login_at, created_at, updated_at, deleted_at, avatar_url, organization_id, user_type, metadata) FROM stdin;
b6f4612b-4cf6-4a2e-87e9-ee9e4ccea95c	Moshfiqur Rahman	moshfiqur.rahman@example.com	$2a$12$LhngxRolZOUMg1TPCDpRE.jJZW2rK9NWeRZE0tqd24l7PuvaiH98G	+8801712345678	active	\N	\N	2026-01-19 17:39:52.45183	2026-01-19 17:39:52.45183	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
c2bfb817-613d-4ec8-be89-52d7367171e6	Nasty Munna	nasty.munna@example.com	$2a$12$WbZFKePM6ysTIkEMVihALeTTNNKfALZLOuyXFFBSqG5IIF.Fu9YFC	+8801915632646	active	\N	\N	2026-01-19 17:42:16.620691	2026-01-19 17:42:16.620691	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
8a343305-2cdd-4883-bb66-2765aa2fb2f5	Saif Chowdhury	saif.chowdhury@example.com	$2a$12$m5cwLj.yVy8ceawxuNVnzOmqeKatauR/ndBUWUAbSPSzY1FYiQLmW	+8801915632646	active	\N	\N	2026-01-21 13:31:13.3203	2026-01-21 13:31:13.3203	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
3c77374a-29b2-4bc3-93dd-011277c861bd	Admin User	admin@example.com	$2a$12$q3ksXodb9GV62wxeOY1CG.Cbsnm6kuE/I/haHXERmeoNJ5h97/6eO	+1234567890	active	\N	\N	2026-04-28 13:52:50.245589	2026-04-28 13:52:50.245589	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
d1a7e67d-90ff-42b6-abe5-9e54bbc4624f	John Doe	john@example.com	$2a$12$FNpLiTztRspieYjLLcVggOZCd/yCHxOZltsXBO6kN/9obWMcCciEq	+0987654321	active	\N	\N	2026-04-28 13:52:50.632345	2026-04-28 13:52:50.632345	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
262f0064-eaed-4104-a761-6301e717f8aa	Jane Smith	jane@example.com	$2a$12$67M8pzu..z0ddR8KGM4AoO9BdjZAhDmu5E4T1IVhikr/QcGOSfh6K	\N	active	\N	\N	2026-04-28 13:52:51.012968	2026-04-28 13:52:51.012968	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
7bd17c85-5ded-4feb-a692-76422bc5dab2	Super Admin	admin@ums.com	$2a$10$FRYOdepkr5IzvRL4cMrXCufSMz0Z88Tab/lzIa0aYL2w2y2rFPa2a	+8801700000001	active	\N	\N	2026-02-03 00:20:49.612293	2026-02-03 00:20:49.612293	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
0ae42eb2-8765-4fe6-8b09-2273a7736fe4	Alice Rahman	alice@ums.com	$2a$10$ZK/Q7TiNHcE1vJanuV9lmOakOf2KnG3JcD58To3GVad0ddBpGO4cm	+8801700000002	active	\N	\N	2026-02-13 00:20:49.683264	2026-02-13 00:20:49.683264	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
7dbbc23d-197f-464b-9288-5eb6dfde8e61	David Kim	david@ums.com	$2a$10$Zyv21nrrXUsj3Bogr4ETe.e4mvFfgBb9qlPs/t2Y5MrOFxWwSUnL.	+8801700000010	active	\N	\N	2026-02-18 00:20:49.752948	2026-02-18 00:20:49.752948	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
14f9cca5-62d3-42ce-b47f-5cedea7c1622	Priya Sharma	priya@ums.com	$2a$10$1lApOsabUG6GQKWA9fL2F.Si39ZMKNLP./G7ipJ7Hoq9pMAAHHyJC	+8801700000005	active	\N	\N	2026-03-05 00:20:49.817283	2026-03-05 00:20:49.817283	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
ea2b32dd-0c0b-4278-9002-260fa8548de3	Omar Hassan	omar@ums.com	$2a$10$Q4f16SHug0Vw43avNQimcuLaWevqxFv/.CVsF2a3MTwysEY.CepQ2	+8801700000006	active	\N	\N	2026-03-10 00:20:49.879795	2026-03-10 00:20:49.879795	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
f1071943-d948-458e-97bb-05caa03c447b	Sophie Turner	sophie@ums.com	$2a$10$VhpJPrxEigaG140DMQ1TKOF37PAYfTzi4ALGhQDMWvOjRY.UMq2xK	+8801700000007	active	\N	\N	2026-03-20 00:20:49.944328	2026-03-20 00:20:49.944328	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
c37f81aa-79bf-4541-a36a-97cc94d92d8e	James Lee	james@ums.com	$2a$10$cJuwkfDRqqHB3aLJUqB8Jekrp0c8QBdrxLLVL14fVyLI6DRi6I3Oq	+8801700000008	active	\N	\N	2026-03-25 00:20:50.007963	2026-03-25 00:20:50.007963	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
9c542918-66cb-409d-bf53-22a8c95bb608	Nina Patel	nina@ums.com	$2a$10$Zwu//4X83RWJZdB6pNXq.eKsipJjOfJOfuJ3/EGn8UPOqtZtq2KAe	+8801700000009	active	\N	\N	2026-04-04 00:20:50.070507	2026-04-04 00:20:50.070507	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
2c05e61c-c6e0-4de3-8a5c-44a8e22ac4a8	Bob Chen	bob@ums.com	$2a$10$teMiVTo.8rbi9koh6p8aku.RWtUb/wIsUs6ZgZMMdWAUmS.J.vNg2	+8801700000003	active	\N	\N	2026-02-23 00:20:50.133182	2026-02-23 00:20:50.133182	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
94fbc977-f19c-4723-9793-44975d51e237	Charlie Davis	charlie@ums.com	$2a$10$tY7fs4IKJh9aqnL94uwky.J382UcDUY2tvoluYRRQqjDwhxvhG5Vm	+8801700000004	active	\N	\N	2026-02-28 00:20:50.196339	2026-02-28 00:20:50.196339	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
63ac4ac0-abd4-4f64-81d3-8c6e0a034539	Eve Wilson	eve@ums.com	$2a$10$kV/zc9PAgst0W34AP3NbIug.64UXWgzCsnuSDI5ie/rZAOUZTXjaK	+8801700000011	inactive	\N	\N	2026-04-14 00:20:50.258969	2026-04-14 00:20:50.258969	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
69f8f2d7-8ed4-4fdc-b7e9-715dac5f54aa	Frank Bruno	frank@ums.com	$2a$10$sx6ceHKjtm7DpRDMPt96Re12m1eojpIiKOHRrOs/fW9t.IO2zWHHS	+8801700000012	blocked	\N	\N	2026-04-24 00:20:50.318935	2026-04-24 00:20:50.318935	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
3ad648cb-1cf9-452f-8eeb-7aa38ebf5eb5	Grace Liu	grace@ums.com	$2a$10$dlPbIyp5VgMTuTij/cJGAuUAOfp0CuElSd5thn2v5jpK.yLnpjIVm	+8801700000013	active	\N	\N	2026-04-29 00:20:50.380046	2026-04-29 00:20:50.380046	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
4e35d96d-0cac-4004-95f3-0a22f2011c15	Henry Park	henry@ums.com	$2a$10$yoOI36NOk6kJ30br2BllJ.vin.KIns1w6T2s4jJkaeoH4o0my0ku2	+8801700000014	active	\N	\N	2026-05-02 00:20:50.441391	2026-05-02 00:20:50.441391	\N	\N	00000000-0000-0000-0000-000000000001	member	{}
f869a9b9-b625-49f2-ad34-424548e2d9ee	MASTERTONY	tony@gmail.com	cb03b96e-81ab-456e-9fe3-cd0a6a410998	\N	active	\N	\N	2026-05-18 08:59:34.297448	2026-05-18 08:59:34.297448	\N	\N	00000000-0000-0000-0000-000000000001	guest	{}
78fad8d4-56fb-4c80-be8f-4c83c17ba737	Test	test@test.com	c24a901b-2e49-46c4-98f5-4fb9bebe8506	\N	active	\N	\N	2026-05-18 12:52:07.570674	2026-05-18 12:52:07.570674	\N	\N	00000000-0000-0000-0000-000000000001	guest	{}
\.


--
-- Name: audit_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.audit_logs_id_seq', 18, true);


--
-- Name: permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.permissions_id_seq', 74, true);


--
-- Name: role_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.role_permissions_id_seq', 136, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.roles_id_seq', 16, true);


--
-- Name: user_hierarchy_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_hierarchy_id_seq', 13, true);


--
-- Name: user_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_permissions_id_seq', 9, true);


--
-- Name: user_roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_roles_id_seq', 15, true);


--
-- Name: audit_logs audit_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT audit_logs_pkey PRIMARY KEY (id);


--
-- Name: login_sessions login_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.login_sessions
    ADD CONSTRAINT login_sessions_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_slug_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.organizations
    ADD CONSTRAINT organizations_slug_key UNIQUE (slug);


--
-- Name: permissions permissions_code_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_code_key UNIQUE (code);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: role_permissions role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permissions_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: service_clients service_clients_client_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.service_clients
    ADD CONSTRAINT service_clients_client_id_key UNIQUE (client_id);


--
-- Name: service_clients service_clients_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.service_clients
    ADD CONSTRAINT service_clients_pkey PRIMARY KEY (id);


--
-- Name: role_permissions uq_role_permission; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT uq_role_permission UNIQUE (role_id, permission_id);


--
-- Name: roles uq_roles_code_org; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uq_roles_code_org UNIQUE (code, organization_id);


--
-- Name: roles uq_roles_name_org; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uq_roles_name_org UNIQUE (name, organization_id);


--
-- Name: user_hierarchy uq_user_hierarchy; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_hierarchy
    ADD CONSTRAINT uq_user_hierarchy UNIQUE (parent_user_id, child_user_id);


--
-- Name: user_permissions uq_user_permission; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT uq_user_permission UNIQUE (user_id, permission_id);


--
-- Name: user_roles uq_user_role; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT uq_user_role UNIQUE (user_id, role_id);


--
-- Name: users uq_users_email_org; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uq_users_email_org UNIQUE (email, organization_id);


--
-- Name: user_hierarchy user_hierarchy_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_hierarchy
    ADD CONSTRAINT user_hierarchy_pkey PRIMARY KEY (id);


--
-- Name: user_permissions user_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT user_permissions_pkey PRIMARY KEY (id);


--
-- Name: user_roles user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (id);


--
-- Name: users users_org_email_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_org_email_unique UNIQUE (organization_id, email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_audit_logs_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_created_at ON public.audit_logs USING btree (created_at);


--
-- Name: idx_audit_logs_entity; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_entity ON public.audit_logs USING btree (entity, entity_id);


--
-- Name: idx_audit_logs_org_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_org_id ON public.audit_logs USING btree (organization_id);


--
-- Name: idx_audit_logs_service; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_service ON public.audit_logs USING btree (service);


--
-- Name: idx_audit_logs_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_user_id ON public.audit_logs USING btree (user_id);


--
-- Name: idx_login_sessions_logged_in_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_login_sessions_logged_in_at ON public.login_sessions USING btree (logged_in_at);


--
-- Name: idx_login_sessions_logged_out_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_login_sessions_logged_out_at ON public.login_sessions USING btree (logged_out_at);


--
-- Name: idx_login_sessions_org_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_login_sessions_org_id ON public.login_sessions USING btree (organization_id);


--
-- Name: idx_login_sessions_refresh_token; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_login_sessions_refresh_token ON public.login_sessions USING btree (refresh_token_hash);


--
-- Name: idx_login_sessions_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_login_sessions_user_id ON public.login_sessions USING btree (user_id);


--
-- Name: idx_organizations_deleted; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_organizations_deleted ON public.organizations USING btree (deleted_at);


--
-- Name: idx_organizations_domain; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_organizations_domain ON public.organizations USING btree (domain);


--
-- Name: idx_organizations_slug; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_organizations_slug ON public.organizations USING btree (slug);


--
-- Name: idx_permissions_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_permissions_code ON public.permissions USING btree (code);


--
-- Name: idx_permissions_organization_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_permissions_organization_id ON public.permissions USING btree (organization_id);


--
-- Name: idx_role_permissions_permission_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_role_permissions_permission_id ON public.role_permissions USING btree (permission_id);


--
-- Name: idx_role_permissions_role_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_role_permissions_role_id ON public.role_permissions USING btree (role_id);


--
-- Name: idx_roles_org_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_roles_org_id ON public.roles USING btree (organization_id);


--
-- Name: idx_service_clients_client_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_service_clients_client_id ON public.service_clients USING btree (client_id);


--
-- Name: idx_service_clients_org_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_service_clients_org_id ON public.service_clients USING btree (organization_id);


--
-- Name: idx_user_hierarchy_child; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_hierarchy_child ON public.user_hierarchy USING btree (child_user_id);


--
-- Name: idx_user_hierarchy_parent; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_hierarchy_parent ON public.user_hierarchy USING btree (parent_user_id);


--
-- Name: idx_user_permissions_permission_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_permissions_permission_id ON public.user_permissions USING btree (permission_id);


--
-- Name: idx_user_permissions_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_permissions_user_id ON public.user_permissions USING btree (user_id);


--
-- Name: idx_user_roles_role_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_roles_role_id ON public.user_roles USING btree (role_id);


--
-- Name: idx_user_roles_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_roles_user_id ON public.user_roles USING btree (user_id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: idx_users_org_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_org_id ON public.users USING btree (organization_id);


--
-- Name: idx_users_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_status ON public.users USING btree (status);


--
-- Name: idx_users_user_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_user_type ON public.users USING btree (user_type);


--
-- Name: audit_logs audit_logs_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT audit_logs_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON DELETE SET NULL;


--
-- Name: audit_logs fk_audit_logs_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT fk_audit_logs_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: login_sessions fk_login_sessions_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.login_sessions
    ADD CONSTRAINT fk_login_sessions_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: role_permissions fk_role_permissions_permission; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: role_permissions fk_role_permissions_role; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;


--
-- Name: user_hierarchy fk_user_hierarchy_child; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_hierarchy
    ADD CONSTRAINT fk_user_hierarchy_child FOREIGN KEY (child_user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_hierarchy fk_user_hierarchy_parent; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_hierarchy
    ADD CONSTRAINT fk_user_hierarchy_parent FOREIGN KEY (parent_user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_permissions fk_user_permissions_permission; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT fk_user_permissions_permission FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: user_permissions fk_user_permissions_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT fk_user_permissions_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_roles fk_user_roles_role; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;


--
-- Name: user_roles fk_user_roles_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: login_sessions login_sessions_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.login_sessions
    ADD CONSTRAINT login_sessions_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON DELETE CASCADE;


--
-- Name: roles roles_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON DELETE CASCADE;


--
-- Name: service_clients service_clients_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.service_clients
    ADD CONSTRAINT service_clients_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON DELETE CASCADE;


--
-- Name: users users_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict CgAcI4osYmb8GbTMQez4NCENw49tWbBXWc0bV12RepDEOmuBCXYW5oM03OVkBpv

