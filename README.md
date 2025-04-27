# vibecheckr

PocketBase backend for [vibe chuck](www.github.com/we-be/vibe-check)

## Overview

vibecheckr is the backend service for the VIBE CHUCK social media app. It provides:
- User authentication
- Image storage
- Post management
- Automated ranking of holiday decoration posts
- Event management

## Getting Started

To start the backend server:

```bash
# Run the server
go run main.go

# Default admin UI is available at http://127.0.0.1:8090/_/
```

The server includes a cron job that recalculates post rankings every 10 minutes.

## Setting Up Collections

If you're starting with a fresh PocketBase instance or need to recreate collections:

1. Start the PocketBase server first: `go run main.go`
2. In a separate terminal, run the setup script:

```bash
cd scripts
./run-setup.sh
```

This script will:
- Create the required collections (events and posts)
- Add a default "Christmas 2024" event
- Set up proper relations between collections
