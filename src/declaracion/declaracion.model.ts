import { model, Document, Schema, Types } from 'mongoose';

/**
 * Deposito.
 */
interface Deposito {
  tipoEntidad: String;
  entidad: String;
  tipo: String;
  pais: String;
  importe: Number;
}

const DepositoSchema = new Schema({
  tipoEntidad: { type: String, required: true },
  entidad: { type: String, required: true },
  tipo: { type: String, required: true },
  pais: { type: String, require: true},
  importe: { type: Number, required: true },
}, { _id : false });

/**
 * Deudor.
 */
interface Deudor {
  nombre: String;
  clase: String;
  plazo: Number;
  importe: Number;
}

const DeudorSchema = new Schema({
  nombre: { type: String, required: true },
  clase: { type: String, required: true },
  plazo: { type: Number, required: true },
  importe: { type: Number, required: true },
}, { _id : false });

/**
 * Immueble.
 */
interface Inmueble {
  padron: String;
  uso: String;
  pais: String;
  distrito: String;
  adquisicion: Number;
  tipoAdquisicion: String;
  superficieTerreno: Number;
  valorTerreno: Number;
  superficieConstruccion: Number;
  valorConstruccion: Number;
  importe: Number;
}

const InmuebleSchema = new Schema({
  padron: { type: String, required: true },
  uso: { type: String, required: true },
  pais: { type: String, required: true },
  distrito: { type: String, required: true },
  adquisicion: { type: Number, required: true },
  tipoAdquisicion: { type: String, required: true },
  superficie: { type: Number, required: true },
  valorTerreno: { type: Number, required: true },
  superficieConstruccion: { type: Number, required: true },
  valorConstruccion: { type: Number, required: true },
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
 * Mueble.
 */
interface Mueble {
  tipo: String;
  importe: Number;
}

const MuebleSchema = new Schema({
  tipo: { type: String, required: true },
  importe: { type: Number, required: true },
}, { _id : false });

/**
 * Otros Activo.
 */
interface OtroActivo {
  descripcion: String;
  empresa: String;
  ruc: String;
  pais: String;
  cantidad: Number;
  precio: Number;
  importe: Number;
}

const OtroActivoSchema = new Schema({
  descripcion: { type: String, required: true },
  empresa: { type: String, required: true },
  ruc: { type: String, required: true },
  pais: { type: String, required: true },
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
  cuota: Number;
  total: Number;
  saldo: Number;
}

const DeudaSchema = new Schema({
  tipo: { type: String, required: true },
  empresa: { type: String, required: true },
  plazo: { type: Number, required: true },
  cuota: { type: Number, required: true },
  total: { type: Number, required: true },
  saldo: { type: Number, required: true },
}, { _id : false });

/**
 * An declaration represents a declaration for a given year.
 */
export interface Declaracion {
  // Information about the public official
  fecha: Date;
  cedula: Number;
  nombre: String;
  apellido: String;
  nombreCompleto: String;
  cargo: String;
  institucion: String;

  // Activos
  depositos?: Deposito[];
  deudores?: Deudor[];
  inmuebles?: Inmueble[];
  vehiculos?: Vehiculo[];
  actividadesAgropecuarias?: Agropecuaria[];
  muebles?: Mueble[];
  otrosActivos?: OtroActivo[];

  deudas?: Deuda[];

  ingresosMensual: Number;
  ingresosAnual: Number;
  egresosMensual: Number;
  egresosAnual: Number;

  activos: Number;
  pasivos: Number;
  patrimonioNeto: Number;
}

interface DeclaracionDocument extends Declaracion, Document {}

/**
 * Schema.
 */
const DeclaracionSchema = new Schema({

  cedula: { type: Number, required: true },
  nombre: { type: String, required: true },
  apellido: { type: String, required: true },
  cargo: { type: String, required: true },
  institucion: { type: String, required: true },

  fecha: { type: Date, required: true },

  depositos: [DepositoSchema],
  deudores: [DeudorSchema],
  inmuebles: [InmuebleSchema],
  vehiculos: [VehiculoSchema],
  actividadesAgropecuarias: [AgropecuariaSchema],
  muebles: [MuebleSchema],
  otrosActivos: [OtroActivoSchema],

  deudas: [DeudaSchema],

  ingresosMensual: { type: Number, required: true },
  ingresosAnual: { type: Number, required: true },
  egresosMensual: { type: Number, required: true },
  egresosAnual: { type: Number, required: true },

  activos: { type: Number, required: true },
  pasivos: { type: Number, required: true },
  patrimonioNeto: { type: Number, required: true },
});

/**
 * Indexes.
 */
DeclaracionSchema.index({cedula: 1});
DeclaracionSchema.index({fecha: 1});

export const DeclaracionModel = model<DeclaracionDocument>('Declaration', DeclaracionSchema);
