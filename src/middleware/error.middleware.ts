import { Request, Response, NextFunction } from 'express';

import HttpException from '../common/http-exception';
import { logger } from '../common/logger';

/**
 * Four arguments are needed to identify a function as an error-handling
 * middleware function in Express.
 *
 * @see https://expressjs.com/en/guide/using-middleware.html#middleware.error-handling
 */
export const errorHandler = (
  error: HttpException,
  request: Request,
  response: Response,
  next: NextFunction // eslint-disable-line
): void => {

  logger.error('errorHandler called.');
  logger.error(error.message);

  const status = 500;
  const message = "It's not you. It's us. We are having some problems.";

  response.status(status).send(message);
};
