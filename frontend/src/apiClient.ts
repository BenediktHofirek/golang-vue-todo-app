import axios from 'axios'
// import { auth } from '@/firebase/config'; // Assume you export firebase auth here

export const apiClient = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
})

apiClient.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    return Promise.reject(error)
  },
)

// REQUEST Interceptor for Auth Token
// apiClient.interceptors.request.use(async (config) => {
//   const user = auth.currentUser;
//   if (user) {
//     const token = await user.getIdToken();
//     config.headers.Authorization = `Bearer ${token}`;
//   }
//   return config;
// }, (error) => {
//   return Promise.reject(error);
// });
//
// // RESPONSE Interceptor (optional, for returning data directly)
// apiClient.interceptors.response.use(
//     (response) => response.data,
//     (error) => Promise.reject(error)
// );
