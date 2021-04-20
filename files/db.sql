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
    "email" TEXT NOT NULL CHECK (char_length(email) <= 100),
    "name" TEXT NOT NULL CHECK (char_length(name) <= 100),
    "profile_image_id" uuid REFERENCES files(id),
    "gender" TEXT NOT NULL CHECK (char_length(gender) <= 50),
    "phone" TEXT NOT NULL CHECK (char_length(phone) <= 100),
    "city_id" uuid REFERENCES cities(id),
    "address" TEXT NOT NULL CHECK (char_length(address) <= 500),
    "status" TEXT NOT NULL CHECK (char_length(status) <= 50),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "users"
(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "code" TEXT NOT NULL CHECK (char_length(code) <= 50),
    "email" TEXT NOT NULL CHECK (char_length(email) <= 100),
    "name" TEXT NOT NULL CHECK (char_length(name) <= 100),
    "profile_image_id" uuid REFERENCES files(id),
    "gender" TEXT NOT NULL CHECK (char_length(gender) <= 50),
    "phone" TEXT NOT NULL CHECK (char_length(phone) <= 100),
    "city_id" uuid REFERENCES cities(id),
    "address" TEXT NOT NULL CHECK (char_length(address) <= 500),
    "status" TEXT NOT NULL CHECK (char_length(status) <= 50),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);