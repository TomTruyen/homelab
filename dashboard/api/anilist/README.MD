# AniList API

This API provides endpoints to fetch anime information from your AniList account, including currently watching/planning anime and upcoming episodes.

## What It Does

This service integrates with [AniList](https://anilist.co/) to display:
- **Currently Watching**: Anime you're actively watching with progress tracking
- **Available Episodes**: Anime you're currently watching or planning to watch, sorted by progress and airing date
- **Upcoming Episodes**: Upcoming episodes for your anime list with airing times

The API uses AniList's OAuth 2.0 client credentials flow to securely authenticate and fetch your personal anime list data.

## Setup

### 1. Create AniList OAuth Application

1. Go to https://anilist.co/settings/developer
2. Create a new OAuth application:
   - **Name**: `Homelab Dashboard` (or your preferred name)
   - **Redirect URL**: `http://localhost:3000/callback` (or your service URL)
3. Copy the **Client ID** and **Client Secret**

### 2. Configure Environment Variables

Rename and configure the environment file:

```bash
mv .env.example .env
```

| Variable | Required | Description |
|----------|----------|-------------|
| `ANILIST_USERNAME` | Yes | Your AniList username (the one displayed on your profile) |
| `ANILIST_CLIENT_ID` | Yes | OAuth Client ID from your AniList developer application |
| `ANILIST_CLIENT_SECRET` | Yes | OAuth Client Secret from your AniList developer application |

### 3. Start the Service

```bash
docker compose up -d --build --force-recreate
```

## API Endpoints

### GET `/watching`

Returns a list of anime you're currently watching with episode progress information.

**Response:**
```json
{
  "items": [
    {
      "title": "Anime Title",
      "nextEpisode": 4,
      "totalEpisodes": 12,
      "url": "https://anilist.co/anime/12345"
    }
  ]
}
```

### GET `/available`

Returns a list of anime you're currently watching or planning to watch, with episode progress information.

**Response:**
```json
{
  "items": [
    {
      "id": 12345,
      "title": "Anime Title",
      "nextEpisode": 5,
      "totalEpisodes": 12,
      "watched": 4,
      "url": "https://anilist.co/anime/12345",
      "status": "RELEASING",
      "airingAt": 1733702400,
      "formattedAiring": "Dec 09, 10:00"
    }
  ]
}
```

### GET `/upcoming`

Returns upcoming episodes from your anime list, sorted by airing date.

**Response:**
```json
{
  "items": [
    {
      "id": 12345,
      "title": "Anime Title",
      "episode": 5,
      "airingAt": 1733702400,
      "formattedAiring": "Dec 09, 10:00",
      "status": "RELEASING",
      "url": "https://anilist.co/anime/12345"
    }
  ]
}
```

## (Optional) Using in Glance Dashboard

You can integrate this API into your Glance dashboard using the `custom-api` widget to display your anime lists.

### Basic Configuration

Add the following widgets to your Glance page configuration (e.g., `config/pages/watchlist.yml`):

```yaml
- type: custom-api
  title: Upcoming Anime
  cache: 1h
  url: http://192.168.0.150:3000/upcoming
  template: |
    <ul class="list list-gap-10 collapsible-container" data-collapse-after="5">
      {{ range .JSON.Array "items" }}
        <li>
          <a class="size-h4 color-highlight block line-clamp-2" href="{{ .String "url" }}" target="_blank">
            {{ .String "title" }}
          </a>
          <ul class="list-vertical-text">
            {{ if .Int "episode" }}
              <li>Episode {{ .Int "episode" }}</li>
            {{ end }}
            <li>Airing: {{ .String "formattedAiring" }}</li>
            {{ $status := .String "status" }}
            {{ if ne $status "NOT_YET_RELEASED" }}
              <li class="color-positive">{{ .String "status" }}</li>
            {{ else }}
              <li class="color-negative">{{ .String "status" }}</li>
            {{ end }}
          </ul>
        </li>
      {{ end }}
    </ul>

- type: custom-api
  title: Available Anime
  cache: 1h
  url: http://192.168.0.150:3000/available
  template: |
    <ul class="list list-gap-10 collapsible-container" data-collapse-after="5">
      {{ range .JSON.Array "items" }}
        <li>
          <a class="size-h4 color-highlight block line-clamp-2" href="{{ .String "url" }}" target="_blank">
            {{ .String "title" }}
          </a>
          <ul class="list-vertical-text">
            <li>Progress: {{ .Int "watched" }}/{{ .Int "totalEpisodes" }}</li>
            {{ if .Int "nextEpisode" }}
              <li>Next: Episode {{ .Int "nextEpisode" }}</li>
            {{ end }}
            <li>Airing: {{ .String "formattedAiring" }}</li>
            {{ $status := .String "status" }}
            {{ if eq $status "FINISHED" }}
              <li class="color-positive">{{ .String "status" }}</li>
            {{ else }}
              <li class="color-negative">{{ .String "status" }}</li>
            {{ end }}
          </ul>
        </li>
      {{ end }}
    </ul>
```

### Configuration Notes

- **`url`**: Update the IP address and port to match your deployment (default: `http://192.168.0.150:3000`)
- **`cache`**: Set how often the widget should refresh data (e.g., `1h`, `30m`)
- **`template`**: The Glance template language allows you to customize the display of anime information
  - `{{ .String "title" }}` - Anime title
  - `{{ .Int "watched" }}` / `{{ .Int "totalEpisodes" }}` - Episode progress
  - `{{ .String "formattedAiring" }}` - Formatted airing time
  - `{{ .String "status" }}` - Anime status (RELEASING, FINISHED, etc.)