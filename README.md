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
go run cmd/main.go serve --dev

# Default admin UI is available at http://127.0.0.1:8090/_/
```

The server includes a cron job that recalculates post rankings every 10 minutes.
