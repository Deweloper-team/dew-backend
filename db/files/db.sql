CREATE extension IF NOT EXISTS "uuid-ossp";

CREATE TABLE "files" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	type TEXT NOT NULL CHECK (char_length(type) <= 50),
	url TEXT NOT NULL CHECK (char_length(url) <= 255),
	user_upload TEXT NOT NULL CHECK (char_length(user_upload) <= 100),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "roles"
(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "code" TEXT NOT NULL CHECK (char_length(code) <= 50),
    "name" TEXT NOT NULL CHECK (char_length(name) <= 100),
    "status" TEXT CHECK (char_length(status) <= 50),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "users"
(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" TEXT NOT NULL CHECK (char_length(name) <= 100),
    "email" TEXT NOT NULL CHECK (char_length(email) <= 100),
    "bio" TEXT CHECK (char_length(bio) <= 500),
    "profile_image_id" uuid REFERENCES files(id),
    "gender" TEXT NOT NULL CHECK (char_length(gender) <= 50),
    "address" TEXT CHECK (char_length(address) <= 500),
    "status" TEXT NOT NULL CHECK (char_length(status) <= 50),
    "details" JSONB CHECK (char_length(details) <= 5000),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);