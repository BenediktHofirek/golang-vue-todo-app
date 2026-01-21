import { NEW_TODO_TEMP_ID } from './constants'
import type { Todo } from './models'

export function isTodoBlank(todo: Todo) {
  if (todo.id === NEW_TODO_TEMP_ID) return true

  return (
    todo.title === null &&
    todo.description === null &&
    todo.starred === false &&
    todo.completed === false &&
    todo.scheduledDate === null &&
    todo.dueDate === null
  )
}
