import * as dotenv from 'dotenv';
dotenv.config();

import { db } from '../../src/db/connection';

import { getDeclaraciones } from './declaraciones';
import { DeclaracionModel } from '../../src/declaracion/declaracion.model';

(async (): Promise<void> => {
  if (!process.env.DB_HOST) {
    console.log('DB_HOST not specified, setting localhost by default.');

    process.env.DB_HOST = 'localhost';
  }

  await db.start();
  await db.drop();

  await Promise.all([
    // Use create to trigger pre-save functions.
    DeclaracionModel.create(getDeclaraciones()),
  ]);

  console.log('Data successfully inserted');

  process.exit(0);
})();
