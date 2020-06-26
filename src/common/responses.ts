
interface HttpResponse {
  status: number;
  message?: string;
}

export const InvalidRequest: HttpResponse = {
  status: 400,
  message: 'Invalid request or parameters',
};

export const getInvalidRequestResponse = (message: string): HttpResponse => {
  return {
    status: 400,
    message: message,
  }
}
