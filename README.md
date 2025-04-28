# vibecheckr

PocketBase backend for [vibe chuck](www.github.com/we-be/vibe-check)

## Overview

vibecheckr is the backend service for the VIBE CHUCK social media app. It provides:
- User authentication
- Image storage
- Post management
- Automated ranking of holiday decoration posts
 - Event management

## Entities

The PocketBase backend defines the following collections (entities):

- **users**
  - Description: Built-in PocketBase authentication and user profiles collection.
  - Fields:
    - `id` (string): Unique identifier.
    - `email` (string): User email address.
    - `username` (string): Display name or username.
    - `avatar` (file): Uploaded user avatar (optional).
    - `created` (datetime): Record creation timestamp.
    - `updated` (datetime): Record last update timestamp.

- **events**
  - Description: Holiday events that users can contribute posts to.
  - Fields:
    - `id` (string): Unique identifier.
    - `displayName` (string): Name of the event (e.g., "Christmas 2024").
    - `description` (text): Event description.
    - `start` (datetime): Event start date.
    - `end` (datetime): Event end date.
    - `location` (string): Event location.

- **posts**
  - Description: User-submitted holiday decoration posts.
  - Fields:
    - `id` (string): Unique identifier.
    - `title` (string): Title of the post.
    - `description` (text): Decoration description.
    - `imgs` (array of files): Uploaded image files.
    - `rank` (integer): Current post ranking within the event (auto-calculated).
    - `votes` (integer): Number of votes/likes received.
    - `event` (relation): Reference to the parent event (`events` collection).
    - `op` (relation): Reference to the original poster (`users` collection).
    - `created` (datetime): Record creation timestamp.

## Getting Started

To start the backend server:

```bash
# Run the server
go run cmd/main.go

# Default admin UI is available at http://127.0.0.1:8090/_/
```

The server includes a cron job that recalculates post rankings every 10 minutes.

## Setting Up Collections

If you're starting with a fresh PocketBase instance or need to recreate collections:

1. Start the PocketBase server first: `go run cmd/main.go`
2. In a separate terminal, run the setup script:

```bash
cd scripts
./run-setup.sh
```

This script will:
- Create the required collections (events and posts)
- Add a default "Christmas 2024" event
- Set up proper relations between collections
