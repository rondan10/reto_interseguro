const { calculateStatistics } = require('../services/matrixService');

describe('Matrix Service Tests', () => {
  test('should calculate statistics for square matrices', () => {
    const q = [
      [1, 0],
      [0, 1]
    ];
    const r = [
      [2, 0],
      [0, 2]
    ];

    const result = calculateStatistics(q, r);

    expect(result).toHaveProperty('qStats');
    expect(result).toHaveProperty('rStats');
    expect(result).toHaveProperty('productStats');

    expect(result.qStats).toHaveProperty('trace');
    expect(result.qStats).toHaveProperty('norm');
    expect(result.qStats).toHaveProperty('rank');
  });

  test('should handle rectangular matrices', () => {
    const q = [
      [1, 0, 0],
      [0, 1, 0],
      [0, 0, 1]
    ];
    const r = [
      [2, 0],
      [0, 2],
      [0, 0]
    ];

    const result = calculateStatistics(q, r);

    expect(result.qStats.dimensions).toEqual([3, 3]);
    expect(result.rStats.dimensions).toEqual([3, 2]);
  });
});
