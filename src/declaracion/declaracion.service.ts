import mongoose = require('mongoose');
import { Declaracion, DeclaracionModel } from './declaracion.model';

export interface DeclaracionAno {

}

class Service {

  async getDeclarations(cedula: Number): Promise<Declaracion[]> {

    return DeclaracionModel.find({cedula: cedula}).sort({ano: 1});

  }

}

export const DeclaracionService = new Service();
