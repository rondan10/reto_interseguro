const express = require('express');
const router = express.Router();
const { calculateStatistics } = require('../services/matrixService');

router.post('/process-qr', (req, res) => {
    try {
        const { q, r } = req.body;
        
        if (!q || !r) {
            return res.status(400).json({ error: 'Missing Q or R matrix' });
        }

        const statistics = calculateStatistics(q, r);
        res.json(statistics);
    } catch (error) {
        console.error('Error processing matrices:', error);
        res.status(500).json({ error: 'Error processing matrices' });
    }
});

router.get('/health', (req, res) => {
    res.send('OK');
});

module.exports = router;
