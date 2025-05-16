const math = require('mathjs');
const { calculateRank, isSquare, safeTrace } = require('../utils/matrixUtils');

function safeMatrixStats(matrix) {
    const stats = {
        trace: safeTrace(matrix),
        norm: math.sqrt(math.sum(math.dotMultiply(matrix, matrix))),
        rank: calculateRank(matrix),
        dimensions: math.size(matrix)
    };

    if (isSquare(matrix)) {
        stats.determinant = math.abs(math.det(matrix));
    }

    return stats;
}

function calculateStatistics(q, r) {
    const qStats = safeMatrixStats(q);
    const rStats = safeMatrixStats(r);

    const qSize = math.size(q);
    const rSize = math.size(r);
    const productStats = {
        dimensions: [qSize[0], rSize[1]]
    };

    if (qSize[1] === rSize[0]) {
        const qr = math.multiply(q, r);
        productStats.trace = safeTrace(qr);
        productStats.norm = math.sqrt(math.sum(math.dotMultiply(qr, qr)));
        productStats.rank = calculateRank(qr);
        
        if (isSquare(qr)) {
            productStats.determinant = math.abs(math.det(qr));
        }
    }

    return {
        qStats,
        rStats,
        productStats
    };
}

module.exports = {
    calculateStatistics
};
