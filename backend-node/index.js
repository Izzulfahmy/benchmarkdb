const express = require('express');
const { Pool } = require('pg');

const app = express();
const pool = new Pool({
  user: 'postgres',
  host: 'localhost',
  database: 'benchmarkdb',
  password: '@Vinceru2',
  port: 5432,
});

// Endpoint ambil 10 user
app.get('/users', async (req, res) => {
  try {
    const result = await pool.query('SELECT id, name, email FROM users LIMIT 10');
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'DB error' });
  }
});

app.listen(8081, () => {
  console.log('Node.js server running at http://127.0.0.1:8081');
});
