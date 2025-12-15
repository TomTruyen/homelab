import fetch from 'node-fetch';
import { getAccessToken } from './token.js';

export async function fetchAvailable() {
    const token = await getAccessToken();

    const query = `
    query {
      MediaListCollection(userName: "${process.env.ANILIST_USERNAME}", type: ANIME, status_in: [PLANNING]) {
        lists {
          entries {
            media {
              id
              title { romaji english }
              episodes
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
            const watched = entry.progress || 0; // number of episodes you've watched
            const totalEpisodes = media.episodes || null; // total episodes of the anime
            const nextEpisodeNumber = watched + 1;

            // If the anime is finished or no episodes left to watch, skip
            if (media.status === "FINISHED" && nextEpisodeNumber > totalEpisodes) return null;

            // If there's no next airing episode and the anime isn't finished, skip
            if (media.status === "NOT_YET_RELEASED") return null;

            return {
                id: media.id,
                title: media.title.english || media.title.romaji,
                nextEpisode: nextEpisodeNumber,
                totalEpisodes,
                watched,
                url: `https://anilist.co/anime/${media.id}`,
                watchUrl: `https://anikai.to/browser?keyword=${media.title.english || media.title.romaji}`,
                status: media.status,
                airingAt: media.nextAiringEpisode?.airingAt,
                formattedAiring: media.nextAiringEpisode
                    ? new Date(media.nextAiringEpisode.airingAt * 1000).toLocaleString("en-US", {
                        month: "short",
                        day: "numeric",
                        hour: "2-digit",
                        minute: "2-digit",
                    })
                    : "Unknown",
            };
        })
        .filter(Boolean)
        .sort((a, b) => {
            // First: sort by progress (higher progress first)
            if (b.watched !== a.watched) {
                return b.watched - a.watched;
            }

            // Second: sort by airing time if progress is equal

            const aUnknown = a.airingAt === undefined || a.airingAt === null;
            const bUnknown = b.airingAt === undefined || b.airingAt === null;
            if (aUnknown && bUnknown) return 0;
            if (aUnknown) return 1;
            if (bUnknown) return -1;
            return a.airingAt - b.airingAt;
        });
}