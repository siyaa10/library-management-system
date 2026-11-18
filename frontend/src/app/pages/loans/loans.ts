
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Book } from '../../models/book';
import { Member } from '../../models/member';
import { Loan } from '../../models/loan';
import { LibraryService } from '../../services/library.service';

@Component({
  selector: 'app-loans',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './loans.html',
  styleUrl: './loans.scss'
})
export class Loans implements OnInit {
  books: Book[] = [];
  members: Member[] = [];
  loans: Loan[] = [];
  selectedBookId = 0;
  selectedMemberId = 0;
  dueDate = '';
  message = '';
  error = '';

  constructor(
  private library: LibraryService,
  private cdr: ChangeDetectorRef
) {}

  ngOnInit(): void {
    this.loadData();
  }

  
loadData(): void {
  this.library.getBooks().subscribe({
    next: data => {
      this.books = data;
      this.cdr.markForCheck();
    },
    error: () => {
      this.error = 'Could not load books.';
      this.cdr.markForCheck();
    }
  });

  this.library.getMembers().subscribe({
    next: data => {
      this.members = data;
      this.cdr.markForCheck();
    },
    error: () => {
      this.error = 'Could not load members.';
      this.cdr.markForCheck();
    }
  });

  this.library.getLoans().subscribe({
    next: data => {
      this.loans = data;
      this.cdr.markForCheck();
    },
    error: () => {
      this.error = 'Could not load loans.';
      this.cdr.markForCheck();
    }
  });
}


  issueBook(): void {
    this.message = '';
    this.error = '';

    if (!this.selectedBookId || !this.selectedMemberId || !this.dueDate) {
      this.error = 'Please select a book, member and due date.';
      return;
    }

    const book = this.books.find(b => b.id === this.selectedBookId);
    if (!book || book.availableCopies < 1) {
      this.error = 'This book is currently unavailable.';
      return;
    }

    this.library.issueBook(
      this.selectedBookId,
      this.selectedMemberId,
      this.dueDate
    ).subscribe({
      next: () => {
        this.message = 'Book issued successfully!';
        this.selectedBookId = 0;
        this.selectedMemberId = 0;
        this.dueDate = '';
        this.loadData();
      },
      error: err => {
        this.error = err.error?.error || 'Unable to issue book.';
      }
    });
  }

  returnBook(id: number): void {
    this.message = '';
    this.error = '';

    if (!confirm('Confirm that this book is being returned?')) return;

    this.library.returnBook(id).subscribe({
      next: () => {
        this.message = 'Book returned successfully!';
        this.loadData();
      },
      error: err => {
        this.error = err.error?.error || 'Unable to return book.';
      }
    });
  }
}
