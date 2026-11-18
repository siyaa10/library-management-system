import { Router } from '@angular/router';
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Book } from '../../models/book';
import { LibraryService } from '../../services/library.service';

@Component({
  selector: 'app-books',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './books.html',
  styleUrl: './books.scss'
})
export class Books implements OnInit {
  books: Book[] = [];
  searchText = '';
  selectedCategory = 'All';
  sortBy = 'title';
  showForm = false;
  editingId: number | null = null;
  message = '';
  error = '';

  form: Book = this.emptyForm();

  constructor(
  private library: LibraryService,
  private cdr: ChangeDetectorRef,
  private router: Router
) {}

  ngOnInit(): void {
    this.loadBooks();
  }

  emptyForm(): Book {
    return {
      title: '',
      author: '',
      category: '',
      isbn: '',
      totalCopies: 1,
      availableCopies: 1,
      description: ''
    };
  }

  loadBooks(): void {
  this.library.getBooks().subscribe({
    next: data => {
      this.books = data;
      this.cdr.markForCheck();
    },
    error: () => {
      this.error = 'Could not load books. Check that the Go server is running.';
      this.cdr.markForCheck();
    }
  });
}

  get filteredBooks(): Book[] {
  const term = this.searchText.toLowerCase();

  return this.books
    .filter(book => {
      const matchesSearch =
        book.title.toLowerCase().includes(term) ||
        book.author.toLowerCase().includes(term) ||
        book.isbn.toLowerCase().includes(term);

      const matchesCategory =
        this.selectedCategory === 'All' ||
        book.category === this.selectedCategory;

      return matchesSearch && matchesCategory;
    })
    .sort((a, b) => {
      switch (this.sortBy) {
        case 'title':
          return a.title.localeCompare(b.title);
        case 'author':
          return a.author.localeCompare(b.author);
        case 'category':
          return a.category.localeCompare(b.category);
        case 'available':
          return b.availableCopies - a.availableCopies;
        default:
          return 0;
      }
    });
}

  get categories(): string[] {
    return ['All', ...new Set(this.books.map(book => book.category))];
  }

  openAddForm(): void {
    this.form = this.emptyForm();
    this.editingId = null;
    this.showForm = true;
    this.message = '';
    this.error = '';
  }

  openEditForm(book: Book): void {
    this.form = { ...book };
    this.editingId = book.id!;
    this.showForm = true;
    this.message = '';
    this.error = '';
  }

  closeForm(): void {
    this.showForm = false;
  }

  saveBook(): void {
    this.error = '';
this.message = '';
    if (!this.form.title.trim() || !this.form.author.trim() ||
        !this.form.category.trim() || !this.form.isbn.trim() ||
        this.form.totalCopies < 1) {
      this.error = 'Please complete all required fields correctly.';
      return;
    }
   

    const payload: Book = {
      ...this.form,
      availableCopies: this.editingId
        ? this.form.availableCopies
        : this.form.totalCopies
    };

    const request = this.editingId
      ? this.library.updateBook(this.editingId, payload)
      : this.library.addBook(payload);

    request.subscribe({
      next: () => {
        this.message = this.editingId
          ? 'Book updated successfully!'
          : 'Book added successfully!';
        this.showForm = false;
        this.loadBooks();
      },
      error: err => {
        this.error = err.error?.error || 'Unable to save book. Check the details and try again.';
      }
    });
  }
   clearValidationError(): void {
  if (
    this.form.title.trim() &&
    this.form.author.trim() &&
    this.form.category.trim() &&
    this.form.isbn.trim() &&
    this.form.totalCopies >= 1
  ) {
    this.error = '';
  }
}
clearMessages(): void {
  this.error = '';
  this.message = '';
}
  deleteBook(id: number): void {
    if (!confirm('Are you sure you want to delete this book?')) return;

    this.library.deleteBook(id).subscribe({
      next: () => {
        this.message = 'Book deleted successfully!';
        this.loadBooks();
      },
      error: err => {
  this.error = err.error?.error || 'Unable to delete this book.';
  this.message = '';
  this.cdr.markForCheck();
}
    });
  }
  viewBook(id: number): void {
  this.router.navigate(['/books', id]);
}
}
