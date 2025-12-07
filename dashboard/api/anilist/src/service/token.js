import fetch from 'node-fetch';
import dotenv from 'dotenv';

dotenv.config();
const { ANILIST_CLIENT_ID, ANILIST_CLIENT_SECRET } = process.env;

export async function getAccessToken() {
    if (!ANILIST_CLIENT_ID || !ANILIST_CLIENT_SECRET) {
        throw new Error('ANILIST_CLIENT_ID and ANILIST_CLIENT_SECRET must be set in env');
    }


    const params = new URLSearchParams();
    params.append('grant_type', 'client_credentials');
    params.append('client_id', ANILIST_CLIENT_ID);
    params.append('client_secret', ANILIST_CLIENT_SECRET);


    const res = await fetch('https://anilist.co/api/v2/oauth/token', {
        method: 'POST',
        body: params,
    });


    if (!res.ok) {
        const body = await res.text();
        throw new Error(`Failed to get access token: ${res.status} - ${body}`);
    }


    const { access_token } = await res.json();
    return access_token;
}