import pino = require('pino');

export const logger = pino({
  level: process.env.DEBUG_LEVEL,
  prettyPrint: process.env.NODE_ENV === 'development',
});
