# Custom APIs for Glance Dashboard

A collection of custom APIs designed to extend your Glance dashboard with personalized data feeds and integrations.

## Available APIs

### 🎌 [AniList API](./anilist)

Fetch anime information from your AniList account. Display upcoming episodes and your current watching/planning list on your dashboard.

- **Endpoints**: `/available`, `/upcoming`
- **Features**: Episode tracking, airing dates, watch progress
- **Setup**: Requires AniList OAuth credentials

## Structure

Each API is organized in its own subdirectory with:
- `Dockerfile` - Container configuration
- `.env.example` - Environment variables template
- `src/` - Application source code
- `README.MD` - Detailed setup and usage documentation

## Adding New APIs

To add a new custom API:

1. Create a new subdirectory with your API name
2. Add a `Dockerfile` and `package.json`
3. Create a `.env.example` with required environment variables
4. Document setup instructions in a `README.MD`
5. Update this file with an overview of your new API