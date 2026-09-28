const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

const PROTO_PATH = path.resolve(__dirname, '../proto/arbforge.proto');

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const arbforgeProto = grpc.loadPackageDefinition(packageDefinition).arbforge;

const client = new arbforgeProto.ArbForgeService(
  'localhost:50051',
  grpc.credentials.createInsecure()
);

function callRPC(methodName, payload) {
  return new Promise((resolve, reject) => {
    client[methodName](payload, (err, response) => {
      if (err) return reject(err);
      resolve(response);
    });
  });
}

module.exports = {
  validateDeal: (productName, buyPrice) =>
    callRPC('ValidateDeal', {
      product_name: productName,
      buy_price: buyPrice,
    }),

  askMonica: (prompt) =>
    callRPC('AskMonica', {
      prompt,
    }),

  getConfidenceScore: (buyPrice, sellPrice) =>
    callRPC('GetConfidenceScore', {
      buy_price: buyPrice,
      sell_price: sellPrice,
    }),

  generateListing: (productName, buyPrice, sellPrice) =>
    callRPC('GenerateListing', {
      product_name: productName,
      buy_price: buyPrice,
      sell_price: sellPrice,
    }),
};
