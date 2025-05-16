const math = require('mathjs');

function findPivotRow(A, col, rank, rows, eps) {
    let pivot_row = rank;
    while (pivot_row < rows && Math.abs(A[pivot_row][col]) < eps) {
        pivot_row++;
    }
    return pivot_row;
}

function eliminateRows(A, col, rank, rows, cols) {
    for (let row = rank; row < rows; row++) {
        const factor = A[row][col] / A[rank - 1][col];
        for (let col2 = col; col2 < cols; col2++) {
            A[row][col2] -= factor * A[rank - 1][col2];
        }
    }
}

function calculateRank(matrix) {
    const m = math.matrix(matrix);
    const rows = m.size()[0];
    const cols = m.size()[1];
    
    let A = m.toArray();
    let rank = 0;
    const eps = 1e-10;

    for (let col = 0; col < cols && rank < rows; col++) {
        let pivot_row = findPivotRow(A, col, rank, rows, eps);
        
        if (pivot_row < rows) {
            rank++;
            
            if (pivot_row !== rank - 1) {
                [A[rank - 1], A[pivot_row]] = [A[pivot_row], A[rank - 1]];
            }
            
            eliminateRows(A, col, rank, rows, cols);
        }
    }
    
    return rank;
}

const isSquare = (matrix) => {
    const size = math.size(matrix);
    return size[0] === size[1];
};

const safeTrace = (matrix) => {
    const size = math.size(matrix);
    const minDim = Math.min(size[0], size[1]);
    let trace = 0;
    for (let i = 0; i < minDim; i++) {
        trace += matrix[i][i];
    }
    return trace;
};

module.exports = {
    calculateRank,
    isSquare,
    safeTrace
};
