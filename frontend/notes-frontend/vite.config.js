import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/signup': 'http://localhost:8080',
      '/login': 'http://localhost:8080',
      '/notes': 'http://localhost:8080',
    },
  },
});
