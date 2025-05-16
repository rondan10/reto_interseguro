import { defineStore } from 'pinia'
import axios from 'axios'

// Token JWT fijo para las APIs (en producción debería venir de variables de entorno)
const API_TOKEN = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpLWFjY2VzcyJ9.0q6QFJ4YMoKC1RiHtO-PKzIcOqhdU2qZVhb7735kcX4'

// Crear instancias de axios para cada API
const goAPI = axios.create({
  baseURL: 'http://localhost:3000',
  headers: {
    'Authorization': `Bearer ${API_TOKEN}`
  }
})

const nodeAPI = axios.create({
  baseURL: 'http://localhost:3001',
  headers: {
    'Authorization': `Bearer ${API_TOKEN}`
  }
})

export const useMatrixStore = defineStore('matrix', {
  state: () => ({
    result: null,
    error: null,
    loading: false,
    statistics: null
  }),

  actions: {
    async calculateQR(matrix) {
      this.loading = true
      this.error = null;
      try {        // Primero obtengo  la factorización QR de la api go
        console.log('Enviando matriz a API Go:', matrix);
        const qrResponse = await goAPI.post('/api/qr', { matrix });
        console.log('Respuesta de API Go:', qrResponse.data);
        
        // Luego envio el  Q y R a la API del node para las estadísticas
        const dataForNode = {
          q: qrResponse.data.q,
          r: qrResponse.data.r
        };
        console.log('Enviando a API Node:', dataForNode);
        
        const statsResponse = await nodeAPI.post('api/process-qr', dataForNode);
        console.log('Respuesta de API Node:', statsResponse.data);

        this.result = qrResponse.data;
        this.statistics = statsResponse.data;
        console.log('Estado final:', {
          result: this.result,
          statistics: this.statistics
        });

        return {
          qr: qrResponse.data,
          stats: statsResponse.data
        }
      } catch (err) {
        this.error = err.message
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
