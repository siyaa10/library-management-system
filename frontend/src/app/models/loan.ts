
export interface Loan {
  id?: number;
  bookId: number;
  memberId: number;
  issueDate: string;
  dueDate: string;
  returnedAt?: string | null;
  status: string;
}
