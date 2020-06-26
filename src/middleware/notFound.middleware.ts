import { Request, Response } from "express";

export const notFoundHandler = (
  request: Request,
  response: Response,
): void => {

  const message = "Resource not found";

  response.status(404).send(message);
};
