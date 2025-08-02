import axios from 'axios';

const API = axios.create({ baseURL: import.meta.env.VITE_API_URL });

API.interceptors.request.use((req) => {
  const token = localStorage.getItem("token");
  if (token) req.headers.Authorization = token;
  return req;
});

export const createUser = (userData) => API.post('/users', userData);
export const login = (formData) => API.post('/users/login', formData);
export const getItems = () => API.get('/items');
export const getCarts = () => API.get('/carts');
export const createCart = (cartData) => API.post('/carts', cartData);
export const getOrders = () => API.get('/orders');
export const createOrder = (orderData) => API.post('/orders', orderData);
export const logout = () => API.post('/logout');
