import fetch from 'node-fetch';
import dotenv from 'dotenv';
import express from 'express';
dotenv.config();

const app = express();
const PORT = process.env.PORT || 3000;

const { ANILIST_CLIENT_ID, ANILIST_CLIENT_SECRET } = process.env;

// 1️⃣ Function to get dynamic access token
async function getAccessToken() {
  const params = new URLSearchParams();
  params.append('grant_type', 'client_credentials');
  params.append('client_id', ANILIST_CLIENT_ID);
  params.append('client_secret', ANILIST_CLIENT_SECRET);

  const response = await fetch('https://anilist.co/api/v2/oauth/token', {
    method: 'POST',
    body: params
  });

  if (!response.ok) {
    throw new Error(`Failed to get access token: ${response.status}`);
  }

  const data = await response.json();
  return data.access_token;
}

// 2️⃣ Fetch upcoming episodes
async function fetchUpcoming() {
  const token = await getAccessToken(); // <-- get a fresh token dynamically

  const query = `
    query {
      MediaListCollection(userName: "${process.env.ANILIST_USERNAME}", type: ANIME, status_in: [CURRENT, PLANNING]) {
        lists {
          entries {
            media {
              id
              title { romaji english }
              nextAiringEpisode { id episode airingAt }
              status
            }
            progress
          }
        }
      }
    }
  `;

  const response = await fetch('https://graphql.anilist.co', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify({ query }),
  });

  const data = await response.json();

  return data.data.MediaListCollection.lists
  .flatMap(list => list.entries)
  .map(entry => {
    const media = entry.media;
    let airingAt = media.nextAiringEpisode?.airingAt;
    let formattedAiring = "Unknown";

    if (airingAt && typeof airingAt === "number") {
      const date = new Date(airingAt * 1000);
      formattedAiring = date.toLocaleString("en-US", {
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    }

    if (!media.nextAiringEpisode) {
      switch (media.status) {
        case "NOT_YET_RELEASED":
          formattedAiring = "Unknown";
          break;
        default:
          return null;
      }
    }

    return {
      id: media.id,
      title: media.title.english || media.title.romaji,
      episode: media.nextAiringEpisode?.episode,
      airingAt: airingAt,
      formattedAiring,
      status: media.status,
      url: `https://anilist.co/anime/${media.id}`
    };
  })
  .filter(Boolean)
  .sort((a, b) => {
    const aUnknown = a.airingAt === "Unknown" || a.airingAt === undefined;
    const bUnknown = b.airingAt === "Unknown" || b.airingAt === undefined;

    if (aUnknown && bUnknown) return 0;   // both unknown, keep original order
    if (aUnknown) return 1;               // a unknown, b known → a after b
    if (bUnknown) return -1;              // b unknown, a known → b after a

    // both known, sort numerically
    return a.airingAt - b.airingAt;
  });
}

// 3️⃣ API endpoints
app.get('/upcoming', async (req, res) => {
  try {
    const upcoming = await fetchUpcoming();
    res.json({ items: upcoming });
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Failed to fetch upcoming episodes' });
  }
});

// 4️⃣ Start server
app.listen(PORT, () => {
  console.log(`AniList API running on http://localhost:${PORT}`);
});