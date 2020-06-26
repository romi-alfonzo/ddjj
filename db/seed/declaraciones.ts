import mongoose = require('mongoose');
import { Declaracion } from '../../src/declaracion/declaracion.model';

export const getDeclaraciones = (): Declaracion[] => {

  return [
    CarlosPalacios,
    CarlosPalacios2,
  ];
};

const CarlosPalacios: Declaracion = {

  cedula: 495050,
  nombre: 'CARLOS NESTOR',
  apellido: 'PALACIOS OCAMPOS',
  nombreCompleto: 'CARLOS NESTOR PALACIOS OCAMPOS',

  ano: 2016,

  creditos: [{
    deudor: 'VARIOS DEUDORES',
    clase: 'A LA VISTA',
    plazo: 1,
    importe: 1100000
  }],

  inmuebles: [{
    finca: 'DATOS PROTEGIDOS',
    padron: '3192',
    uso: 'ESTANCIA',
    superficie: 10000000,
    valorCompra: 120000000000,
    valorAgregado: 500000000,
    importe: 12500000000,
    pais: 'PARAGUAY',
    distrito: 'CHACO',
    adquisicion: 2015,
    tipoAdquisicion: 'COMPRA',
  }],

  vehiculos: [{
    tipo: 'CAMIONETA',
    marca: 'TOYOTA HILUX',
    modelo: 'DOBLE CABINA',
    fabricacion: 2014,
    adquisicion: 2014,
    importe: 200000000,
  }],

  agropecuaria: [{
    tipo: 'VACUNO',
    ubicacion: 'CHACO',
    especie: 'BRANGUS',
    cantidad: 1000,
    precio: 2500000,
    importe: 25000000000,
  }],

  resumen: {
    activo: 35980000000,
    pasivo: 1661000000,
  }

}

const CarlosPalacios2: Declaracion = {

  cedula: 495050,
  nombre: 'CARLOS NESTOR',
  apellido: 'PALACIOS OCAMPOS',
  nombreCompleto: 'CARLOS NESTOR PALACIOS OCAMPOS',

  ano: 2017,

  creditos: [{
    deudor: 'VARIOS DEUDORES',
    clase: 'A LA VISTA',
    plazo: 1,
    importe: 1100000
  }],

  inmuebles: [{
    finca: 'DATOS PROTEGIDOS',
    padron: '3192',
    uso: 'ESTANCIA',
    superficie: 10000000,
    valorCompra: 120000000000,
    valorAgregado: 500000000,
    importe: 12500000000,
    pais: 'PARAGUAY',
    distrito: 'CHACO',
    adquisicion: 2015,
    tipoAdquisicion: 'COMPRA',
  }],

  vehiculos: [{
    tipo: 'CAMIONETA',
    marca: 'TOYOTA HILUX',
    modelo: 'DOBLE CABINA',
    fabricacion: 2014,
    adquisicion: 2014,
    importe: 200000000,
  }],

  agropecuaria: [{
    tipo: 'VACUNO',
    ubicacion: 'CHACO',
    especie: 'BRANGUS',
    cantidad: 1000,
    precio: 2500000,
    importe: 25000000000,
  }],

  resumen: {
    activo: 35980000000,
    pasivo: 1661000000,
  }

}

