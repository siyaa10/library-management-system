
export interface Book {
  id?: number;
  title: string;
  author: string;
  category: string;
  isbn: string;
  totalCopies: number;
  availableCopies: number;
  description?: string | null;
}
