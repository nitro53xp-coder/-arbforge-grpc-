const express = require('express');
const grpcRoutes = require('./routes/test-grpc');

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// gRPC Routes
app.use('/api', grpcRoutes);

// Health check
app.get('/health', (req, res) => {
  res.json({ status: 'ok', service: 'arbforge-grpc-bridge' });
});

// 404 Handler
app.use((req, res) => {
  res.status(404).json({ error: 'Route not found' });
});

// Start server
app.listen(PORT, () => {
  console.log(`🚀 Express server running on http://localhost:${PORT}`);
  console.log(`📡 gRPC endpoint: localhost:50051`);
  console.log(`✅ gRPC bridge initialized`);
});

module.exports = app;
