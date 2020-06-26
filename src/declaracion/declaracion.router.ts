import express from 'express';
import { DeclaracionService } from './declaracion.service';
import { getInvalidRequestResponse } from '../common/responses';

export const declaracionRouter = express.Router();

// Create a new assignment.
declaracionRouter.get('/', async (req, res) => {

  const { cedula } = req.query

  try {

    const parsedCedula = parseInt(cedula as string);
    const declarations = await DeclaracionService.getDeclarations(parsedCedula);

    res.json(declarations);

  }
  catch {
    res.status(400).json(getInvalidRequestResponse("Cédula inválida"));
  }
});
