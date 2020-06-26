import mongoose = require('mongoose');
import { logger } from '../common/logger';

export class Database {
  db: string;
  client: mongoose.Mongoose;

  constructor() {
    this.db = process.env.DB_NAME;
  }

  /**
   * Start a connection to the database.
   */
  async start(): Promise<void> {
    const uri = this.getUri();

    logger.info('Trying to connect to database at ' + uri);

    this.client = await mongoose.connect(uri, {
      autoIndex: (process.env.DB_AUTOINDEX === 'true'),
      useCreateIndex: true,
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });

    logger.info('Connected to database');
  }

  async drop(): Promise<void> {
    return this.client.connection.dropDatabase();
  }

  /**
   * Generate the URI value to connect to the MongoDB cluster.
   */
  private getUri(): string {
    const { DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD} = process.env;

    let uri = `${DB_HOST}:${DB_PORT}/${this.db}`;

    if (DB_USERNAME) {
      uri = `${DB_USERNAME}:${DB_PASSWORD}@${uri}`;
    }

    uri = `mongodb://${uri}`

    return uri;
  }

}

