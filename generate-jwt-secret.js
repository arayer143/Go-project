const crypto = require('crypto');

function generateJWTSecret(length = 64) {
  return crypto.randomBytes(length).toString('hex');
}

const jwtSecret = generateJWTSecret();
console.log('Your JWT Secret:');
console.log(jwtSecret);
console.log('\nMake sure to keep this secret secure and don\'t share it publicly!');