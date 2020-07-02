import express from 'express';
import { DeclaracionService } from './declaracion.service';
import { getInvalidRequestResponse } from '../common/responses';

export const declaracionRouter = express.Router();

// Create a new assignment.
declaracionRouter.get('/:id', async (req, res) => {

  const { id } = req.params
  console.log(id)
  try {

    const declaration = await DeclaracionService.getDeclaration(id);

    res.json(declaration);

  }
  catch {
    res.status(400).json(getInvalidRequestResponse("Declaracion no válida"));
  }
});
