import axios from 'axios'
import { getCurrentUser } from './firebase'
import type { Todo } from './models'

const apiClient = axios.create({
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
apiClient.interceptors.request.use(
  async (config) => {
    const user = await getCurrentUser()

    if (user) {
      const token = await user.getIdToken()
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

const getTodoList = ({ signal }: { signal: AbortSignal }) => {
  return apiClient.get('/todos', { signal }) as Promise<Todo[]>
}

const updateTodo = ({
  id,
  ...updatedTodo
}: Partial<Todo>) => {
  return apiClient.patch(`/todos/${id}`, updatedTodo) as Promise<Todo>
}

const deleteTodo = (id: number) => {
  return apiClient.delete(`/todos/${id}`) as Promise<void>
}

const createBlankTodo = () => {
  return apiClient.post('/todos/create-blank') as Promise<number>
}

export const api = {
  getTodoList,
  updateTodo,
  deleteTodo,
  createBlankTodo,
}
