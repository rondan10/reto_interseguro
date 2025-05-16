const request = require('supertest');
const express = require('express');
const matrixRoutes = require('../routes/matrixRoutes');

const app = express();
app.use(express.json());
app.use('/api', matrixRoutes);

describe('Matrix API Integration Tests', () => {
  test('POST /api/process-qr should process QR factorization', async () => {
    const response = await request(app)
      .post('/api/process-qr')
      .send({
        q: [
          [1, 0],
          [0, 1]
        ],
        r: [
          [2, 0],
          [0, 2]
        ]
      });

    expect(response.status).toBe(200);
    expect(response.body).toHaveProperty('qStats');
    expect(response.body).toHaveProperty('rStats');
    expect(response.body).toHaveProperty('productStats');
  });

  test('should handle invalid input', async () => {
    const response = await request(app)
      .post('/api/process-qr')
      .send({
        q: [],
        r: []
      });

    expect(response.status).toBe(400);
  });
});
