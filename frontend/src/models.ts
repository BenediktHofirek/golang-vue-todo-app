export interface Todo {
  id: number;
  userId: string;
  title: string;
  description: string | null;
  completed: boolean;
  dueDate: string | null;
  createdAt: string;
  updatedAt: string;
}
