const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');
const matrixRoutes = require('./src/routes/matrixRoutes');

const app = express();
const port = 3001;

app.use(cors());
app.use(bodyParser.json());

// Usar las rutas de matriz
app.use('/api', matrixRoutes);

app.listen(port, () => {
    console.log(`Node.js API listening at http://localhost:${port}`);
});
