import { model, Document, Schema, Types } from 'mongoose';

/**
 * Credito.
 */
interface Credito {
  deudor: String;
  clase: String;
  plazo: Number;
  importe: Number;
}

const CreditoSchema = new Schema({
  deudor: { type: String, required: true },
  clase: { type: String, required: true },
  plazo: { type: Number, required: true },
  importe: { type: Number, required: true },
}, { _id : false });

/**
 * Immueble.
 */
interface Inmueble {
  finca: String;
  padron: String;
  uso: String;
  pais: String;
  distrito: String;
  adquisicion: Number;
  tipoAdquisicion: String;
  superficie: Number;
  valorCompra: Number;
  valorAgregado: Number;
  importe: Number;
}

const InmuebleSchema = new Schema({
  finca: { type: String, required: true },
  padron: { type: String, required: true },
  uso: { type: String, required: true },
  pais: { type: String, required: true },
  distrito: { type: String, required: true },
  adquisicion: { type: Number, required: true },
  tipoAdquisicion: { type: String, required: true },
  superficie: { type: Number, required: true },
  valorCompra: { type: Number, required: true },
  valorAgregado: { type: Number, required: true },
  importe: { type: Number, required: true },
}, { _id : false });

/**
 * Vehiculo.
 */
interface Vehiculo {
  tipo: String;
  marca: String;
  modelo: String;
  adquisicion: Number;
  fabricacion: Number;
  importe: Number;
}

const VehiculoSchema = new Schema({
  tipo: { type: String, required: true },
  marca: { type: String, required: true },
  modelo: { type: String, required: true },
  adquisicion: { type: Number, required: true },
  fabricacion: { type: Number, required: true },
  importe: { type: Number, required: true },
}, { _id : false });


/**
 * Actividad Agropecuaria.
 */
interface Agropecuaria {
  tipo: String;
  ubicacion: String;
  especie: String;
  cantidad: Number;
  precio: Number;
  importe: Number;
}

const AgropecuariaSchema = new Schema({
  tipo: { type: String, required: true },
  ubicacion: { type: String, required: true },
  especie: { type: String, required: true },
  cantidad: { type: Number, required: true },
  precio: { type: Number, required: true },
  importe: { type: Number, required: true },
}, { _id : false });


/**
 * Deuda.
 */
interface Deuda {
  tipo: String;
  empresa: String;
  plazo: Number;
  cuotaMensual: Number;
  total: Number;
  saldo: Number;
}

const DeudaSchema = new Schema({
  tipo: { type: String, required: true },
  empresa: { type: String, required: true },
  plazo: { type: Number, required: true },
  cuotaMensual: { type: Number, required: true },
  total: { type: Number, required: true },
  saldo: { type: Number, required: true },
}, { _id : false });

interface Resumen {
  activo: Number;
  pasivo: Number;
}

/**
 * An declaration represents a declaration for a given year.
 */
export interface Declaracion {
  // Information about the public official
  cedula: Number;
  nombre: String;
  apellido: String;
  nombreCompleto: String;
  funcion?: String;
  institucion?: String;

  ano: Number;

  // Activos
  creditos?: Credito[];
  inmuebles?: Inmueble[];
  vehiculos?: Vehiculo[];
  agropecuaria?: Agropecuaria[];

  deuda?: Deuda[];

  resumen: Resumen;
}

interface DeclaracionDocument extends Declaracion, Document {}

/**
 * Schema.
 */
const DeclaracionSchema = new Schema({

  cedula: { type: Number, required: true },
  nombre: { type: String, required: true },
  apellido: { type: String, required: true },
  nombreCompleto: { type: String, require: true },
  funcion: { type: String, required: false },
  institucion: { type: String, required: false },

  ano: { type: Number, required: true },

  creditos: [CreditoSchema],
  inmuebles: [InmuebleSchema],
  vehiculos: [VehiculoSchema],
  agropecuaria: [AgropecuariaSchema],

  deudas: [DeudaSchema],

  resumen: { type: Object, required: true},
});

/**
 * Indexes.
 */
DeclaracionSchema.index({cedula: 1});
DeclaracionSchema.index({ano: 1});
DeclaracionSchema.index({nombreCompleto: 1});

export const DeclaracionModel = model<DeclaracionDocument>('Declaration', DeclaracionSchema);
