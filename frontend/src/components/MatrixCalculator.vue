<template>
  <div class="calculator-container">
    <div class="input-section">
      <div class="dimension-controls">
        <div class="input-group">
          <label for="rows">Número de Filas:</label>
          <select id="rows" v-model="rows" @change="updateMatrix">
            <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
          </select>
        </div>
        
        <div class="input-group">
          <label for="columns">Número de Columnas:</label>
          <select id="columns" v-model="columns" @change="updateMatrix">
            <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
          </select>
        </div>
      </div>

      <div class="matrix-section">
        <h3>Matrix A:</h3>
        <div class="matrix-grid">          <div v-for="(row, i) in matrix" :key="i" class="matrix-row">
            <input
              v-for="(cell, j) in row"
              :key="j"
              type="number"
              v-model="matrix[i][j]"
              class="matrix-input"
            >
          </div>
        </div>
      </div>

      <div class="button-group">
        <button class="calculate-btn" @click="calculate" :disabled="loading">
          {{ loading ? 'Calculando...' : 'Calcular' }}
        </button>
        <button class="clear-btn" @click="clearMatrix">Limpiar Todo</button>
      </div>

      <div v-if="error" class="error-message">
        {{ error }}
      </div>
    </div>

    <div v-if="result" class="result-section">
      <div class="matrix-result">
        <h3>Matriz Q:</h3>
        <div class="matrix-grid">
          <div v-for="(row, i) in result.q" :key="i" class="matrix-row">
            <div v-for="(value, j) in row" :key="j" class="matrix-cell">
              {{ Number(value).toFixed(4) }}
            </div>
          </div>
        </div>
      </div>

      <div class="matrix-result">
        <h3>Matriz R:</h3>
        <div class="matrix-grid">
          <div v-for="(row, i) in result.r" :key="i" class="matrix-row">
            <div v-for="(value, j) in row" :key="j" class="matrix-cell">
              {{ Number(value).toFixed(4) }}
            </div>
          </div>
        </div>      </div>

      <div v-if="statistics" class="statistics-section">
        <h3>Estadísticas de la Matriz:</h3>
        <div class="stats-grid">
          <div class="stat-item">
            <span class="stat-label">Rango:</span>
            <span class="stat-value">{{ Number(statistics.productStats.rank).toFixed(4) }}</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">Cuadrado o no :</span>
            <span class="stat-value">{{ Number(statistics.qStats.determinant).toFixed(4) }}</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">Traza:</span>
            <span class="stat-value">{{ statistics.rStats.trace }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useMatrixStore } from '../stores/matrix'

export default {
  name: 'MatrixCalculator',
  
  data() {
    return {
      rows: 3,
      columns: 3,      matrix: Array(3).fill().map(() => Array(3).fill(0)),
      error: null,
      loading: false,
      matrixStore: null
    }
  },
  computed: {
    result() {
      return this.matrixStore?.result
    },
    statistics() {
      return this.matrixStore?.statistics
    }
  },

  created() {
    this.matrixStore = useMatrixStore()
  },

  methods: {
    updateMatrix() {
      this.matrix = Array(this.rows).fill().map(() => 
        Array(this.columns).fill(0)
      )
    },

    async calculate() {
      try {
        this.loading = true
        await this.matrixStore.calculateQR(this.matrix)
        this.error = null
      } catch (err) {
        this.error = err.message
        console.error('Error:', err)
      } finally {
        this.loading = false
      }
    },

    clearMatrix() {
      this.updateMatrix()
      this.matrixStore.result = null
      this.error = null
    }
  }
}
</script>

<style scoped>
.calculator-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.input-section {
  margin-bottom: 30px;
}

.dimension-controls {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.input-group label {
  font-weight: bold;
  color: #333;
}

select {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  width: 100px;
}

.matrix-section {
  margin-bottom: 20px;
}

h3 {
  margin-bottom: 10px;
  color: #2c3e50;
}

.matrix-grid {
  display: inline-block;
  padding: 15px;
  background: white;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
}

.matrix-row {
  display: flex;
  gap: 5px;
  margin-bottom: 5px;
}

.matrix-row:last-child {
  margin-bottom: 0;
}

.matrix-input {
  width: 60px;
  height: 40px;
  text-align: center;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

.matrix-cell {
  width: 60px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #eee;
  background: #f9f9f9;
  font-family: monospace;
}

.button-group {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}

.calculate-btn, .clear-btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.calculate-btn:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.calculate-btn {
  background-color: #4CAF50;
  color: white;
}

.calculate-btn:not(:disabled):hover {
  background-color: #45a049;
}

.clear-btn {
  background-color: #f44336;
  color: white;
}

.clear-btn:hover {
  background-color: #da190b;
}

.result-section {
  display: flex;
  gap: 30px;
  flex-wrap: wrap;
}

.matrix-result {
  flex: 1;
  min-width: 300px;
}

.error-message {
  color: #f44336;
  margin-top: 10px;
  padding: 10px;
  background-color: #ffebee;
  border: 1px solid #ef9a9a;
  border-radius: 4px;
}

@media (max-width: 768px) {
  .dimension-controls {
    flex-direction: column;
  }
  
  .matrix-input, .matrix-cell {
    width: 50px;
    height: 35px;
    font-size: 14px;
  }
}

.statistics-section {
  margin-top: 30px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-top: 15px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  padding: 15px;
  background: white;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.stat-label {
  font-size: 14px;
  color: #6c757d;
  margin-bottom: 5px;
}

.stat-value {
  font-size: 18px;
  font-weight: bold;
  color: #2c3e50;
  font-family: monospace;
}
</style>
