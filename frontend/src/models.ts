export interface Todo {
  id: number
  title: string | null
  description: string | null
  completed: boolean
  dueDate: string | null
  starred: boolean
  scheduledDate: string | null
  createdAt: string
  updatedAt: string
}
