
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap, catchError, of } from 'rxjs';
import { Book } from '../models/book';
import { Member } from '../models/member';
import { Loan } from '../models/loan';

@Injectable({
  providedIn: 'root'
})
export class LibraryService {
  private apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  // Dashboard
  getDashboard(): Observable<any> {
    return this.http.get(`${this.apiUrl}/dashboard`);
  }

  // Books
  getBooks(): Observable<Book[]> {
  return this.http.get<Book[]>(`${this.apiUrl}/books`).pipe(
    tap(data => localStorage.setItem('books', JSON.stringify(data))),
    catchError(() => {
      const cached = localStorage.getItem('books');
      return of(cached ? JSON.parse(cached) as Book[] : []);
    })
  );
  
}
getBookById(id: number): Observable<Book> {
  return this.http.get<Book>(`${this.apiUrl}/books/${id}`);
}

  addBook(book: Book): Observable<Book> {
    return this.http.post<Book>(`${this.apiUrl}/books`, book);
  }

  updateBook(id: number, book: Book): Observable<any> {
    return this.http.put(`${this.apiUrl}/books/${id}`, book);
  }

  deleteBook(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/books/${id}`);
  }

  searchBooks(query: string): Observable<Book[]> {
    return this.http.get<Book[]>(
      `${this.apiUrl}/books?search=${encodeURIComponent(query)}`
    );
  }

  // Members
  getMembers(): Observable<Member[]> {
  return this.http.get<Member[]>(`${this.apiUrl}/members`).pipe(
    tap(data => localStorage.setItem('members', JSON.stringify(data))),
    catchError(() => {
      const cached = localStorage.getItem('members');
      return of(cached ? JSON.parse(cached) as Member[] : []);
    })
  );
}
getMemberById(id: number): Observable<Member> {
  return this.http.get<Member>(`${this.apiUrl}/members/${id}`);
}

  addMember(member: Member): Observable<Member> {
    return this.http.post<Member>(`${this.apiUrl}/members`, member);
  }

  updateMember(id: number, member: Member): Observable<any> {
    return this.http.put(`${this.apiUrl}/members/${id}`, member);
  }

  deleteMember(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/members/${id}`);
  }

  // Loans
  getLoans(): Observable<Loan[]> {
  return this.http.get<Loan[]>(`${this.apiUrl}/loans`).pipe(
    tap(data => localStorage.setItem('loans', JSON.stringify(data))),
    catchError(() => {
      const cached = localStorage.getItem('loans');
      return of(cached ? JSON.parse(cached) as Loan[] : []);
    })
  );
}

  issueBook(bookId: number, memberId: number, dueDate: string): Observable<Loan> {
    return this.http.post<Loan>(`${this.apiUrl}/loans`, {
      bookId,
      memberId,
      dueDate
    });
  }

  returnBook(id: number): Observable<any> {
    return this.http.put(`${this.apiUrl}/loans/${id}/return`, {});
  }

  getOverdueLoans(): Observable<Loan[]> {
    return this.http.get<Loan[]>(`${this.apiUrl}/loans/overdue`);
  }

  getMemberHistory(id: number): Observable<Loan[]> {
    return this.http.get<Loan[]>(`${this.apiUrl}/members/${id}/history`);
  }

  getBookHistory(id: number): Observable<Loan[]> {
    return this.http.get<Loan[]>(`${this.apiUrl}/books/${id}/history`);
  }
}
