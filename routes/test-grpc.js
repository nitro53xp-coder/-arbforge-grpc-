const express = require('express');
const grpcClient = require('../services/grpcClient');
const router = express.Router();

router.post('/test-grpc', async (req, res) => {
  try {
    const { action, productName, buyPrice, sellPrice } = req.body;
    let result;

    switch (action) {
      case 'validate':
        result = await grpcClient.validateDeal(productName || '2026 Hot Wheels', buyPrice || 1.99);
        break;

      case 'monica':
        result = await grpcClient.askMonica(productName || 'Best flips this week in Lafayette');
        break;

      case 'confidence':
        result = await grpcClient.getConfidenceScore(buyPrice || 5, sellPrice || 35);
        break;

      case 'generate':
        result = await grpcClient.generateListing(productName || 'Hot Wheels STH', buyPrice || 2, sellPrice || 35);
        break;

      default:
        return res.status(400).json({
          error: 'Invalid action. Use: validate, monica, confidence, or generate',
        });
    }

    res.json({
      success: true,
      action,
      result,
      message: 'gRPC call successful from Node.js',
    });
  } catch (err) {
    console.error('gRPC test error:', err);
    res.status(500).json({ error: err.message });
  }
});

module.exports = router;
