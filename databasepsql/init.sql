CREATE DATABASE TodoDB;
\c TodoDB;

CREATE TABLE public."Items"
(
    "Id" text NOT NULL,
    "Title" text,
    "Description" text,
    "Created" date,
    "IsDone" boolean,
    "Image" text,
    "Color" text,
    PRIMARY KEY ("Id")
);

ALTER TABLE IF EXISTS public."Items"
    OWNER to root;