import fetch from 'node-fetch';
import { getAccessToken } from './token.js';

export async function fetchWatching() {
    const token = await getAccessToken();

    const query = `
    query {
      MediaListCollection(
        userName: "${process.env.ANILIST_USERNAME}",
        type: ANIME,
        status: CURRENT
      ) {
        lists {
          entries {
            media {
              id
              title { romaji english }
              episodes
            }
            progress
          }
        }
      }
    }
    `;

    const response = await fetch("https://graphql.anilist.co", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({ query }),
    });

    const data = await response.json();

    return data.data.MediaListCollection.lists
        .flatMap(list => list.entries)
        .map(entry => {
            const media = entry.media;

            return {
                title: media.title.english || media.title.romaji,
                watched: (entry.progress || 0) + 1, // Next episode to watch
                totalEpisodes: media.episodes || null,
                url: `https://anilist.co/anime/${media.id}`
            };
        })
        .sort((a, b) => b.watched - a.watched); // ← DESC sort
}
