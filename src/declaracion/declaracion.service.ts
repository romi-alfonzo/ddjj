import mongoose = require('mongoose');
import { Declaracion, DeclaracionModel } from './declaracion.model';

export interface DeclaracionAno {

}

class Service {

  async getDeclaration(id: string): Promise<Declaracion> {

    mongoose.set('debug', true);
    return DeclaracionModel.findById(id);

  }

}

export const DeclaracionService = new Service();
