import fetch from 'node-fetch';
import { getAccessToken } from "./token.js";

export async function fetchUpcoming() {
  const token = await getAccessToken(); // <-- get a fresh token dynamically

  console.log(token);

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
            status
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
      status: entry.status, // Personal Status. `media.status` is the overall anime status.
      url: `https://anilist.co/anime/${media.id}`,
      watchUrl: `https://anikai.to/browser?keyword=${media.title.english || media.title.romaji}`,
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