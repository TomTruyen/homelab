import dotenv from 'dotenv';
import express from 'express';
import { fetchAvailable } from './service/available.js';
import { fetchUpcoming } from './service/upcoming.js';

const app = express();
const PORT = process.env.PORT || 3000;
app.get('/available', async (req, res) => {
  try {
    const available = await fetchAvailable();
    res.json({ items: available });
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Failed to fetch available episodes' });
  }
});

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