import axios from 'axios'
import { getCurrentUser } from './firebase'

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
apiClient.interceptors.request.use(async (config) => {
  const user = await getCurrentUser();

  if (user) {
    const token = await user.getIdToken();
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});
